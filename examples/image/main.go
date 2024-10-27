package main

import (
	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	gameContext := woutils.NewContext("Image Rendering Example")

	gameContext.Start()
	defer gameContext.Destroy()

	image := woutils.NewImage(&gameContext, "img/img.png")
	defer image.Destroy()

	image2 := woutils.NewImage(&gameContext, "img/img.png")
	defer image2.Destroy()

	width, height := gameContext.GetWindowSize()
	image.FillArea(0, 0, width, height)

	centerX, centerY := gameContext.GetWindowCenter()
	image2.SetSize(150, 190)
	image2.CenterOn(centerX, centerY)

	gameContext.AddRenderable(&image)
	gameContext.AddRenderable(&image2)

	gameContext.MainLoop()
}

func main() {
	// You don't need to destroy the Game, I know this is sad, so most of the memory is freed automatically
	// when game.Run() returns
	game := woengine.NewGame()
	game.SetEntrypoint(gameLogic)
	game.Run()
}
