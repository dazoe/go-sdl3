package sdl

/*
#include <SDL3/SDL_error.h>
#include <stdlib.h>

void GoSetError(const char *str) {
	SDL_SetError("%s", str);
}
*/
import "C"
import (
	"fmt"
	"io"
	"unsafe"
)

// # CategoryError
// (https://wiki.libsdl.org/SDL3/CategoryError)

type errorSDL string

func (e errorSDL) Error() string {
	return string(e)
}

// Set the SDL error message for the current thread.
// (https://wiki.libsdl.org/SDL3/SDL_SetError)
func SetError(format string, a ...any) {
	// ClearError()
	errStr := fmt.Sprintf(format, a...)
	cstr := C.CString(errStr)
	defer C.free(unsafe.Pointer(cstr))
	C.GoSetError(cstr)
}

// Set an error indicating that memory allocation failed.
// (https://wiki.libsdl.org/SDL3/SDL_OutOfMemory)
func OutOfMemory() {
	C.SDL_OutOfMemory()
}

// Retrieve a message about the last error that occurred on the current thread.
// (https://wiki.libsdl.org/SDL3/SDL_GetError)
func GetError() error {
	ret := C.SDL_GetError()
	// always returns a string
	if *ret != 0 {
		return errorSDL(C.GoString(ret))
	}
	return nil
}

// Clear any previous error message for this thread.
// (https://wiki.libsdl.org/SDL3/SDL_ClearError)
func ClearError() {
	C.SDL_ClearError()
}

// Internal sdl error functions...
func Unsupported() error {
	return errorSDL("That operation is not supported")
}

func InvalidParamError(param string) error {
	return errorSDL(fmt.Sprintf("Parameter '%s' is invalid", param))
}

var ErrorEOF = io.EOF
