package converter

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func ToChatUserFromDesc(user *desc.User) model.User {
	return model.User{
		ID:   user.Id,
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

func ToChatMessageFromDesc(req *desc.SendRequest) *model.Message {
	return &model.Message{
		ChatID:    req.Id,
		UserID:    req.Message.UserId,
		Text:      req.Message.Text,
		CreatedAt: req.Message.CreatedAt.AsTime(),
	}
}
