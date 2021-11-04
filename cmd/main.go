package main

import (
	"go-cmdpb/progressbar"
	"time"
)

func main() {
	b := progressbar.NewCmdProgressBar()
	b.SetLength(30)
	step := (b.Max() - b.Min()) * 0.1
	for b.Value() < b.Max() {
		b.SetValue(b.Value() + step)
		b.Print("")
		time.Sleep(time.Second * 1)
	}
}
