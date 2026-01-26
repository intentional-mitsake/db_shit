package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
	Type     string //mysql, pg
}

func LoadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     viper.GetString("host"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
		Database: viper.GetString("name"),
		Port:     viper.GetInt("port"),
		Type:     viper.GetString("type"),
	}
}
