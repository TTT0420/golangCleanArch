package repository

import (
	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
	AddPost(*entity.Post) error
}

type PostRepositoryImpl struct {
	DB *gorm.DB
}

// NewPostRepositoryImpl は新しいPostRepositoryImplを生成します。
func NewPostRepositoryImpl(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

// 投稿全件取得
func (r *PostRepositoryImpl) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post
	result := r.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

// 投稿登録
func (r *PostRepositoryImpl) AddPost(post *entity.Post) error {
	p := &entity.Post{
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
		Status:  1,
	}

	if err := r.DB.Omit("CreatedDate", "UpdatedDate").Create(&p).Error; err != nil {
		return err
	}
	return nil
}
