package pkg

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func GodotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Info("Error loading .env file")
	}
	return os.Getenv(key)
}
