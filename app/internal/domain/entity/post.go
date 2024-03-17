package entity

import (
	"time"
)

type Post struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title" `
	Content     string    `json:"content"`
	Status      int       `json:"status"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
