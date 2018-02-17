
.PHONY: prepare
prepare:
		brew install protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go


.PHONY: applereversi/gen
applereversi/gen:
		protoc -I . apple_reversi.proto --go_out=plugins=grpc:.

.PHONY: applereversi/server
applereversi/server:
	@go run ./cmd/applereversi-server/main.go

.PHONY: applereversi/client/host
applereversi/client/host:
	@go run ./cmd/applereversi-client/main.go

.PHONY: applereversi/client/guest
applereversi/client/guest:
	@go run ./cmd/applereversi-client/main.go --guest --gameId ${GAME_ID}
