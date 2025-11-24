package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email  string
	Db     DbConfig
	MyUser MyUserConfig
}

type MyUserConfig struct {
	Secret string
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
		MyUser: MyUserConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}

func LoadTestConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, using dafault config")
	}

	// return &Config{
	// 	Email: "aalexanderbn@google.com",
	// 	Db: DbConfig{
	// 		Dsn: os.Getenv("TEST_DSN"), // например: "postgres://user:pass@localhost/myapp_test"
	// 	},
	// 	MyUser: MyUserConfig{
	// 		Secret: "TEST_JWT_SECRET",
	// 	},
	// }
	return &Config{
		Email: "aalexanderbn@google.com",
		Db: DbConfig{
			Dsn: "host=localhost user=aleksandr password=my_strong_password dbname=product_test port=5432 sslmode=disable", // например: "postgres://user:pass@localhost/myapp_test"
		},
		MyUser: MyUserConfig{
			Secret: "/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=",
		},
	}
}
