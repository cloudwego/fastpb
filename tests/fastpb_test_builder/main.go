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
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"code.byted.org/wangtieju/utils"
)

var IDLDir string

func init() {
	flag.StringVar(&IDLDir, "search", ".", "directory of thrift files")
}

func IDLSearcher() <-chan (string) {
	// read file path from stdin
	if IDLDir == "" {
		return utils.LineReader(os.Stdin)
	}
	// search ThriftDir to find thrift files
	return utils.FileSearcher(IDLDir, func(name string) bool {
		return strings.HasSuffix(name, ".proto")
	})
}

func main() {
	flag.Parse()
	for protoFile := range IDLSearcher() {
		dir, err := os.MkdirTemp("", "")
		if err != nil {
			log.Fatalf("create temp dir failed: %s\n", err.Error())
		}

		content, err := ioutil.ReadFile(protoFile)
		if err != nil {
			log.Fatalf("read %s file failed: %s\n", protoFile, err.Error())
		}
		if bytes.Contains(content, []byte(`"proto2"`)) {
			continue
		}
		var cmds []*exec.Cmd
		cmd := exec.Command("protoc", "--go_opt=paths=source_relative", "--go_out=.", "--fastpb_opt=paths=source_relative", "--fastpb_out=.", "-I", filepath.Dir(protoFile), protoFile)
		cmds = append(cmds, cmd)
		cmd = exec.Command("go", "mod", "init", "mydemo")
		cmds = append(cmds, cmd)
		cmd = exec.Command("go", "mod", "tidy")
		cmds = append(cmds, cmd)
		cmd = exec.Command("go", "test", "./...")
		cmds = append(cmds, cmd)

		for _, cmd := range cmds {
			cmd.Dir = dir
			if err := cmd.Run(); err != nil {
				log.Println("CWD:", dir)
				log.Fatalln(cmd.Args, "run failed:", err.Error())
			}
		}
		log.Println(protoFile, " build success")

		os.RemoveAll(dir)
	}
	log.Println("build finished")
}
