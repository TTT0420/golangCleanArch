package repository

import (
	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
	AddPost(*entity.Post) error
	UpdatePostById(*entity.Post) error
	IsPostExist(int) bool
	DeletePostById(*entity.Post) error
}
