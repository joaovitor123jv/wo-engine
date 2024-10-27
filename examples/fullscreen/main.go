package main

import (
	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	context := woutils.NewContext("Full Screen Example")
	defer context.Destroy()

	context.Start()

	isFullScreen := false
	fullScreenButton := woutils.NewButtonWithText(&context, "Toggle Full Screen")

	centerX, centerY := context.GetWindowCenter()
	fullScreenButton.CenterOn(centerX, centerY)

	fullScreenButton.OnClick(func() {
		if isFullScreen {
			context.ExitFullScreen()
			context.SetWindowSize(800, 600)
			fullScreenButton.CenterOn(centerX, centerY)
			isFullScreen = false
		} else {
			maxWidth, maxHeight := context.GetTotalDisplaySize()

			context.SetWindowSize(maxWidth, maxHeight)
			fullScreenButton.CenterOn(maxWidth/2, maxHeight/2)
			context.EnterFullScreen()

			isFullScreen = true
		}
	})

	fullScreenButton.AddListeners(&context)

	context.AddRenderable(&fullScreenButton)

	context.MainLoop()
}

func main() {
	game := woengine.NewGame()
	game.SetEntrypoint(gameLogic)
	game.Run()
}
