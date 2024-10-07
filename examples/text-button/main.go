package main

import (
	"log"

	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	game := woengine.NewGame("Warrior's Odyssey")

	game.SetEntrypoint(func() bool {
		gameContext := woutils.NewContext("Init screen")
		defer gameContext.Destroy()

		gameContext.Start() // The start method creates the window and renderer, you can also set the window size before calling it
		renderer := gameContext.GetRenderer()

		button := woutils.NewButtonWithText(renderer, "Press me")
		defer button.Destroy()

		button2 := woutils.NewButtonWithText(renderer, "Exit")
		defer button2.Destroy()

		running := true
		clickedExit := false

		button.OnClick(func() {
			log.Println("Button Pressed")
		})

		button2.OnClick(func() {
			log.Println("Button2 Pressed")
			clickedExit = true
		})

		button2.SetPosition(30, 90)

		button2.AddListeners(&gameContext)
		button.AddListeners(&gameContext)

		gameContext.AddRenderable(&button)
		gameContext.AddRenderable(&button2)

		for running && !clickedExit {
			// Processa eventos
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				running = gameContext.HandleEvent(&event)
			}

			gameContext.Render()

			sdl.Delay(16)
		}
		return false
	})
	game.Run()

}
