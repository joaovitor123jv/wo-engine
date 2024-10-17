package woutils

import (
	"log"
	"slices"

	wointerfaces "github.com/joaovitor123jv/wo-engine/wo-interfaces"
	"github.com/veandco/go-sdl2/sdl"
)

type GameContext struct {
	mouseMovementListeners []func(x, y int32) bool
	mouseClickListeners    []func(x, y int32, button uint8, isPressed bool) bool
	renderQueue            []wointerfaces.Renderable
	windowWidth            int32
	windowHeight           int32
	gameName               string
	window                 *Window
	renderer               *sdl.Renderer
	shouldExit             bool
	targetFramerate        uint32
	lastFrameTime          uint64
}

func NewContext(gameName string) GameContext {
	return GameContext{
		mouseMovementListeners: nil,
		mouseClickListeners:    nil,
		renderQueue:            nil,
		windowWidth:            INIT_SCREEN_WINDOW_WIDTH,
		windowHeight:           INIT_SCREEN_WINDOW_HEIGHT,
		gameName:               gameName,
		window:                 nil,
		renderer:               nil,
		shouldExit:             false,
		targetFramerate:        30,
		lastFrameTime:          0,
	}
}

func (gc *GameContext) Start() {
	var err error
	gc.window = NewWindow(gc.gameName, gc.windowWidth, gc.windowHeight)

	if gc.renderer, err = sdl.CreateRenderer(gc.window.AsSDLWindow(), -1, sdl.RENDERER_ACCELERATED); err != nil {
		log.Fatalf("Failed to create renderer: %s", err)
	}
}

func (gc *GameContext) GetTargetFramerate() uint32 {
	return gc.targetFramerate
}

func (gc *GameContext) SetTargetFramerate(framesPerSecond uint32) {
	if framesPerSecond == 0 {
		log.Fatalf("Invalid target framerate: %d. Should be greater than 0", framesPerSecond)
	}

	gc.targetFramerate = framesPerSecond
}

func (gc *GameContext) GetWindowSize() (int32, int32) {
	width, height := gc.window.GetSize()
	return width, height
}

func (gc *GameContext) GetWindowCenter() (int32, int32) {
	return gc.window.GetCenter()
}

func (gc *GameContext) GetRenderer() *sdl.Renderer {
	return gc.renderer
}

func (gc *GameContext) StopExecution() {
	gc.shouldExit = true
}

// Set the window size and update the window size if the window is already created
// The struct variables are needed to keep the desired size if the window is not created yet
func (gc *GameContext) SetWindowSize(width, height int32) {
	gc.windowWidth = width
	gc.windowHeight = height

	if gc.window != nil {
		gc.window.SetSize(width, height)
	}
}

func (gc *GameContext) AddMouseMovementListener(listener func(x, y int32) bool) {
	gc.mouseMovementListeners = append(gc.mouseMovementListeners, listener)
}

func (gc *GameContext) AddMouseClickListener(listener func(x, y int32, button uint8, isPressed bool) bool) {
	gc.mouseClickListeners = append(gc.mouseClickListeners, listener)
}

// The events are handled in the main loop in the "backward" order, to
// match the order in which they were added
//
// This is important because the last added listener should have the priority
// to handle the event (e.g. a button click over another button)
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
		for _, listener := range slices.Backward(gc.mouseMovementListeners) {
			if listener(t.X, t.Y) {
				return keepRunning
			}
		}
	case *sdl.MouseButtonEvent:
		for _, listener := range slices.Backward(gc.mouseClickListeners) {
			// Listener returns true to stop iteration
			if listener(t.X, t.Y, t.Button, t.State == sdl.PRESSED) {
				return keepRunning
			}
		}
	}

	return keepRunning
}

func (gc *GameContext) AddRenderable(thingToRender wointerfaces.Renderable) {
	gc.renderQueue = append(gc.renderQueue, thingToRender)
}

func (gc *GameContext) Render() {
	if gc.renderer == nil {
		panic("Renderer not initialized. Did you run Start()?")
	}

	// Limpa a tela com uma cor (neste caso, preta)
	if err := gc.renderer.SetDrawColor(20, 0, 20, 255); err != nil {
		log.Fatalf("Falha ao definir cor de desenho: %s", err)
	}

	if err := gc.renderer.Clear(); err != nil {
		log.Fatalf("Falha ao limpar o renderer: %s", err)
	}

	for _, renderable := range gc.renderQueue {
		renderable.Render(gc.renderer)
	}

	// Atualiza a janela com o frame atual
	gc.renderer.Present()
}

func (gc *GameContext) Destroy() {
	if gc.renderer != nil {
		gc.renderer.Destroy()
	}

	if gc.window != nil {
		gc.window.Destroy()
	}
}

func (gc *GameContext) runDelay() {
	if gc.lastFrameTime == 0 {
		return
	}

	sdl.Delay((1000 / gc.targetFramerate) - uint32(sdl.GetTicks64()-gc.lastFrameTime))
}

func (gc *GameContext) MainLoop() {
	running := true

	for running && !gc.shouldExit {
		// Processa eventos
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			running = gc.HandleEvent(&event)
		}

		gc.Render()

		gc.runDelay()
	}
}

func (gc *GameContext) EnterFullScreen() {
	if gc.window == nil {
		log.Fatalf("Cannot enter fullscreen mode without a window")
	}

	gc.window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
}

func (gc *GameContext) ExitFullScreen() {
	if gc.window == nil {
		log.Fatalf("Cannot enter fullscreen mode without a window")
	}

	gc.window.SetFullscreen(0)
}
