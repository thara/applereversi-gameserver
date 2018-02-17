
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go

gen/helloworld:
		protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

.PHONY: runserver
helloworld/server:  ## Build helloworld server
	@go run ./cmd/helloworld-server/main.go

.PHONY: runclient
helloworld/client:  ## Build helloworld client
	@go run ./cmd/helloworld-client/main.go Tomochika
