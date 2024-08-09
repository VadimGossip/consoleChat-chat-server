package chat

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/converter"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	if err := i.chatService.SendMessage(ctx, converter.ToChatMessageFromDesc(req)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
