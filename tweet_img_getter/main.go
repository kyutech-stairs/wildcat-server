package main

import (
	"fmt"
	"github.com/kyutech-stairs/wildcat-server/config"
	"net/url"
)

func main() {
	api := config.GetTwitterApi()
	v := url.Values{}
	v.Set("count", "10")

	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		fmt.Println("tweet: ", tweet.Text)
	}
}
