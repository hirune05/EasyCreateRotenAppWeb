package main

import (
	"backend/app/server"
	"log"
)

func main() {
	log.Fatalf("%+v", server.Run())
}
