package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port       string `json:"port"`
	BasePath   string `json:"basepath"`
	BrowserURL string `json:"browserurl" validate:"required"`
}

var Configuration = Config{
	Port:     "8008",
	BasePath: "/",
}

// LoadConfig loads the config values into the struct from the config.json
func LoadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&Configuration)
}
