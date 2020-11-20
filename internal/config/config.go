package config

import "os"

type Config struct {
	DBConfig  DBConfig
	APIConfig APIConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type APIConfig struct {
	Port string
}

func LoadConfig() Config {
	return Config{
		DBConfig: DBConfig{
			Host:     os.Getenv(`DB_HOST`),
			Port:     os.Getenv(`DB_PORT`),
			Name:     os.Getenv(`DB_NAME`),
			User:     os.Getenv(`DB_USER`),
			Password: os.Getenv(`DB_PASSWORD`),
		},
		APIConfig: APIConfig{
			Port: os.Getenv(`API_PORT`),
		},
	}
}