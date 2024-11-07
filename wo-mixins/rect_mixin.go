package womixins

import "github.com/veandco/go-sdl2/sdl"

type RectMixin sdl.Rect

func NewRectMixin(x, y, w, h int32) RectMixin {
	return RectMixin{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (r *RectMixin) SdlRect() *sdl.Rect {
	return (*sdl.Rect)(r)
}

func (r *RectMixin) GetCenter() (x, y int32) {
	x = r.X + r.W/2
	y = r.Y + r.H/2
	return x, y
}

func (r *RectMixin) CenterOn(x, y int32) {
	r.X = x - r.W/2
	r.Y = y - r.H/2
}

func (r *RectMixin) Contains(x, y int32) bool {
	return (x >= r.X) && (x <= r.X+r.W) && (y >= r.Y) && (y <= r.Y+r.H)
}

func (r *RectMixin) SetPosition(x, y int32) {
	r.X = x
	r.Y = y
}

func (r *RectMixin) SetSize(width, height int32) {
	r.W = width
	r.H = height
}

func (r *RectMixin) GetSize() (width, height int32) {
	return r.W, r.H
}

func (r *RectMixin) HasArea() bool {
	return r.W > 0 && r.H > 0
}

func (r *RectMixin) FillArea(x, y, width, height int32) {
	r.X = x
	r.Y = y
	r.W = width
	r.H = height
}
