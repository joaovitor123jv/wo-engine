package woutils

import (
	"embed"
	"log"

	womixins "github.com/joaovitor123jv/wo-engine/wo-mixins"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	womixins.HideMixin
	womixins.RectMixin
	womixins.ColorMixin
	text         string
	renderedText *sdl.Texture
	font         *ttf.Font
	rwops        *sdl.RWops // rwops is used to keep the font data open (when using in-memory fonts)
}

//go:embed assets/fonts/default.ttf
var fontData embed.FS

func NewText(context *GameContext, text string) Text {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture
	var font *ttf.Font
	color := womixins.NewColorMixin(255, 255, 255, 255)

	fontBytes, err := fontData.ReadFile("assets/fonts/default.ttf")
	if err != nil {
		panic(err)
	}

	rwops, err := sdl.RWFromMem(fontBytes)
	if err != nil {
		panic(err)
	}
	// DO NOT defer rwops.Close() because it will close the font data and cause a panic when rendering text

	font, err = ttf.OpenFontRW(rwops, 0, 16)
	if err != nil {
		panic(err)
	}

	if surfaceText, err = font.RenderUTF8Blended(text, color.SdlColor()); err != nil {
		panic(err)
	}
	defer surfaceText.Free()

	if renderedText, err = context.GetRenderer().CreateTextureFromSurface(surfaceText); err != nil {
		panic(err)
	}

	_, _, width, height, err := renderedText.Query()

	return Text{
		HideMixin:    womixins.NewHideMixin(),
		ColorMixin:   color,
		text:         text,
		renderedText: renderedText,
		font:         font,
		rwops:        rwops,
		RectMixin: womixins.RectMixin{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
	}
}

func NewTextWithCustomFont(context *GameContext, customFont string, text string) Text {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture
	var font *ttf.Font
	color := womixins.NewColorMixin(255, 255, 255, 255)

	if font, err = ttf.OpenFont(customFont, 16); err != nil {
		panic(err)
	}

	if surfaceText, err = font.RenderUTF8Blended(text, color.SdlColor()); err != nil {
		panic(err)
	}
	defer surfaceText.Free()

	if renderedText, err = context.GetRenderer().CreateTextureFromSurface(surfaceText); err != nil {
		panic(err)
	}

	_, _, width, height, err := renderedText.Query()

	return Text{
		HideMixin:    womixins.NewHideMixin(),
		ColorMixin:   color,
		text:         text,
		renderedText: renderedText,
		font:         font,
		rwops:        nil, // rwops is nil because we are using a file font
		RectMixin: womixins.RectMixin{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
	}
}

func (t *Text) SetText(context *GameContext, newText string) {
	var err error
	var surfaceText *sdl.Surface
	var renderedText *sdl.Texture

	if t.font == nil {
		log.Fatal("Font is nil")
	}

	if surfaceText, err = t.font.RenderUTF8Blended(newText, t.SdlColor()); err != nil {
		log.Fatal(err)
	}
	defer surfaceText.Free()

	if renderedText, err = context.GetRenderer().CreateTextureFromSurface(surfaceText); err != nil {
		log.Fatal(err)
	}

	_, _, width, height, err := renderedText.Query()
	if err != nil {
		log.Fatal(err)
	}

	t.text = newText

	if t.renderedText != nil {
		aux := t.renderedText
		t.renderedText = renderedText
		t.SetSize(width, height)
		aux.Destroy()
	} else {
		t.renderedText = renderedText
		t.SetSize(width, height)
	}
}

func (t *Text) Refresh(context *GameContext) {
	t.SetText(context, t.text)
}

func (t *Text) Destroy() {
	if t.rwops != nil {
		t.rwops.Close()
	}

	if t.font != nil {
		t.font.Close()
	}

	if t.renderedText != nil {
		t.renderedText.Destroy()
	}
}

func (t *Text) Render(context *GameContext) {
	if t.renderedText == nil || !t.HasArea() {
		return
	}

	context.GetRenderer().Copy(t.renderedText, nil, t.SdlRect())
}
