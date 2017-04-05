package geometry

type Cube struct {
	Geometry
	X        float64
	Y        float64
	Z        float64
	Position Position
}

func NewCube(x, y, z float64) Cube {
	return Cube{
		X:        x,
		Y:        y,
		Z:        z,
		Position: Position{0, 0, 0},
	}
}
