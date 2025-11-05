package config

import (

	// Deprecated, use os.ReadFile instead
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	// Config file path should be in the home
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return (homeDir + "/" + configFileName), nil
}

// Function to write the config back to the file
func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Marshal the config into JSON format
	fileBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	// Write the file content
	err = os.WriteFile(configFilePath, fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// SetUser method to update the current user in the config
func (cfg *Config) SetUser(userName string) error {
	// Update the current user name in the config
	cfg.CurrentUserName = userName

	// Write the updated config back to the file
	return write(*cfg)
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Read the file content
	fileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create an instance of your struct
	var config Config
	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return config, nil
}
