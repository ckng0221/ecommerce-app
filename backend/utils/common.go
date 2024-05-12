package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(requiredEnv []string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file")
	}
	for _, envName := range requiredEnv {
		env := os.Getenv(envName)
		if env == "" {
			log.Fatalf("environment variable '%s' is required", envName)
		}
	}
}

// Get string from environmnet variable.
// If empty, assign the value with the fallback string.
func Getenv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func RandString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
