package validator

import (
	"fmt"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

func CreateValidation(usernames []string) error {
	if len(usernames) == 0 {
		return fmt.Errorf("can't create chat without users")
	}
	return nil
}

func SendValidation(msg *model.Message) error {
	if msg.Text == "" {
		return fmt.Errorf("can't send empty message")
	}
	return nil
}
