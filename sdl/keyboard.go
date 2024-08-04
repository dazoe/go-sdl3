package sdl

//#include <SDL3/SDL_keyboard.h>
import "C"
import "unsafe"

// This is a unique ID for a keyboard for the time it is connected to the
// system, and is never reused for the lifetime of the application.
// (https://wiki.libsdl.org/SDL3/SDL_KeyboardID)

type KeyboardID C.SDL_KeyboardID

/* Function prototypes */

// Return whether a keyboard is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasKeyboard)
func HasKeyboard() bool {
	return C.SDL_HasKeyboard() != 0
}

// Get a list of currently connected keyboards.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboards)
func GetKeyboards() ([]KeyboardID, error) {
	var count C.int
	ret := C.SDL_GetKeyboards(&count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((*KeyboardID)(ret), int(count)), nil
}

// Get the name of a keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID)
func (instance_id KeyboardID) GetName() (name string, err error) {
	name = C.GoString(C.SDL_GetKeyboardNameForID(C.SDL_KeyboardID(instance_id)))
	if len(name) == 0 {
		err = GetError()
	}
	return
}

// Query the window which currently has keyboard focus.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus)
func GetKeyboardFocus() *Window {
	return (*Window)(C.SDL_GetKeyboardFocus())
}

// Get a snapshot of the current state of the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState)
func GetKeyboardState() []uint8 {
	var numkeys C.int
	ret := C.SDL_GetKeyboardState(&numkeys)
	return unsafe.Slice((*uint8)(ret), int(numkeys))
}

// Clear the state of the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard)
func ResetKeyboard() {
	C.SDL_ResetKeyboard()
}

// Get the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetModState)
func GetModState() Keymod {
	return Keymod(C.SDL_GetModState())
}

// Set the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_SetModState)
func SetModState(modstate Keymod) {
	C.SDL_SetModState(C.SDL_Keymod(modstate))
}

// Get the key code corresponding to the given scancode according to a default
// en_US keyboard layout.
// (https://wiki.libsdl.org/SDL3/SDL_GetDefaultKeyFromScancode)
func GetDefaultKeyFromScancode(scancode Scancode, modstate Keymod) Keycode {
	return Keycode(C.SDL_GetDefaultKeyFromScancode(C.SDL_Scancode(scancode), C.SDL_Keymod(modstate)))
}

// Get the key code corresponding to the given scancode according to the
// current keyboard layout.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode)
func GetKeyFromScancode(scancode Scancode, modstate Keymod) Keycode {
	return Keycode(C.SDL_GetKeyFromScancode(C.SDL_Scancode(scancode), C.SDL_Keymod(modstate)))
}

// Get the scancode corresponding to the given key code according to a default
// en_US keyboard layout.
// (https://wiki.libsdl.org/SDL3/SDL_GetDefaultScancodeFromKey)
func GetDefaultScancodeFromKey(key Keycode, modstate *Keymod) Scancode {
	return Scancode(C.SDL_GetDefaultScancodeFromKey(C.SDL_Keycode(key), (*C.SDL_Keymod)(modstate)))
}

// Get the scancode corresponding to the given key code according to the
// current keyboard layout.
// (https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey)
func GetScancodeFromKey(key Keycode, modstate *Keymod) Scancode {
	return Scancode(C.SDL_GetScancodeFromKey(C.SDL_Keycode(key), (*C.SDL_Keymod)(modstate)))
}

// Set a human-readable name for a scancode.
// (https://wiki.libsdl.org/SDL3/SDL_SetScancodeName)
func SetScancodeName(scancode Scancode, name string) (err error) {
	ret := C.SDL_SetScancodeName(C.SDL_Scancode(scancode), C.CString(name))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get a human-readable name for a scancode.
// (https://wiki.libsdl.org/SDL3/SDL_GetScancodeName)
func GetScancodeName(scancode Scancode) string {
	return C.GoString(C.SDL_GetScancodeName(C.SDL_Scancode(scancode)))
}

// Get a scancode from a human-readable name.
// (https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName)
func GetScancodeFromName(name string) (code Scancode, err error) {
	code = Scancode(C.SDL_GetScancodeFromName(C.CString(name)))
	if code == SCANCODE_UNKNOWN {
		err = GetError()
	}
	return
}

// Get a human-readable name for a key.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyName)
func GetKeyName(key Keycode) string {
	return C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(key)))
}

// Get a key code from a human-readable name.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName)
func GetKeyFromName(name string) (code Keycode, err error) {
	code = Keycode(C.SDL_GetKeyFromName(C.CString(name)))
	if code == K_UNKNOWN {
		err = GetError()
	}
	return
}

// Start accepting Unicode text input events in a window.
// (https://wiki.libsdl.org/SDL3/SDL_StartTextInput)
func (window *Window) StartTextInput() (err error) {
	ret := C.SDL_StartTextInput(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Check whether or not Unicode text input events are enabled for a window.
// (https://wiki.libsdl.org/SDL3/SDL_TextInputActive)
func (window *Window) TextInputActive() bool {
	return C.SDL_TextInputActive(window.cptr()) != 0
}

// Stop receiving any text input events in a window.
// (https://wiki.libsdl.org/SDL3/SDL_StopTextInput)
func (window *Window) StopTextInput() (err error) {
	ret := C.SDL_StopTextInput(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Dismiss the composition window/IME without disabling the subsystem.
// (https://wiki.libsdl.org/SDL3/SDL_ClearComposition)
func (window *Window) ClearComposition() (err error) {
	ret := C.SDL_ClearComposition(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the area used to type Unicode text input.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea)
func (window *Window) SetTextInputArea(rect Rect, cursor int32) (err error) {
	ret := C.SDL_SetTextInputArea(window.cptr(), rect.cptr(), C.int(cursor))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the area used to type Unicode text input.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea)
func (window *Window) GetTextInputArea() (rect Rect, cursor int32, err error) {
	ret := C.SDL_GetTextInputArea(window.cptr(), rect.cptr(), (*C.int)(&cursor))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Check whether the platform has screen keyboard support.
// (https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport)
func HasScreenKeyboardSupport() bool {
	return C.SDL_HasScreenKeyboardSupport() != 0
}

// Check whether the screen keyboard is shown for given window.
// (https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown)
func (window *Window) ScreenKeyboardShown() bool {
	return C.SDL_ScreenKeyboardShown(window.cptr()) != 0
}
