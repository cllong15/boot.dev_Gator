package config

import "fmt"

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func SetUser(cfg Config, currentUser string) error {
	cfg.CurrentUserName = currentUser
	ok := writeConfig(cfg)
	if ok != nil {
		return fmt.Errorf("Config SetUser: %v\n", ok)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	return "./" + configFileName, nil
}
