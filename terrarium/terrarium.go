package terrarium

import (
)

type Terrarium struct {
	TotalBeastNumber 	int
	CurrentBeastNumber	int
	Beasts			[]Beast
	Width			int
	Height			int
}

func (t Terrarium) RunOneTurn() {
}
