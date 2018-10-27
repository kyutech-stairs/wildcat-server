package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	TwitterApiKeys `toml:"twitterApiKeys"`
}

type TwitterApiKeys struct {
	ConsumerKey    string `toml:"consumerKey"`
	ConsumerSecret string `toml:"consumerSecret"`
	AccessToken    string `toml:"accessToken"`
	AccessSecret   string `toml:"accessSecret"`
}

func GetTwitterApi() *anaconda.TwitterApi {
	var config Config
	_, err := toml.DecodeFile("../config/config.toml", &config)
	if err != nil {
		panic(err)
	}
	consumerKey := config.ConsumerKey
	consumerSecret := config.TwitterApiKeys.ConsumerSecret
	accessToken := config.TwitterApiKeys.AccessToken
	accessSecret := config.TwitterApiKeys.AccessSecret
	fmt.Printf("Consumer Key: %s\n", consumerKey)
	fmt.Printf("Consumer Secret Key: %s\n", consumerSecret)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)
	return api
}
