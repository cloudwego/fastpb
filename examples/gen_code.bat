

protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ user/user.proto
protoc --go_opt=paths=source_relative --go_out=./fastpb_gen --fastpb_opt=paths=source_relative --fastpb_out=./fastpb_gen -I ./idl/ pet/pet.proto