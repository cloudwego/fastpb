// Copyright 2022 ByteDance Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var IDLDir string

func IDLSearcher() <-chan (string) {
	out := make(chan (string))
	// read file path from stdin
	if IDLDir == "" {
		br := bufio.NewReader(os.Stdin)
		go func() {
			str, err := br.ReadString('\n')
			for ; err == nil; str, err = br.ReadString('\n') {
				out <- str
			}
			if err == io.EOF {
				close(out)
			}
			log.Fatalln(fmt.Errorf("read from stdin failed: %w", err))
		}()
		return out
	}
	// search ThriftDir to find thrift files
	dirs := make(chan (string))
	var searchDir func(string)
	searchDir = func(dir string) {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Println(fmt.Errorf("read directory %s failed: %w", dir, err))
		}
		dirs <- dir
		for _, file := range files {
			if file.IsDir() {
				searchDir(filepath.Join(IDLDir, file.Name()))
			}
		}
	}
	go func() {
		searchDir(IDLDir)
		close(dirs)
	}()
	go func() {
		for dir := range dirs {
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Println(fmt.Errorf("read directory %s failed: %w", dir, err))
			}
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".proto") {
					out <- filepath.Join(dir, file.Name())
				}
			}
		}
		close(out)
	}()
	return out
}

func main() {
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalf("create temp dir failed: %s", err.Error())
	}
	os.RemoveAll(dir)
	for protoFile := range IDLSearcher() {
		content, err := ioutil.ReadFile(protoFile)
		if err != nil {
			log.Fatalf("read %s file failed: %s", protoFile, err.Error())
		}
		if bytes.Contains(content, []byte(`"proto2"`)) {
			continue
		}

	}
}
