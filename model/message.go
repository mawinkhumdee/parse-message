package model

import "time"

type Message struct {
	ID        string
	UserID    string
	Content   string
	Source    string // "text" | "voice" | "image"
	Status    string // "pending" | "success" | "failed"
	CreatedAt time.Time
}
