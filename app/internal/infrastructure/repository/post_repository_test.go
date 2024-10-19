package repository

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"strings"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/app/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type GormMock struct {
	mock.Mock
}

func (g *GormMock) GetAllPosts() ([]entity.Post, error) {
	args := g.Called()
	return args.Get(0).([]entity.Post), args.Error(1)
}
func TestNewPostRepositoryImpl(t *testing.T) {

	db := pkg.NewTestDb("test")
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("ロガーの初期化に失敗しました: %v", err)
	}

	// NewPostRepositoryImpl関数をテスト
	repo := NewPostRepositoryImpl(db, logger)

	// 戻り値がnilでないことを確認
	assert.NotNil(t, repo, "nilです")
	// repo.DBとrepo.Loggerがそれぞれ期待通りに設定されていることを確認する
	assert.Equal(t, db, repo.DB, "DBが正しく注入されていません")
	assert.Equal(t, logger, repo.Logger, "Loggerが正しく注入されていません")
}

func TestGetAllPosts(t *testing.T) {

	logger, _ := zap.NewDevelopment()
	db := pkg.NewTestDb("test")
	repo := NewPostRepositoryImpl(db, logger)
	var dbMock *GormMock

	var p *[]entity.Post
	tests := []struct {
		name  string
		setup func()
		want  *[]entity.Post
		err   error
	}{
		{
			name: "一件もない場合",
			setup: func() {
				db.Delete(&entity.Post{})
			},
			want: &[]entity.Post{},
		},
		{
			name: "複数件ある場合",
			setup: func() {
				p = setMultiplePosts()
				db.Create(&p)
			},
			want: setMultiplePosts(),
		},
		{
			name: "異常系",
			setup: func() {
				dbMock = new(GormMock)
				var posts []entity.Post
				expectedError := errors.New("database error")

				dbMock.On("GetAllPosts", &posts, mock.Anything).Return(nil, expectedError)
			},
			err: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.setup != nil {
				tt.setup()
			}
			if strings.Contains(tt.name, "正常") {
				posts, err := repo.GetAllPosts()
				assert.NoError(t, err)
				assert.Equal(t, tt.want, &posts)
				assert.Equal(t, len(*tt.want), len(posts))
			} else if strings.Contains(tt.name, "異常") {
				_, err := dbMock.GetAllPosts()
				assert.EqualError(t, err, fmt.Sprintf("%s", tt.err))
			}
		})
	}
}

func TestAddPost(t *testing.T) {

	logger, _ := zap.NewDevelopment()
	db := pkg.NewTestDb("test")
	repo := NewPostRepositoryImpl(db, logger)

	tests := []struct {
		name  string
		setup func()
		post  *entity.Post
		want  int
		err   error
	}{
		{
			name: "正常系",
			post: &entity.Post{
				UserID:      1,
				Title:       "Test Post",
				Content:     "Content",
				Status:      1,
				IsDeleted:   false,
				CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
				UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "異常系",
			post: &entity.Post{CreatedDate: time.Date(time.Now().Year()+9999, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)},
			want: -1,
			err:  errors.New("Incorrect datetime value: '12023-03-24T00:00:00Z'"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := repo.AddPost(tt.post)
			if strings.Contains(tt.name, "正常") {
				assert.NoError(t, err)
			} else if strings.Contains(tt.name, "異常") {
				assert.Error(t, err, tt.err)
			}

			assert.Equal(t, tt.want, id)
		})
	}

}

func TestUpdatePostByID(t *testing.T) {

	logger, _ := zap.NewDevelopment()
	db := pkg.NewTestDb("test")
	repo := NewPostRepositoryImpl(db, logger)

	tests := []struct {
		name  string
		setup func()
		arg   *entity.Post
		want  error
	}{
		{
			name: "正常系",
			setup: func() {
				db.Where("1=1").Delete(&entity.Post{})
				post := entity.Post{
					ID:          1,
					UserID:      1,
					Title:       "Test Post",
					Content:     "Content",
					Status:      1,
					IsDeleted:   false,
					CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
					UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
				}
				db.Create(&post)
			},
			arg: &entity.Post{
				ID:    1,
				Title: "Test Post was updated",
			},
			want: nil,
		},
		{
			name: "異常系 存在しないID",
			setup: func() {
				db.Where("1=1").Delete(&entity.Post{})
			},
			arg: &entity.Post{
				ID:          1,
				UserID:      1,
				Title:       "Test Post 1",
				Content:     "Content 1",
				Status:      1,
				IsDeleted:   false,
				CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
				UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
			},
			want: errors.New("no rows affected"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			err := repo.UpdatePostByID(tt.arg)
			if strings.Contains(tt.name, "正常") {
				var ac entity.Post
				assert.NoError(t, err)
				db.First(tt.arg.ID, &ac)
				assert.Equal(t, *tt.arg, ac)
			} else if strings.Contains(tt.name, "異常") {
				assert.Error(t, err, tt.want)
			}
		})
	}
}

func TestIsPostExist(t *testing.T) {

	logger, _ := zap.NewDevelopment()
	db := pkg.NewTestDb("test")
	repo := NewPostRepositoryImpl(db, logger)

	tests := []struct {
		name  string
		setup func()
		arg   int
		want  bool
	}{
		{
			name: "正常系:存在するID",
			setup: func() {
				db.Where("1=1").Delete(&entity.Post{})
				post := entity.Post{
					ID:          1,
					UserID:      1,
					Title:       "Test Post",
					Content:     "Content",
					Status:      1,
					IsDeleted:   false,
					CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
					UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
				}
				db.Create(&post)
			},
			arg:  1,
			want: true,
		},
		{
			name: "準正常系:存在しないID",
			setup: func() {
				db.Where("1=1").Delete(&entity.Post{})
			},
			arg:  1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			result := repo.IsPostExist(tt.arg)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestDeletePostByID(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db := pkg.NewTestDb("test")
	repo := NewPostRepositoryImpl(db, logger)
	test := []struct {
		name  string
		setup func()
		arg   int
		want  error
	}{
		{
			name: "正常系",
			setup: func() {
				db.Where("1=1").Delete(&entity.Post{})
				post := entity.Post{
					ID:          1,
					UserID:      1,
					Title:       "Test Post",
					Content:     "Content",
					Status:      1,
					IsDeleted:   false,
					CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
					UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
				}
				db.Create(&post)
			},
			arg:  1,
			want: nil,
		},
		{
			name: "異常系 存在しないID",
			setup: func() {
				db.Where("1=1").Delete(&entity.Post{})
			},
			arg:  1,
			want: errors.New("record not found"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			err := repo.DeletePostByID(tt.arg)
			if strings.Contains(tt.name, "正常") {
				assert.NoError(t, err)
				var post entity.Post
				db.First(&post, tt.arg)
				assert.Equal(t, true, post.IsDeleted)
			} else if strings.Contains(tt.name, "異常") {
				assert.EqualError(t, err, fmt.Sprintf("%s", tt.want))
			}
		})
	}
}

func setMultiplePosts() *[]entity.Post {
	return &[]entity.Post{
		{
			ID:          1,
			UserID:      1,
			Title:       "Test Post 1",
			Content:     "Content 1",
			Status:      1,
			IsDeleted:   false,
			CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
			UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          2,
			UserID:      1,
			Title:       "Test Post 2",
			Content:     "Content 2",
			Status:      1,
			IsDeleted:   false,
			CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
			UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
		},
	}
}
