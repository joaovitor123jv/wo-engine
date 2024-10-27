package woutils

import "github.com/veandco/go-sdl2/sdl"

type GameCamera struct {
	translationX int32
	translationY int32
	zoom         float32
}

func NewGameCamera() GameCamera {
	return GameCamera{
		translationX: 0,
		translationY: 0,
		zoom:         1,
	}
}

func (gc *GameCamera) SetTranslation(x, y int32) {
	gc.translationX = x
	gc.translationY = y
}

func (gc *GameCamera) Translate(x, y int32) {
	gc.translationX += x
	gc.translationY += y
}

// SetZoom alters the zoom in wich the map is rendered.
// 1 is the default zoom (100%).
// Param "zoom" must be between 0.1 and 10.0
// If zoom is less than 0.1, it will be set to 0.1
// If zoom is greater than 10.0, it will be set to 10.0
func (gc *GameCamera) SetZoom(zoom float32) {
	if zoom < 0.1 {
		zoom = 0.1
	}
	if zoom > 10.0 {
		zoom = 10.0
	}
	gc.zoom = zoom
}

func (gc *GameCamera) GetZoom() float32 {
	return gc.zoom
}

func (gc *GameCamera) GetTranslation() (int32, int32) {
	return gc.translationX, gc.translationY
}

func (gc *GameCamera) ApplyTranslation(x, y int32) (int32, int32) {
	return x + gc.translationX, y + gc.translationY
}

func (gc *GameCamera) TranslateSDLRect(sdlRect *sdl.Rect) {
	sdlRect.X += gc.translationX
	sdlRect.Y += gc.translationY
}
