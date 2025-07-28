package main

import (
	"github.com/joho/godotenv"
	"log"
	"movies_online/internal/application"
)

func main() {
	application.Run()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
