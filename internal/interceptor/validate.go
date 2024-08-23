package interceptor

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	descGrpc "github.com/VadimGossip/consoleChat-chat-server/internal/client/grpc"
)

func BuildInterceptor(authGRPCClient descGrpc.AuthClient) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logrus.Infof("Intercepted request %+v method %+v", req, info.FullMethod)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("metadata is not provided")
		}
		ptCtx := metadata.NewOutgoingContext(ctx, md)

		err := authGRPCClient.Check(ptCtx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
