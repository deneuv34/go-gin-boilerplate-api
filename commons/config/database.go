package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "rangga:adhitya@root/base_api")
	if err != nil {
		log.Panic("Cannot connect to database: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(2)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
