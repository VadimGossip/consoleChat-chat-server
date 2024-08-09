package chat

import (
	"context"

	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := i.chatService.Delete(ctx, req.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
