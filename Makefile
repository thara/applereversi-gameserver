
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go

gen/helloworld:
		protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

.PHONY: runserver
runserver:  ## Build server
	@go run ./cmd/server/main.go

.PHONY: runclient
runclient:  ## Build client
	@go run ./cmd/client/main.go
