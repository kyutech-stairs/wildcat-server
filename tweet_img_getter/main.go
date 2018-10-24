package main

import (
	"fmt"
	"net/url"
)

func main() {
	Hello()
	api := GetTwitterApi()
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
