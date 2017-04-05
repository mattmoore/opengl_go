package main

import (
	"fmt"
	"github.com/mattmoore/opengl_go/carbon"
	"github.com/mattmoore/opengl_go/carbon/geometry"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	engine := carbon.Engine{}
	engine.Graphics = &carbon.Graphics{
		GraphicsOptions: &carbon.GraphicsOptions{1280, 720},
	}
	engine.Init()

	sdl.Delay(1000)
	engine.Quit()
}

func makeGeometry() []geometry.Geometry {
	var objects []geometry.Geometry

	cube := geometry.NewCube(50, 50, 50)
	cube.Position = geometry.Position{X: 5, Y: 5, Z: 5}
	fmt.Println(cube)
	objects = append(objects, cube)

	return objects
}
