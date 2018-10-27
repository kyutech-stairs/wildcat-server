package models

func (imageInfo *ImageInfo) Create() (*ImageInfo, error) {
	if err := db.Create(&imageInfo).Error; err != nil {
		return &ImageInfo{ImageURL: ""}, err
	}
	return imageInfo, nil
}
