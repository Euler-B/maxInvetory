package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string `string:"host"`
	Port     string `string:"port"`
	User     string `string:"user"`
	Password string `string:"password"`
	Name     string `string:"name"`
}

type Settings struct {
	Host string         `string:"host"`
	Port string         `int64:"port"`
	DB   DatabaseConfig `string:"database"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func New() *Settings {
	dbConf := DatabaseConfig{
		Host:     os.Getenv("D_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
	s := Settings{
		Host: os.Getenv("HOST_SERVER"),
		Port: os.Getenv("PORT_SERVER"),
		DB:   dbConf,
	}
	return &s
}
