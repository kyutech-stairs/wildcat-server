package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kyutech-stairs/wildcat-server/models/services"
)

// Create ImageInfoをCreate
func (imageInfo *ImageInfo) Create() (*ImageInfo, error) {
	if err := db.Create(&imageInfo).Error; err != nil {
		return &ImageInfo{ImageURL: ""}, err
	}
	return imageInfo, nil
}

// GetOffsetImageInfos オフセットされたImageをsize枚だけ取得できる
func GetOffsetImageInfos(size, offset uint) *services.ImageInfos {
	imageInfos := &services.ImageInfos{}
	db.Limit(size).Offset(size * offset).Find(&imageInfos)
	// db.Find(&imageInfos)
	return imageInfos
}

func ImageInfoIsRetweeted(imageUrl string) bool {
	imageInfo := &ImageInfo{ImageURL: ""}
	if err := db.Where("image_url = ?", imageUrl).First(&imageInfo).Error; gorm.IsRecordNotFoundError(err) {
		// record not found
		log.Println("I m OK")
		return false
	} else {
		return true
	}
}
