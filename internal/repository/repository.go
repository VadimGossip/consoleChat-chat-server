package repository

import (
	"context"
	"github.com/jackc/pgx/v4"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

type ChatRepository interface {
	BeginTxSerializable(ctx context.Context) (pgx.Tx, error)
	StopTx(ctx context.Context, tx pgx.Tx, err error) error
	CreateChat(ctx context.Context, tx pgx.Tx, name string) (int64, error)
	CreateChatUser(ctx context.Context, tx pgx.Tx, chatId int64, user model.User) error
	Delete(ctx context.Context, tx pgx.Tx, chatId int64) error
	SendMessage(ctx context.Context, tx pgx.Tx, msg *model.Message) error
}
