package woutils

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
}

func NewWindow(title string, width int32, height int32) *Window {
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)

	if err != nil {
		log.Fatalf("Failed to create window: %s", err)
	}

	window.SetMinimumSize(800, 600)
	window.SetResizable(true)

	return &Window{window}
}

func (w *Window) GetCenter() (int32, int32) {
	width, height := w.Window.GetSize()
	return width / 2, height / 2
}

func (w *Window) AsSDLWindow() *sdl.Window {
	return w.Window
}
