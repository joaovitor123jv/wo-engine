package main

import (
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Settings struct {
	id            uint32
	settingsLabel woutils.Text
	background    woutils.Image
	closeButton   woutils.Button
	canRender     bool
}

func NewSettings(renderer *sdl.Renderer, backgroundPath string) Settings {
	background := woutils.NewImage(renderer, backgroundPath)
	settingsLabel := woutils.NewText(renderer, "Settings")
	closeButton := woutils.NewButtonWithText(renderer, "Close")
	viewport := renderer.GetViewport() // Get the rendering viewport

	// Viewport Center X and Y
	vpcX, vpcY := viewport.W/2, viewport.H/2

	background.SetSize(vpcX, vpcY)  // Set the image size to half of the viewport size
	background.CenterOn(vpcX, vpcY) // Centralize the image on the viewport
	background.SetSrcRect(2000, 350, 1750, 1200)

	closeButton.CenterOn(vpcX, vpcY+(2*(vpcY/5)))   // Positions the button near the bottom of the background
	settingsLabel.CenterOn(vpcX, vpcY-(2*(vpcY/5))) // Positions the text near the top of the background

	closeButton.Hide()
	background.Hide()
	settingsLabel.Hide()

	return Settings{
		settingsLabel: settingsLabel,
		background:    background,
		closeButton:   closeButton,
	}
}

func (s *Settings) GetId() uint32 {
	return s.id
}

// Used by the game context to manage listeners and renderables list
func (s *Settings) SetId(id uint32) {
	s.id = id
}

func (s *Settings) ToggleVisibility() {
	s.canRender = !s.canRender
	s.background.ToggleVisibility()
	s.closeButton.ToggleVisibility() // This also disables button listeners
	s.settingsLabel.ToggleVisibility()
}

func (s *Settings) Render(renderer *sdl.Renderer) {
	if !s.canRender {
		return
	}

	s.background.Render(renderer)
	s.settingsLabel.Render(renderer)
	s.closeButton.Render(renderer)
}

func (s *Settings) AddListeners(gameContext *woutils.GameContext) {
	s.closeButton.OnClick(func() {
		s.ToggleVisibility()
	})

	stealFocus := func(x, y int32) bool {
		stealFocus := false

		if s.canRender {
			stealFocus = s.background.Contains(x, y)
		}

		return stealFocus
	}

	gameContext.AddMouseMovementListener(stealFocus)
	gameContext.AddMouseClickListener(func(x, y int32, button uint8, isPressed bool) bool {
		return stealFocus(x, y)
	})
	s.closeButton.AddListeners(gameContext)
}

func (s *Settings) Destroy() {
	s.background.Destroy()
	s.closeButton.Destroy()
	s.settingsLabel.Destroy()
}
