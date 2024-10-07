package main

import (
	"log"

	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	gameContext := woutils.NewContext("Text Button Example")
	defer gameContext.Destroy()

	gameContext.Start() // The start method creates the window and renderer, you can also set the window size before calling it
	renderer := gameContext.GetRenderer()

	button := woutils.NewButtonWithText(renderer, "Press me")
	defer button.Destroy()

	button2 := woutils.NewButtonWithText(renderer, "Exit")
	defer button2.Destroy()

	button.OnClick(func() {
		log.Println("Button Pressed")
	})

	button2.OnClick(func() {
		log.Println("Button2 Pressed")
		gameContext.StopExecution()
	})

	button2.SetPosition(30, 90)

	button2.AddListeners(&gameContext)
	button.AddListeners(&gameContext)

	gameContext.AddRenderable(&button)
	gameContext.AddRenderable(&button2)

	gameContext.MainLoop()
}

func main() {
	game := woengine.NewGame()

	game.SetEntrypoint(gameLogic)
	game.Run()

}
