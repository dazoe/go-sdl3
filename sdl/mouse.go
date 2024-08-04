package sdl

//#include <SDL3/SDL_mouse.h>
import "C"
import "unsafe"

// # CategoryMouse
// (https://wiki.libsdl.org/SDL3/CategoryMouse)

type MouseID C.SDL_MouseID
type Cursor C.SDL_Cursor

func (c *Cursor) cptr() *C.SDL_Cursor {
	return (*C.SDL_Cursor)(c)
}

// Cursor types for SDL_CreateSystemCursor().
// (https://wiki.libsdl.org/SDL3/SDL_SystemCursor)
type SystemCursor C.SDL_SystemCursor

const (
	SYSTEM_CURSOR_DEFAULT     = SystemCursor(C.SDL_SYSTEM_CURSOR_DEFAULT)     /**< Default cursor. Usually an arrow. */
	SYSTEM_CURSOR_TEXT        = SystemCursor(C.SDL_SYSTEM_CURSOR_TEXT)        /**< Text selection. Usually an I-beam. */
	SYSTEM_CURSOR_WAIT        = SystemCursor(C.SDL_SYSTEM_CURSOR_WAIT)        /**< Wait. Usually an hourglass or watch or spinning ball. */
	SYSTEM_CURSOR_CROSSHAIR   = SystemCursor(C.SDL_SYSTEM_CURSOR_CROSSHAIR)   /**< Crosshair. */
	SYSTEM_CURSOR_PROGRESS    = SystemCursor(C.SDL_SYSTEM_CURSOR_PROGRESS)    /**< Program is busy but still interactive. Usually it's WAIT with an arrow. */
	SYSTEM_CURSOR_NWSE_RESIZE = SystemCursor(C.SDL_SYSTEM_CURSOR_NWSE_RESIZE) /**< Double arrow pointing northwest and southeast. */
	SYSTEM_CURSOR_NESW_RESIZE = SystemCursor(C.SDL_SYSTEM_CURSOR_NESW_RESIZE) /**< Double arrow pointing northeast and southwest. */
	SYSTEM_CURSOR_EW_RESIZE   = SystemCursor(C.SDL_SYSTEM_CURSOR_EW_RESIZE)   /**< Double arrow pointing west and east. */
	SYSTEM_CURSOR_NS_RESIZE   = SystemCursor(C.SDL_SYSTEM_CURSOR_NS_RESIZE)   /**< Double arrow pointing north and south. */
	SYSTEM_CURSOR_MOVE        = SystemCursor(C.SDL_SYSTEM_CURSOR_MOVE)        /**< Four pointed arrow pointing north, south, east, and west. */
	SYSTEM_CURSOR_NOT_ALLOWED = SystemCursor(C.SDL_SYSTEM_CURSOR_NOT_ALLOWED) /**< Not permitted. Usually a slashed circle or crossbones. */
	SYSTEM_CURSOR_POINTER     = SystemCursor(C.SDL_SYSTEM_CURSOR_POINTER)     /**< Pointer that indicates a link. Usually a pointing hand. */
	SYSTEM_CURSOR_NW_RESIZE   = SystemCursor(C.SDL_SYSTEM_CURSOR_NW_RESIZE)   /**< Window resize top-left. This may be a single arrow or a double arrow like NWSE_RESIZE. */
	SYSTEM_CURSOR_N_RESIZE    = SystemCursor(C.SDL_SYSTEM_CURSOR_N_RESIZE)    /**< Window resize top. May be NS_RESIZE. */
	SYSTEM_CURSOR_NE_RESIZE   = SystemCursor(C.SDL_SYSTEM_CURSOR_NE_RESIZE)   /**< Window resize top-right. May be NESW_RESIZE. */
	SYSTEM_CURSOR_E_RESIZE    = SystemCursor(C.SDL_SYSTEM_CURSOR_E_RESIZE)    /**< Window resize right. May be EW_RESIZE. */
	SYSTEM_CURSOR_SE_RESIZE   = SystemCursor(C.SDL_SYSTEM_CURSOR_SE_RESIZE)   /**< Window resize bottom-right. May be NWSE_RESIZE. */
	SYSTEM_CURSOR_S_RESIZE    = SystemCursor(C.SDL_SYSTEM_CURSOR_S_RESIZE)    /**< Window resize bottom. May be NS_RESIZE. */
	SYSTEM_CURSOR_SW_RESIZE   = SystemCursor(C.SDL_SYSTEM_CURSOR_SW_RESIZE)   /**< Window resize bottom-left. May be NESW_RESIZE. */
	SYSTEM_CURSOR_W_RESIZE    = SystemCursor(C.SDL_SYSTEM_CURSOR_W_RESIZE)    /**< Window resize left. May be EW_RESIZE. */
	NUM_SYSTEM_CURSORS        = SystemCursor(C.SDL_NUM_SYSTEM_CURSORS)
)

