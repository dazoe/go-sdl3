package sdl

/*
#include <SDL3/SDL_dialog.h>
#include "stdlib.h"
*/
import "C"
import "unsafe"

// # CategoryDialog
// (https://wiki.libsdl.org/SDL3/CategoryDialog)

// An entry for filters for file dialogs.
// (https://wiki.libsdl.org/SDL3/SDL_DialogFileFilter)
type cDialogFileFilter struct {
	name    *C.char
	pattern *C.char
}

func (dff *cDialogFileFilter) Free() {
	C.free(unsafe.Pointer(dff.name))
	C.free(unsafe.Pointer(dff.pattern))
}

type DialogFileFilter struct {
	Name    string
	Pattern string
}

func (dff *DialogFileFilter) toC() *cDialogFileFilter {
	return &cDialogFileFilter{
		name:    C.CString(dff.Name),
		pattern: C.CString(dff.Pattern),
	}
}

// Callback used by file dialog functions.
// (https://wiki.libsdl.org/SDL3/SDL_DialogFileCallback)
type DialogFileCallback func(userdata any, filelist []string, filter int)

//  typedef void (SDLCALL *SDL_DialogFileCallback)(void *userdata, const char * const *filelist, int filter);

// Displays a dialog that lets the user select a file on their filesystem.
// (https://wiki.libsdl.org/SDL3/SDL_DialogFile)
func ShowOpenFileDialog(callback DialogFileCallback, userdata any, window *Window, filters []DialogFileFilter, default_location string, allow_many bool) {
	panic("not implemented")
}

//extern SDL_DECLSPEC void SDLCALL SDL_ShowOpenFileDialog(SDL_DialogFileCallback callback, void *userdata, SDL_Window *window, const SDL_DialogFileFilter *filters, int nfilters, const char *default_location, bool allow_many);

// Displays a dialog that lets the user choose a new or existing file on their
// filesystem.
// (https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog)
func ShowSaveFileDialog(callback DialogFileCallback, userdata any, window *Window, filters []DialogFileFilter, default_location string) {
	panic("not implemented")
}

//extern SDL_DECLSPEC void SDLCALL SDL_ShowSaveFileDialog(SDL_DialogFileCallback callback, void *userdata, SDL_Window *window, const SDL_DialogFileFilter *filters, int nfilters, const char *default_location);

// Displays a dialog that lets the user select a folder on their filesystem.
// (https://wiki.libsdl.org/SDL3/SDL_ShowSelectFolderDialog)
func ShowSelectFolderDialog(callback DialogFileCallback, userdata any, window *Window, default_location string, allow_many bool) {
	panic("not implemented")
}

//extern SDL_DECLSPEC void SDLCALL SDL_ShowOpenFolderDialog(SDL_DialogFileCallback callback, void *userdata, SDL_Window *window, const char *default_location, bool allow_many);
