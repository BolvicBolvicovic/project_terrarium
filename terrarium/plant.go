package terrarium

import (
	"math/rand"

	"github.com/BolvicBolvicovic/project_terrarium/lib"
)

type Plant struct {
	Position	*lib.Position	`json:"position"`
	Generation	int		`json:"generation"`

	food		float64
	Alive		bool
}

// Food interface
func (p Plant) GetPosition() *lib.Position { return p.Position }
func (p Plant) AsFood(herbivorRate float64) float64 { return p.food * herbivorRate }
func (p Plant) AsSelf() interface{} { return p }

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

	newPlants := make([]*Plant, 2)
	newPlants0 := &Plant{
		Position: &lib.Position{
			X: p.Position.X + 0.1,
			Y: p.Position.Y,
		},
		Generation: generation,
		food: food,
		Alive: true,
	}

	newPlants1 := &Plant{
		Position: &lib.Position{
			X: p.Position.X - 0.1,
			Y: p.Position.Y,
		},
		Generation: generation,
		food: food,
		Alive: true,
	}

	newPlants2 := &Plant{
		Position: &lib.Position{
			X: p.Position.X,
			Y: p.Position.Y + 0.1,
		},
		Generation: generation,
		food: food,
		Alive: true,
	}

	newPlants3 := &Plant{
		Position: &lib.Position{
			X: p.Position.X,
			Y: p.Position.Y - 0.1,
		},
		Generation: generation,
		food: food,
		Alive: true,
	}
	newPlants[0] = lib.RandomChoice(newPlants0, newPlants1)
	newPlants[1] = lib.RandomChoice(newPlants2, newPlants3)
	
	newPlants[0] = lib.RandomChoice(newPlants[0], nil)
//	newPlants[1] = lib.RandomChoice(newPlants[1], nil)

	return newPlants
}

func NewRandomPlant() *Plant {
	return &Plant{
		Position: lib.RandomPosition(),
		Generation: 0,
		food: rand.Float64(),
		Alive: true,
	}
}
