package chat

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/converter"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, req.Usernames)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := i.chatService.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	if err := i.chatService.SendMessage(ctx, req.Id, converter.ToChatMessageFromDesc(req.Message)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
