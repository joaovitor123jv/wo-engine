package woutils

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// LoadTexture loads a PNG image from a file and converts it into an SDL texture.
// It takes a renderer (to which the texture will be bound) and the filename of the image.
// Returns a pointer to the created SDL texture and an error if any occurs during loading or texture creation.
func LoadTexture(renderer *sdl.Renderer, filename string) (*sdl.Texture, error) {
	// Load the image using SDL_image
	surface, err := img.Load(filename)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	// Create a texture from the surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	return texture, nil
}

// LoadTextureFromEmbedFs loads a PNG image from embedded filesystem data and converts it into an SDL texture.
// It accepts a renderer (to which the texture will be bound) and the image data as a byte slice.
// Returns a pointer to the created SDL texture. It does NOT return an error because the data is already in memory
// and should be ok.
func LoadTextureFromEmbedFs(renderer *sdl.Renderer, data []byte) *sdl.Texture {
	// Create an RWops from the memory data
	rwops, err := sdl.RWFromMem(data)
	if err != nil {
		panic("Failed to load texture file from embed.FS")
	}
	defer rwops.Close()

	// Load the surface using SDL_image
	surface, err := img.LoadRW(rwops, false)
	if err != nil {
		panic("Failed to build texture surface from embed.FS")
	}
	defer surface.Free()

	// Create a texture from the surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic("Failed to build texture from surface from embed.FS")
	}

	return texture
}
