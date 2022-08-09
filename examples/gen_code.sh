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
