package repository

import (
	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
	AddPost(*entity.Post) (int, error)
	UpdatePostById(*entity.Post) (int, error)
	IsPostExist(int) bool
	DeletePostById(*entity.Post) (int, error)
}
