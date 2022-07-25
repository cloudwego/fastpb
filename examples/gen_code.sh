# fastpb-gen:

# FIXME: delete
echo "now in $(pwd)"
rm $(go env GOPATH)/bin/protoc-gen-fastpb
cd $(go env GOPATH)/src/github.com/cloudwego/fastpb/protoc-gen-fastpb/
go install
which protoc-gen-fastpb
cd -

rm -rf ./fastpb_gen && mkdir ./fastpb_gen
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen -I ./idl/ ./idl/*.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_out=./fastpb_gen -I ./idl/ ./idl/*.proto
#protoc --go_opt=paths=source_relative --go_out=./fastpb_gen -I ./idl/ ./idl/*.proto
