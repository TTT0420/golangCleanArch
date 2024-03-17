package entity

import "time"

type Users struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	UserType    int       `json:"user_type"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
