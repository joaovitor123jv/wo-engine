package main

import (
	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	context := woutils.NewContext("Components Demonstration")
	defer context.Destroy()

	context.Start()

	settings := NewSettings(&context, "assets/settings-background.png")
	defer settings.Destroy()

	centerX, centerY := context.GetWindowCenter()
	button := woutils.NewButtonWithText(&context, "Open Settings")
	defer button.Destroy()
	button.CenterOn(centerX, centerY)

	context.AddRenderable(&button)
	context.AddRenderable(&settings)

	button.AddListeners(&context)
	settings.AddListeners(&context)

	button.OnClick(func() {
		settings.ToggleVisibility()
	})

	context.MainLoop()
}

func main() {
	game := woengine.NewGame()
	game.SetEntrypoint(gameLogic)
	game.Run()
}
