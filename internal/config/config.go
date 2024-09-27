package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBConnectionString string `json:"db_url"`
	Username           string `json:"current_user_name"`
}

func Read() (Config, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("error getting home directory: %v", err)
	}
	configPath := filepath.Join(userHome, configFileName)

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	if err = json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("error unmarshalling json: %v", err)
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.Username = username

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshalling to json: %v", err)
	}

	userHome, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %v", err)
	}

	file, err := os.OpenFile(filepath.Join(userHome, configFileName), os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening config file for writing: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}
