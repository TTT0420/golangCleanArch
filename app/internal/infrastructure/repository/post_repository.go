package repository

import (
	"errors"

	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/pkg"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

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
func (r *PostRepositoryImpl) AddPost(post *entity.Post) (int, error) {
	if err := r.DB.Omit("CreatedDate", "UpdatedDate").Create(post).Error; err != nil {
		return pkg.FailedId, err
	}

	return post.Id, nil
}

// 投稿編集
func (r *PostRepositoryImpl) UpdatePostById(post *entity.Post) (int, error) {
	if err := r.DB.Model(&entity.Post{}).Where("id = ?", post.Id).Updates(post).Error; err != nil {
		return pkg.FailedId, err
	}
	return post.Id, nil
}

// idより投稿の存在確認
func (r *PostRepositoryImpl) IsPostExist(id int) bool {
	var p entity.Post
	if err := r.DB.Model(&entity.Post{}).Where("id = ?", id).
		First(&p).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// 投稿削除
func (r *PostRepositoryImpl) DeletePostById(post *entity.Post) (int, error) {
	post.IsDeleted = true
	if err := r.DB.Model(&entity.Post{}).Updates(&post).Error; err != nil {
		return pkg.FailedId, err
	}
	return post.Id, nil
}
