// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 28 Jul 2022 15:21:50 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

#include "_cgo_export.h"
#include "cgo_helpers.h"

unsigned int STREAMPROC_c741188f(unsigned int handle, void* buffer, unsigned int length, void* user) {
	return streamprocC741188F(handle, buffer, length, user);
}

void FILECLOSEPROC_16713455(void* user) {
	filecloseproc16713455(user);
}

unsigned long int FILELENPROC_69f2a6d6(void* user) {
	return filelenproc69F2A6D6(user);
}

unsigned int FILEREADPROC_4d4b6dca(void* buffer, unsigned int length, void* user) {
	return filereadproc4D4B6DCA(buffer, length, user);
}

int FILESEEKPROC_9820b893(unsigned long int offset, void* user) {
	return fileseekproc9820B893(offset, user);
}

void DOWNLOADPROC_d99fa183(void* buffer, unsigned int length, void* user) {
	downloadprocD99FA183(buffer, length, user);
}

void SYNCPROC_1ec67831(unsigned int handle, unsigned int channel, unsigned int data, void* user) {
	syncproc1EC67831(handle, channel, data, user);
}

void DSPPROC_a3208d1a(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, void* user) {
	dspprocA3208D1A(handle, channel, buffer, length, user);
}

int RECORDPROC_f0b593c(unsigned int handle, void* buffer, unsigned int length, void* user) {
	return recordprocF0B593C(handle, buffer, length, user);
}

void IOSNOTIFYPROC_c7638a61(unsigned int status) {
	iosnotifyprocC7638A61(status);
}

void ENCODEPROC_a0c39546(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, void* user) {
	encodeprocA0C39546(handle, channel, buffer, length, user);
}

void ENCODEPROCEX_8f215ff2(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, unsigned long int offset, void* user) {
	encodeprocex8F215FF2(handle, channel, buffer, length, offset, user);
}

unsigned int ENCODERPROC_ca4a3c53(unsigned int handle, unsigned int channel, void* buffer, unsigned int length, unsigned int maxout, void* user) {
	return encoderprocCA4A3C53(handle, channel, buffer, length, maxout, user);
}

int ENCODECLIENTPROC_541d704(unsigned int handle, int connect, char* client, char* headers, void* user) {
	return encodeclientproc541D704(handle, connect, client, headers, user);
}

void ENCODENOTIFYPROC_fdef6823(unsigned int handle, unsigned int status, void* user) {
	encodenotifyprocFDEF6823(handle, status, user);
}

