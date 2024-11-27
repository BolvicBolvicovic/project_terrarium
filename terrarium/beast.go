package terrarium

import (
	"log"

	"github.com/BolvicBolvicovic/project_terrarium/terrarium/items"
)

type Beast struct {
	Name		string
	Stamina		int
	Strenght	int
	Armor		int
	X		int
	Y		int
	Notebook	[]Beast
	Bag		[]items.Item
	Brain		NeuralNetwork
}

func (b Beast) Think() Thought {
	thought := b.Brain.Think()
	log.Println(b.Name, "thinks:", thought.Idea)
	return thought
}

func (b Beast) Learn() {
	b.Brain.Learn();
}
