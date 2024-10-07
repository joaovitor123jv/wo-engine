package woutils

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	texture   *sdl.Texture
	destRect  sdl.Rect
	imagePath string
	canRender bool
}

func NewImage(renderer *sdl.Renderer, imagePath string) Image {
	texture, err := LoadTexture(renderer, imagePath)
	if err != nil {
		log.Fatalf("Failed to load image (%s) and convert to texture: %s", imagePath, err)
	}

	_, _, width, height, err := texture.Query()
	if err != nil {
		log.Fatalf("Failed to get texture information (%s): %s", imagePath, err)
	}

	return Image{
		imagePath: imagePath,
		texture:   texture,
		destRect: sdl.Rect{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
		canRender: true,
	}
}

func (this *Image) Destroy() {
	this.texture.Destroy()
}

func (this *Image) Render(renderer *sdl.Renderer) {
	if !this.canRender {
		return
	}

	renderer.Copy(this.texture, nil, &this.destRect)
}

func (this *Image) Hide() {
	this.canRender = false
}

func (this *Image) Show() {
	this.canRender = true
}

func (this *Image) ToggleVisibility() {
	this.canRender = !this.canRender
}

func (this *Image) FollowCursor(x, y int32) bool {
	this.destRect.X = x
	this.destRect.Y = y
	return false
}
