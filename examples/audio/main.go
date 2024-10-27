package main

import (
	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
)

func gameLogic() {
	context := woutils.NewContext("Audio Player Example")
	defer context.Destroy()

	context.Start()

	var selectedAudio *woutils.Audio = nil

	audioA := woutils.NewAudio("assets/audio-a.mp3")
	defer audioA.Destroy()

	audioB := woutils.NewAudio("assets/audio-b.mp3")
	defer audioB.Destroy()

	playButton := woutils.NewButtonWithText(&context, "Play")
	defer playButton.Destroy()

	stopButton := woutils.NewButtonWithText(&context, "Stop")
	defer stopButton.Destroy()

	selectAudioAButton := woutils.NewButtonWithText(&context, "Select Audio A")
	defer selectAudioAButton.Destroy()

	selectAudioBButton := woutils.NewButtonWithText(&context, "Select Audio B")
	defer selectAudioBButton.Destroy()

	selectedMusicLabel := woutils.NewText(&context, "Selected Audio: None")
	defer selectedMusicLabel.Destroy()

	centerX, centerY := context.GetWindowCenter()

	playButton.CenterOn(centerX+100, centerY)
	stopButton.CenterOn(centerX-100, centerY)

	selectAudioAButton.CenterOn(centerX, centerY+100)
	selectAudioBButton.CenterOn(centerX, centerY+200)

	selectedMusicLabel.CenterOn(centerX, centerY-150)

	context.AddRenderable(&playButton)
	context.AddRenderable(&stopButton)
	context.AddRenderable(&selectAudioAButton)
	context.AddRenderable(&selectAudioBButton)
	context.AddRenderable(&selectedMusicLabel)

	selectAudioAButton.OnClick(func() {
		selectedMusicLabel.SetText(&context, "Selected Audio A")

		selectedAudio = &audioA
	})

	selectAudioBButton.OnClick(func() {
		selectedMusicLabel.SetText(&context, "Selected Audio B")

		selectedAudio = &audioB
	})

	playButton.OnClick(func() {
		if selectedAudio != nil {
			selectedAudio.Play()
		}
	})

	stopButton.OnClick(func() {
		if selectedAudio != nil {
			selectedAudio.Stop()
		}
	})

	playButton.AddListeners(&context)
	stopButton.AddListeners(&context)
	selectAudioAButton.AddListeners(&context)
	selectAudioBButton.AddListeners(&context)

	context.MainLoop()
}

func main() {
	game := woengine.NewGame()
	game.SetEntrypoint(gameLogic)
	game.Run()
}
