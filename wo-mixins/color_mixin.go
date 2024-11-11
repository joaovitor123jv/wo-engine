package womixins

import "github.com/veandco/go-sdl2/sdl"

type ColorMixin sdl.Color

func NewColorMixin(r, g, b, a uint8) ColorMixin {
	return ColorMixin{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (c *ColorMixin) SdlColor() sdl.Color {
	return sdl.Color(*c)
}

func (c *ColorMixin) SetColor(r, g, b uint8) {
	c.R = r
	c.G = g
	c.B = b
}

func (c *ColorMixin) SetAlpha(a uint8) {
	c.A = a
}
