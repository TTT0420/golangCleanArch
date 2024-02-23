package entity

import "time"

type Post struct {
	Id          int    `jdon:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Status      int
	IsDeleted   bool `json:"is_deleted"`
	CreatedDate time.Time
	UpdatedDate time.Time
}
