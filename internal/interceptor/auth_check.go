package interceptor

import (
	"context"
	"fmt"

	descGrpc "github.com/VadimGossip/consoleChat-chat-server/internal/client/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthCheckInterceptor interface {
	Hook() grpc.UnaryServerInterceptor
}

type authCheckInterceptor struct {
	authGRPCClient descGrpc.AuthClient
}

func NewInterceptor(authGRPCClient descGrpc.AuthClient) *authCheckInterceptor {
	return &authCheckInterceptor{authGRPCClient: authGRPCClient}
}

func (i *authCheckInterceptor) Hook() grpc.UnaryServerInterceptor {
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
