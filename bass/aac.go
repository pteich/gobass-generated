//go:build windows || linux

package bass

/*
#include "../include/bassenc_aac.h"
#include "../include/bass_aac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// BassErrorMp4Nostream as defined in include/bass_aac.h:15
const BassErrorMp4Nostream = 6000

// BassConfigMp4Video as defined in include/bass_aac.h:18
const BassConfigMp4Video = 0x10700

// BassConfigAacMp4 as defined in include/bass_aac.h:19
const BassConfigAacMp4 = 0x10701

// BassConfigAacPrescan as defined in include/bass_aac.h:20
const BassConfigAacPrescan = 0x10702

// BassAacFrame960 as defined in include/bass_aac.h:23
const BassAacFrame960 = 0x1000

// BassAacStereo as defined in include/bass_aac.h:24
const BassAacStereo = 0x400000

// BassCtypeStreamAac as defined in include/bass_aac.h:27
const BassCtypeStreamAac = 0x10b00

// BassCtypeStreamMp4 as defined in include/bass_aac.h:28
const BassCtypeStreamMp4 = 0x10b01

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

// BassAacStreamcreatefile function as declared in include/bass_aac.h:31
func BassAacStreamcreatefile(mem int32, file unsafe.Pointer, offset uint64, length uint64, flags uint32) Hstream {
	cmem, cmemAllocMap := (C.BOOL)(mem), cgoAllocsUnknown
	cfile, cfileAllocMap := file, cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.QWORD)(offset), cgoAllocsUnknown
	clength, clengthAllocMap := (C.QWORD)(length), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	__ret := C.BASS_AAC_StreamCreateFile(cmem, cfile, coffset, clength, cflags)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(clengthAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(cfileAllocMap)
	runtime.KeepAlive(cmemAllocMap)
	__v := (Hstream)(__ret)
	return __v
}

// BassAacStreamcreateurl function as declared in include/bass_aac.h:32
func BassAacStreamcreateurl(url string, offset uint32, flags uint32, proc Downloadproc, user unsafe.Pointer) Hstream {
	curl, curlAllocMap := unpackPCharString(url)
	coffset, coffsetAllocMap := (C.DWORD)(offset), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cproc, cprocAllocMap := proc.PassRef()
	cuser, cuserAllocMap := user, cgoAllocsUnknown
	__ret := C.BASS_AAC_StreamCreateURL(curl, coffset, cflags, cproc, cuser)
	runtime.KeepAlive(cuserAllocMap)
	runtime.KeepAlive(cprocAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(curlAllocMap)
	__v := (Hstream)(__ret)
	return __v
}

// BassAacStreamcreatefileuser function as declared in include/bass_aac.h:33
func BassAacStreamcreatefileuser(system uint32, flags uint32, procs *BassFileprocs, user unsafe.Pointer) Hstream {
	csystem, csystemAllocMap := (C.DWORD)(system), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cprocs, cprocsAllocMap := procs.PassRef()
	cuser, cuserAllocMap := user, cgoAllocsUnknown
	__ret := C.BASS_AAC_StreamCreateFileUser(csystem, cflags, cprocs, cuser)
	runtime.KeepAlive(cuserAllocMap)
	runtime.KeepAlive(cprocsAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(csystemAllocMap)
	__v := (Hstream)(__ret)
	return __v
}

// BassMp4Streamcreatefile function as declared in include/bass_aac.h:34
func BassMp4Streamcreatefile(mem int32, file unsafe.Pointer, offset uint64, length uint64, flags uint32) Hstream {
	cmem, cmemAllocMap := (C.BOOL)(mem), cgoAllocsUnknown
	cfile, cfileAllocMap := file, cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.QWORD)(offset), cgoAllocsUnknown
	clength, clengthAllocMap := (C.QWORD)(length), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	__ret := C.BASS_MP4_StreamCreateFile(cmem, cfile, coffset, clength, cflags)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(clengthAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(cfileAllocMap)
	runtime.KeepAlive(cmemAllocMap)
	__v := (Hstream)(__ret)
	return __v
}

// BassMp4Streamcreatefileuser function as declared in include/bass_aac.h:35
func BassMp4Streamcreatefileuser(system uint32, flags uint32, procs *BassFileprocs, user unsafe.Pointer) Hstream {
	csystem, csystemAllocMap := (C.DWORD)(system), cgoAllocsUnknown
	cflags, cflagsAllocMap := (C.DWORD)(flags), cgoAllocsUnknown
	cprocs, cprocsAllocMap := procs.PassRef()
	cuser, cuserAllocMap := user, cgoAllocsUnknown
	__ret := C.BASS_MP4_StreamCreateFileUser(csystem, cflags, cprocs, cuser)
	runtime.KeepAlive(cuserAllocMap)
	runtime.KeepAlive(cprocsAllocMap)
	runtime.KeepAlive(cflagsAllocMap)
	runtime.KeepAlive(csystemAllocMap)
	__v := (Hstream)(__ret)
	return __v
}
