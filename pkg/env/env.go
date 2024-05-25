package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CheckEnvFile() {
	errEnv := godotenv.Load("../../.env")
	if errEnv != nil {
		panic(fmt.Sprintf("The .env file could not be loaded: %v", errEnv))
	}
}

func GetEnvValue(key string) string {
	value := os.Getenv(key)
	return value
}
