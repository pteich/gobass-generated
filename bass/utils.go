package bass

/*
#include "../include/bass.h"
#include "../include/bassenc.h"
#include "../include/bassenc_mp3.h"
#include "../include/bassenc_flac.h"
#include "../include/bassenc_ogg.h"
#include "../include/bassmix.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"unsafe"
)

// GetHighWord returns the high word from an integer
func GetHighWord(v uint32) uint16 {
	return uint16(v >> 16)
}

// GetLowWord returns the high word from an integer
func GetLowWord(v uint32) uint16 {
	return uint16(v & 0x0000FFFF)
	//	ptr := (*uint16)(unsafe.Pointer(&v))
	//	return *ptr
}

func ChannelSetGain(handle int, gain float32) int32 {
	var params C.BASS_FX_VOLUME_PARAM

	params.fTarget = C.float(gain)
	params.fCurrent = C.float(-1)
	params.fTime = C.float(0)
	params.lCurve = C.DWORD(0)

	if C.BASS_FXSetParameters(C.HFX(handle), unsafe.Pointer(&params)) != 1 {
		return int32(C.BASS_ErrorGetCode())
	}

	return 0
}
