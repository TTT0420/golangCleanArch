package repository

import (
	"fmt"

	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/pkg"
	"gorm.io/gorm"
)

type UserRepostiroy struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepostiroy {
	return &UserRepostiroy{DB: db}
}

// ユーザー登録
func (r *UserRepostiroy) AddUser(user *entity.Users) (int, error) {
	if err := r.DB.Model(&entity.Users{}).Create(&user).Error; err != nil {
		fmt.Println("asdfsdfs", err)
		return pkg.FailedID, err
	}

	return user.UserID, nil
}

// ユーザー情報取得
func (r *UserRepostiroy) GetUserByID(id int) (*entity.Users, error) {
	var user entity.Users
	if err := r.DB.Model(&entity.Users{}).Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ユーザー情報更新
func (r *UserRepostiroy) UpdateUserByID(user *entity.Users) error {
	if err := r.DB.Model(&entity.Users{}).Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザー削除
func (r *UserRepostiroy) DeleteUserByID(id int) error {
	if err := r.DB.Model(&entity.Users{}).Where("user_id= ?", id).Update("is_deleted", true).Error; err != nil {
		return err
	}
	return nil
}
