//go:build windows || linux

package bass

/*
#include "../include/bassenc_aac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// BassEncodeAacGetversion function as declared in include/bassenc_aac.h:25
func BassEncodeAacGetversion() uint32 {
	__ret := C.BASS_Encode_AAC_GetVersion()
	__v := (uint32)(__ret)
	return __v
}

// BassEncodeAacStart function as declared in include/bassenc_aac.h:27
func BassEncodeAacStart(handle uint32, options string, flags uint32, proc Encodeprocex, user unsafe.Pointer) Hencode {
	chandle, chandleAllocMap := (C.DWORD)(handle), cgoAllocsUnknown
	coptions, coptionsAllocMap := unpackPCharString(options)
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cproc, cprocAllocMap := proc.PassRef()
	cuser, cuserAllocMap := user, cgoAllocsUnknown
	__ret := C.BASS_Encode_AAC_Start(chandle, coptions, cflags, cproc, cuser)
	runtime.KeepAlive(cuserAllocMap)
	runtime.KeepAlive(cprocAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(coptionsAllocMap)
	runtime.KeepAlive(chandleAllocMap)
	__v := (Hencode)(__ret)
	return __v
}

// BassEncodeAacStartfile function as declared in include/bassenc_aac.h:28
func BassEncodeAacStartfile(handle uint32, options string, flags uint32, filename string) Hencode {
	chandle, chandleAllocMap := (C.DWORD)(handle), cgoAllocsUnknown
	coptions, coptionsAllocMap := unpackPCharString(options)
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cfilename, cfilenameAllocMap := unpackPCharString(filename)
	__ret := C.BASS_Encode_AAC_StartFile(chandle, coptions, cflags, cfilename)
	runtime.KeepAlive(cfilenameAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(coptionsAllocMap)
	runtime.KeepAlive(chandleAllocMap)
	__v := (Hencode)(__ret)
	return __v
}
