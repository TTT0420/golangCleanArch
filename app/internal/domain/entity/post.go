package domain

import "time"

type Post struct {
	ID          int
	UserID      int
	Title       string
	Content     string
	Status      string
	CreatedDate time.Time
	UpdatedDate time.Time
}
