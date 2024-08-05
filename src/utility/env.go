package utility

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func setModeFromEnv() {
	switch os.Getenv("ENV") {
	case "PRODUCTION":
		gin.SetMode(gin.ReleaseMode)
	}
}

func SetupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	setModeFromEnv()
}
