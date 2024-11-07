package woutils

import (
	"strconv"

	womixins "github.com/joaovitor123jv/wo-engine/wo-mixins"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type FramerateViewer struct {
	womixins.HideMixin
	lastTicks uint64
}

func NewFramerateViewer() FramerateViewer {
	return FramerateViewer{
		HideMixin: womixins.NewHideMixin(),
		lastTicks: 0,
	}
}

func (fv *FramerateViewer) Render(context *GameContext) {
	currentTicks := sdl.GetTicks64()

	frame_time := currentTicks - fv.lastTicks
	var fps uint64

	if frame_time > 0 {
		fps = 1000 / frame_time
	} else {
		fps = 0
	}

	fv.lastTicks = currentTicks

	gfx.StringColor(context.GetRenderer(), 16, 16, "FPS: "+strconv.FormatUint(fps, 10), sdl.Color{R: 255, G: 255, B: 0, A: 255})
}
