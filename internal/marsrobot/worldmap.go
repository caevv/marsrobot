package marsrobot

type WorldMap struct {
	TopRightCoordinate Coordinate
	Scents             []Coordinate
}

func (w *WorldMap) IsInbound(coordinate Coordinate) bool {
	if !coordinate.isInBoundNorth(w.TopRightCoordinate) || !coordinate.isInBoundEast(w.TopRightCoordinate) {
		return false
	}

	if !coordinate.IsValid() {
		return false
	}

	return true
}

func (w *WorldMap) HasScent(coordinate Coordinate) bool {
	for _, v := range w.Scents {
		if isSame := v.IsEqual(coordinate); isSame {
			return true
		}
	}

	return false
}

func (w *WorldMap) AddScent(coordinate Coordinate) {
	w.Scents = append(w.Scents, coordinate)
}
