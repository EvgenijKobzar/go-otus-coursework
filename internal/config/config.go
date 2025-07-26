package config

import (
	"encoding/json"
	"errors"
	"movies_online/internal/model"
	"movies_online/internal/model/catalog"
	"os"
)

type FilePathConfig struct {
	Serial  string `json:"serial"`
	Season  string `json:"season"`
	Episode string `json:"episode"`
	Account string `json:"account"`
}

type Config struct {
	FilePath FilePathConfig `json:"filepath"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ResolvePathByEntityType[T catalog.HasId](entity T) (string, error) {

	configApp, _ := LoadConfig("config.json")
	//configApp, _ := LoadConfig("../../config_test.json")

	var path string
	switch any(entity).(type) {
	case *catalog.Serial:
		path = configApp.FilePath.Serial
	case *catalog.Season:
		path = configApp.FilePath.Season
	case *catalog.Episode:
		path = configApp.FilePath.Episode
	case *model.Account:
		path = configApp.FilePath.Account
	default:
		return "", errors.New("invalid entity type")
	}
	return path, nil
}
