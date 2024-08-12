package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	chatImpl "github.com/VadimGossip/consoleChat-chat-server/internal/api/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service"
	serviceMocks "github.com/VadimGossip/consoleChat-chat-server/internal/service/mocks"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func TestDelete(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()

		id         = gofakeit.Int64()
		serviceErr = fmt.Errorf("some service error")

		req = &desc.DeleteRequest{Id: id}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			chatServiceMock := test.chatServiceMock(mc)

			impl := chatImpl.NewImplementation(chatServiceMock)
			actualRes, err := impl.Delete(test.args.ctx, test.args.req)

			assert.Equal(t, test.want, actualRes)
			assert.Equal(t, test.err, err)
		})
	}
}
