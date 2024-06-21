package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ServiceConfiguration struct {
	Minecraft Minecraft `json:"minecraft"`
}

type Minecraft struct {
	Env     []Env    `json:"env"`
	Ports   []Port   `json:"ports"`
	Volumes []Volume `json:"volumes"`
}

type Env struct {
	Name        string  `json:"name"`
	Key         string  `json:"key"`
	Required    bool    `json:"required"`
	Description string  `json:"description"`
	Default     string  `json:"default"`
	Values      []Value `json:"values"`
}

type Value struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Port struct {
	Host      int64 `json:"host"`
	Container int64 `json:"container"`
}

type Volume struct {
	Host        string `json:"host"`
	Destination string `json:"destination"`
}

func GetServiceConfiguration() (*ServiceConfiguration, error) {
	// Open the JSON file
	jsonFile, err := os.Open("service-configuration.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open json file: %v", err)
	}
	defer jsonFile.Close()

	// Read the JSON file content
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read json file: %v", err)
	}

	// Unmarshal the JSON content into ServiceConfiguration struct
	var config ServiceConfiguration
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json content: %v", err)
	}

	return &config, nil
}
