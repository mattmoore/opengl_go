package carbon

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	Graphics *Graphics
	Sound    *Sound
}

func (e *Engine) Init() {
	sdl.Init(sdl.INIT_EVERYTHING)
	e.Graphics.Init()
	e.Sound.Init()
}

func (e *Engine) Quit() {
	sdl.Quit()
}
