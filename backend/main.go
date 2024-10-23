package main

import (
	"backend/app/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("OS environment variables are read.")
	} else {
		log.Printf("Environment variables in the .env file are read.")
	}

	log.Fatalf("%+v", server.Run())
}
