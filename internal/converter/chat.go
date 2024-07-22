package converter

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func ToChatUserFromDesc(user *desc.User) model.User {
	return model.User{
		Id:   user.Id,
		Name: user.Name,
	}
}

func ToChatUsersFromDesc(users []*desc.User) []model.User {
	var result []model.User
	for _, user := range users {
		result = append(result, ToChatUserFromDesc(user))
	}
	return result
}

func ToChatFromDesc(chat *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		Name:  chat.Name,
		Users: ToChatUsersFromDesc(chat.Users),
	}
}

func ToChatMessageFromDesc(msg *desc.ChatMessage) *model.Message {
	return &model.Message{
		User:      ToChatUserFromDesc(msg.User),
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt.AsTime(),
	}
}
