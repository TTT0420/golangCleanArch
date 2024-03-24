package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"github.com/dolthub/go-mysql-server/driver"
	"github.com/dolthub/go-mysql-server/memory"
	sqle "github.com/dolthub/go-mysql-server/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewTestDb(dbname string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 標準出力にログを出力
		logger.Config{
			SlowThreshold: time.Second, // 実行に1秒以上かかるクエリを遅いクエリとしてログに記録
			LogLevel:      logger.Info, // Infoレベルでログを出力（すべてのSQLクエリを出力）
			Colorful:      true,        // ログの色付けを有効にする
		},
	)
	db, _ :=
		gorm.Open(mysql.New(mysql.Config{Conn: New(dbname)}), nil, &gorm.Config{Logger: newLogger})

	db.AutoMigrate(&entity.Post{})

	return db
}

type dbs []sqle.Database

var _ driver.Provider = dbs{}

func (d dbs) Resolve(name string, options *driver.Options) (string, sqle.DatabaseProvider, error) {
	return name, memory.NewDBProvider(d...), nil
}

func New(dbNames ...string) *sql.DB {
	var memdbs dbs
	for _, dbName := range dbNames {
		memdb := memory.NewDatabase(dbName)
		memdb.EnablePrimaryKeyIndexes()
		memdbs = append(memdbs, memdb)
	}
	drv := driver.New(memdbs, nil)
	conn, err := drv.OpenConnector(strings.Join(dbNames, ";"))
	if err != nil {
		fmt.Println("asfd")
	}
	db := sql.OpenDB(conn)
	if len(dbNames) > 0 {
		db.Exec("USE " + dbNames[0])
	}
	return db
}
