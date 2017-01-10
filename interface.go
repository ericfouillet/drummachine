package drummachine

import "io"

// Song represents a song
type Song struct {
	title string
	bpm   int
	beat  int // number of beats per measure
	bar   int // number of bars per beat
	drum  grid
}

// grid represents the grid of a song, as a slice of steps.
type grid struct {
	steps []step
}

// A step is a group of sounds.
type step struct {
	sounds []Instrument
}

// Instrument is an object that can play a sound to an output.
type Instrument interface {
	play(io.Writer)
}

// DrumSound is an Instrument that can play a drum sound.
type DrumSound struct {
	name  string
	sound string
}