// Scroll direction types for the Scroll event
// (https://wiki.libsdl.org/SDL3/SDL_MouseWheelDirection)
type MouseWheelDirection C.SDL_MouseWheelDirection

const (
	MOUSEWHEEL_NORMAL  = C.SDL_MOUSEWHEEL_NORMAL  /**< The scroll direction is normal */
	MOUSEWHEEL_FLIPPED = C.SDL_MOUSEWHEEL_FLIPPED /**< The scroll direction is flipped / natural */
)

// A bitmask of pressed mouse buttons, as reported by SDL_GetMouseState, etc.
// (https://wiki.libsdl.org/SDL3/SDL_MouseWheelDirection)
type MouseButtonFlags C.SDL_MouseButtonFlags

const (
	BUTTON_LEFT   = MouseButtonFlags(C.SDL_BUTTON_LEFT)
	BUTTON_MIDDLE = MouseButtonFlags(C.SDL_BUTTON_MIDDLE)
	BUTTON_RIGHT  = MouseButtonFlags(C.SDL_BUTTON_RIGHT)
	BUTTON_X1     = MouseButtonFlags(C.SDL_BUTTON_X1)
	BUTTON_X2     = MouseButtonFlags(C.SDL_BUTTON_X2)

	BUTTON_LMASK  = MouseButtonFlags(C.SDL_BUTTON_LMASK)
	BUTTON_MMASK  = MouseButtonFlags(C.SDL_BUTTON_MMASK)
	BUTTON_RMASK  = MouseButtonFlags(C.SDL_BUTTON_RMASK)
	BUTTON_X1MASK = MouseButtonFlags(C.SDL_BUTTON_X1MASK)
	BUTTON_X2MASK = MouseButtonFlags(C.SDL_BUTTON_X2MASK)
)

/* Function prototypes */

// Return whether a mouse is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasMouse)
func HasMouse() bool {
	return C.SDL_HasMouse() != 0
}

