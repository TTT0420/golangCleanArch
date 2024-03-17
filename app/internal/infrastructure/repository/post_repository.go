package repository

import (
	"errors"

	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
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
	result := r.DB.Model(&entity.Post{}).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

// 投稿登録
func (r *PostRepositoryImpl) AddPost(post *entity.Post) (int, error) {
	// CreatedDate, UpdatedDateはsqlで設定されるため、Omitで除外
	if err := r.DB.Model(&entity.Post{}).Create(&post).Error; err != nil {
		return pkg.FailedID, err
	}

	return post.ID, nil
}

// 投稿編集
func (r *PostRepositoryImpl) UpdatePostByID(post *entity.Post) error {
	if err := r.DB.Model(&entity.Post{}).Where("ID = ?", post.ID).Updates(post).Error; err != nil {
		return err
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
	if err := r.DB.Model(&entity.Post{}).Where("id = ?", id).Update("IsDeleted", true).Error; err != nil {
		return err
	}
	return nil
}
