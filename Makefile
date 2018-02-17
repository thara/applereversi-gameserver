
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: helloworld/gen
helloworld/gen:
		protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

.PHONY: helloworld/server
helloworld/server:  ## Build helloworld server
	@go run ./cmd/helloworld-server/main.go

.PHONY: helloworld/client
helloworld/client:  ## Build helloworld client
	@go run ./cmd/helloworld-client/main.go Tomochika


.PHONY: route_guide/gen
routeguide/gen:
		protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide

.PHONY: route_guide/server
routeguide/server:
	@go run ./cmd/routeguide-server/main.go

.PHONY: route_guide/client
routeguide/client:
	@go run ./cmd/routeguide-client/main.go
