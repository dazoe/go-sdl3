package sdl

/*
#include <SDL3/SDL_clipboard.h>
#include "clipboard.h"
*/
import "C"
import "unsafe"

// # CategoryClipboard
// (https://wiki.libsdl.org/SDL3/CategoryClipboard)

// Put UTF-8 text into the clipboard.
// (https://wiki.libsdl.org/SDL3/SDL_SetClipboardText)
func SetClipboardText(text string) error {
	cstr := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(cstr))
	ret := C.SDL_SetClipboardText(cstr)
	if !ret {
		return GetError()
	}
	return nil
}

// Get UTF-8 text from the clipboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetClipboardText)
func GetClipboardText() (string, error) {
	ret := C.SDL_GetClipboardText()
	// Alwyas returns a string
	defer C.SDL_free(unsafe.Pointer(ret))
	if *ret == 0 {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Query whether the clipboard exists and contains a non-empty text string.
// (https://wiki.libsdl.org/SDL3/SDL_HasClipboardText)
func HasClipboardText() bool {
	return bool(C.SDL_HasClipboardText())
}

// Put UTF-8 text into the primary selection.
// (https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText)
func SetPrimarySelectionText(text string) error {
	cstr := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(cstr))
	ret := C.SDL_SetPrimarySelectionText(cstr)
	if !ret {
		return GetError()
	}
	return nil
}

// Get UTF-8 text from the primary selection.
// (https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText)
func GetPrimarySelectionText() (string, error) {
	ret := C.SDL_GetPrimarySelectionText()
	// Always returns a string
	defer C.SDL_free(unsafe.Pointer(ret))
	if *ret == 0 {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Query whether the primary selection exists and contains a non-empty text
// string.
// (https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText)
func HasPrimarySelectionText() bool {
	return bool(C.SDL_HasPrimarySelectionText())
}

// TODO: This needs finished
// Callback function that will be called when data for the specified mime-type
// is requested by the OS.
// (https://wiki.libsdl.org/SDL3/SDL_ClipboardDataCallback)
type ClipboardDataCallback func(userdata any, mime_type string, size int) []byte

//export goClipboardDataCallback
func goClipboardDataCallback(userdata unsafe.Pointer, mime_type *C.char, size *C.size_t) unsafe.Pointer {
	panic("not implemented")
}

// Callback function that will be called when the clipboard is cleared, or new
// data is set.
// (https://wiki.libsdl.org/SDL3/SDL_ClipboardCleanupCallback)
type ClipboardCleanupCallback func(userdata any)

//export goClipboardCleanupCallback
func goClipboardCleanupCallback(userdata unsafe.Pointer) {
	panic("not implemented")
}

// TODO: needs completed
// Offer clipboard data to the OS.
// (https://wiki.libsdl.org/SDL3/SDL_SetClipboardData)
func SetClipboardData(dataCB ClipboardDataCallback, cleanupCB ClipboardCleanupCallback, userdata any, mime_types []string) {
	panic("not implemented")
}

// Clear the clipboard data.
// (https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData)
func ClearClipboardData() error {
	ret := C.SDL_ClearClipboardData()
	if !ret {
		return GetError()
	}
	return nil
}

// TODO: Test me
// Get the data from clipboard for a given mime type.
// (https://wiki.libsdl.org/SDL3/SDL_GetClipboardData)
func GetClipboardData(mimeType string) ([]byte, error) {
	var dataLen C.size_t
	cstr := C.CString(mimeType)
	defer C.SDL_free(unsafe.Pointer(cstr))
	ret := C.SDL_GetClipboardData(cstr, &dataLen)
	if ret == nil {
		return nil, GetError()
	}
	// ret needs to be freed so we need to make a copy
	defer C.SDL_free(ret)
	cSlice := unsafe.Slice((*byte)(ret), dataLen)
	goSlice := make([]byte, dataLen)
	copy(goSlice, cSlice)
	return goSlice, nil
}

// Query whether there is data in the clipboard for the provided mime type.
// (https://wiki.libsdl.org/SDL3/SDL_HasClipboardData)
func HasClipboardData(mimeType string) bool {
	cstr := C.CString(mimeType)
	defer C.SDL_free(unsafe.Pointer(cstr))
	return bool(C.SDL_HasClipboardData(cstr))
}

// Retrieve the list of mime types available in the clipboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes)
func GetClipboardMimeTypes() ([]string, error) {
	var count C.size_t
	ret := C.SDL_GetClipboardMimeTypes(&count)
	if ret == nil {
		return nil, GetError()
	}
	// ret needs to be freed so we need to make a copy
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice(ret, count)
	mimeTypes := make([]string, count)
	for i := 0; i < int(count); i++ {
		mimeTypes[i] = C.GoString(cSlice[i])
	}
	return mimeTypes, nil
}
