package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, menggunakan environment variable yang ada")
	}
}

func GetDeepseekAPIKey() string {
	apiKey := ""
	if v := getenv("DEEPSEK_API_KEY"); v != "" {
		apiKey = v
	}
	if apiKey == "" {
		log.Fatal("DEEPSEK_API_KEY belum di-set di environment variable")
	}
	return apiKey
}

func getenv(key string) string {
	return os.Getenv(key)
}
