package repository

import (
	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/pkg"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewUserRepositoryImpl(db *gorm.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{DB: db, Logger: logger}
}

// ユーザー登録
func (r *UserRepository) AddUser(user *entity.Users) (int, error) {
	if err := r.DB.Model(&entity.Users{}).Create(&user).Error; err != nil {
		return pkg.FailedID, err
	}

	return user.UserID, nil
}

// ユーザー情報取得
func (r *UserRepository) GetUserByID(id int) (*entity.Users, error) {
	var user entity.Users
	if err := r.DB.Model(&entity.Users{}).Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ユーザー情報更新
func (r *UserRepository) UpdateUserByID(user *entity.Users) error {
	if err := r.DB.Model(&entity.Users{}).Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザー削除
func (r *UserRepository) DeleteUserByID(id int) error {
	if err := r.DB.Model(&entity.Users{}).Where("user_id= ?", id).Update("is_deleted", true).Error; err != nil {
		return err
	}
	return nil
}
