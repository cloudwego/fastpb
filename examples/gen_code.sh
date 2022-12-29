# Copyright 2022 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# fastpb-gen:

# install
if [ "$(which protoc-gen-fastpb)" == "" ]; then
  if [ "$(go version | grep -E '1.16|1.17|1.18|1.19')" != "" ]; then
    go install github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
  else
    go get github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
  fi
fi

# clean
rm -rf ./fastpb_gen && mkdir -p ./fastpb_gen/nested

# gen code
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ ./idl/base.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ ./idl/echo.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen/nested --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen/nested -I ./idl/ ./idl/nested.proto

protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ user/user.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ pet/pet.proto
