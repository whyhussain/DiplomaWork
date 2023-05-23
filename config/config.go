package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database DB `json:"database"`
}

func NewConfig() *Config {
	a, err := os.ReadFile("D:/diploma/DiplomaWork/config/congig.json")
	if err != nil {
		fmt.Println("asaaaaaaaa")
	}

	cfg := Config{}
	err = json.Unmarshal(a, &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}

type DB struct {
	Primary string `json:"primary"`
}
