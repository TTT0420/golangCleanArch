package dto

// add_userリクエスト
type AddUserReq struct {
	UserID   int    `json:"user_id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
}

// edit_userリクエスト
type EditUserReq struct {
	UserID   int    `json:"user_id" binding:"required"`
	UserName string `json:"user_name"`
	UserType int    `json:"user_type"`
}
