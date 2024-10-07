package woutils

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// loadTexture carrega uma imagem PNG e a converte em uma textura SDL
func LoadTexture(renderer *sdl.Renderer, filename string) (*sdl.Texture, error) {
	// Carrega a imagem usando SDL_image
	surface, err := img.Load(filename)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	// Cria uma textura a partir da surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return texture, nil
}

// loadTexture carrega uma imagem PNG e a converte em uma textura SDL
func LoadTextureFromEmbedFs(renderer *sdl.Renderer, data []byte) (*sdl.Texture, error) {
	// Cria um RWops a partir dos dados em memória
	rwops, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}
	defer rwops.Close()

	// Carrega a surface usando SDL_image
	surface, err := img.LoadRW(rwops, false)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	// Cria uma textura a partir da surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return texture, nil
}
