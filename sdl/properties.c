#include "_cgo_export.h"

#include <SDL3/SDL_properties.h>
#include "properties.h"

void cgoCleanupPropertyCallback(void *userdata, void *value) {
    goCleanupPropertyCallback(userdata, value);
}

void cgoEnumeratePropertiesCallback(void *userdata, SDL_PropertiesID props, char *name) {
    goEnumeratePropertiesCallback(userdata, props, name);
}
