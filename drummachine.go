package drummachine

import (
	"fmt"
	"io"
)

func newGrid(beats, sub int) grid {
	return grid{
		steps: make([]step, beats*sub),
	}
}

func (d *DrumSound) play(out io.Writer) {
	fmt.Fprint(out, d.sound)
}

// NewKick creates a kick drum instrument
func NewKick() Instrument {
	return &DrumSound{
		name:  "Kick",
		sound: "Kick",
	}
}

// NewSnare creates a snare drum instrument
func NewSnare() Instrument {
	return &DrumSound{
		name:  "Snare",
		sound: "Snare",
	}
}

// NewHiHat creates a hi-hat drum instrument
func NewHiHat() Instrument {
	return &DrumSound{
		name:  "HiHat",
		sound: "HiHat",
	}
}
