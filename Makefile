.PHONY: protoc

protoc:
    protoc --go_out=protobuf/protos/proto --go_opt=paths=source_relative \
           --go-grpc_out=protobuf/protos/proto --go-grpc_opt=paths=source_relative \
           -I=protos/proto protos/proto/order.proto