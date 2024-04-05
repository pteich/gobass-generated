//go:build darwin

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
	"runtime"
	"unsafe"
)

// BassEncodeGetcaref function as declared in include/bassenc.h:181
func BassEncodeGetcaref(handle uint32) unsafe.Pointer {
	chandle, chandleAllocMap := (C.DWORD)(handle), cgoAllocsUnknown
	__ret := C.BASS_Encode_GetCARef(chandle)
	runtime.KeepAlive(chandleAllocMap)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// BassEncodeStartca function as declared in include/bassenc.h:179
func BassEncodeStartca(handle uint32, ftype uint32, atype uint32, flags uint32, bitrate uint32, proc Encodeprocex, user unsafe.Pointer) Hencode {
	chandle, chandleAllocMap := (C.DWORD)(handle), cgoAllocsUnknown
	cftype, cftypeAllocMap := (C.DWORD)(ftype), cgoAllocsUnknown
	catype, catypeAllocMap := (C.DWORD)(atype), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cbitrate, cbitrateAllocMap := (C.DWORD)(bitrate), cgoAllocsUnknown
	cproc, cprocAllocMap := proc.PassRef()
	cuser, cuserAllocMap := user, cgoAllocsUnknown
	__ret := C.BASS_Encode_StartCA(chandle, cftype, catype, cflags, cbitrate, cproc, cuser)
	runtime.KeepAlive(cuserAllocMap)
	runtime.KeepAlive(cprocAllocMap)
	runtime.KeepAlive(cbitrateAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(catypeAllocMap)
	runtime.KeepAlive(cftypeAllocMap)
	runtime.KeepAlive(chandleAllocMap)
	__v := (Hencode)(__ret)
	return __v
}

// BassEncodeStartcafile function as declared in include/bassenc.h:180
func BassEncodeStartcafile(handle uint32, ftype uint32, atype uint32, flags uint32, bitrate uint32, filename string) Hencode {
	chandle, chandleAllocMap := (C.DWORD)(handle), cgoAllocsUnknown
	cftype, cftypeAllocMap := (C.DWORD)(ftype), cgoAllocsUnknown
	catype, catypeAllocMap := (C.DWORD)(atype), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cbitrate, cbitrateAllocMap := (C.DWORD)(bitrate), cgoAllocsUnknown
	cfilename, cfilenameAllocMap := unpackPCharString(filename)
	__ret := C.BASS_Encode_StartCAFile(chandle, cftype, catype, cflags, cbitrate, cfilename)
	runtime.KeepAlive(cfilenameAllocMap)
	runtime.KeepAlive(cbitrateAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(catypeAllocMap)
	runtime.KeepAlive(cftypeAllocMap)
	runtime.KeepAlive(chandleAllocMap)
	__v := (Hencode)(__ret)
	return __v
}
