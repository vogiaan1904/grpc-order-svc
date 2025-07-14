.PHONY: protoc run update-proto sync-proto

# Variables
APP_NAME=order-svc

run-server:
	@echo "Starting $(APP_NAME)..."
	go run cmd/server/main.go

protoc-all:
	$(MAKE) protoc PROTO=protos-submodule/proto/order.proto OUT_DIR=protogen/golang/order
	$(MAKE) protoc PROTO=protos-submodule/proto/product.proto OUT_DIR=protogen/golang/product

protoc:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	-I=protos-submodule $(PROTO)
	
gen-proto:
	@echo "Updating git submodule..."
	git submodule update --remote protos-submodule

	@echo "Regenerating proto code..."
	make protoc-all

	@echo "Proto code regenerated."


