package main

import (
	"github.com/mattmoore/opengl_go/carbon"
	"github.com/mattmoore/opengl_go/carbon/geometry"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	engine := carbon.Engine{}
	engine.Graphics = &carbon.Graphics{
		GraphicsOptions: &carbon.GraphicsOptions{1280, 720},
	}

	engine.Init()

	for _, actor := range loadActors() {
		engine.Actors = append(engine.Actors, actor)
	}

	engine.Graphics.Render()
	engine.Quit()
}

func loadActors() []*carbon.Actor {
	var actors []*carbon.Actor

	cubeGeometry := geometry.NewCube(50, 50, 50)

	cube := carbon.Actor{}
	cube.Geometry = cubeGeometry
	cube.Position = carbon.Position{X: 5, Y: 5, Z: 5}

	actors = append(actors, &cube)

	return actors
}
