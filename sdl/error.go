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

type errorSDL string

func (e errorSDL) Error() string {
	return string(e)
}

// SetError Set the SDL error message
// (https://wiki.libsdl.org/SDL3/SDL_SetError)
func SetError(err error) {
	ClearError()
	if err != nil {
		errMsg := C.CString(err.Error())
		defer C.free(unsafe.Pointer(errMsg))
		C.GoSetError(errMsg)
	}
}

// GetError Retrieve a message about the last error that occurred
// (https://wiki.libsdl.org/SDL3/SDL_GetError)
func GetError() error {
	errStr := C.SDL_GetError()
	if errStr != nil {
		goStr := C.GoString(errStr)
		if len(goStr) > 0 {
			return errorSDL(goStr)
		}
	}
	return nil
}

// ClearError Clear any previous error message
// (https://wiki.libsdl.org/SDL3/SDL_ClearError)
func ClearError() {
	C.SDL_ClearError()
}

func Unsupported() error {
	return errorSDL("That operation is not supported")
}

func InvalidParamError(param string) error {
	return errorSDL(fmt.Sprintf("Parameter '%s' is invalid", param))
}

var ErrorEOF = io.EOF
