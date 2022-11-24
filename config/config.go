package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alexander256/shop/models"
)

func getFilePath() string {
	return os.Getenv("config_path")
}
func InitCinfig() (*models.Config, error) {
	filePath := getFilePath()
	configuration, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("couldn't load configuration file : %v", err)
	}

	var config models.Config
	if err := json.Unmarshal(configuration, &config); err != nil {
		return nil, fmt.Errorf("could't unmarshal configuration file : %v", err)
	}

	if config.Postgres.Password == "" {
		config.Postgres.Password = os.Getenv("postgres_password")
	}
	return &config, nil
}
