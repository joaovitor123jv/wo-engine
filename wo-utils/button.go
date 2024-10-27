package woutils

import (
	"embed"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type ButtonSize uint8
type ButtonState uint8
type ButtonBehavior uint8

const (
	SmallButton ButtonSize = iota + 1
	MediumButton
	LargeButton
)

const (
	Idle ButtonState = iota
	Pressed
	Hover
	Active
	Disabled
)

//go:embed assets/images/buttons/*.png
var buttonImages embed.FS

type Button struct {
	text               *Text
	destRect           Rect
	idleTexture        *sdl.Texture
	pressedTexture     *sdl.Texture
	hoverTexture       *sdl.Texture
	disabledTexture    *sdl.Texture
	collisionRect      Rect
	collisionThreshold Rect
	state              ButtonState
	size               ButtonSize
	behaviour          ButtonBehavior
	canRender          bool
	canListenEvents    bool
	onClick            func()
}

func NewButton() Button {
	destRect := NewRect(0, 0, 100, 100)
	collisionThreshold := NewRect(0, 0, 0, 0)
	return Button{
		text:               nil,
		destRect:           destRect,
		collisionRect:      destRect,
		collisionThreshold: collisionThreshold,
		idleTexture:        nil,
		pressedTexture:     nil,
		hoverTexture:       nil,
		disabledTexture:    nil,
		state:              Idle,
		size:               MediumButton,
		canRender:          true,
		canListenEvents:    true,
		onClick:            nil,
	}
}

func NewButtonWithText(context *GameContext, text string) Button {
	button := NewButton()
	uiText := NewText(context, text)
	button.text = &uiText

	button.setDefaultIdle(context.GetRenderer())
	button.setDefaultHover(context.GetRenderer())
	button.setDefaultPressed(context.GetRenderer())
	button.setDefaultDisabled(context.GetRenderer())

	button.updateDimensions()
	button.setDefaultCollisionThreshold()

	return button
}

func (b *Button) setDefaultCollisionThreshold() {
	b.SetCollisionThreshold(5, 10, 8, 18)
}

// SetCollisionThreshold sets the collision threshold for the button.
// The values are collision percentages from the button's dimensions.
// For example, if the button has a width of 100 and height of 100, and the
// collision threshold is set to 10, 20, 15, 12, the collision area will be
// 10% from the left, 20% from the top, 15% from the right and 12% from the bottom.
func (b *Button) SetCollisionThreshold(xStart, yStart, xEnd, yEnd uint16) {
	b.collisionThreshold = NewRect(int32(xStart), int32(yStart), int32(xEnd), int32(yEnd))

	b.calcCollisionRect()
}

func (b *Button) calcCollisionRect() {
	b.collisionRect = NewRect(
		b.destRect.X+int32((float32(b.destRect.W)*float32(b.collisionThreshold.X))/100),
		b.destRect.Y+int32((float32(b.destRect.H)*float32(b.collisionThreshold.Y))/100),
		b.destRect.W-int32((float32(b.destRect.W)*float32(b.collisionThreshold.X+b.collisionThreshold.W))/100),
		b.destRect.H-int32((float32(b.destRect.H)*float32(b.collisionThreshold.Y+b.collisionThreshold.H))/100),
	)
}

func (b *Button) updateDimensions() {
	if b.text != nil {
		width, height := b.text.GetDimensions()
		b.destRect.W = width + 40
		b.destRect.H = height + 40
		b.text.SetPosition(b.destRect.X+15, b.destRect.Y+15)
	}
	b.calcCollisionRect()
}

func (b *Button) Render(context *GameContext) {
	var texture *sdl.Texture = nil

	switch b.state {
	case Idle:
		texture = b.idleTexture
	case Pressed:
		texture = b.pressedTexture
	case Hover:
		texture = b.hoverTexture
	case Active:
		texture = b.pressedTexture
	case Disabled:
		texture = b.disabledTexture
	}

	context.GetRenderer().Copy(texture, nil, b.destRect.AsSdlRect())

	if b.text != nil {
		b.text.Render(context)
	}
}

func (b *Button) Hide() {
	b.canRender = false
}

func (b *Button) Show() {
	b.canRender = true
}

func (b *Button) ToggleVisibility() {
	b.canRender = !b.canRender
}

func (b *Button) DisableEvents() {
	b.canListenEvents = false
}

func (b *Button) EnableEvents() {
	b.canListenEvents = true
}

func (b *Button) MouseMovementListener(x, y int32) bool {
	if !b.canListenEvents || !b.canRender || b.state == Disabled {
		return false
	}

	if b.state == Pressed {
		return true
	}

	if b.collisionRect.Contains(x, y) {
		b.state = Hover
		return true
	}

	b.state = Idle
	return false
}

func (b *Button) MouseClickListener(x, y int32, button uint8, isPressed bool) bool {
	if !b.canListenEvents || !b.canRender || b.state == Disabled {
		return false
	}

	if button == sdl.BUTTON_LEFT {
		if isPressed {
			if b.collisionRect.Contains(x, y) {
				b.state = Pressed
				return true
			}
		}

		if b.state == Pressed && !isPressed {
			if b.collisionRect.Contains(x, y) {
				b.state = Hover
			} else {
				b.state = Idle
			}

			if b.onClick != nil {
				b.onClick()
				return true
			}
		}
	}

	return false
}

func (b *Button) OnClick(onClick func()) {
	b.onClick = onClick
}

func (b *Button) AddListeners(screenContext *GameContext) {
	screenContext.AddMouseMovementListener(b.MouseMovementListener)
	screenContext.AddMouseClickListener(b.MouseClickListener)
}

func (b *Button) SetPosition(x, y int32) {
	b.destRect.SetPosition(x, y)
	b.calcCollisionRect()

	if b.text != nil {
		b.text.SetPosition(x+15, y+15)
	}
}

func (b *Button) GetCenter() (x, y int32) {
	x = b.destRect.X + b.destRect.W/2
	y = b.destRect.Y + b.destRect.H/2
	return x, y
}

func (b *Button) CenterOn(x, y int32) {
	b.SetPosition(x-b.collisionRect.W/2, y-b.collisionRect.H/2)
}

func (b *Button) SetText(context *GameContext, text string) {
	b.text.SetText(context, text)
	b.updateDimensions()
}

func getTextureFromEmbedFs(renderer *sdl.Renderer, path string) *sdl.Texture {
	data, err := buttonImages.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load default button image (%s): %s", path, err)
	}

	return LoadTextureFromEmbedFs(renderer, data)
}

