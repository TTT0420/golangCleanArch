package usecase

import (
	"fmt"

	entity "github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/internal/domain/repository"
	"github.com/gin-gonic/gin"
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

func (u *PostUsecase) AddPost(c *gin.Context) error {

	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		fmt.Printf("era-dayon:%v", err)
		return err
	}

	fmt.Printf("bindけっか%+v", &post)
	return u.PostRepo.AddPost(&post)
}
