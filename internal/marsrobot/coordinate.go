package marsrobot

const MINX = 0
const MINY = 0

const MAXX = 50
const MAXY = 50

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) IsValid() bool {
	lowestCoordinate := Coordinate{X: MINX, Y: MINY}
	highestCoordinate := Coordinate{X: MAXX, Y: MAXY}

	return c.isInBoundWest(lowestCoordinate) || c.isInBoundSouth(lowestCoordinate) ||
		c.isInBoundEast(highestCoordinate) || c.isInBoundNorth(highestCoordinate)
}

func (c Coordinate) NewDirection(d Direction) Coordinate {
	switch d.Direction {
	case "N":
		return Coordinate{X: c.X, Y: c.Y + 1}
	case "E":
		return Coordinate{X: c.X + 1, Y: c.Y}
	case "W":
		return Coordinate{X: c.X - 1, Y: c.Y}
	case "S":
		return Coordinate{X: c.X, Y: c.Y - 1}
	}

	return Coordinate{}
}

func (c Coordinate) isInBoundNorth(newCoordinate Coordinate) bool {
	return c.Y <= newCoordinate.Y
}

func (c Coordinate) isInBoundSouth(newCoordinate Coordinate) bool {
	return c.Y >= newCoordinate.Y
}

func (c Coordinate) isInBoundWest(newCoordinate Coordinate) bool {
	return c.X >= newCoordinate.X
}

func (c Coordinate) isInBoundEast(newCoordinate Coordinate) bool {
	return c.X <= newCoordinate.X
}

func (c Coordinate) IsEqual(newCoordinate Coordinate) bool {
	return c.X == newCoordinate.X && c.Y == newCoordinate.Y
}
