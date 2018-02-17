
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go


gen/helloworld:
		go generate github.com/thara/applereversi-gameserver
