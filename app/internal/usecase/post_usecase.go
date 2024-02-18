package usecase

import (
	"errors"
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

// 全件取得
func (u *PostUsecase) GetAllPosts() ([]entity.Post, error) {
	return u.PostRepo.GetAllPosts()
}

// 新規追加
func (u *PostUsecase) AddPost(c *gin.Context) error {

	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		fmt.Printf("era-dayon:%v", err)
		return err
	}

	fmt.Printf("bindけっか%+v", &post)
	return u.PostRepo.AddPost(&post)
}

// 投稿編集
func (u *PostUsecase) EditPost(c *gin.Context) error {

	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return err
	}

	// 存在確認
	if !u.PostRepo.IsPostExist(post.ID) {
		return errors.New("there is no post")
	}

	return u.PostRepo.UpdatePostById(&post)
}
