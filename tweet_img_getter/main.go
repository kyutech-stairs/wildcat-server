package main

import (
	"net/url"

	"github.com/kyutech-stairs/wildcat-server/models"

	"github.com/kyutech-stairs/wildcat-server/config"
)

func main() {
	config := config.GetConfig()
	api := config.GetTwitterApi()
	v := url.Values{}
	v.Set("count", "300")

	result, err := api.GetSearch(config.HashTagNames[0], v)
	if err != nil {
		panic(err)
	}
	tweets := result.Statuses

	for _, tweet := range tweets {
		tweetInfo := new(models.TweetInfo)
		var imageInfos models.ImageInfos

		// Retweetされたものは除外
		if tweet.Retweeted == false {
			// json, _ := json.Marshal(tweet)
			// log.Println("TWEET_DATA:  ", string(json))
			for _, media := range tweet.ExtendedEntities.Media {
				if media.Type == "photo" {
					imageURL := media.Media_url
					imageInfos = append(imageInfos, models.ImageInfo{ImageURL: imageURL})
				}
			}
			tweetInfo.Images = imageInfos
			if len(tweetInfo.Images) > 0 {
				tweetInfo.Create()
			}
		}
	}
}
