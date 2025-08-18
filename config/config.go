package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppPort    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

var AppConfig *Config

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	AppConfig = &Config{
		AppPort:    os.Getenv("APP_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
	}
}
