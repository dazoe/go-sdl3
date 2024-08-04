package sdl

/*
#include <SDL3/SDL_hints.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Set a hint with normal priority.
// (https://wiki.libsdl.org/SDL3/SDL_SetHint)
func SetHint(name, value string) error {
	cName := C.CString(name)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cName))
	defer C.free(unsafe.Pointer(cValue))
	ret := C.SDL_SetHint(cName, cValue)
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Get the value of a hint.
// (https://wiki.libsdl.org/SDL3/SDL_GetHint)
func GetHint(name string) (string, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.SDL_GetHint(cName)
	if cValue == nil {
		return "", GetError()
	}
	return C.GoString(cValue), nil
}

// TODO: incomplete
