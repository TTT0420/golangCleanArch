package usecase

import (
	"time"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/app/internal/domain/repository"
	"github.com/TTT0420/golangCleanArch/app/internal/interfaces/dto"
	"github.com/TTT0420/golangCleanArch/app/pkg"
	"go.uber.org/zap"
)

type PostUsecase struct {
	PostRepo repository.PostRepository
	Logger   *zap.Logger
}

func NewPostUsecase(repo repository.PostRepository, logger *zap.Logger) *PostUsecase {
	return &PostUsecase{
		PostRepo: repo,
		Logger:   logger,
	}
}

// 全件取得
func (u *PostUsecase) GetAllPosts() ([]entity.Post, error) {
	return u.PostRepo.GetAllPosts()
}

// 新規追加
func (u *PostUsecase) AddPost(postReq dto.AddPostReq) (int, error) {
	post := entity.Post{
		UserID:      postReq.UserID,
		Title:       postReq.Title,
		Content:     postReq.Content,
		IsDeleted:   false,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	return u.PostRepo.AddPost(&post)
}

// 投稿編集
func (u *PostUsecase) EditPostByID(postReq dto.EditPostReq) error {

	post := entity.Post{
		ID:          postReq.ID,
		Title:       postReq.Title,
		Content:     postReq.Content,
		UpdatedDate: time.Now(),
	}

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
