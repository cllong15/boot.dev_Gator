package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	// read json file
	cfgFile, ok := getConfigFilePath()
	if ok != nil {
		return Config{}, fmt.Errorf("Read UserHomeDir: %v\n", ok)
	}
	dat, ok := os.ReadFile(cfgFile)
	if ok != nil {
		return Config{}, fmt.Errorf("Read dat: %v\n", ok)
	}

	jsonShit := Config{}
	ok = json.Unmarshal(dat, &jsonShit)
	if ok != nil {
		return Config{}, fmt.Errorf("Read jsonShit: %v\n", ok)
	}
	return jsonShit, nil
}

func writeConfig(cfg Config) error {
	// write to json
	jsonShit, ok := json.Marshal(cfg)
	if ok != nil {
		return fmt.Errorf("Read writeConfig Marshal: %v\n", ok)
	}
	file, ok := getConfigFilePath()
	// fmt.Println(file)
	if ok != nil {
		return fmt.Errorf("Read writeConfig file: %v\n", ok)
	}
	ok = os.WriteFile(file, jsonShit, 0666)
	if ok != nil {
		return fmt.Errorf("Read writeConfig: %v\n", ok)
	}
	return nil
}
