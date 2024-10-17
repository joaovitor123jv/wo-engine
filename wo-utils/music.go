package woutils

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
)

type Music struct {
	path    string
	file    *mix.Music
	canStop bool
}

// Use this for background tracks that will be paused or stopped frequently.
// This structure does not allow multiple instances of the same music to be played at the same time.
func NewMusic(path string) Music {
	var err error
	var music *mix.Music

	if music, err = mix.LoadMUS(path); err != nil {
		log.Fatalln(err)
	}

	return Music{
		path: path,
		file: music,
	}
}

func (a *Music) PlayOnce() {
	if err := a.file.Play(0); err != nil {
		log.Fatalln(err)
	}
}

func (a *Music) PlayForever() {
	if err := a.file.Play(-1); err != nil {
		log.Fatalln(err)
	}
}

func (a *Music) Loaded() bool {
	if a.file == nil {
		return false
	}
	return true
}

func (a *Music) Stop() {
	mix.HaltMusic()
}

func (a *Music) Destroy() {
	if a.file != nil {
		a.file.Free()
	}

}
