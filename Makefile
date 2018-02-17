
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go


gen/helloworld:
		protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld



.PHONY: runserver
runserver:  ## Build server
	@cd ./cmd/server && go build -o ../../_dist/server -v
	@./_dist/server

.PHONY: runclient
runclient:  ## Build client
	@cd ./cmd/client && go build -o ../../_dist/client -v
	@./_dist/client Tomochika
