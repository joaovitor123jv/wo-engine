package woutils

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	texture       *sdl.Texture
	destRect      sdl.Rect
	srcRect       sdl.Rect
	customSrcRect bool
	imagePath     string
	canRender     bool
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
		srcRect:       sdl.Rect{},
		customSrcRect: false,
		canRender:     true,
	}
}

func (i *Image) FillArea(x, y, width, height int32) {
	i.destRect.X = x
	i.destRect.Y = y
	i.destRect.W = width
	i.destRect.H = height
}

func (i *Image) CentralizeOn(x, y int32) {
	i.destRect.X = x - i.destRect.W/2
	i.destRect.Y = y - i.destRect.H/2
}

func (i *Image) SetSrcRect(x, y, w, h int32) {
	i.srcRect = sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
	i.customSrcRect = true
}

func (i *Image) SetSize(width, height int32) {
	i.destRect.W = width
	i.destRect.H = height
}

func (i *Image) Destroy() {
	if i.texture != nil {
		i.texture.Destroy()
	}
}

func (i *Image) Render(renderer *sdl.Renderer) {
	if !i.canRender {
		return
	}

	if i.customSrcRect {
		renderer.Copy(i.texture, &i.srcRect, &i.destRect)
	} else {
		renderer.Copy(i.texture, nil, &i.destRect)
	}
}

func (i *Image) Hide() {
	i.canRender = false
}

func (i *Image) Show() {
	i.canRender = true
}

func (i *Image) ToggleVisibility() {
	i.canRender = !i.canRender
}

func (i *Image) FollowCursor(x, y int32) bool {
	i.destRect.X = x
	i.destRect.Y = y
	return false
}

func (i *Image) Contains(x, y int32) bool {
	r := Rect(i.destRect)
	return r.Contains(x, y)
}
