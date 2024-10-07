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
		windowWidth := woutils.INIT_SCREEN_WINDOW_WIDTH
		windowHeight := woutils.INIT_SCREEN_WINDOW_HEIGHT

		var err error
		var window *sdl.Window
		var renderer *sdl.Renderer

		// Cria a janela
		if window, err = sdl.CreateWindow(
			"Init Screen",
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			windowWidth,
			windowHeight,
			sdl.WINDOW_SHOWN,
		); err != nil {
			log.Fatalf("Failed to create window: %s", err)
		}
		defer window.Destroy()

		window.SetResizable(true)

		// Cria o renderer
		if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
			log.Fatalf("Failed to create renderer: %s", err)
		}
		defer renderer.Destroy()

		gameContext := woutils.NewContext()

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

			gameContext.Render(renderer)

			sdl.Delay(16)
		}
		return false
	})
	game.Run()

}
