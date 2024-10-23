package main

import (
	"backend/app/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	log.Fatalf("%+v", server.Run())
}
