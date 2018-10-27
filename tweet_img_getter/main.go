package main

import (
	"log"
	"net/url"

	"github.com/kyutech-stairs/wildcat-server/models"

	"github.com/kyutech-stairs/wildcat-server/config"
)

func main() {
	config := config.GetConfig()
	api := config.GetTwitterApi()
	v := url.Values{}
	v.Set("count", "300")

	//tweets, err := api.GetHomeTimeline(v)
	result, err := api.GetSearch(config.HashTagNames[0], v)
	if err != nil {
		panic(err)
	}
	tweets := result.Statuses

	for _, tweet := range tweets {
		tweetInfo := new(models.TweetInfo)
		var imageInfos models.ImageInfos
		log.Println("MEDIA_DATA: ", tweet.Entities.Media)
		log.Println("MEDIA_START")
		for _, media := range tweet.Entities.Media {
			if media.Type == "photo" {
				log.Println("TYPE", media.Type)
				log.Println("URL", media.Media_url)
				imageURL := media.Media_url
				imageInfos = append(imageInfos, models.ImageInfo{ImageURL: imageURL})
			}
		}
		log.Println("LOG", imageInfos)
		tweetInfo.Images = imageInfos
		if len(tweetInfo.Images) > 0 {
			tweetInfo.Create()
		}

		log.Println("MEDIA_END")
	}
}
