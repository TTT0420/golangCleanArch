package repository

import (
	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
	AddPost(*entity.Post) (int, error)
	UpdatePostByID(*entity.Post) error
	IsPostExist(int) bool
	DeletePostByID(int) error
}
