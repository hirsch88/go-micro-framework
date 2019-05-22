package lib

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
}

func Env(key string, defaultValue string) string{
	var value = os.Getenv(key)
	if len(strings.TrimSpace(value)) == 0 {
		return defaultValue
	}
	return value
}

func EnvNumber(key string, defaultValue int) int{
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}
	return defaultValue
}

func EnvBool(key string, defaultValue bool) bool{
	var value = os.Getenv(key)
	if len(strings.TrimSpace(value)) == 0 {
		return defaultValue
	}
	return value == "true"
}