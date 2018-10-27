package main

import (
	"github.com/kyutech-stairs/wildcat-server/models"

	_ "github.com/jinzhu/gorm"
)

func main() {
	db := models.GetDBConn()

	// Migrate
	db.DropTableIfExists(&models.TweetInfo{})
	db.DropTableIfExists(&models.ImageInfo{})

	db.CreateTable(&models.TweetInfo{})
	db.CreateTable(&models.ImageInfo{})

}
