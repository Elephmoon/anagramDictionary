package config

import "os"

type Debug string

const On Debug = "true"

type Config struct {
	AppConfig AppConfig
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
	Host string
	Port string
}

type AppConfig struct {
	Debug Debug
}

func LoadConfig() Config {
	return Config{
		AppConfig: AppConfig{
			Debug: Debug(os.Getenv(`DEBUG`)),
		},
		DBConfig: DBConfig{
			Host:     os.Getenv(`DB_HOST`),
			Port:     os.Getenv(`DB_PORT`),
			Name:     os.Getenv(`DB_NAME`),
			User:     os.Getenv(`DB_USER`),
			Password: os.Getenv(`DB_PASSWORD`),
		},
		APIConfig: APIConfig{
			Host: os.Getenv(`API_HOST`),
			Port: os.Getenv(`API_PORT`),
		},
	}
}
