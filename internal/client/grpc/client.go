package grpc

import "context"

type AuthClient interface {
	Check(ctx context.Context, endpointAddress string) error
}
