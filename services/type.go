package services

type ImageInfo struct {
	ID       uint   `json:"id"`
	ImageURL string `json:"url"`
}

type ImageInfos []ImageInfo
