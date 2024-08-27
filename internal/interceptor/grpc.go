package interceptor

import (
	"context"
	"fmt"

	descGrpc "github.com/VadimGossip/consoleChat-chat-server/internal/client/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCInterceptor interface {
	Hook() grpc.UnaryServerInterceptor
}

type interceptor struct {
	authGRPCClient descGrpc.AuthClient
}

func NewInterceptor(authGRPCClient descGrpc.AuthClient) *interceptor {
	return &interceptor{authGRPCClient: authGRPCClient}
}

type validator interface {
	Validate() error
}

func (i *interceptor) Hook() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("metadata is not provided")
		}
		ptCtx := metadata.NewOutgoingContext(ctx, md)

		err := i.authGRPCClient.Check(ptCtx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
