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

func (t *Text) ChangeText(renderer *sdl.Renderer, newText string) {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture

	if surfaceText, err = t.font.RenderUTF8Blended(newText, sdl.Color{R: 255, G: 255, B: 255, A: 255}); err != nil {
		panic(err)
	}
	defer surfaceText.Free()

	if renderedText, err = renderer.CreateTextureFromSurface(surfaceText); err != nil {
		panic(err)
	}

	_, _, width, height, err := renderedText.Query()

	t.text = newText
	t.renderedText.Destroy()
	t.renderedText = renderedText
	t.rect.W = width
	t.rect.H = height
}

func (t *Text) Destroy() {
	if t.font != nil {
		t.font.Close()
	}

	if t.renderedText != nil {
		t.renderedText.Destroy()
	}
}

func (t *Text) Render(renderer *sdl.Renderer) {
	if t.renderedText != nil {
		renderer.Copy(t.renderedText, nil, &t.rect)
	}
}

func (t *Text) SetPosition(x, y int32) {
	t.rect.X = x
	t.rect.Y = y
}

func (t *Text) CenterOn(x, y int32) {
	t.rect.X = x - t.rect.W/2
	t.rect.Y = y - t.rect.H/2
}

func (t *Text) Hide() {
	t.canRender = false
}

func (t *Text) Show() {
	t.canRender = true
}

func (t *Text) ToggleVisibility() {
	t.canRender = !t.canRender
}

func (t *Text) GetDimensions() (width, height int32) {
	return t.rect.W, t.rect.H
}
