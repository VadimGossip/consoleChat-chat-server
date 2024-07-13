package app

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/api/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service"
	chatService "github.com/VadimGossip/consoleChat-chat-server/internal/service/chat"
)

type Factory struct {
	dbAdapter *DBAdapter

	chatService service.ChatService

	chatImpl *chat.Implementation
}

var factory *Factory

func newFactory(dbAdapter *DBAdapter) *Factory {
	factory = &Factory{dbAdapter: dbAdapter}
	factory.chatService = chatService.NewService(dbAdapter.chatRepo)
	factory.chatImpl = chat.NewImplementation(factory.chatService)
	return factory
}
