package woengine

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
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
	// Initializes SDL2
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO); err != nil {
		log.Fatalf("Failed SDL initialization: %s", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("Failed SDL_ttf initialization (text loading): %s", err)
	}
	defer ttf.Quit()

	// Prepares SDL2_image to load PNG files
	if err := img.Init(img.INIT_PNG); err != nil {
		log.Fatalf("Failed SDL_image initialization (png loading): %s", sdl.GetError())
	}
	defer img.Quit()

	// Prepares SDL2_mixer to load MP3 files
	if err := mix.Init(mix.INIT_MP3 | mix.INIT_OGG); err != nil {
		log.Println(err)
		return
	}
	defer mix.Quit()

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	g.entrypoint()
}
