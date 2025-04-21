.PHONY: protoc run update-proto

# Variables
APP_NAME=order-svc

protoc:
    protoc --go_out=protogen/golang/order --go_opt=paths=source_relative \                                                                    1 ↵
           --go-grpc_out=protogen/golang/order --go-grpc_opt=paths=source_relative \
           -I=protos/proto protos/proto/order.proto

# Run command
run:
    @echo "Starting $(APP_NAME)..."
    go run cmd/server/main.go
    
# Update proto submodule
update-proto:
    @echo "Updating proto submodule..."
    git submodule update --remote protos
    @echo "Proto update complete."