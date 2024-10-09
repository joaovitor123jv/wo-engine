package woutils

import (
	"embed"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	text         string
	renderedText *sdl.Texture
	font         *ttf.Font
	rect         sdl.Rect
	canRender    bool
}

//go:embed assets/fonts/default.ttf
var fontData embed.FS

func NewText(renderer *sdl.Renderer, text string) Text {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture
	var font *ttf.Font

	fontBytes, err := fontData.ReadFile("assets/fonts/default.ttf")
	if err != nil {
		panic(err)
	}

	rwops, err := sdl.RWFromMem(fontBytes)
	if err != nil {
		panic(err)
	}
	defer rwops.Close()

	font, err = ttf.OpenFontRW(rwops, 0, 16)
	if err != nil {
		panic(err)
	}

	if surfaceText, err = font.RenderUTF8Blended(text, sdl.Color{R: 255, G: 255, B: 255, A: 255}); err != nil {
		panic(err)
	}
	defer surfaceText.Free()

	if renderedText, err = renderer.CreateTextureFromSurface(surfaceText); err != nil {
		panic(err)
	}

	_, _, width, height, err := renderedText.Query()

	return Text{
		text:         text,
		renderedText: renderedText,
		font:         font,
		rect: sdl.Rect{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
		canRender: true,
	}
}

func NewTextWithCustomFont(renderer *sdl.Renderer, customFont string, text string) Text {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture
	var font *ttf.Font

	if font, err = ttf.OpenFont(customFont, 16); err != nil {
		panic(err)
	}

	if surfaceText, err = font.RenderUTF8Blended(text, sdl.Color{R: 255, G: 255, B: 255, A: 255}); err != nil {
		panic(err)
	}
	defer surfaceText.Free()

	if renderedText, err = renderer.CreateTextureFromSurface(surfaceText); err != nil {
		panic(err)
	}

	_, _, width, height, err := renderedText.Query()

	return Text{
		text:         text,
		renderedText: renderedText,
		font:         font,
		rect: sdl.Rect{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
		canRender: true,
	}
}

func (this *Text) ChangeText(renderer *sdl.Renderer, newText string) {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture

	if surfaceText, err = this.font.RenderUTF8Blended(newText, sdl.Color{R: 255, G: 255, B: 255, A: 255}); err != nil {
		panic(err)
	}
	defer surfaceText.Free()

	if renderedText, err = renderer.CreateTextureFromSurface(surfaceText); err != nil {
		panic(err)
	}

	_, _, width, height, err := renderedText.Query()

	this.text = newText
	this.renderedText.Destroy()
	this.renderedText = renderedText
	this.rect.W = width
	this.rect.H = height
}

func (this *Text) Destroy() {
	if this.font != nil {
		this.font.Close()
	}

	if this.renderedText != nil {
		this.renderedText.Destroy()
	}
}

func (this *Text) Render(renderer *sdl.Renderer) {
	if this.renderedText != nil {
		renderer.Copy(this.renderedText, nil, &this.rect)
	}
}

func (this *Text) SetPosition(x, y int32) {
	this.rect.X = x
	this.rect.Y = y
}

func (this *Text) CenterOn(x, y int32) {
	this.rect.X = x - this.rect.W/2
	this.rect.Y = y - this.rect.H/2
}

func (this *Text) Hide() {
	this.canRender = false
}

func (this *Text) Show() {
	this.canRender = true
}

func (this *Text) ToggleVisibility() {
	this.canRender = !this.canRender
}

func (this *Text) GetDimensions() (width, height int32) {
	return this.rect.W, this.rect.H
}
