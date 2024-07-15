package app

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat"
)

type DBAdapter struct {
	chatRepo repository.ChatRepository
}

func NewDBAdapter() *DBAdapter {
	return &DBAdapter{chatRepo: chat.NewRepository()}
}