func (b *Button) setDefaultIdle(renderer *sdl.Renderer) {
	b.idleTexture = getTextureFromEmbedFs(renderer, "assets/images/buttons/idle.png")
}

func (b *Button) setDefaultPressed(renderer *sdl.Renderer) {
	b.pressedTexture = getTextureFromEmbedFs(renderer, "assets/images/buttons/pressed.png")
}

func (b *Button) setDefaultHover(renderer *sdl.Renderer) {
	b.hoverTexture = getTextureFromEmbedFs(renderer, "assets/images/buttons/hover.png")
}

func (b *Button) setDefaultDisabled(renderer *sdl.Renderer) {
	b.disabledTexture = getTextureFromEmbedFs(renderer, "assets/images/buttons/disabled.png")
}

func getTextureFromFile(renderer *sdl.Renderer, path string) *sdl.Texture {
	texture, err := LoadTexture(renderer, path)
	if err != nil {
		log.Fatalf("Failed to load button image \"%s\" and convert to texture: %s", path, err)
	}

	return texture
}

func (b *Button) SetIdle(context *GameContext, path string) {
	b.idleTexture = getTextureFromFile(context.GetRenderer(), path)
}

func (b *Button) SetPressed(context *GameContext, path string) {
	b.pressedTexture = getTextureFromFile(context.GetRenderer(), path)
}

func (b *Button) SetHover(context *GameContext, path string) {
	b.hoverTexture = getTextureFromFile(context.GetRenderer(), path)
}

func (b *Button) SetDisabled(context *GameContext, path string) {
	b.disabledTexture = getTextureFromFile(context.GetRenderer(), path)
}

func (b *Button) Disable() {
	b.state = Disabled
}

func (b *Button) Enable() {
	b.state = Idle
}

func (b *Button) Destroy() {
	if b.idleTexture != nil {
		b.idleTexture.Destroy()
	}

	if b.pressedTexture != nil {
		b.pressedTexture.Destroy()
	}

	if b.hoverTexture != nil {
		b.hoverTexture.Destroy()
	}

	if b.disabledTexture != nil {
		b.disabledTexture.Destroy()
	}

	if b.text != nil {
		b.text.Destroy()
	}
}
