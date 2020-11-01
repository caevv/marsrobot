package marsrobotcmd

import (
	"fmt"
	"github.com/caevv/marsrobot/internal/marsrobot"
	"github.com/caevv/marsrobot/internal/marsrobot/robot"
	"strconv"
	"strings"
)

func Execute(lines []string) ([]string, error) {
	var (
		worldMap          *marsrobot.WorldMap
		robots            []*robot.Robot
		instructions      []robot.Instruction
		currentCoordinate marsrobot.Coordinate
		currentDirection  marsrobot.Direction
	)

	for k, i := range lines {
		if isFirstLine(k) {
			values := strings.Split(i, " ")

			x, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, fmt.Errorf("failed to get coordinate: %w", err)
			}

			y, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, fmt.Errorf("failed to get coordinate: %w", err)
			}

			worldMap = &marsrobot.WorldMap{
				TopRightCoordinate: marsrobot.Coordinate{
					X: x,
					Y: y,
				},
				Scents: make([]marsrobot.Coordinate, 0),
			}

			continue
		}

		if isLineWithCoordinate(k) {
			values := strings.Split(i, " ")

			x, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, fmt.Errorf("failed to get coordinate: %w", err)
			}

			y, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, fmt.Errorf("failed to get coordinate: %w", err)
			}

			currentCoordinate = marsrobot.Coordinate{
				X: x,
				Y: y,
			}

			currentDirection = marsrobot.Direction{
				Direction: values[2],
			}
		} else {
			values := strings.Split(i, "")

			for _, v := range values {
				newInstruction, err := robot.NewInstruction(v, worldMap)
				if err != nil {
					return nil, fmt.Errorf("failed to get instruction: %w", err)
				}

				instructions = append(instructions, newInstruction)
			}

			robots = append(robots, robot.New(currentCoordinate, currentDirection, instructions))
			currentCoordinate = marsrobot.Coordinate{}
			instructions = make([]robot.Instruction, 0)
			currentDirection = marsrobot.Direction{}
		}
	}

	var printOutput []string

	for _, r := range robots {
		r.ExecuteInstructions()

		line := fmt.Sprintf(
			"%d %d %s",
			r.Coordinate.X,
			r.Coordinate.Y,
			r.Direction.Direction,
		)
		if r.IsLost {
			line += " LOST"
		}

		printOutput = append(printOutput, line)
	}

	return printOutput, nil
}

func isLineWithCoordinate(k int) bool {
	return (k+1)%2 == 0
}

func isFirstLine(k int) bool {
	return k == 0
}
