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

	toggleIsOn := true
	toggleButton := woutils.NewButtonWithText(renderer, "Click Me (on)")
	defer toggleButton.Destroy()

	pressMeButton := woutils.NewButtonWithText(renderer, "Press me")
	defer pressMeButton.Destroy()

	exitButton := woutils.NewButtonWithText(renderer, "Exit")
	defer exitButton.Destroy()

	pressMeButton.OnClick(func() {
		log.Println("Button Pressed")
	})

	exitButton.OnClick(func() {
		log.Println("exitButton Pressed")
		gameContext.StopExecution()
	})

	toggleButton.OnClick(func() {
		if toggleIsOn {
			toggleButton.SetText(renderer, "Click Me (Off)")
			toggleIsOn = false
		} else {
			toggleButton.SetText(renderer, "Click Me (On)")
			toggleIsOn = true
		}
	})

	exitButton.SetPosition(30, 90)
	toggleButton.CenterOn(300, 300)

	exitButton.AddListeners(&gameContext)
	pressMeButton.AddListeners(&gameContext)
	toggleButton.AddListeners(&gameContext)

	gameContext.AddRenderable(&pressMeButton)
	gameContext.AddRenderable(&exitButton)
	gameContext.AddRenderable(&toggleButton)

	gameContext.MainLoop()
}

func main() {
	game := woengine.NewGame()

	game.SetEntrypoint(gameLogic)
	game.Run()

}
