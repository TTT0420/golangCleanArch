package repository

import (
	"errors"

	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
	AddPost(*entity.Post) error
	UpdatePostById(*entity.Post) error
	IsPostExist(int) bool
	DeletePostById(*entity.Post) error
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
	if err := r.DB.Omit("CreatedDate", "UpdatedDate").Create(&post).Error; err != nil {
		return err
	}
	return nil
}

// 投稿編集
func (r *PostRepositoryImpl) UpdatePostById(post *entity.Post) error {
	if err := r.DB.Where(post.ID).Updates(&post).Error; err != nil {
		return err
	}
	return nil
}

// idより投稿の存在確認
func (r *PostRepositoryImpl) IsPostExist(id int) bool {
	var p entity.Post
	if err := r.DB.Where(id).First(&p).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// 投稿削除
func (r *PostRepositoryImpl) DeletePostById(post *entity.Post) error {
	post.DeleteFlg = true
	if err := r.DB.Updates(&post).Error; err != nil {
		return err
	}
	return nil
}
