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

	fmt.Println("HashTag: ", config.HashTagNames[0])
	//tweets, err := api.GetHomeTimeline(v)
	result, err := api.GetSearch(config.HashTagNames[0], v)
	//result, err := api.GetSearch("golang", v)
	if err != nil {
		panic(err)
	}
	tweets := result.Statuses

	for _, tweet := range tweets {
		fmt.Println("tweet: ", tweet.Text)
	}
}
