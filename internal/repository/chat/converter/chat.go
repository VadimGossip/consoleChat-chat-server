package converter

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	repoModel "github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/model"
)

func ToRepoFromMessage(msg *model.Message) repoModel.Message {
	return repoModel.Message{
		ChatID:    msg.ChatID,
		UserID:    msg.UserID,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt,
	}
}

func ToRepoFromUser(user model.User) repoModel.User {
	return repoModel.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
