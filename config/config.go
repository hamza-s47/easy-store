package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() map[string]string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return nil
	}

	env := make(map[string]string)

	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		env[kv[0]] = kv[1]
	}

	return env
}
