package converter

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	repoModel "github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/model"
)

func ToRepoFromMessage(msg *model.Message) repoModel.Message {
	return repoModel.Message{
		From:      msg.From,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt,
	}
}
