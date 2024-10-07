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
	text            *Text
	destRect        Rect
	idleTexture     *sdl.Texture
	pressedTexture  *sdl.Texture
	hoverTexture    *sdl.Texture
	disabledTexture *sdl.Texture
	state           ButtonState
	size            ButtonSize
	behaviour       ButtonBehavior
	canRender       bool
	onClick         func()
}

func NewButton() Button {
	return Button{
		text: nil,
		destRect: NewRect(
			0,
			0,
			100,
			100,
		),
		idleTexture:     nil,
		pressedTexture:  nil,
		hoverTexture:    nil,
		disabledTexture: nil,
		state:           Idle,
		size:            MediumButton,
		canRender:       true,
		onClick:         nil,
	}
}

func NewButtonWithText(renderer *sdl.Renderer, text string) Button {
	button := NewButton()
	uiText := NewText(renderer, text)
	button.text = &uiText

	button.setDefaultIdle(renderer)
	button.setDefaultHover(renderer)
	button.setDefaultPressed(renderer)
	button.setDefaultDisabled(renderer)

	button.updateDimensions()
	return button
}

func (b *Button) updateDimensions() {
	if b.text != nil {
		width, height := b.text.GetDimensions()
		b.destRect.W = width + 40
		b.destRect.H = height + 40
		b.text.SetPosition(b.destRect.X+15, b.destRect.Y+15)
	}
}

func (b *Button) Render(renderer *sdl.Renderer) {
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

	renderer.Copy(texture, nil, b.destRect.AsSdlRect())

	if b.text != nil {
		b.text.Render(renderer)
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

func (b *Button) MouseMovementListener(x, y int32) bool {
	if b.state == Pressed {
		return true
	}

	if b.destRect.IsPointInside(x, y) {
		b.state = Hover
		return true
	}

	b.state = Idle
	return false
}

func (b *Button) MouseClickListener(x, y int32, button uint8, isPressed bool) bool {
	if b.state == Disabled {
		return false
	}

	if button == sdl.BUTTON_LEFT {
		if isPressed {
			if b.destRect.IsPointInside(x, y) {
				b.state = Pressed
				return true
			}
		}

		if b.state == Pressed && !isPressed {
			if b.destRect.IsPointInside(x, y) {
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

	if b.text != nil {
		b.text.SetPosition(x+15, y+15)
	}
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

func (b *Button) SetIdle(renderer *sdl.Renderer, path string) {
	b.idleTexture = getTextureFromFile(renderer, path)
}

func (b *Button) SetPressed(renderer *sdl.Renderer, path string) {
	b.pressedTexture = getTextureFromFile(renderer, path)
}

func (b *Button) SetHover(renderer *sdl.Renderer, path string) {
	b.hoverTexture = getTextureFromFile(renderer, path)
}

func (b *Button) SetDisabled(renderer *sdl.Renderer, path string) {
	b.disabledTexture = getTextureFromFile(renderer, path)
}

func (b *Button) Disable() {
	b.state = Disabled
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
