package usecase

import (
	"errors"

	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/internal/domain/repository"
	"github.com/TTT0420/golangCleanArch/pkg"
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
func (u *PostUsecase) AddPost(c *gin.Context) (int, error) {

	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return pkg.FailedId, err
	}
	return u.PostRepo.AddPost(&post)
}

// 投稿編集
func (u *PostUsecase) EditPost(c *gin.Context) (int, error) {

	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return pkg.FailedId, err
	}

	// 存在確認
	if !u.PostRepo.IsPostExist(post.Id) {
		return pkg.FailedId, errors.New("there is no post")
	}

	return u.PostRepo.UpdatePostById(&post)
}

// 投稿削除
func (u *PostUsecase) DeletePost(c *gin.Context) (int, error) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return pkg.FailedId, err
	}

	// 存在確認
	if !u.PostRepo.IsPostExist(post.Id) {
		return pkg.FailedId, errors.New("there is no post")
	}

	return u.PostRepo.DeletePostById(&post)
}
