package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email string
	Db    DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig(email string) *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, using dafault config")
	}
	return &Config{
		Email: email,
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
