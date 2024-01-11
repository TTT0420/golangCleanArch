package usecase

import (
	"fmt"

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
	fmt.Println("usecase層")
	return u.PostRepo.GetAllPosts()
}
