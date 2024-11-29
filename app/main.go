package main

import (
	"log"

	//"github.com/BolvicBolvicovic/project_terrarium/server"
	"github.com/BolvicBolvicovic/project_terrarium/terrarium"
	env "github.com/joho/godotenv"
)

func main() {

	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}
	genom, err := terrarium.NewGenom(0.01, 0.9, 0.4, 0.5, 0.5, 0.9, 0.1, 2, 5, true)
	if err != nil {
		log.Fatal(err)
	}
	beast := terrarium.NewBeastRandomGenre("Victor's", genom)

	terrarium := terrarium.NewTerrarium(5, 4, 64, 50000, 200.0, 200.0, beast)

	for terrarium.CurrentIteration < terrarium.MaxIteration {
		terrarium.RunOneTurn()
		log.Printf(`
Current iteration: %d
Current beast num: %d
Current plant num: %d`, terrarium.CurrentIteration, terrarium.CurrentBeastNumber, terrarium.CurrentPlantNumber)
	}
	//server.Run()
}
