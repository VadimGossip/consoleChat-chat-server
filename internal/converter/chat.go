package converter

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func ToChatMessageFromDesc(msg *desc.ChatMessage) *model.Message {
	return &model.Message{
		From:      msg.From,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt.AsTime(),
	}
}
