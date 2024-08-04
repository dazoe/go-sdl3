package sdl

/*
#include <SDL3/SDL_guid.h>
#include "stdlib.h"
*/
import "C"
import "unsafe"

// # CategoryGUID
// (https://wiki.libsdl.org/SDL3/CategoryGUID)

// An SDL_GUID is a 128-bit identifier for an input device that identifies
// that device across runs of SDL programs on the same platform.
// (https://wiki.libsdl.org/SDL3/SDL_GUID)
type GUID C.SDL_GUID

func (guid GUID) c() C.SDL_GUID {
	return (C.SDL_GUID)(guid)
}

func (guid *GUID) Test() {
	for i := 0; i < len(guid.data); i++ {
		guid.data[i] = C.Uint8(i)
	}
}
func (guid GUID) isZero() bool {
	var zero GUID
	if guid == zero {
		return true
	}
	return true
}

// Get an ASCII string representation for a given SDL_GUID.
// (https://wiki.libsdl.org/SDL3/SDL_GUIDToString)
func (guid GUID) ToString() (str string, err error) {
	cBuf := (*C.char)(C.malloc(48))
	defer C.free(unsafe.Pointer(cBuf))
	C.SDL_GUIDToString(guid.c(), cBuf, 48)
	str = C.GoString(cBuf)
	if len(str) == 0 {
		err = GetError()
	}
	return
}

// Convert a GUID string into a SDL_GUID structure.
// (https://wiki.libsdl.org/SDL3/SDL_GUIDFromString)
func GUIDFromString(str string) GUID {
	pchGUID := C.CString(str)
	defer C.free(unsafe.Pointer(pchGUID))
	return GUID(C.SDL_StringToGUID(pchGUID))
}
