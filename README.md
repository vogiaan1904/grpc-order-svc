# Order Service

This service handles order management for the e-commerce platform.

## Proto Management

This service uses a git submodule (`protos-submodule`) to track the shared proto definitions. The workflow for managing proto files is:

1. **Update proto files from the central repository**:
   ```
   make sync-proto
   ```
   This will:
   - Update the `protos-submodule` from the remote repository
   - Copy the proto files to the local `protos` directory

2. **Generate code from proto files**:
   ```
   make protoc-all
   ```
   This will generate Go code from the proto definitions.

3. **Development workflow**:
   - Always run `make sync-proto` before starting development to ensure you have the latest proto definitions
   - After updating protos, run `make protoc-all` to regenerate code
   - Commit both the updated proto files and generated code

## Why this approach?

We use this approach (copying from submodule to local directory) because:
1. Docker builds can't directly access git submodules
2. It ensures reproducible builds with a "frozen" state of proto files
3. It allows developers to review proto changes before incorporating them

## Running the service

```
make run
```

## Building the service

```
docker build -t order-svc .
``` 