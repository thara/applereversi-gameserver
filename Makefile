
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go


gen/helloworld:
		protoc -I helloworld/ helloworld.proto --go_out=plugins=grpc:helloworld
