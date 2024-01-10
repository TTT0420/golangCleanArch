package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DB struct {
	Host       string
	Username   string
	PassWord   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		PassWord: c.DB.Production.PassWord,
		DBName:   c.DB.Production.DBName,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open("mysql", d.Username+":"+d.PassWord+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Loca")
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	fmt.Println()
	return d
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
