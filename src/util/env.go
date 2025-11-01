package util

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("⚠️  Warning: .env file not found, using system environment")
	}
}

func GetEnv(s string) string {
	return os.Getenv(s)
}

func GetEnvWithComma(s string, comma string) []string {
	var val string
	var arrs []string
	val = os.Getenv(s)
	arrs = strings.Split(val, comma)
	return arrs
}
