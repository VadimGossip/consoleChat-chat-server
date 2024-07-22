package service

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, id int64, msg *model.Message) error
}
