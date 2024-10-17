package main

import (
	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	context := woutils.NewContext("Music Player Example")
	defer context.Destroy()

	context.Start()

	renderer := context.GetRenderer()

	var music woutils.Music
	defer music.Destroy()

	playButton := woutils.NewButtonWithText(renderer, "Play")
	defer playButton.Destroy()

	stopButton := woutils.NewButtonWithText(renderer, "Stop")
	defer stopButton.Destroy()

	selectMusicAButton := woutils.NewButtonWithText(renderer, "Select Music A")
	defer selectMusicAButton.Destroy()

	selectMusicBButton := woutils.NewButtonWithText(renderer, "Select Music B")
	defer selectMusicBButton.Destroy()

	selectedMusicLabel := woutils.NewText(renderer, "Selected Music: None")
	defer selectedMusicLabel.Destroy()

	centerX, centerY := context.GetWindowCenter()

	playButton.CenterOn(centerX+100, centerY)
	stopButton.CenterOn(centerX-100, centerY)

	selectMusicAButton.CenterOn(centerX, centerY+100)
	selectMusicBButton.CenterOn(centerX, centerY+200)

	selectedMusicLabel.CenterOn(centerX, centerY-150)

	context.AddRenderable(&playButton)
	context.AddRenderable(&stopButton)
	context.AddRenderable(&selectMusicAButton)
	context.AddRenderable(&selectMusicBButton)
	context.AddRenderable(&selectedMusicLabel)

	selectMusicAButton.OnClick(func() {
		selectedMusicLabel.ChangeText(renderer, "Selected Music A")

		if music.Loaded() {
			music.Destroy()
		}

		music = woutils.NewMusic("assets/music-a.mp3")
	})

	selectMusicBButton.OnClick(func() {
		selectedMusicLabel.ChangeText(renderer, "Selected Music B")

		if music.Loaded() {
			music.Destroy()
		}
		music = woutils.NewMusic("assets/music-b.mp3")
	})

	playButton.OnClick(func() {
		if music.Loaded() {
			music.PlayOnce()
		}
	})

	stopButton.OnClick(func() {
		if music.Loaded() {
			music.Stop()
		}
	})

	playButton.AddListeners(&context)
	stopButton.AddListeners(&context)
	selectMusicAButton.AddListeners(&context)
	selectMusicBButton.AddListeners(&context)

	context.MainLoop()
}

func main() {
	game := woengine.NewGame()
	game.SetEntrypoint(gameLogic)
	game.Run()
}
