package app

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	CachePort   int `json:"cache_port"`
	ListenerPort int `json:"listener_port"`
}

func GetConfig() Config {
	jsonFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal([]byte(jsonFile), &config)
	if err != nil {
		panic(err)
	}
	return config
}
