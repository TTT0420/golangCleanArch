package repository

import (
	"errors"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/app/pkg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewPostRepositoryImpl(db *gorm.DB, logger *zap.Logger) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db, Logger: logger}
}

// 投稿全件取得
func (r *PostRepositoryImpl) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post
	result := r.DB.Model(&entity.Post{}).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

// 投稿登録
func (r *PostRepositoryImpl) AddPost(post *entity.Post) (int, error) {
	if err := r.DB.Model(&entity.Post{}).Create(&post).Error; err != nil {
		return pkg.FailedID, err
	}

	return post.ID, nil
}

// 投稿編集
func (r *PostRepositoryImpl) UpdatePostByID(post *entity.Post) error {
	result := r.DB.Model(&entity.Post{}).Where("id = ?", post.ID).Updates(post)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

// IDより投稿の存在確認
func (r *PostRepositoryImpl) IsPostExist(ID int) bool {
	var p entity.Post
	if err := r.DB.Model(&entity.Post{}).Where("ID = ?", ID).
		First(&p).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// 投稿削除
func (r *PostRepositoryImpl) DeletePostByID(id int) error {
	// 論理削除はis_deletedのみの更新で良いので、Updateを使用
	result := r.DB.Model(&entity.Post{}).Where("id = ?", id).Update("IsDeleted", true)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
