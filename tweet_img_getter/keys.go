package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	TwitterApi TwitterApi `toml:"twitterApi"`
}

type TwitterApi struct {
	ConsumerKey    string `toml:"consumerKey"`
	ConsumerSecret string `toml:"consumerSecret"`
	AccessToken    string `toml:"accessToken"`
	AccessSecret   string `toml:"accessSecret"`
}

func Hello() {
	fmt.Println("Hello my project!!")
}

func GetTwitterApi() *anaconda.TwitterApi {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}
	consumerKey := config.TwitterApi.ConsumerKey
	consumerSecret := config.TwitterApi.ConsumerSecret
	accessToken := config.TwitterApi.AccessToken
	accessSecret := config.TwitterApi.AccessSecret
	fmt.Printf("Consumer Key: %s\n", config.TwitterApi.ConsumerKey)
	fmt.Printf("Consumer Secret Key: %s\n", config.TwitterApi.ConsumerSecret)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)
	//fmt.Println("vim-go")
	return api
}
