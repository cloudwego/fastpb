# fastpb-gen:

# FIXME: delete
echo "now in $(pwd)"
rm $(go env GOPATH)/bin/protoc-gen-fastpb
cd $(go env GOPATH)/src/github.com/cloudwego/fastpb/protoc-gen-fastpb/
go install
which protoc-gen-fastpb
cd -

rm -rf ./fastpb_gen && mkdir -p ./fastpb_gen/nested
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ ./idl/base.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ ./idl/echo.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen/nested --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen/nested -I ./idl/ ./idl/nested.proto
#protoc --go_opt=paths=source_relative --go_out=./fastpb_gen -I ./idl/ ./idl/*.proto
