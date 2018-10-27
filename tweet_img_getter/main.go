package main

import (
	"fmt"
	"github.com/kyutech-stairs/wildcat-server/config"
	"net/url"
)

func main() {
	config := config.GetConfig()
	api := config.GetTwitterApi()
	v := url.Values{}
	v.Set("count", "30")

	//tweets, err := api.GetHomeTimeline(v)
	result, err := api.GetSearch(config.HashTagNames[0], v)
	if err != nil {
		panic(err)
	}
	tweets := result.Statuses

	for _, tweet := range tweets {
		fmt.Println("MEDIA_DATA: ", tweet.Entities.Media)
		fmt.Println("MEDIA_START")
		for _, media := range tweet.Entities.Media {
			fmt.Println("TYPE", media.Type)
			fmt.Println("URL", media.Media_url)
		}
		fmt.Println("MEDIA_END")
	}
}
