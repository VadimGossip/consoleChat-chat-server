package app

import (
	"google.golang.org/grpc"

	"github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func initGrpcRouter(app *App) func(*grpc.Server) {
	return func(s *grpc.Server) {
		chat_v1.RegisterChatV1Server(s, app.chatImpl)
	}
}
