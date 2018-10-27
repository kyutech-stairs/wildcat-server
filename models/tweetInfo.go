package models

func (tweetInfo *TweetInfo) Create() (*TweetInfo, error) {
	if err := db.Create(&tweetInfo).Error; err != nil {
		return &TweetInfo{}, err
	}
	return tweetInfo, nil
}
