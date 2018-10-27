package models

func (tweetImage *TweetImage) Create() (*TweetImage, error) {
	if err := db.Create(&tweetImage).Error; err != nil {
		return &TweetImage{}, err
	}
	return tweetImage, nil
}
