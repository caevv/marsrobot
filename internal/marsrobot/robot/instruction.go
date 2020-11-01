package robot

import (
	"errors"
	"github.com/caevv/marsrobot/internal/marsrobot"
)

type Instruction interface {
	Execute(robot *Robot)
}

type Forward struct{ worldMap *marsrobot.WorldMap }

type Left struct{}

type Right struct{}

func NewInstruction(input string, worldMap *marsrobot.WorldMap) (Instruction, error) {
	switch input {
	case "F":
		return Forward{worldMap: worldMap}, nil
	case "L":
		return Left{}, nil
	case "R":
		return Right{}, nil
	}

	return nil, errors.New("unknown instruction: " + input)
}

func (i Forward) Execute(robot *Robot) {
	coordinate := robot.Coordinate.NewDirection(robot.Direction)

	if !i.worldMap.IsInbound(coordinate) {
		if i.worldMap.HasScent(coordinate) {
			return
		}

		i.worldMap.AddScent(coordinate)
		robot.IsLost = true
	}

	robot.Coordinate = coordinate
}

func (i Left) Execute(robot *Robot) {
	moves := map[string]string{
		"N": "W",
		"E": "N",
		"S": "E",
		"W": "S",
	}
	robot.Direction = marsrobot.Direction{Direction: moves[robot.Direction.Direction]}
}

func (i Right) Execute(robot *Robot) {
	moves := map[string]string{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}

	robot.Direction = marsrobot.Direction{Direction: moves[robot.Direction.Direction]}
}
