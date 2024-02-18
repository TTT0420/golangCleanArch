package domain

import "time"

type Post struct {
	ID          int    `jdon:"id"`
	UserID      int    `json:"user_id" binding:"required"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Status      int
	CreatedDate time.Time
	UpdatedDate time.Time
}
