package configs

type Config struct {
	Email string
}

func LoadConfig(email string) *Config {

	return &Config{
		Email: email,
	}
}
