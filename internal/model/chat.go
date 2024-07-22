package model

import "time"

type Chat struct {
	Name  string
	Users []User
}

type User struct {
	ID   int64
	Name string
}

type Message struct {
	ChatID    int64
	UserID    int64
	Text      string
	CreatedAt time.Time
}