// Get a list of currently connected mice.
// (https://wiki.libsdl.org/SDL3/SDL_GetMice)
func GetMice() ([]MouseID, error) {
	var count C.int
	ret := C.SDL_GetMice(&count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((*MouseID)(ret), int(count)), nil
}

// Get the name of a mouse.
// (https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID)
func (instance_id MouseID) GetName() (name string, err error) {
	name = C.GoString(C.SDL_GetMouseNameForID(C.SDL_MouseID(instance_id)))
	if len(name) == 0 {
		err = GetError()
	}
	return
}

// Get the window which currently has mouse focus.
// (https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus)
func GetMouseFocus() *Window {
	return (*Window)(C.SDL_GetMouseFocus())
}

// Retrieve the current state of the mouse.
// (https://wiki.libsdl.org/SDL3/SDL_GetMouseState)
func GetMouseState() (x, y float32, flags MouseButtonFlags) {
	flags = MouseButtonFlags(C.SDL_GetMouseState((*C.float)(&x), (*C.float)(&y)))
	return
}

// Get the current state of the mouse in relation to the desktop.
// (https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState)
func GetGlobalMouseState() (x, y float32, flags MouseButtonFlags) {
	flags = MouseButtonFlags(C.SDL_GetGlobalMouseState((*C.float)(&x), (*C.float)(&y)))
	return
}

// Retrieve the relative state of the mouse.
// (https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState)
func GetRelativeMouseState() (x, y float32, flags MouseButtonFlags) {
	flags = MouseButtonFlags(C.SDL_GetRelativeMouseState((*C.float)(&x), (*C.float)(&y)))
	return
}

// Move the mouse cursor to the given position within the window.
// (https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow)
func (window *Window) WarpMouseInWindow(x, y float32) {
	C.SDL_WarpMouseInWindow(window.cptr(), C.float(x), C.float(y))
}

// Move the mouse to the given position in global screen space.
// (https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal)
func WarpMouseGlobal(x, y float32) (err error) {
	ret := C.SDL_WarpMouseGlobal(C.float(x), C.float(y))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set relative mouse mode.
// (https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseMode)
func SetRelativeMouseMode(enabled bool) (err error) {
	ret := C.SDL_SetRelativeMouseMode(sdlBool(enabled))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Capture the mouse and to track input outside an SDL window.
// (https://wiki.libsdl.org/SDL3/SDL_CaptureMouse)
func CaptureMouse(enabled bool) (err error) {
	ret := C.SDL_CaptureMouse(sdlBool(enabled))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Query whether relative mouse mode is enabled.
// (https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseMode)
func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() != 0
}

// Create a cursor using the specified bitmap data and mask (in MSB format).
// (https://wiki.libsdl.org/SDL3/SDL_CreateCursor)
func CreateCursor(data unsafe.Pointer, mask unsafe.Pointer, w, h, hot_x, hot_y int32) (cursor *Cursor, err error) {
	cursor = (*Cursor)(C.SDL_CreateCursor(
		(*C.Uint8)(data), (*C.Uint8)(mask), C.int(w), C.int(h), C.int(hot_x), C.int(hot_y)))
	if cursor == nil {
		err = GetError()
	}
	return
}

// Create a color cursor.
// (https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor)
func CreateColorCursor(surface *Surface, hot_x, hot_y int) (cursor *Cursor, err error) {
	cursor = (*Cursor)(C.SDL_CreateColorCursor(surface.cptr(), C.int(hot_x), C.int(hot_y)))
	if cursor == nil {
		err = GetError()
	}
	return
}

// Create a system cursor.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor)
func CreateSystemCursor(id SystemCursor) (cursor *Cursor, err error) {
	cursor = (*Cursor)(C.SDL_CreateSystemCursor(C.SDL_SystemCursor(id)))
	if cursor == nil {
		err = GetError()
	}
	return
}

// Set the active cursor.
// (https://wiki.libsdl.org/SDL3/SDL_SetCursor)
func SetCursor(cursor *Cursor) (err error) {
	ret := C.SDL_SetCursor(cursor.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the active cursor.
// (https://wiki.libsdl.org/SDL3/SDL_GetCursor)
func GetCursor() *Cursor {
	return (*Cursor)(C.SDL_GetCursor())
}

// Get the default cursor.
// (https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor)
func GetDefaultCursor() (cursor *Cursor, err error) {
	cursor = (*Cursor)(C.SDL_GetDefaultCursor())
	if cursor == nil {
		err = GetError()
	}
	return
}

// Free a previously-created cursor.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyCursor)
func DestroyCursor(cursor *Cursor) {
	C.SDL_DestroyCursor(cursor.cptr())
}

// Show the cursor.
// (https://wiki.libsdl.org/SDL3/SDL_ShowCursor)
func ShowCursor() (err error) {
	ret := C.SDL_ShowCursor()
	if ret != 0 {
		err = GetError()
	}
	return
}

// Hide the cursor.
// (https://wiki.libsdl.org/SDL3/SDL_HideCursor)
func HideCursor() (err error) {
	ret := C.SDL_HideCursor()
	if ret != 0 {
		err = GetError()
	}
	return
}

// Return whether the cursor is currently being shown.
// (https://wiki.libsdl.org/SDL3/SDL_CursorVisible)
func CursorVisible() bool {
	return C.SDL_CursorVisible() != 0
}
