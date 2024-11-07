package woutils

import (
	"log"

	womixins "github.com/joaovitor123jv/wo-engine/wo-mixins"
	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	womixins.HideMixin
	womixins.RectMixin
	texture       *sdl.Texture
	srcRect       sdl.Rect
	customSrcRect bool
	imagePath     string
}

func NewImage(context *GameContext, imagePath string) Image {
	texture, err := LoadTexture(context.GetRenderer(), imagePath)
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
		RectMixin: womixins.RectMixin{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
		srcRect:       sdl.Rect{},
		customSrcRect: false,
		HideMixin:     womixins.NewHideMixin(),
	}
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

func (i *Image) Destroy() {
	if i.texture != nil {
		i.texture.Destroy()
	}
}

func (i *Image) Render(context *GameContext) {
	if i.texture == nil || !i.HasArea() {
		return
	}

	if i.customSrcRect {
		context.GetRenderer().Copy(i.texture, &i.srcRect, i.SdlRect())
	} else {
		context.GetRenderer().Copy(i.texture, nil, i.SdlRect())
	}
}

func (i *Image) FollowCursor(x, y int32) bool {
	i.SetPosition(x, y)
	return false
}
