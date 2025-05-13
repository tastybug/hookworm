package hookworm

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Task represents a single hook configuration from .hookworm.yml
type Task struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

// TaskBook represents the structure of .hookworm.yml
type TaskBook struct {
	Task []Task `yaml:"tasks"`
}

// InitializeTaskBook reads and parses .hookworm.yml
func InitializeTaskBook(configFilePath string) (*TaskBook, error) {
	log.Printf("Reading config from %s\n", configFilePath)
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config TaskBook
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %v", configFilePath, err)
	}

	if len(config.Task) == 0 {
		log.Println("No hooks defined in .hookworm.yml")
	}

	return &config, nil
}
