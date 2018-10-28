package main

import (
	"net/url"

	"log"

	"github.com/kyutech-stairs/wildcat-server/models"

	"github.com/kyutech-stairs/wildcat-server/config"
)

func main() {
	config := config.GetConfig()
	api := config.GetTwitterApi()
	v := url.Values{}
	v.Set("count", "450")

	result, err := api.GetSearch(config.HashTagNames[0], nil)
	if err != nil {
		panic(err)
	}
	tweets := result.Statuses

	log.Println("TWEET_COUNTS: ", len(tweets))
	for _, tweet := range tweets {
		tweetInfo := new(models.TweetInfo)
		var imageInfos models.ImageInfos

		// json, _ := json.Marshal(tweet)
		// log.Println("TWEET_DATA:  ", string(json))
		for _, media := range tweet.ExtendedEntities.Media {
			if media.Type == "photo" {
				imageURL := media.Media_url
				if !models.ImageInfoIsRetweeted(imageURL) {
					imageInfos = append(imageInfos, models.ImageInfo{ImageURL: imageURL})
				}
			}
		}
		if len(imageInfos) > 0 {
			tweetInfo.Images = imageInfos
			tweetInfo.Create()
		}
	}
}
