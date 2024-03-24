package test

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/TTT0420/golangCleanArch/app/internal/domain/entity"
	"github.com/dolthub/go-mysql-server/driver"
	"github.com/dolthub/go-mysql-server/memory"
	sqle "github.com/dolthub/go-mysql-server/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTestDb(dbname string) *gorm.DB {
	db, _ :=
		gorm.Open(mysql.New(mysql.Config{Conn: New(dbname)}), nil)

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
