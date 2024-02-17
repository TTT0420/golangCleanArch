package usecase

import (
	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/internal/domain/repository"
)

type PostUsecase struct {
	PostRepo repository.PostRepository
}

func NewPostUsecase(repo repository.PostRepository) *PostUsecase {
	return &PostUsecase{
		PostRepo: repo,
	}
}

func (u *PostUsecase) GetAllPosts() ([]entity.Post, error) {
	return u.PostRepo.GetAllPosts()
}
