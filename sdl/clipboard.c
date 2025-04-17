#include "_cgo_export.h"
#include <SDL3/SDL_clipboard.h>
#include "clipboard.h"

//TODO: need to figure this out....
void *cgoClipboardDataCallback(void *userdata, char *mime_type, size_t *size) {
  goClipboardDataCallback(userdata, mime_type, size);
  return NULL;
}
void cgoClipboardCleanupCallback(void *userdata) {
  goClipboardCleanupCallback(userdata);
}
