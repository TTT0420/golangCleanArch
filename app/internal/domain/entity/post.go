package domain

import "time"

type Post struct {
	ID          int    `jdon:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Status      int
	DeleteFlg   bool `json:"delete_flg"`
	CreatedDate time.Time
	UpdatedDate time.Time
}
