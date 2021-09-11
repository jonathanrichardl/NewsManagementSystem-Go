package config

import "os"

type databaseConfig struct {
	Username     string
	Password     string
	Address      string
	DatabaseName string
}

func LoadDatabaseConfiguration() (databaseConfig, error) {
	var result databaseConfig
	result.Username = os.Getenv("USER_NAME")
	result.Password = os.Getenv("PASSWORD")
	result.Address = os.Getenv("ADRESS")
	result.Address = os.Getenv("DATABASE_NAME")
	return result, nil

}
