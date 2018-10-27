package config

import (
	_ "fmt"
	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	TwitterApiKeys `toml:"twitterApiKeys"`
	HashTagNames   []string `toml:"hashtagNames"`
}

type TwitterApiKeys struct {
	ConsumerKey    string `toml:"consumerKey"`
	ConsumerSecret string `toml:"consumerSecret"`
	AccessToken    string `toml:"accessToken"`
	AccessSecret   string `toml:"accessSecret"`
}

func GetConfig() *Config {
	var config *Config
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		panic(err)
	}
	return config
}

func (config *Config) GetTwitterApi() *anaconda.TwitterApi {
	consumerKey := config.ConsumerKey
	consumerSecret := config.ConsumerSecret
	accessToken := config.AccessToken
	accessSecret := config.AccessSecret

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)
	return api
}
