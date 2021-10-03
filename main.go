package main

import (
	"unsafe"

	"github.com/pteich/gobass/bass"
)

func main()  {
	win := 0
	dsguid := 0
	bass.BassInit(-1, 44100, 0, unsafe.Pointer(&win), unsafe.Pointer(&dsguid))
}
