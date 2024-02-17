package repository

import (
	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
}

type PostRepositoryImpl struct {
	DB *gorm.DB
}

// NewPostRepositoryImpl は新しいPostRepositoryImplを生成します。
func NewPostRepositoryImpl(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

func (r *PostRepositoryImpl) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post
	result := r.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
