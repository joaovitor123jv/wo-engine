package woutils

import "github.com/veandco/go-sdl2/sdl"

type Rect sdl.Rect

func NewRect(x, y, w, h int32) Rect {
	return Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (this *Rect) AsSdlRect() *sdl.Rect {
	return (*sdl.Rect)(this)
}

func (this *Rect) IsPointInside(x, y int32) bool {
	return (x >= this.X) && (x <= this.X+this.W) && (y >= this.Y) && (y <= this.Y+this.H)
}

func (this *Rect) SetPosition(x, y int32) {
	this.X = x
	this.Y = y
}
