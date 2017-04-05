package carbon

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Graphics struct {
	GraphicsOptions *GraphicsOptions
	Window          *sdl.Window
	Renderer        *sdl.Renderer
}

type GraphicsOptions struct {
	Width  int
	Height int
}

func (g *Graphics) Init() {
	var err error = nil

	g.Window, err = sdl.CreateWindow(
		"Carbon",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		g.GraphicsOptions.Width,
		g.GraphicsOptions.Height,
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		panic(err)
	}

	g.Renderer, err = sdl.CreateRenderer(g.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}
}
