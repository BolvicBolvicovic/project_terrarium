package terrarium

import (
)

type NeuralNetwork struct {}
type Thought struct {
	Idea	string
}


func buildThought() Thought {
	return Thought { "an idea" }
}

func (n NeuralNetwork) Think() Thought {
	thought := buildThought()
	return thought
}

func (n NeuralNetwork) Learn() {
}
