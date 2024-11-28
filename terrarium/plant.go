package terrarium

import (
	"math/rand"

	"github.com/BolvicBolvicovic/project_terrarium/lib"
)

type Plant struct {
	Position	*lib.Position
	Generation	int

	food		float64
}

func (p *Plant) AsFood(herbivorRate float64) float64 {
	return p.food * herbivorRate
}

func (p *Plant) Propagate() []*Plant {

	generation := p.Generation + 1
	food := func() float64 {
		sign := lib.RandomChoice(-0.1, 0.1)
		newFood := rand.Float64() * sign
		newFood = p.food + newFood
		if newFood > 1 { newFood = 1 }
		if newFood <= 0 { newFood = 0.01 }
		return newFood
	}()

	newPlants := make([]*Plant, 4)
	newPlants[0] = &Plant{
		Position: &lib.Position{
			X: p.Position.X + 0.1,
			Y: p.Position.Y,
		},
		Generation: generation,
		food: food,
	}

	newPlants[1] = &Plant{
		Position: &lib.Position{
			X: p.Position.X - 0.1,
			Y: p.Position.Y,
		},
		Generation: generation,
		food: food,
	}

	newPlants[2] = &Plant{
		Position: &lib.Position{
			X: p.Position.X,
			Y: p.Position.Y + 0.1,
		},
		Generation: generation,
		food: food,
	}

	newPlants[3] = &Plant{
		Position: &lib.Position{
			X: p.Position.X,
			Y: p.Position.Y - 0.1,
		},
		Generation: generation,
		food: food,
	}

	return newPlants
}

func NewRandomPlant() *Plant {
	return &Plant{
		Position: lib.RandomPosition(),
		Generation: 0,
		food: rand.Float64(),
	}
}
