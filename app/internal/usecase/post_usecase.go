package usecase

import (
	"github.com/TTT0420/golangCleanArch/internal/domain"
)

type PostUsecase struct {
	PostRepo domain.PostRepository
}

func NewPostUsecase(repo domain.PostRepository) *PostUsecase {
	return &PostUsecase{
		PostRepo: repo,
	}
}

func (u *PostUsecase) GetAllPosts() ([]domain.Post, error) {
	return u.PostRepo.GetAllPosts()
}
