package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func CloseDBConn() {
	db.Close()
}

func GetDBConfig() (string, string) {
	//return "mysql", "root:password@tcp(kyotohack2018-f-1.cm68cjyvpekc.ap-northeast-1.rds.amazonaws.com:3306)/kyotohack2018_f?charset=utf8&parsetime=true&loc=local"
	return "sqlite3", "gorm.db"
}
