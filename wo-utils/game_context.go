package woutils

import (
	"log"

	wointerfaces "github.com/joaovitor123jv/wo-engine/wo-interfaces"
	"github.com/veandco/go-sdl2/sdl"
)

type GameContext struct {
	mouseMovementListeners []func(x, y int32) bool
	mouseClickListeners    []func(x, y int32, button uint8, isPressed bool) bool
	renderQueue            []wointerfaces.Renderable
}

func NewContext() GameContext {
	return GameContext{
		mouseMovementListeners: nil,
		mouseClickListeners:    nil,
		renderQueue:            nil,
	}
}

func (gc *GameContext) AddMouseMovementListener(listener func(x, y int32) bool) {
	gc.mouseMovementListeners = append(gc.mouseMovementListeners, listener)
}

func (gc *GameContext) AddMouseClickListener(listener func(x, y int32, button uint8, isPressed bool) bool) {
	gc.mouseClickListeners = append(gc.mouseClickListeners, listener)
}

func (gc *GameContext) HandleEvent(event *sdl.Event) bool {
	keepRunning := true

	switch t := (*event).(type) {
	case *sdl.QuitEvent:
		keepRunning = false
	case *sdl.KeyboardEvent:
		// Se o botão "ESC" for pressionado, fecha o programa
		if t.Keysym.Sym == sdl.K_ESCAPE && t.State == sdl.PRESSED {
			keepRunning = false
		}
	case *sdl.MouseMotionEvent:
		for _, listener := range gc.mouseMovementListeners {
			if listener(t.X, t.Y) {
				return keepRunning
			}
		}
	case *sdl.MouseButtonEvent:
		for _, listener := range gc.mouseClickListeners {
			// Listener returns true to stop iteration
			if listener(t.X, t.Y, t.Button, t.State == sdl.PRESSED) {
				return keepRunning
			}
		}
	}

	return keepRunning
}

func (gc *GameContext) AddRenderable(thingToRender Renderable) {
	gc.renderQueue = append(gc.renderQueue, thingToRender)
}

func (gc *GameContext) Render(renderer *sdl.Renderer) {
	// Limpa a tela com uma cor (neste caso, preta)
	if err := renderer.SetDrawColor(20, 0, 20, 255); err != nil {
		log.Fatalf("Falha ao definir cor de desenho: %s", err)
	}

	if err := renderer.Clear(); err != nil {
		log.Fatalf("Falha ao limpar o renderer: %s", err)
	}

	for _, renderable := range gc.renderQueue {
		renderable.Render(renderer)
	}

	// Atualiza a janela com o frame atual
	renderer.Present()
}
