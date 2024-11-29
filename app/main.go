package main

import (
//	"fmt"
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
	genom, err := terrarium.NewGenom(0.01, 0.9, 0.4, 0.5, 0.5, 0.4, 0.6, 2, 5, true)
	if err != nil {
		log.Fatal(err)
	}
	beast := terrarium.NewBeastRandomGenre("Victor's", genom)

	terrarium := terrarium.NewTerrarium(0, 50, 1000, 1000, 1000, 10.0, 10.0, beast)
	beastPreviousState := *terrarium.Beasts[0]
	genomPreviousState := *terrarium.Beasts[0].Genom

	for terrarium.CurrentIteration < terrarium.MaxIteration {
		terrarium.RunOneTurn()
//		beasts := ""
//		for _, beast := range terrarium.Beasts {
//			beasts += fmt.Sprintf(`
//Name: %s
//Generation: %d
//Health: %f
//IsGestating: %t`, beast.Name, beast.Generation, beast.Health, beast.Genom.IsGestating)
//		}
		log.Printf(`
Current iteration: %d
Current beast num: %d
Current plant num: %d`, terrarium.CurrentIteration, terrarium.CurrentBeastNumber, terrarium.CurrentPlantNumber)
		if terrarium.CurrentBeastNumber == 0 || terrarium.CurrentPlantNumber == 0 { 
			log.Println(beastPreviousState, genomPreviousState)
			break 
		}
		beastPreviousState = *terrarium.Beasts[0]
		genomPreviousState = *terrarium.Beasts[0].Genom
	}
	//server.Run()
}
