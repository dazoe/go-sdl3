void cgoAudioStreamGetCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount);
void cgoAudioStreamPutCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount);
void cgoAudioStreamUnknownCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount);
void cgoAudioPostmixCallback(void *userdata, SDL_AudioSpec *spec, float *buffer, int buflen);

