package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validator interface {
	Validate() bool
}

func ValidationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if v, ok := req.(validator); ok {
		if !v.Validate() {
			return nil, status.Errorf(codes.InvalidArgument, "Bad request")
		}
	}

	return handler(ctx, req)
}
