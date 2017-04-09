package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mattmoore/opengl_go/carbon"
	"github.com/mattmoore/opengl_go/carbon/geometry"
	"github.com/veandco/go-sdl2/sdl"
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

	render(engine.Graphics)

	sdl.Delay(1000)
	engine.Quit()
}

func render(g *carbon.Graphics) {
	gl.Viewport(0, 0, int32(g.GraphicsOptions.Width), int32(g.GraphicsOptions.Height))
	gl.ClearColor(0.0, 0.1, 0.0, 1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

  var vbo uint32
  gl.GenBuffers(1, &vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
}

func makeGeometry() []geometry.Geometry {
	var objects []geometry.Geometry

	cube := geometry.NewCube(50, 50, 50)
	cube.Position = geometry.Position{X: 5, Y: 5, Z: 5}
	fmt.Println(cube)
	objects = append(objects, cube)

	return objects
}

var cubeVertices = []float32{
  //  X, Y, Z, U, V
  // Bottom
  -1.0, -1.0, -1.0, 0.0, 0.0,
  1.0, -1.0, -1.0, 1.0, 0.0,
  -1.0, -1.0, 1.0, 0.0, 1.0,
  1.0, -1.0, -1.0, 1.0, 0.0,
  1.0, -1.0, 1.0, 1.0, 1.0,
  -1.0, -1.0, 1.0, 0.0, 1.0,

  // Top
  -1.0, 1.0, -1.0, 0.0, 0.0,
  -1.0, 1.0, 1.0, 0.0, 1.0,
  1.0, 1.0, -1.0, 1.0, 0.0,
  1.0, 1.0, -1.0, 1.0, 0.0,
  -1.0, 1.0, 1.0, 0.0, 1.0,
  1.0, 1.0, 1.0, 1.0, 1.0,

  // Front
  -1.0, -1.0, 1.0, 1.0, 0.0,
  1.0, -1.0, 1.0, 0.0, 0.0,
  -1.0, 1.0, 1.0, 1.0, 1.0,
  1.0, -1.0, 1.0, 0.0, 0.0,
  1.0, 1.0, 1.0, 0.0, 1.0,
  -1.0, 1.0, 1.0, 1.0, 1.0,

  // Back
  -1.0, -1.0, -1.0, 0.0, 0.0,
  -1.0, 1.0, -1.0, 0.0, 1.0,
  1.0, -1.0, -1.0, 1.0, 0.0,
  1.0, -1.0, -1.0, 1.0, 0.0,
  -1.0, 1.0, -1.0, 0.0, 1.0,
  1.0, 1.0, -1.0, 1.0, 1.0,

  // Left
  -1.0, -1.0, 1.0, 0.0, 1.0,
  -1.0, 1.0, -1.0, 1.0, 0.0,
  -1.0, -1.0, -1.0, 0.0, 0.0,
  -1.0, -1.0, 1.0, 0.0, 1.0,
  -1.0, 1.0, 1.0, 1.0, 1.0,
  -1.0, 1.0, -1.0, 1.0, 0.0,

  // Right
  1.0, -1.0, 1.0, 1.0, 1.0,
  1.0, -1.0, -1.0, 1.0, 0.0,
  1.0, 1.0, -1.0, 0.0, 0.0,
  1.0, -1.0, 1.0, 1.0, 1.0,
  1.0, 1.0, -1.0, 0.0, 0.0,
  1.0, 1.0, 1.0, 0.0, 1.0,
}
