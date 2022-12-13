package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv(key string, file string) string {
	// load .env file
	err := godotenv.Load(file)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)

}

func GetIntEnv(key string, file string) int {
	val := GetEnv(key, file)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Error converting string to int in GetIntEnv")
	}
	return ret
}
