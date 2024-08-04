#include "_cgo_export.h"

#include <SDL3/SDL_events.h>
#include "events.h"

int cgoEventFilter(void *userdata, SDL_Event *event) {
	return goEventFilter(event);
}

int cgoEventWatcher(void *userdata, SDL_Event *event) {
	return goEventWatcher(userdata, event);
}

int cgoEventFilterOnce(void *userdata, SDL_Event *event) {
	return goEventFilterOnce(event);
}