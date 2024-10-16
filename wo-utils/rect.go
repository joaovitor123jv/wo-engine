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

func (r *Rect) AsSdlRect() *sdl.Rect {
	return (*sdl.Rect)(r)
}

func (r *Rect) IsPointInside(x, y int32) bool {
	return (x >= r.X) && (x <= r.X+r.W) && (y >= r.Y) && (y <= r.Y+r.H)
}

func (r *Rect) Contains(x, y int32) bool {
	return r.IsPointInside(x, y)
}

func (this *Rect) SetPosition(x, y int32) {
	this.X = x
	this.Y = y
}
