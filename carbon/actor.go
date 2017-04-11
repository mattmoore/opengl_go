package carbon

import (
	"github.com/mattmoore/opengl_go/carbon/geometry"
)

type Actor struct {
	Geometry geometry.Geometry
	Position Position
}

type Position struct {
	X float64
	Y float64
	Z float64
}

func (actor *Actor) Add() {

}
