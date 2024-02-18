package domain

import "time"

type Post struct {
	ID          int
	UserID      int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Status      int
	CreatedDate time.Time
	UpdatedDate time.Time
}
