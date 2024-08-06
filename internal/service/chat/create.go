package chat

import (
	"context"
	"fmt"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service/chat/validator"
)

func (s *service) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	if err := validator.CreateValidation(chat.Users); err != nil {
		return 0, err
	}
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		id, txErr = s.chatRepository.CreateChat(ctx, chat.Name)
		if txErr != nil {
			return txErr
		}

		for _, user := range chat.Users {
			if txErr = s.chatRepository.CreateChatUser(ctx, id, user); txErr != nil {
				return txErr
			}
		}

		txErr = s.auditService.Create(ctx, &model.Audit{
			Action:     "create chat",
			CallParams: fmt.Sprintf("chat %+v", chat),
		})
		if txErr != nil {
			return txErr
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
