package main

import (
	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	context := woutils.NewContext("Text Rendering Example")
	defer context.Destroy()

	context.Start()

	centerX, centerY := context.GetWindowCenter()

	textTopLeft := woutils.NewText(&context, "Text Rendered on Top Left")
	defer textTopLeft.Destroy()

	textCentered := woutils.NewText(&context, "Text Rendered on Window Center")
	defer textCentered.Destroy()

	textCentered.CenterOn(centerX, centerY)

	textTopLeft.SetColor(255, 0, 255)
	textTopLeft.Refresh(&context)

	context.AddRenderable(&textTopLeft)
	context.AddRenderable(&textCentered)

	context.MainLoop()
}

func main() {
	game := woengine.NewGame()
	game.SetEntrypoint(gameLogic)
	game.Run()
}
