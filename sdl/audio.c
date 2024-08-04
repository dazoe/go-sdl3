#include "_cgo_export.h"

#include <SDL3/SDL_audio.h>
#include "audio.h"

void cgoAudioStreamGetCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount) {
	goAudioStreamGetCallback(stream, additional_amount, total_amount);
}

void cgoAudioStreamPutCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount) {
	goAudioStreamPutCallback(stream, additional_amount, total_amount);
}

void cgoAudioStreamUnknownCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount) {
	goAudioStreamUnknownCallback(stream, additional_amount, total_amount);
}

void cgoAudioPostmixCallback(void *userdata, SDL_AudioSpec *spec, float *buffer, int buflen) {
	goAudioPostmixCallback(userdata, spec, buffer, buflen);
}
