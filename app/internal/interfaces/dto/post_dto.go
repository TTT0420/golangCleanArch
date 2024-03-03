package dto

// add_postリクエスト
type AddPostReq struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required,ContentsCheck"`
	Content string `json:"content" binding:"required,ContentsCheck"`
}

// edit_postリクエスト
type EditPostReq struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
}
