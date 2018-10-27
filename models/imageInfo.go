package models

import "github.com/kyutech-stairs/wildcat-server/models/services"

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
