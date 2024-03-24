package repository

import (
	"testing"
	"time"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"github.com/TTT0420/golangCleanArch/app/internal/infrastructure/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewPostRepositoryImpl(t *testing.T) {
	// テスト用のデータベース接続を作成（ここではSQLiteを使用）
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("テスト用データベースの接続に失敗しました: %v", err)
	}

	// テスト用のロガーを作成
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
	db := test.NewTestDB("test")
	repo := NewPostRepositoryImpl(db, logger)

	// ケース1: 0件の場合
	posts, err := repo.GetAllPosts()
	assert.NoError(t, err)
	assert.Len(t, posts, 0)

	// テストデータを挿入
	testPosts := []entity.Post{
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
			CreatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
			UpdatedDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
		},
	}
	for _, p := range testPosts {
		db.Create(&p)
	}

	// ケース2: 複数件の場合
	posts, err = repo.GetAllPosts()
	assert.NoError(t, err)
	assert.Len(t, posts, len(testPosts))

	// ポストの内容を確認
	for i, post := range posts {
		assert.Equal(t, testPosts[i].Title, post.Title)
		assert.Equal(t, testPosts[i].Content, post.Content)
		// タイムスタンプの比較にはWithinDurationを使用
		assert.WithinDuration(t, testPosts[i].CreatedDate, post.CreatedDate, time.Second)
		assert.WithinDuration(t, testPosts[i].UpdatedDate, post.UpdatedDate, time.Second)
	}
}
