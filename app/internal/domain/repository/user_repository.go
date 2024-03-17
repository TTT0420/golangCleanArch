package repository

import "github.com/TTT0420/golangCleanArch/internal/domain/entity"

type UserRepository interface {
	AddUser(*entity.Users) (int, error)
	GetUserByID(int) (*entity.Users, error)
	UpdateUserByID(*entity.Users) error
	DeleteUserByID(int) error
}
