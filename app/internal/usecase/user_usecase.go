package usecase

import (
	"time"

	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/internal/domain/repository"
	"go.uber.org/zap"
)

type UserUseCase struct {
	UserRepo repository.UserRepository
	Logger   *zap.Logger
}

func NewUserUsecase(repo repository.UserRepository, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{
		UserRepo: repo,
		Logger:   logger,
	}
}

// ユーザー登録
func (u *UserUseCase) AddUser(user *entity.Users) (int, error) {
	return u.UserRepo.AddUser(user)
}

// ユーザー情報取得
func (u *UserUseCase) GetUserByID(id int) (*entity.Users, error) {
	return u.UserRepo.GetUserByID(id)
}

// ユーザー情報更新
func (u *UserUseCase) UpdateUserByID(user *entity.Users) error {
	user.UpdatedDate = time.Now()
	return u.UserRepo.UpdateUserByID(user)
}

// ユーザー削除
func (u *UserUseCase) DeleteUserByID(id int) error {
	return u.UserRepo.DeleteUserByID(id)
}
