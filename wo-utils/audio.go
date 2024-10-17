package woutils

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
)

type Audio struct {
	path             string
	file             *mix.Chunk
	isPlaying        bool
	playingOnChannel int
	preferredChannel int
}

// Use this for long songs that will not be paused or stopped frequently
func NewAudio(path string) Audio {
	var err error
	var audio *mix.Chunk

	if audio, err = mix.LoadWAV(path); err != nil {
		log.Fatalln(err)
	}

	return Audio{
		path:             path,
		file:             audio,
		isPlaying:        false,
		playingOnChannel: -1, // -1 means it's not playing
		preferredChannel: -1, // -1 means it can play in any free channel
	}
}

func NewUIAudio(path string) Audio {
	var err error
	var audio *mix.Chunk

	if audio, err = mix.LoadWAV(path); err != nil {
		log.Fatalln(err)
	}

	return Audio{
		path:             path,
		file:             audio,
		isPlaying:        false,
		playingOnChannel: -1, // -1 means it's not playing
		preferredChannel: UI_AUDIO_CHANNEL,
	}
}

func NewPlayerAudio(path string) Audio {
	audio := NewAudio(path)
	audio.preferredChannel = PLAYER_AUDIO_CHANNEL
	return audio
}

func (a *Audio) Play() {
	if channel, err := a.file.Play(a.preferredChannel, 0); err == nil {
		a.playingOnChannel = channel
		a.isPlaying = true
	} else {
		log.Println(err)
	}
}

func (a *Audio) PlayNTimes(n int) {
	if channel, err := a.file.Play(a.preferredChannel, n-1); err == nil {
		a.playingOnChannel = channel
		a.isPlaying = true
	} else {
		log.Println(err)
	}
}

func (a *Audio) Loaded() bool {
	if a.file == nil {
		return false
	}
	return true
}

func (a *Audio) Stop() {
	a.isPlaying = false
	mix.HaltChannel(a.playingOnChannel)
}

func (a *Audio) Destroy() {
	a.Stop()
	a.file.Free()
}
