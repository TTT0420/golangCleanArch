package usecase

import (
	"github.com/TTT0420/golangCleanArch/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/internal/domain/repository"
	"github.com/TTT0420/golangCleanArch/pkg"
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
func (u *PostUsecase) AddPost(post entity.Post) (int, error) {
	return u.PostRepo.AddPost(&post)
}

// 投稿編集
func (u *PostUsecase) EditPostByID(post entity.Post) error {
	// 存在確認
	if !u.PostRepo.IsPostExist(post.ID) {
		return pkg.ErrRecordNotFound(post.ID)
	}
	return u.PostRepo.UpdatePostByID(&post)
}

// 投稿削除
func (u *PostUsecase) DeletePostByID(id int) error {
	// 存在確認
	if !u.PostRepo.IsPostExist(id) {
		return pkg.ErrRecordNotFound(id)
	}
	return u.PostRepo.DeletePostByID(id)
}
