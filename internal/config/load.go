package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err == nil {
		log.Println("Loaded .env successfully from current directory")
		return
	}

	if err := godotenv.Load("../env"); err == nil {
		log.Println("Loaded .env successfully from parent directory")
	}

	log.Println("No .env file found")
}
