package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/pteich/gobass-generated/bass"
)

func main() {
	win := 0
	dsguid := 0
	bass.BassInit(-1, 44100, 0, unsafe.Pointer(&win), unsafe.Pointer(&dsguid))

	hstream := bass.BassStreamcreateurl("https://stream.sunshine-live.de/sp8/mp3-192", 0, 0, nil, nil)
	if hstream == 0 {
		errCode := bass.BassErrorgetcode()
		fmt.Println(errCode)
		os.Exit(1)
	}
}
