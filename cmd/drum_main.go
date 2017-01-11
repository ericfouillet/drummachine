package main

import (
	"context"
	"os"
	"time"

	"github.com/ericfouillet/drummachine"
)

func main() {
	song := drummachine.NewSong("Sample song", 120, "4/4")
	song.AddInstrument([]int{0, 4, 8, 12}, drummachine.NewKick())
	song.AddInstrument([]int{4, 12}, drummachine.NewSnare())
	song.AddInstrument([]int{2, 6}, drummachine.NewHiHat())
	ctx, cancel := context.WithCancel(context.Background())
	out := &drummachine.StdDevice{Out: os.Stdout}
	go song.Play(ctx, out)
	time.Sleep(10 * time.Second)
	cancel()
}
