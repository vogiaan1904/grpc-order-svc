package interceptors

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/vogiaan1904/order-svc/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// redact redacts sensitive information from the data
func redact(data interface{}, fields []string) string {
	if data == nil {
		return "null"
	}

	// Convert data to bytes
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "Error marshaling data"
	}

	// Unmarshal to map
	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataBytes, &dataMap); err != nil {
		return string(dataBytes)
	}

	// Redact sensitive fields
	for _, field := range fields {
		if _, exists := dataMap[field]; exists {
			dataMap[field] = "[Redacted]"
		}
	}

	// Convert back to JSON
	redactedBytes, err := json.MarshalIndent(dataMap, "", "  ")
	if err != nil {
		return "Error marshaling redacted data"
	}

	return string(redactedBytes)
}

// GrpcClientLoggingInterceptor creates a UnaryClientInterceptor that logs requests and responses
func GrpcClientLoggingInterceptor(logger log.Logger, redactedFields []string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Log the request using fmt instead of logger
		fmt.Printf("\n➡️  gRPC client request - Method: %s\n", method)
		// fmt.Printf("📤 Request: %s\n", redact(req, redactedFields))
		fmt.Printf("📤 Request: %+v\n", req)

		// Record start time
		startTime := time.Now()

		// Call the RPC method
		err := invoker(ctx, method, req, reply, cc, opts...)

		// Calculate duration
		duration := time.Since(startTime)

		// Log the response or error using fmt instead of logger
		if err != nil {
			st, _ := status.FromError(err)
			fmt.Printf("❌ gRPC client error - Method: %s, Code: %s, Message: %s, Duration: %v\n",
				method, st.Code(), st.Message(), duration)
		} else {
			fmt.Printf("✅ gRPC client response - Method: %s, Duration: %v\n", method, duration)
			// fmt.Printf("📥 Response: %s\n\n", redact(reply, redactedFields))
			fmt.Printf("📥 Response: %+v\n", reply)
		}

		return err
	}
}
