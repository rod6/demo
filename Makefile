.PHONY: pb user things gateway

all: pb user things gateway

pb:
	@rm -f pb/*.pb.go
	protoc -I=.:pb/extra/src --go_out=plugins=grpc:. pb/*.proto

user:
	go build -o bin/user ./user

things:
	go build -o bin/things ./things

gateway:
	go build -o bin/gateway ./gateway