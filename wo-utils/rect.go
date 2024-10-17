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

func (r *Rect) GetCenter() (x, y int32) {
	x = r.X + r.W/2
	y = r.Y + r.H/2
	return x, y
}

func (r *Rect) Contains(x, y int32) bool {
	return (x >= r.X) && (x <= r.X+r.W) && (y >= r.Y) && (y <= r.Y+r.H)
}

func (this *Rect) SetPosition(x, y int32) {
	this.X = x
	this.Y = y
}
