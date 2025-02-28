package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func loadEnv() Config {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}

	return Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

func BuildDSN() string {
	var config = loadEnv()

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
