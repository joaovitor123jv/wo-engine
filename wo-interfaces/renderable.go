package wointerfaces

import "github.com/veandco/go-sdl2/sdl"

type Renderable interface {
	Render(renderer *sdl.Renderer)
}
