package repository

import (
	"fmt"

	"github.com/TTT0420/golangCleanArch/internal/domain"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

// NewPostRepositoryImpl は新しいPostRepositoryImplを生成します。
func NewPostRepositoryImpl(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

func (r *PostRepositoryImpl) GetAllPosts() ([]domain.Post, error) {
	fmt.Println("GetAllPosts!!")
	var posts []domain.Post
	result := r.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
