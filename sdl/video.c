#include "_cgo_export.h"

#include <SDL3/SDL_video.h>
#include "video.h"

SDL_HitTestResult cgoHitTestCallback(SDL_Window *win, SDL_Point *area, void *data) {
    return goHitTestCallback(win, area);
}