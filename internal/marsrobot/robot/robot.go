package robot

import (
	"github.com/caevv/marsrobot/internal/marsrobot"
)

type Robot struct {
	Coordinate   marsrobot.Coordinate
	Direction    marsrobot.Direction
	Instructions []Instruction
	IsLost       bool
}

func New(currentCoordinate marsrobot.Coordinate, direction marsrobot.Direction, instructions []Instruction) *Robot {
	return &Robot{
		Coordinate:   currentCoordinate,
		Direction:    direction,
		Instructions: instructions,
	}
}

func (r *Robot) ExecuteInstructions() {
	for _, v := range r.Instructions {
		v.Execute(r)
	}
}
