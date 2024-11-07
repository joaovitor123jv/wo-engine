package main

import (
	womixins "github.com/joaovitor123jv/wo-engine/wo-mixins"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

type Settings struct {
	id            uint32
	settingsLabel woutils.Text
	background    woutils.Image
	closeButton   woutils.Button
	womixins.HideMixin
}

func NewSettings(context *woutils.GameContext, backgroundPath string) Settings {
	background := woutils.NewImage(context, backgroundPath)
	settingsLabel := woutils.NewText(context, "Settings")
	closeButton := woutils.NewButtonWithText(context, "Close")
	viewport := context.GetRenderer().GetViewport() // Get the rendering viewport

	// Viewport Center X and Y
	vpcX, vpcY := viewport.W/2, viewport.H/2

	background.SetSize(vpcX, vpcY)  // Set the image size to half of the viewport size
	background.CenterOn(vpcX, vpcY) // Centralize the image on the viewport
	background.SetSrcRect(2000, 350, 1750, 1200)

	closeButton.CenterOn(vpcX, vpcY+(2*(vpcY/5)))   // Positions the button near the bottom of the background
	settingsLabel.CenterOn(vpcX, vpcY-(2*(vpcY/5))) // Positions the text near the top of the background

	// Useful for disabling listeners and rendering
	hideMixin := womixins.NewHideMixin()

	hideMixin.AddDependant(&closeButton)
	hideMixin.AddDependant(&background)
	hideMixin.AddDependant(&settingsLabel)

	hideMixin.Hide()

	return Settings{
		HideMixin:     hideMixin,
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

func (s *Settings) Render(context *woutils.GameContext) {
	s.background.Render(context)
	s.settingsLabel.Render(context)
	s.closeButton.Render(context)
}

func (s *Settings) AddListeners(gameContext *woutils.GameContext) {
	s.closeButton.OnClick(func() {
		s.ToggleVisibility()
	})

	stealFocus := func(x, y int32) (stealFocus bool) {
		if s.IsVisible() {
			return s.background.Contains(x, y)
		} else {
			return false
		}
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
