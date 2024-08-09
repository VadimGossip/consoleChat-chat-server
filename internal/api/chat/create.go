package chat

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/converter"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
