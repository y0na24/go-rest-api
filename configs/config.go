package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Dsn    string
	Secret string
}

func GenerateConfig() *Config {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Не удалось загрузить переменные окружения, используем дефолтный конфиг")
	}

	return &Config{
		Dsn:    os.Getenv("DSN"),
		Secret: os.Getenv("TOKEN"),
	}
}
