package main

import (
	"log"

	"github.com/BolvicBolvicovic/project_terrarium/server"
	env "github.com/joho/godotenv"
)

func main() {

	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
