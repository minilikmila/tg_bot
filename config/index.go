package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)


var (

	PORT string
	TOKEN string
)

func init() {
	// os.env
	godotenv.Load()

	logrus.Info("ENV loader")
	PORT = os.Getenv("PORT")
	TOKEN = os.Getenv("TOKEN")
}