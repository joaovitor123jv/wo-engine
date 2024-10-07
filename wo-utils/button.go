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

const (
	ClickBehavior ButtonBehavior = iota
	ToggleBehavior
)

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
		behaviour:       ClickBehavior,
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

	width, height := button.text.GetDimensions()
	button.destRect.W = width + 40
	button.destRect.H = height + 40
	button.text.SetPosition(button.destRect.X+15, button.destRect.Y+15)
	return button
}

func (this *Button) Render(renderer *sdl.Renderer) {
	if this.canRender == false {
		return
	}

	var texture *sdl.Texture = nil

	switch this.state {
	case Idle:
		texture = this.idleTexture
	case Pressed:
		texture = this.pressedTexture
	case Hover:
		texture = this.hoverTexture
	case Active:
		texture = this.pressedTexture
	case Disabled:
		texture = this.disabledTexture
	}

	renderer.Copy(texture, nil, this.destRect.AsSdlRect())

	if this.text != nil {
		this.text.Render(renderer)
	}
}

func (this *Button) Hide() {
	this.canRender = false
}

func (this *Button) Show() {
	this.canRender = true
}

func (this *Button) ToggleVisibility() {
	this.canRender = !this.canRender
}

func (this *Button) MouseMovementListener(x, y int32) bool {
	if this.state == Pressed {
		return true
	}

	if this.destRect.IsPointInside(x, y) {
		this.state = Hover
		return true
	}

	this.state = Idle
	return false
}

func (this *Button) MouseClickListener(x, y int32, button uint8, isPressed bool) bool {
	if this.state == Disabled {
		return false
	}

	if button == sdl.BUTTON_LEFT {
		if isPressed {
			if this.destRect.IsPointInside(x, y) {
				this.state = Pressed
				return true
			}
		}

		if this.state == Pressed && !isPressed {
			if this.destRect.IsPointInside(x, y) {
				this.state = Hover
			} else {
				this.state = Idle
			}

			if this.onClick != nil {
				this.onClick()
				return true
			}
		}
	}

	return false
}

func (this *Button) OnClick(onClick func()) {
	this.onClick = onClick
}

func (this *Button) AddListeners(screenContext *GameContext) {
	screenContext.AddMouseMovementListener(this.MouseMovementListener)
	screenContext.AddMouseClickListener(this.MouseClickListener)
}

func (this *Button) SetPosition(x, y int32) {
	if this.text != nil {
		this.text.SetPosition(x+15, y+15)
	}

	this.destRect.SetPosition(x, y)
}

func (this *Button) setDefaultIdle(renderer *sdl.Renderer) {
	data, err := buttonImages.ReadFile("assets/images/buttons/idle.png")
	if err != nil {
		log.Fatalf("Failed to load default idle button image (%s): %s", err)
	}
	texture, err := LoadTextureFromEmbedFs(renderer, data)
	if err != nil {
		log.Fatalf("Failed to load default idle button texture (%s): %s", err)
	}
	this.idleTexture = texture
}

func (this *Button) setDefaultPressed(renderer *sdl.Renderer) {
	data, err := buttonImages.ReadFile("assets/images/buttons/pressed.png")
	if err != nil {
		log.Fatalf("Failed to load default pressed button image (%s): %s", err)
	}
	texture, err := LoadTextureFromEmbedFs(renderer, data)
	if err != nil {
		log.Fatalf("Failed to load default pressed button texture (%s): %s", err)
	}
	this.pressedTexture = texture
}

func (this *Button) setDefaultHover(renderer *sdl.Renderer) {
	data, err := buttonImages.ReadFile("assets/images/buttons/hover.png")
	if err != nil {
		log.Fatalf("Failed to load default hover button image (%s): %s", err)
	}
	texture, err := LoadTextureFromEmbedFs(renderer, data)
	if err != nil {
		log.Fatalf("Failed to load default hover button texture (%s): %s", err)
	}
	this.hoverTexture = texture
}

func (this *Button) setDefaultDisabled(renderer *sdl.Renderer) {
	data, err := buttonImages.ReadFile("assets/images/buttons/disabled.png")
	if err != nil {
		log.Fatalf("Failed to load default disabled button image (%s): %s", err)
	}
	texture, err := LoadTextureFromEmbedFs(renderer, data)
	if err != nil {
		log.Fatalf("Failed to load default disabled button texture (%s): %s", err)
	}
	this.disabledTexture = texture
}

func (this *Button) SetIdle(renderer *sdl.Renderer, path string) {
	texture, err := LoadTexture(renderer, path)
	if err != nil {
		log.Fatalf("Failed to load idle button image (%s) and convert to texture: %s", path, err)
	}

	this.idleTexture = texture
}

func (this *Button) SetPressed(renderer *sdl.Renderer, path string) {
	texture, err := LoadTexture(renderer, path)
	if err != nil {
		log.Fatalf("Failed to load pressed button image (%s) and convert to texture: %s", path, err)
	}

	this.pressedTexture = texture
}

func (this *Button) SetHover(renderer *sdl.Renderer, path string) {
	texture, err := LoadTexture(renderer, path)
	if err != nil {
		log.Fatalf("Failed to load hover button image (%s) and convert to texture: %s", path, err)
	}

	this.hoverTexture = texture
}

func (this *Button) SetDisabled(renderer *sdl.Renderer, path string) {
	texture, err := LoadTexture(renderer, path)
	if err != nil {
		log.Fatalf("Failed to load disabled button image (%s) and convert to texture: %s", path, err)
	}

	this.disabledTexture = texture
}

func (this *Button) Disable() {
	this.state = Disabled
}

func (this *Button) Destroy() {
	if this.idleTexture != nil {
		this.idleTexture.Destroy()
	}

	if this.pressedTexture != nil {
		this.pressedTexture.Destroy()
	}

	if this.hoverTexture != nil {
		this.hoverTexture.Destroy()
	}

	if this.disabledTexture != nil {
		this.disabledTexture.Destroy()
	}

	if this.text != nil {
		this.text.Destroy()
	}
}
