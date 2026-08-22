package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	Accounts struct {
		Source       string   `json:"source"`
		Destinations []string `json:"destinations"`
	} `json:"accounts"`
}

func LoadConfig() Config {
	var config Config

	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}