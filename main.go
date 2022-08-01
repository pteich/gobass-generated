package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"unsafe"

	"github.com/pteich/gobass-generated/bass"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	win := 0
	dsguid := 0
	ret := bass.BassInit(0, 44100, bass.BassDeviceNospeaker|bass.BassDeviceStereo, unsafe.Pointer(&win), unsafe.Pointer(&dsguid))
	if ret <= 0 {
		log.Fatal(bass.BassErrorgetcode())
	}

	defer bass.BassFree()

	hstream := bass.BassStreamcreateurl("Http://url-of-a-icecast-stream", 0, bass.BassStreamAutofree, nil, nil)
	if hstream == 0 {
		log.Fatal(bass.BassErrorgetcode())
	}

	var syncProc bass.Syncproc = func(handle uint32, channel uint32, data uint32, user unsafe.Pointer) {
		log.Print(bass.BassChannelgettags(channel, bass.BassTagMeta))
	}

	hsync := bass.BassChannelsetsync(hstream, bass.BassSyncMeta, 0, syncProc, nil)
	if hsync == 0 {
		log.Fatal(bass.BassErrorgetcode())
	}
	defer bass.BassStreamfree(hstream)

	hencode := bass.BassEncodeMp3Start(hstream, "-r -s 44100 -b 128", bass.BassEncodeAutofree, nil, nil)
	if hencode == 0 {
		log.Fatal(bass.BassErrorgetcode())
	}

	ret = bass.BassChannelplay(hstream, 1)
	if ret <= 0 {
		log.Fatal(bass.BassErrorgetcode())
	}

	ret = bass.BassEncodeCastinit(hencode,
		"icecast-server:8000/mount",
		"password",
		bass.BassEncodeTypeMp3,
		"Bass Test",
		"url",
		"",
		"",
		"",
		128,
		bass.BassEncodeCastPublic,
	)
	if ret <= 0 {
		log.Fatal(bass.BassErrorgetcode())
	}

	<-ctx.Done()
}
