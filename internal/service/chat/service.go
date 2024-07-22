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

func (s *service) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	if err := validator.CreateValidation(chat.Users); err != nil {
		return 0, err
	}
	tx, err := s.chatRepository.BeginTxSerializable(ctx)
	if err != nil {
		return 0, err
	}
	id, err := s.chatRepository.CreateChat(ctx, tx, chat.Name)
	if err != nil {
		return 0, s.chatRepository.StopTx(ctx, tx, err)
	}

	for _, user := range chat.Users {
		if err = s.chatRepository.CreateChatUser(ctx, tx, id, user); err != nil {
			return 0, s.chatRepository.StopTx(ctx, tx, err)
		}
	}

	return id, s.chatRepository.StopTx(ctx, tx, nil)
}

func (s *service) Delete(ctx context.Context, chatId int64) error {
	tx, err := s.chatRepository.BeginTxSerializable(ctx)
	if err != nil {
		return err
	}

	if err = s.chatRepository.Delete(ctx, tx, chatId); err != nil {
		return s.chatRepository.StopTx(ctx, tx, err)
	}

	return s.chatRepository.StopTx(ctx, tx, nil)
}

func (s *service) SendMessage(ctx context.Context, msg *model.Message) error {
	if err := validator.SendValidation(msg); err != nil {
		return err
	}

	tx, err := s.chatRepository.BeginTxSerializable(ctx)
	if err != nil {
		return err
	}

	if err = s.chatRepository.SendMessage(ctx, tx, msg); err != nil {
		return s.chatRepository.StopTx(ctx, tx, err)
	}

	return s.chatRepository.StopTx(ctx, tx, nil)
}
