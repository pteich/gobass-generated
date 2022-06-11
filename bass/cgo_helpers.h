// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 11 Jun 2022 23:23:46 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

#include "../include/bass.h"
#include "../include/bassenc.h"
#include "../include/bassenc_mp3.h"
#include "../include/bassenc_flac.h"
#include "../include/bassenc_ogg.h"
#include "../include/bassmix.h"
#include <stdlib.h>
#pragma once

#define __CGOGEN 1

// STREAMPROC_c741188f is a proxy for callback STREAMPROC.
unsigned int STREAMPROC_c741188f(unsigned int handle, void* buffer, unsigned int length, void* user);

// FILECLOSEPROC_16713455 is a proxy for callback FILECLOSEPROC.
void FILECLOSEPROC_16713455(void* user);

// FILELENPROC_69f2a6d6 is a proxy for callback FILELENPROC.
unsigned long int FILELENPROC_69f2a6d6(void* user);

// FILEREADPROC_4d4b6dca is a proxy for callback FILEREADPROC.
unsigned int FILEREADPROC_4d4b6dca(void* buffer, unsigned int length, void* user);

// FILESEEKPROC_9820b893 is a proxy for callback FILESEEKPROC.
int FILESEEKPROC_9820b893(unsigned long int offset, void* user);

// DOWNLOADPROC_d99fa183 is a proxy for callback DOWNLOADPROC.
void DOWNLOADPROC_d99fa183(void* buffer, unsigned int length, void* user);

// SYNCPROC_1ec67831 is a proxy for callback SYNCPROC.
void SYNCPROC_1ec67831(unsigned int handle, unsigned int channel, unsigned int data, void* user);

// DSPPROC_a3208d1a is a proxy for callback DSPPROC.
void DSPPROC_a3208d1a(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, void* user);

// RECORDPROC_f0b593c is a proxy for callback RECORDPROC.
int RECORDPROC_f0b593c(unsigned int handle, void* buffer, unsigned int length, void* user);

// IOSNOTIFYPROC_c7638a61 is a proxy for callback IOSNOTIFYPROC.
void IOSNOTIFYPROC_c7638a61(unsigned int status);

// ENCODEPROC_a0c39546 is a proxy for callback ENCODEPROC.
void ENCODEPROC_a0c39546(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, void* user);

// ENCODEPROCEX_8f215ff2 is a proxy for callback ENCODEPROCEX.
void ENCODEPROCEX_8f215ff2(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, unsigned long int offset, void* user);

// ENCODERPROC_ca4a3c53 is a proxy for callback ENCODERPROC.
unsigned int ENCODERPROC_ca4a3c53(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, unsigned int maxout, void* user);

// ENCODECLIENTPROC_541d704 is a proxy for callback ENCODECLIENTPROC.
int ENCODECLIENTPROC_541d704(unsigned int handle, int connect, char* client, char* headers, void* user);

// ENCODENOTIFYPROC_fdef6823 is a proxy for callback ENCODENOTIFYPROC.
void ENCODENOTIFYPROC_fdef6823(unsigned int handle, unsigned int status, void* user);
