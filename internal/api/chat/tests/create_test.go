package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"

	chatImpl "github.com/VadimGossip/consoleChat-chat-server/internal/api/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service"
	serviceMocks "github.com/VadimGossip/consoleChat-chat-server/internal/service/mocks"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		id         = gofakeit.Int64()
		name       = gofakeit.BookAuthor()
		serviceErr = fmt.Errorf("some service error")

		req = &desc.CreateRequest{
			Name: name,
			Users: []*desc.User{
				{
					Id:   1,
					Name: "Name one",
				},
				{
					Id:   2,
					Name: "Name another",
				},
			},
		}

		chat = &model.Chat{
			Name: name,
			Users: []model.User{
				{
					ID:   1,
					Name: "Name one",
				},
				{
					ID:   2,
					Name: "Name another",
				},
			},
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, chat).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, chat).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			chatServiceMock := test.chatServiceMock(mc)

			impl := chatImpl.NewImplementation(chatServiceMock)
			actualRes, err := impl.Create(test.args.ctx, test.args.req)

			assert.Equal(t, test.want, actualRes)
			assert.Equal(t, test.err, err)
		})
	}
}
