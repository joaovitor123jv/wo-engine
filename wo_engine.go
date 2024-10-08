package woengine

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	entrypoint func()
}

func init() {
	// Isso é necessário para gerenciar o contexto SDL na thread principal.
	runtime.LockOSThread()
}

func NewGame() Game {
	return Game{
		entrypoint: nil,
	}
}

func (g *Game) SetEntrypoint(entrypoint func()) {
	g.entrypoint = entrypoint
}

func (g *Game) Run() {
	// Inicializa SDL2
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("Failed SDL initialization: %s", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("Failed SDL_ttf initialization (text loading): %s", err)
	}
	defer ttf.Quit()

	// Inicializa SDL2_image para carregar PNG
	if err := img.Init(img.INIT_PNG); err != nil {
		log.Fatalf("Failed SDL_image initialization (png loading): %s", sdl.GetError())
	}
	defer img.Quit()

	g.entrypoint()
}
