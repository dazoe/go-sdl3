package sdl

//#include <SDL3/SDL_messagebox.h>
import "C"

type MessageBoxFlags C.SDL_MessageBoxFlags

const (
	MESSAGEBOX_ERROR                 = MessageBoxFlags(C.SDL_MESSAGEBOX_ERROR)                 /**< error dialog */
	MESSAGEBOX_WARNING               = MessageBoxFlags(C.SDL_MESSAGEBOX_WARNING)               /**< warning dialog */
	MESSAGEBOX_INFORMATION           = MessageBoxFlags(C.SDL_MESSAGEBOX_INFORMATION)           /**< informational dialog */
	MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT = MessageBoxFlags(C.SDL_MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT) /**< buttons placed left to right */
	MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT = MessageBoxFlags(C.SDL_MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT) /**< buttons placed right to left */
)

// ShowSimpleMessageBox Display a simple modal message box
// (https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox)
func ShowSimpleMessageBox(flags MessageBoxFlags, title string, message string, window *Window) error {
	if !C.SDL_ShowSimpleMessageBox(C.SDL_MessageBoxFlags(flags), C.CString(title),
		C.CString(message), nil) {
		return GetError()
	}
	return nil
}

// TODO: incomplete
