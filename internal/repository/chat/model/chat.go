package model

import "time"

type Message struct {
	ChatID    int64
	UserID    int64
	Text      string
	CreatedAt time.Time
}

type User struct {
	ID   int64
	Name string
}
