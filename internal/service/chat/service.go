package chat

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/service"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service/chat/validator"
)

var _ def.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatRepository
}

func NewService(chatRepository repository.ChatRepository) *service {
	return &service{
		chatRepository: chatRepository,
	}
}

func (s *service) Create(ctx context.Context, usernames []string) (int64, error) {
	if err := validator.CreateValidation(usernames); err != nil {
		return 0, err
	}
	return s.chatRepository.Create(ctx, usernames)
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.chatRepository.Delete(ctx, id)
}

func (s *service) SendMessage(ctx context.Context, id int64, msg *model.Message) error {
	if err := validator.SendValidation(msg); err != nil {
		return err
	}
	return s.chatRepository.SendMessage(ctx, id, msg)
}
