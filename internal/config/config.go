package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir() // get home dir
	if err != nil {
		return "", err
	}
	path := filepath.Join(homeDir, configFileName)
	return path, nil
}

func Read() (*Config, error) {

	//open file logic
	configFilePath, err := getConfigFilePath() // get path
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFilePath) // open file
	if err != nil {
		return nil, err
	}
	defer file.Close() //defer close until function is done

	var cfg Config // init cfg struct
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func write(cfg *Config) error { // helper to write to file
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return err
	}

	// Couldve used:

	// file, err := os.Create(fullPath)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// err = encoder.Encode(cfg)
	// if err != nil {
	// 	return err
	// }
	return nil

}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(c)
}
