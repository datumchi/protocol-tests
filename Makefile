





generate-protocol:
	mkdir -p ./generated/protocol
	protoc -I=../protocol/protobuf --go_out=plugins=grpc:./generated/protocol ../protocol/protobuf/*.proto