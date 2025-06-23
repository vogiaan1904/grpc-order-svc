.PHONY: protoc run update-proto

# Variables
APP_NAME=order-svc

protoc-all:
	$(MAKE) protoc PROTO=protos/proto/order.proto OUT_DIR=protogen/golang/order
	$(MAKE) protoc PROTO=protos/proto/product.proto OUT_DIR=protogen/golang/product

protoc:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	-I=protos/proto $(PROTO)

# Run command
run:
	@echo "Starting $(APP_NAME)..."
	go run cmd/server/main.go

# Update proto submodule
update-proto:
	@echo "Updating proto submodule..."
	git submodule update --remote protos
	@echo "Proto update complete."
