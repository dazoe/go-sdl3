package sdl

/*
#include <SDL3/SDL_video.h>
#include "video.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/**
 * # CategoryVideo
 *
 * SDL video functions.
 *
 * (https://wiki.libsdl.org/SDL3/CategoryVideo)
 */

type DisplayID C.SDL_DisplayID

func (id DisplayID) c() C.SDL_DisplayID {
	return C.SDL_DisplayID(id)
}
func (id *DisplayID) cptr() *C.SDL_DisplayID {
	return (*C.SDL_DisplayID)(unsafe.Pointer(id))
}

type WindowID C.SDL_WindowID

func (id *WindowID) cptr() *C.SDL_WindowID {
	return (*C.SDL_WindowID)(unsafe.Pointer(id))
}

// The pointer to the global `wl_display` object used by the Wayland video backend
// (https://wiki.libsdl.org/SDL3/SDL_PROP_GLOBAL_VIDEO_WAYLAND_WL_DISPLAY_POINTER)
const PROP_GLOBAL_VIDEO_WAYLAND_WL_DISPLAY_POINTER = C.SDL_PROP_GLOBAL_VIDEO_WAYLAND_WL_DISPLAY_POINTER

// System theme
// (https://wiki.libsdl.org/SDL3/SDL_SystemTheme)
type SystemTheme C.SDL_SystemTheme

const (
	SYSTEM_THEME_UNKNOWN = SystemTheme(C.SDL_SYSTEM_THEME_UNKNOWN) /**< Unknown system theme */
	SYSTEM_THEME_LIGHT   = SystemTheme(C.SDL_SYSTEM_THEME_LIGHT)   /**< Light colored system theme */
	SYSTEM_THEME_DARK    = SystemTheme(C.SDL_SYSTEM_THEME_DARK)    /**< Dark colored system theme */
)

// The structure that defines a display mode.
// (https://wiki.libsdl.org/SDL3/SDL_DisplayMode)
type DisplayMode struct {
	DisplayID              DisplayID      /**< the display this mode is associated with */
	PixelFormat            PixelFormat    /**< pixel format */
	Width                  C.int          /**< width */
	Height                 C.int          /**< height */
	PixelDensity           float32        /**< scale converting size to pixels (e.g. a 1920x1080 mode with 2.0 scale would have 3840x2160 pixels) */
	RefreshRate            float32        /**< refresh rate (or 0.0f for unspecified) */
	RefreshRateNumerator   C.int          /**< precise refresh rate numerator (or 0 for unspecified) */
	RefreshRateDenominator C.int          /**< precise refresh rate denominator */
	DriverData             unsafe.Pointer /**< driver-specific data, initialize to 0 */
}

func (mode *DisplayMode) cptr() *C.SDL_DisplayMode {
	return (*C.SDL_DisplayMode)(unsafe.Pointer(mode))
}

// Display orientation values; the way a display is rotated.
// (https://wiki.libsdl.org/SDL3/SDL_DisplayOrientation)
type DisplayOrientation C.SDL_DisplayOrientation

const (
	ORIENTATION_UNKNOWN           = DisplayOrientation(C.SDL_ORIENTATION_UNKNOWN)           /**< The display orientation can't be determined */
	ORIENTATION_LANDSCAPE         = DisplayOrientation(C.SDL_ORIENTATION_LANDSCAPE)         /**< The display is in landscape mode, with the right side up, relative to portrait mode */
	ORIENTATION_LANDSCAPE_FLIPPED = DisplayOrientation(C.SDL_ORIENTATION_LANDSCAPE_FLIPPED) /**< The display is in landscape mode, with the left side up, relative to portrait mode */
	ORIENTATION_PORTRAIT          = DisplayOrientation(C.SDL_ORIENTATION_PORTRAIT)          /**< The display is in portrait mode */
	ORIENTATION_PORTRAIT_FLIPPED  = DisplayOrientation(C.SDL_ORIENTATION_PORTRAIT_FLIPPED)  /**< The display is in portrait mode, upside down */
)

// The struct used as an opaque handle to a window.
// (https://wiki.libsdl.org/SDL3/SDL_Window)
type Window C.SDL_Window

func (window *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(window))
}

// The flags on a window.
// (https://wiki.libsdl.org/SDL3/SDL_WindowFlags)
type WindowFlags C.SDL_WindowFlags

func (f WindowFlags) c() C.SDL_WindowFlags {
	return C.SDL_WindowFlags(f)
}

const (
	WINDOW_FULLSCREEN         = WindowFlags(C.SDL_WINDOW_FULLSCREEN)         /**< window is in fullscreen mode */
	WINDOW_OPENGL             = WindowFlags(C.SDL_WINDOW_OPENGL)             /**< window usable with OpenGL context */
	WINDOW_OCCLUDED           = WindowFlags(C.SDL_WINDOW_OCCLUDED)           /**< window is occluded */
	WINDOW_HIDDEN             = WindowFlags(C.SDL_WINDOW_HIDDEN)             /**< window is neither mapped onto the desktop nor shown in the taskbar/dock/window list; SDL_ShowWindow() is required for it to become visible */
	WINDOW_BORDERLESS         = WindowFlags(C.SDL_WINDOW_BORDERLESS)         /**< no window decoration */
	WINDOW_RESIZABLE          = WindowFlags(C.SDL_WINDOW_RESIZABLE)          /**< window can be resized */
	WINDOW_MINIMIZED          = WindowFlags(C.SDL_WINDOW_MINIMIZED)          /**< window is minimized */
	WINDOW_MAXIMIZED          = WindowFlags(C.SDL_WINDOW_MAXIMIZED)          /**< window is maximized */
	WINDOW_MOUSE_GRABBED      = WindowFlags(C.SDL_WINDOW_MOUSE_GRABBED)      /**< window has grabbed mouse input */
	WINDOW_INPUT_FOCUS        = WindowFlags(C.SDL_WINDOW_INPUT_FOCUS)        /**< window has input focus */
	WINDOW_MOUSE_FOCUS        = WindowFlags(C.SDL_WINDOW_MOUSE_FOCUS)        /**< window has mouse focus */
	WINDOW_EXTERNAL           = WindowFlags(C.SDL_WINDOW_EXTERNAL)           /**< window not created by SDL */
	WINDOW_MODAL              = WindowFlags(C.SDL_WINDOW_MODAL)              /**< window is modal */
	WINDOW_HIGH_PIXEL_DENSITY = WindowFlags(C.SDL_WINDOW_HIGH_PIXEL_DENSITY) /**< window uses high pixel density back buffer if possible */
	WINDOW_MOUSE_CAPTURE      = WindowFlags(C.SDL_WINDOW_MOUSE_CAPTURE)      /**< window has mouse captured (unrelated to MOUSE_GRABBED) */
	WINDOW_ALWAYS_ON_TOP      = WindowFlags(C.SDL_WINDOW_ALWAYS_ON_TOP)      /**< window should always be above others */
	WINDOW_UTILITY            = WindowFlags(C.SDL_WINDOW_UTILITY)            /**< window should be treated as a utility window, not showing in the task bar and window list */
	WINDOW_TOOLTIP            = WindowFlags(C.SDL_WINDOW_TOOLTIP)            /**< window should be treated as a tooltip and does not get mouse or keyboard focus, requires a parent window */
	WINDOW_POPUP_MENU         = WindowFlags(C.SDL_WINDOW_POPUP_MENU)         /**< window should be treated as a popup menu, requires a parent window */
	WINDOW_KEYBOARD_GRABBED   = WindowFlags(C.SDL_WINDOW_KEYBOARD_GRABBED)   /**< window has grabbed keyboard input */
	WINDOW_VULKAN             = WindowFlags(C.SDL_WINDOW_VULKAN)             /**< window usable for Vulkan surface */
	WINDOW_METAL              = WindowFlags(C.SDL_WINDOW_METAL)              /**< window usable for Metal view */
	WINDOW_TRANSPARENT        = WindowFlags(C.SDL_WINDOW_TRANSPARENT)        /**< window with transparent buffer */
	WINDOW_NOT_FOCUSABLE      = WindowFlags(C.SDL_WINDOW_NOT_FOCUSABLE)      /**< window should not be focusable */
)

// TODO: incomplete: skipped macros

// Window flash operation.
// (https://wiki.libsdl.org/SDL3/SDL_FlashOperation)
type FlashOperation C.SDL_FlashOperation

const (
	FLASH_CANCEL        = FlashOperation(C.SDL_FLASH_CANCEL)        /**< Cancel any window flash state */
	FLASH_BRIEFLY       = FlashOperation(C.SDL_FLASH_BRIEFLY)       /**< Flash the window briefly to get attention */
	FLASH_UNTIL_FOCUSED = FlashOperation(C.SDL_FLASH_UNTIL_FOCUSED) /**< Flash the window until it gets focus */
)

// An opaque handle to an OpenGL context.
type GLContext C.SDL_GLContext

// Opaque EGL types.
type EGLDisplay C.SDL_EGLDisplay
type EGLConfig C.SDL_EGLConfig
type EGLSurface C.SDL_EGLSurface
type EGLAttrib C.SDL_EGLAttrib
type EGLint C.SDL_EGLint

// EGL attribute initialization callback types.
type EGLAttribArrayCallback C.SDL_EGLAttribArrayCallback
type EGLIntArrayCallback C.SDL_EGLIntArrayCallback

// An enumeration of OpenGL configuration attributes.
// (https://wiki.libsdl.org/SDL3/SDL_GLattr)
type GLAttr C.SDL_GLattr

const (
	GL_RED_SIZE                   = GLAttr(C.SDL_GL_RED_SIZE)
	GL_GREEN_SIZE                 = GLAttr(C.SDL_GL_GREEN_SIZE)
	GL_BLUE_SIZE                  = GLAttr(C.SDL_GL_BLUE_SIZE)
	GL_ALPHA_SIZE                 = GLAttr(C.SDL_GL_ALPHA_SIZE)
	GL_BUFFER_SIZE                = GLAttr(C.SDL_GL_BUFFER_SIZE)
	GL_DOUBLEBUFFER               = GLAttr(C.SDL_GL_DOUBLEBUFFER)
	GL_DEPTH_SIZE                 = GLAttr(C.SDL_GL_DEPTH_SIZE)
	GL_STENCIL_SIZE               = GLAttr(C.SDL_GL_STENCIL_SIZE)
	GL_ACCUM_RED_SIZE             = GLAttr(C.SDL_GL_ACCUM_RED_SIZE)
	GL_ACCUM_GREEN_SIZE           = GLAttr(C.SDL_GL_ACCUM_GREEN_SIZE)
	GL_ACCUM_BLUE_SIZE            = GLAttr(C.SDL_GL_ACCUM_BLUE_SIZE)
	GL_ACCUM_ALPHA_SIZE           = GLAttr(C.SDL_GL_ACCUM_ALPHA_SIZE)
	GL_STEREO                     = GLAttr(C.SDL_GL_STEREO)
	GL_MULTISAMPLEBUFFERS         = GLAttr(C.SDL_GL_MULTISAMPLEBUFFERS)
	GL_MULTISAMPLESAMPLES         = GLAttr(C.SDL_GL_MULTISAMPLESAMPLES)
	GL_ACCELERATED_VISUAL         = GLAttr(C.SDL_GL_ACCELERATED_VISUAL)
	GL_RETAINED_BACKING           = GLAttr(C.SDL_GL_RETAINED_BACKING)
	GL_CONTEXT_MAJOR_VERSION      = GLAttr(C.SDL_GL_CONTEXT_MAJOR_VERSION)
	GL_CONTEXT_MINOR_VERSION      = GLAttr(C.SDL_GL_CONTEXT_MINOR_VERSION)
	GL_CONTEXT_FLAGS              = GLAttr(C.SDL_GL_CONTEXT_FLAGS)
	GL_CONTEXT_PROFILE_MASK       = GLAttr(C.SDL_GL_CONTEXT_PROFILE_MASK)
	GL_SHARE_WITH_CURRENT_CONTEXT = GLAttr(C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT)
	GL_FRAMEBUFFER_SRGB_CAPABLE   = GLAttr(C.SDL_GL_FRAMEBUFFER_SRGB_CAPABLE)
	GL_CONTEXT_RELEASE_BEHAVIOR   = GLAttr(C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR)
	GL_CONTEXT_RESET_NOTIFICATION = GLAttr(C.SDL_GL_CONTEXT_RESET_NOTIFICATION)
	GL_CONTEXT_NO_ERROR           = GLAttr(C.SDL_GL_CONTEXT_NO_ERROR)
	GL_FLOATBUFFERS               = GLAttr(C.SDL_GL_FLOATBUFFERS)
	GL_EGL_PLATFORM               = GLAttr(C.SDL_GL_EGL_PLATFORM)
)

// Possible values to be set for the SDL_GL_CONTEXT_PROFILE_MASK attribute.
// (https://wiki.libsdl.org/SDL3/SDL_GLprofile)
type GLProfile C.SDL_GLprofile

const (
	GL_CONTEXT_PROFILE_CORE          = GLProfile(C.SDL_GL_CONTEXT_PROFILE_CORE)
	GL_CONTEXT_PROFILE_COMPATIBILITY = GLProfile(C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY)
	GL_CONTEXT_PROFILE_ES            = GLProfile(C.SDL_GL_CONTEXT_PROFILE_ES) /**< GLX_CONTEXT_ES2_PROFILE_BIT_EXT */
)

// Possible values to be set for the SDL_GL_CONTEXT_FLAGS attribute.
// (https://wiki.libsdl.org/SDL3/SDL_GLcontextFlag)
type GLContextFlag C.SDL_GLcontextFlag

const (
	GL_CONTEXT_DEBUG_FLAG              = GLContextFlag(C.SDL_GL_CONTEXT_DEBUG_FLAG)
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = GLContextFlag(C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG)
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = GLContextFlag(C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG)
	GL_CONTEXT_RESET_ISOLATION_FLAG    = GLContextFlag(C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG)
)

// Possible values to be set for the SDL_GL_CONTEXT_RELEASE_BEHAVIOR
// (https://wiki.libsdl.org/SDL3/SDL_GLcontextReleaseFlag)
type GLContextReleaseFlag C.SDL_GLcontextReleaseFlag

const (
	GL_CONTEXT_RELEASE_BEHAVIOR_NONE  = GLContextReleaseFlag(C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR_NONE)
	GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH = GLContextReleaseFlag(C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH)
)

// Possible values to be set SDL_GL_CONTEXT_RESET_NOTIFICATION attribute.
// (https://wiki.libsdl.org/SDL3/SDL_GLContextResetNotification)
type GLContextResetNotification C.SDL_GLContextResetNotification

const (
	GL_CONTEXT_RESET_NO_NOTIFICATION = GLContextResetNotification(C.SDL_GL_CONTEXT_RESET_NO_NOTIFICATION)
	GL_CONTEXT_RESET_LOSE_CONTEXT    = GLContextResetNotification(C.SDL_GL_CONTEXT_RESET_LOSE_CONTEXT)
)

// Get the number of video drivers compiled into SDL.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers)
func GetNumVideoDrivers() (n int, err error) {
	n = int(C.SDL_GetNumVideoDrivers())
	if n != 0 {
		err = GetError()
	}
	return
}

// Get the name of a built in video driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver)
func GetVideoDriver(index int) string {
	return C.GoString(C.SDL_GetVideoDriver(C.int(index)))

}

// Get the name of the currently initialized video driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver)
func GetCurrentVideoDriver() string {
	return C.GoString(C.SDL_GetCurrentVideoDriver())
}

// Get the current system theme.
// (https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme)
func GetSystemTheme() SystemTheme {
	return SystemTheme(C.SDL_GetSystemTheme())
}

// Get a list of currently connected displays.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplays)
func GetDisplays() []DisplayID {
	var count C.int
	ret := C.SDL_GetDisplays(&count)
	if ret == nil || count == 0 {
		return nil
	}

	sdlSlice := unsafe.Slice((*C.SDL_DisplayID)(ret), count)
	result := make([]DisplayID, count)
	for i := range result {
		result[i] = DisplayID(sdlSlice[i])
	}

	C.SDL_free(unsafe.Pointer(ret))
	return result
}

// Return the primary display.
// (https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay)
func GetPrimaryDisplay() (id DisplayID, err error) {
	id = DisplayID(C.SDL_GetPrimaryDisplay())
	if id == 0 {
		err = GetError()
	}
	return
}

// Get the properties associated with a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties)
func (displayID DisplayID) GetDisplayProperties() (id PropertiesID, err error) {
	id = PropertiesID(C.SDL_GetDisplayProperties(C.SDL_DisplayID(displayID)))
	if id == 0 {
		err = GetError()
	}
	return
}

const PROP_DISPLAY_HDR_ENABLED_BOOLEAN = C.SDL_PROP_DISPLAY_HDR_ENABLED_BOOLEAN
const PROP_DISPLAY_KMSDRM_PANEL_ORIENTATION_NUMBER = C.SDL_PROP_DISPLAY_KMSDRM_PANEL_ORIENTATION_NUMBER

// Get the name of a display in UTF-8 encoding.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayName)
func (displayID DisplayID) GetName() (name string, err error) {
	ret := C.SDL_GetDisplayName(C.SDL_DisplayID(displayID))
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Get the desktop area represented by a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds)
func (displayID DisplayID) GetBounds() (rect Rect, err error) {
	ret := C.SDL_GetDisplayBounds(C.SDL_DisplayID(displayID), rect.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the usable desktop area represented by a display, in screen
// coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds)
func (displayID DisplayID) GetUsableBounds() (rect Rect, err error) {
	ret := C.SDL_GetDisplayUsableBounds(C.SDL_DisplayID(displayID), rect.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the orientation of a display when it is unrotated.
// (https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation)
func (displayID DisplayID) GetNaturalOrientation() DisplayOrientation {
	return DisplayOrientation(C.SDL_GetNaturalDisplayOrientation(C.SDL_DisplayID(displayID)))
}

// Get the orientation of a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation)
func (displayID DisplayID) GetCurrentOrientation() DisplayOrientation {
	return DisplayOrientation(C.SDL_GetCurrentDisplayOrientation(C.SDL_DisplayID(displayID)))
}

// Get the content scale of a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale)
func (displayID DisplayID) GetContentScale() (scale float32, err error) {
	scale = float32(C.SDL_GetDisplayContentScale(C.SDL_DisplayID(displayID)))
	if scale == 0.0 {
		err = GetError()
	}
	return
}

// Get a list of fullscreen display modes available on a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes)
func (displayID DisplayID) GetFullscreenModes() ([]DisplayMode, error) {
	var count C.int
	ret := C.SDL_GetFullscreenDisplayModes(C.SDL_DisplayID(displayID), &count)
	if ret == nil {
		return nil, GetError()
	}
	// Make our own copy
	cSlice := unsafe.Slice((**C.SDL_DisplayMode)(ret), count)
	result := make([]DisplayMode, count)
	for i := 0; i < int(count); i++ {
		a := *((*DisplayMode)(unsafe.Pointer(cSlice[i])))
		result[i] = a
	}
	C.SDL_free(unsafe.Pointer(ret))
	return result, nil
}

// Get the closest match to the requested display mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode)
func (displayID DisplayID) GetClosestFullscreenMode(w, h int32, refresh_rate float32, include_high_density_modes bool) (mode DisplayMode, err error) {
	ret := C.SDL_GetClosestFullscreenDisplayMode(
		displayID.c(), C.int(w), C.int(h), C.float(refresh_rate), sdlBool(include_high_density_modes), mode.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get information about the desktop's display mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode)
func (displayID DisplayID) GetDesktopMode() (m *DisplayMode, err error) {
	m = (*DisplayMode)(unsafe.Pointer(C.SDL_GetDesktopDisplayMode(C.SDL_DisplayID(displayID))))
	if m == nil {
		err = GetError()
	}
	return
}

// Get information about the current display mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode)
func (displayID DisplayID) GetCurrentMode() (m *DisplayMode, err error) {
	m = (*DisplayMode)(unsafe.Pointer(C.SDL_GetCurrentDisplayMode(C.SDL_DisplayID(displayID))))
	if m == nil {
		err = GetError()
	}
	return
}

// Get the display containing a point.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint)
func (point *Point) GetDisplay() (id DisplayID, err error) {
	id = DisplayID(C.SDL_GetDisplayForPoint(point.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Get the display primarily containing a rect.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect)
func (rect *Rect) GetDisplay() (id DisplayID, err error) {
	id = DisplayID(C.SDL_GetDisplayForRect(rect.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Get the display associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow)
func (window *Window) GetDisplayForWindow() (id DisplayID, err error) {
	id = DisplayID(C.SDL_GetDisplayForWindow(window.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Get the pixel density of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity)
func (window *Window) GetPixelDensity() (f float32, err error) {
	f = float32(C.SDL_GetWindowPixelDensity(window.cptr()))
	if f == 0 {
		err = GetError()
	}
	return
}

// Get the content display scale relative to a window's pixel size.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale)
func (window *Window) GetDisplayScale() (f float32, err error) {
	f = float32(C.SDL_GetWindowDisplayScale(window.cptr()))
	if f == 0 {
		err = GetError()
	}
	return
}

// Set the display mode to use when a window is visible and fullscreen.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode)
func (window *Window) SetFullscreenMode(mode *DisplayMode) error {
	if C.SDL_SetWindowFullscreenMode(window.cptr(), mode.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Query the display mode to use when a window is visible at fullscreen.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode)
func (window *Window) GetFullscreenMode() *DisplayMode {
	return (*DisplayMode)(unsafe.Pointer(C.SDL_GetWindowFullscreenMode(window.cptr())))
}

// Get the raw ICC profile data for the screen the window is currently on.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile)
func (window *Window) GetICCProfile() (iccProfile unsafe.Pointer, size uintptr, err error) {
	iccProfile = C.SDL_GetWindowICCProfile(window.cptr(), (*C.size_t)(unsafe.Pointer(&size)))
	if iccProfile == nil {
		err = GetError()
	}
	return
}

// Get the pixel format associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat)
func (window *Window) GetPixelFormat() (format PixelFormat, err error) {
	format = PixelFormat(C.SDL_GetWindowPixelFormat(window.cptr()))
	if format == 0 {
		err = GetError()
	}
	return
}

// Get a list of valid windows.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindows)
func GetWindows() ([]*Window, error) {
	var count C.int
	//SDL_Window **
	ret := unsafe.Pointer(C.SDL_GetWindows(&count))
	if ret == nil {
		return nil, GetError()
	}
	cSlice := unsafe.Slice((**Window)(ret), count)
	// Make our own copy
	result := make([]*Window, count)
	for i := 0; i < int(count); i++ {
		result[i] = cSlice[i]
	}
	C.SDL_free(ret)
	return result, nil
}

// Create a window with the specified dimensions and flags.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindow)
func CreateWindow(title string, w, h int, flags WindowFlags) (window *Window, err error) {
	window = (*Window)(unsafe.Pointer(C.SDL_CreateWindow(C.CString(title), C.int(w), C.int(h), flags.c())))
	if window == nil {
		err = GetError()
	}
	return
}

// Create a child popup window of the specified parent window.
// (https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow)
func (parent *Window) CreatePopupWindow(offset_x, offset_y int, w, h int, flags WindowFlags) (window *Window, err error) {
	window = (*Window)(unsafe.Pointer(C.SDL_CreatePopupWindow(
		parent.cptr(), C.int(offset_x), C.int(offset_y), C.int(w), C.int(h), flags.c())))
	if window == nil {
		err = GetError()
	}
	return
}

// Create a window with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties)
func CreateWindowWithProperties(props PropertiesID) (window *Window, err error) {
	window = (*Window)(unsafe.Pointer(C.SDL_CreateWindowWithProperties(C.SDL_PropertiesID(props))))
	if window == nil {
		err = GetError()
	}
	return
}

const (
	PROP_WINDOW_CREATE_ALWAYS_ON_TOP_BOOLEAN               = "always_on_top"
	PROP_WINDOW_CREATE_BORDERLESS_BOOLEAN                  = "borderless"
	PROP_WINDOW_CREATE_FOCUSABLE_BOOLEAN                   = "focusable"
	PROP_WINDOW_CREATE_EXTERNAL_GRAPHICS_CONTEXT_BOOLEAN   = "external_graphics_context"
	PROP_WINDOW_CREATE_FULLSCREEN_BOOLEAN                  = "fullscreen"
	PROP_WINDOW_CREATE_HEIGHT_NUMBER                       = "height"
	PROP_WINDOW_CREATE_HIDDEN_BOOLEAN                      = "hidden"
	PROP_WINDOW_CREATE_HIGH_PIXEL_DENSITY_BOOLEAN          = "high_pixel_density"
	PROP_WINDOW_CREATE_MAXIMIZED_BOOLEAN                   = "maximized"
	PROP_WINDOW_CREATE_MENU_BOOLEAN                        = "menu"
	PROP_WINDOW_CREATE_METAL_BOOLEAN                       = "metal"
	PROP_WINDOW_CREATE_MINIMIZED_BOOLEAN                   = "minimized"
	PROP_WINDOW_CREATE_MODAL_BOOLEAN                       = "modal"
	PROP_WINDOW_CREATE_MOUSE_GRABBED_BOOLEAN               = "mouse_grabbed"
	PROP_WINDOW_CREATE_OPENGL_BOOLEAN                      = "opengl"
	PROP_WINDOW_CREATE_PARENT_POINTER                      = "parent"
	PROP_WINDOW_CREATE_RESIZABLE_BOOLEAN                   = "resizable"
	PROP_WINDOW_CREATE_TITLE_STRING                        = "title"
	PROP_WINDOW_CREATE_TRANSPARENT_BOOLEAN                 = "transparent"
	PROP_WINDOW_CREATE_TOOLTIP_BOOLEAN                     = "tooltip"
	PROP_WINDOW_CREATE_UTILITY_BOOLEAN                     = "utility"
	PROP_WINDOW_CREATE_VULKAN_BOOLEAN                      = "vulkan"
	PROP_WINDOW_CREATE_WIDTH_NUMBER                        = "width"
	PROP_WINDOW_CREATE_X_NUMBER                            = "x"
	PROP_WINDOW_CREATE_Y_NUMBER                            = "y"
	PROP_WINDOW_CREATE_COCOA_WINDOW_POINTER                = "cocoa.window"
	PROP_WINDOW_CREATE_COCOA_VIEW_POINTER                  = "cocoa.view"
	PROP_WINDOW_CREATE_WAYLAND_SURFACE_ROLE_CUSTOM_BOOLEAN = "wayland.surface_role_custom"
	PROP_WINDOW_CREATE_WAYLAND_CREATE_EGL_WINDOW_BOOLEAN   = "wayland.create_egl_window"
	PROP_WINDOW_CREATE_WAYLAND_WL_SURFACE_POINTER          = "wayland.wl_surface"
	PROP_WINDOW_CREATE_WIN32_HWND_POINTER                  = "win32.hwnd"
	PROP_WINDOW_CREATE_WIN32_PIXEL_FORMAT_HWND_POINTER     = "win32.pixel_format_hwnd"
	PROP_WINDOW_CREATE_X11_WINDOW_NUMBER                   = "x11.window"
)

// Get the numeric ID of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowID)
func (window *Window) GetID() (id WindowID, err error) {
	id = WindowID(C.SDL_GetWindowID(window.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Get a window from a stored ID.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID)
func (id WindowID) GetWindow() (window *Window, err error) {
	window = (*Window)(unsafe.Pointer(C.SDL_GetWindowFromID(C.SDL_WindowID(id))))
	if window == nil {
		err = GetError()
	}
	return
}

// Get parent of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowParent)
func (window *Window) GetParent() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetWindowParent(window.cptr())))
}

// Get the properties associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties)
func (window *Window) GetProperties() (id PropertiesID, err error) {
	id = PropertiesID(C.SDL_GetWindowProperties(window.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

const (
	PROP_WINDOW_SHAPE_POINTER                             = "SDL.window.shape"
	PROP_WINDOW_HDR_ENABLED_BOOLEAN                       = "SDL.window.HDR_enabled"
	PROP_WINDOW_SDR_WHITE_LEVEL_FLOAT                     = "SDL.window.SDR_white_level"
	PROP_WINDOW_HDR_HEADROOM_FLOAT                        = "SDL.window.HDR_headroom"
	PROP_WINDOW_ANDROID_WINDOW_POINTER                    = "SDL.window.android.window"
	PROP_WINDOW_ANDROID_SURFACE_POINTER                   = "SDL.window.android.surface"
	PROP_WINDOW_UIKIT_WINDOW_POINTER                      = "SDL.window.uikit.window"
	PROP_WINDOW_UIKIT_METAL_VIEW_TAG_NUMBER               = "SDL.window.uikit.metal_view_tag"
	PROP_WINDOW_UIKIT_OPENGL_FRAMEBUFFER_NUMBER           = "SDL.window.uikit.opengl.framebuffer"
	PROP_WINDOW_UIKIT_OPENGL_RENDERBUFFER_NUMBER          = "SDL.window.uikit.opengl.renderbuffer"
	PROP_WINDOW_UIKIT_OPENGL_RESOLVE_FRAMEBUFFER_NUMBER   = "SDL.window.uikit.opengl.resolve_framebuffer"
	PROP_WINDOW_KMSDRM_DEVICE_INDEX_NUMBER                = "SDL.window.kmsdrm.dev_index"
	PROP_WINDOW_KMSDRM_DRM_FD_NUMBER                      = "SDL.window.kmsdrm.drm_fd"
	PROP_WINDOW_KMSDRM_GBM_DEVICE_POINTER                 = "SDL.window.kmsdrm.gbm_dev"
	PROP_WINDOW_COCOA_WINDOW_POINTER                      = "SDL.window.cocoa.window"
	PROP_WINDOW_COCOA_METAL_VIEW_TAG_NUMBER               = "SDL.window.cocoa.metal_view_tag"
	PROP_WINDOW_VIVANTE_DISPLAY_POINTER                   = "SDL.window.vivante.display"
	PROP_WINDOW_VIVANTE_WINDOW_POINTER                    = "SDL.window.vivante.window"
	PROP_WINDOW_VIVANTE_SURFACE_POINTER                   = "SDL.window.vivante.surface"
	PROP_WINDOW_WINRT_WINDOW_POINTER                      = "SDL.window.winrt.window"
	PROP_WINDOW_WIN32_HWND_POINTER                        = "SDL.window.win32.hwnd"
	PROP_WINDOW_WIN32_HDC_POINTER                         = "SDL.window.win32.hdc"
	PROP_WINDOW_WIN32_INSTANCE_POINTER                    = "SDL.window.win32.instance"
	PROP_WINDOW_WAYLAND_DISPLAY_POINTER                   = "SDL.window.wayland.display"
	PROP_WINDOW_WAYLAND_SURFACE_POINTER                   = "SDL.window.wayland.surface"
	PROP_WINDOW_WAYLAND_EGL_WINDOW_POINTER                = "SDL.window.wayland.egl_window"
	PROP_WINDOW_WAYLAND_XDG_SURFACE_POINTER               = "SDL.window.wayland.xdg_surface"
	PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_POINTER              = "SDL.window.wayland.xdg_toplevel"
	PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_EXPORT_HANDLE_STRING = "SDL.window.wayland.xdg_toplevel_export_handle"
	PROP_WINDOW_WAYLAND_XDG_POPUP_POINTER                 = "SDL.window.wayland.xdg_popup"
	PROP_WINDOW_WAYLAND_XDG_POSITIONER_POINTER            = "SDL.window.wayland.xdg_positioner"
	PROP_WINDOW_X11_DISPLAY_POINTER                       = "SDL.window.x11.display"
	PROP_WINDOW_X11_SCREEN_NUMBER                         = "SDL.window.x11.screen"
	PROP_WINDOW_X11_WINDOW_NUMBER                         = "SDL.window.x11.window"
)

// Get the window flags.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags)
func (window *Window) GetFlags() WindowFlags {
	return WindowFlags(C.SDL_GetWindowFlags(window.cptr()))
}

// Set the title of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle)
func (window *Window) SetTitle(title string) (err error) {
	ret := C.SDL_SetWindowTitle(window.cptr(), C.CString(title))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the title of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle)
func (window *Window) GetTitle() string {
	return C.GoString(C.SDL_GetWindowTitle(window.cptr()))
}

// Set the icon for a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon)
func (window *Window) SetIcon(icon *Surface) (err error) {
	ret := C.SDL_SetWindowIcon(window.cptr(), icon.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the window's position be set.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition)
func (window *Window) SetPosition(x, y int32) (err error) {
	ret := C.SDL_SetWindowPosition(window.cptr(), C.int(x), C.int(y))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the position of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition)
func (window *Window) GetPosition() (x, y int32, err error) {
	ret := C.SDL_GetWindowPosition(window.cptr(), (*C.int)(&x), (*C.int)(&y))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the size of a window's client area be set.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowSize)
func (window *Window) SetSize(w, h int32) (err error) {
	ret := C.SDL_SetWindowSize(window.cptr(), C.int(w), C.int(h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSize)
func (window *Window) GetSize() (w, h int32, err error) {
	ret := C.SDL_GetWindowSize(window.cptr(), (*C.int)(&w), (*C.int)(&h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the aspect ratio of a window's client area be set.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio)
func (window *Window) SetAspectRatio(min_aspect, max_aspect float32) (err error) {
	ret := C.SDL_SetWindowAspectRatio(window.cptr(), C.float(min_aspect), C.float(max_aspect))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio)
func (window *Window) GetAspectRatio() (min_aspect, max_aspect float32, err error) {
	ret := C.SDL_GetWindowAspectRatio(window.cptr(), (*C.float)(&min_aspect), (*C.float)(&max_aspect))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the size of a window's borders (decorations) around the client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize)
func (window *Window) GetBordersSize() (top, left, bottom, right int32, err error) {
	ret := C.SDL_GetWindowBordersSize(window.cptr(), (*C.int)(&top), (*C.int)(&left),
		(*C.int)(&bottom), (*C.int)(&right))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the size of a window's client area, in pixels.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels)
func (window *Window) GetSizeInPixels() (w, h int32, err error) {
	ret := C.SDL_GetWindowSizeInPixels(window.cptr(), (*C.int)(&w), (*C.int)(&h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the minimum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize)
func (window *Window) SetMinimumSize(min_w, min_h int32) (err error) {
	ret := C.SDL_SetWindowMinimumSize(window.cptr(), C.int(min_w), C.int(min_h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the minimum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize)
func (window *Window) GetMinimumSize() (w, h int32, err error) {
	ret := C.SDL_GetWindowMinimumSize(window.cptr(), (*C.int)(&w), (*C.int)(&h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the maximum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize)
func (window *Window) SetMaximumSize(max_w, max_h int32) (err error) {
	ret := C.SDL_SetWindowMaximumSize(window.cptr(), C.int(max_w), C.int(max_h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the maximum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize)
func (window *Window) GetMaximumSize() (w, h int32, err error) {
	ret := C.SDL_GetWindowMaximumSize(window.cptr(), (*C.int)(&w), (*C.int)(&h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the border state of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered)
func (window *Window) SetBordered(bordered bool) (err error) {
	ret := C.SDL_SetWindowBordered(window.cptr(), sdlBool(bordered))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the user-resizable state of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable)
func (window *Window) SetResizable(resizable bool) (err error) {
	ret := C.SDL_SetWindowResizable(window.cptr(), sdlBool(resizable))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the window to always be above the others.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop)
func (window *Window) SetAlwaysOnTop(on_top bool) (err error) {
	ret := C.SDL_SetWindowAlwaysOnTop(window.cptr(), sdlBool(on_top))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Show a window.
// (https://wiki.libsdl.org/SDL3/SDL_ShowWindow)
func (window *Window) Show() (err error) {
	ret := C.SDL_ShowWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Hide a window.
// (https://wiki.libsdl.org/SDL3/SDL_HideWindow)
func (window *Window) Hide() (err error) {
	ret := C.SDL_HideWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that a window be raised above other windows and gain the input
// focus.
// (https://wiki.libsdl.org/SDL3/SDL_RaiseWindow)
func (window *Window) Raise() (err error) {
	ret := C.SDL_RaiseWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the window be made as large as possible.
// (https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow)
func (window *Window) Maximize() (err error) {
	ret := C.SDL_MaximizeWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the window be minimized to an iconic representation.
// (https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow)
func (window *Window) Minimize() (err error) {
	ret := C.SDL_MinimizeWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the size and position of a minimized or maximized window be
// restored.
// (https://wiki.libsdl.org/SDL3/SDL_RestoreWindow)
func (window *Window) Restore() (err error) {
	ret := C.SDL_RestoreWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request that the window's fullscreen state be changed.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen)
func (window *Window) SetFullscreen(fullscreen bool) (err error) {
	ret := C.SDL_SetWindowFullscreen(window.cptr(), sdlBool(fullscreen))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Block until any pending window state is finalized.
// (https://wiki.libsdl.org/SDL3/SDL_SyncWindow)
func (window *Window) Sync() (err error) {
	ret := C.SDL_SyncWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Return whether the window has a surface associated with it.
// (https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface)
func (window *Window) HasSurface() bool {
	return C.SDL_WindowHasSurface(window.cptr()) != 0
}

// Get the SDL surface associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface)
func (window *Window) GetSurface() (surface *Surface, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(window.cptr())))
	if surface == nil {
		err = GetError()
	}
	return
}

// Toggle VSync for the window surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync)
func (window *Window) SetSurfaceVSync(vsync int32) (err error) {
	ret := C.SDL_SetWindowSurfaceVSync(window.cptr(), C.int(vsync))
	if ret != 0 {
		err = GetError()
	}
	return
}

const SDL_WINDOW_SURFACE_VSYNC_DISABLED = 0
const SDL_WINDOW_SURFACE_VSYNC_ADAPTIVE = (-1)

// Get VSync for the window surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync)
func (window *Window) GetSurfaceVSync() (vsync int32, err error) {
	ret := C.SDL_GetWindowSurfaceVSync(window.cptr(), (*C.int)(&vsync))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Copy the window surface to the screen.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface)
func (window *Window) UpdateSurface() (err error) {
	ret := C.SDL_UpdateWindowSurface(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Copy areas of the window surface to the screen.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects)
func (window *Window) UpdateSurfaceRects(rects []Rect) (err error) {
	ret := C.SDL_UpdateWindowSurfaceRects(window.cptr(), rects[0].cptr(), C.int(len(rects)))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Destroy the surface associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface)
func (window *Window) DestroySurface() (err error) {
	ret := C.SDL_DestroyWindowSurface(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set a window's keyboard grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab)
func (window *Window) SetKeyboardGrab(grabbed bool) (err error) {
	ret := C.SDL_SetWindowKeyboardGrab(window.cptr(), sdlBool(grabbed))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set a window's mouse grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab)
func (window *Window) SetMouseGrab(grabbed bool) (err error) {
	ret := C.SDL_SetWindowMouseGrab(window.cptr(), sdlBool(grabbed))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get a window's keyboard grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab)
func (window *Window) GetKeyboardGrab() bool {
	return C.SDL_GetWindowKeyboardGrab(window.cptr()) != 0
}

// Get a window's mouse grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab)
func (window *Window) GetMouseGrab() bool {
	return C.SDL_GetWindowMouseGrab(window.cptr()) != 0
}

// Get the window that currently has an input grab enabled.
// (https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow)
func GetGrabbedWindow() *Window {
	return (*Window)(unsafe.Pointer(C.SDL_GetGrabbedWindow()))
}

// Confines the cursor to the specified area of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect)
func (window *Window) SetMouseRect(rect Rect) (err error) {
	ret := C.SDL_SetWindowMouseRect(window.cptr(), rect.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the mouse confinement rectangle of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect)
func (window *Window) GetMouseRect() *Rect {
	return (*Rect)(unsafe.Pointer(C.SDL_GetWindowMouseRect(window.cptr())))
}

// Set the opacity for a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity)
func (window *Window) SetOpacity(opacity float32) (err error) {
	ret := C.SDL_SetWindowOpacity(window.cptr(), C.float(opacity))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the opacity of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity)
func (window *Window) GetOpacity() (opacity float32, err error) {
	opacity = float32(C.SDL_GetWindowOpacity(window.cptr()))
	if opacity < 0 {
		err = GetError()
	}
	return
}

// Set the window as a modal to a parent window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowModalFor)
func (modal_window *Window) SetModalFor(parent_window *Window) (err error) {
	ret := C.SDL_SetWindowModalFor(modal_window.cptr(), parent_window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set whether the window may have input focus.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable)
func (window *Window) SetFocusable(focusable bool) (err error) {
	ret := C.SDL_SetWindowFocusable(window.cptr(), sdlBool(focusable))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Display the system-level window menu.
// (https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu)
func (window *Window) ShowSystemMenu(x, y int32) (err error) {
	ret := C.SDL_ShowWindowSystemMenu(window.cptr(), C.int(x), C.int(y))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Possible return values from the SDL_HitTest callback.
// (https://wiki.libsdl.org/SDL3/SDL_HitTestResult)
type HitTestResult C.SDL_HitTestResult

const (
	HITTEST_NORMAL             = HitTestResult(C.SDL_HITTEST_NORMAL)             /**< Region is normal. No special properties. */
	HITTEST_DRAGGABLE          = HitTestResult(C.SDL_HITTEST_DRAGGABLE)          /**< Region can drag entire window. */
	HITTEST_RESIZE_TOPLEFT     = HitTestResult(C.SDL_HITTEST_RESIZE_TOPLEFT)     /**< Region is the resizable top-left corner border. */
	HITTEST_RESIZE_TOP         = HitTestResult(C.SDL_HITTEST_RESIZE_TOP)         /**< Region is the resizable top border. */
	HITTEST_RESIZE_TOPRIGHT    = HitTestResult(C.SDL_HITTEST_RESIZE_TOPRIGHT)    /**< Region is the resizable top-right corner border. */
	HITTEST_RESIZE_RIGHT       = HitTestResult(C.SDL_HITTEST_RESIZE_RIGHT)       /**< Region is the resizable right border. */
	HITTEST_RESIZE_BOTTOMRIGHT = HitTestResult(C.SDL_HITTEST_RESIZE_BOTTOMRIGHT) /**< Region is the resizable bottom-right corner border. */
	HITTEST_RESIZE_BOTTOM      = HitTestResult(C.SDL_HITTEST_RESIZE_BOTTOM)      /**< Region is the resizable bottom border. */
	HITTEST_RESIZE_BOTTOMLEFT  = HitTestResult(C.SDL_HITTEST_RESIZE_BOTTOMLEFT)  /**< Region is the resizable bottom-left corner border. */
	HITTEST_RESIZE_LEFT        = HitTestResult(C.SDL_HITTEST_RESIZE_LEFT)        /**< Region is the resizable left border. */
)

// Callback used for hit-testing.
// (https://wiki.libsdl.org/SDL3/SDL_HitTest)
type HitTestCallback func(window *Window, area *Point, data any) HitTestResult

type hitTestCallbackData struct {
	callback HitTestCallback
	data     any
}

var windowHitTestCallbacks = make(map[*Window]hitTestCallbackData)

//export goHitTestCallback
func goHitTestCallback(window *C.SDL_Window, area *C.SDL_Point) C.SDL_HitTestResult {
	cd, ok := windowHitTestCallbacks[(*Window)(window)]
	if !ok {
		fmt.Printf("goHitTestCallback: no callback for window\n")
		return C.SDL_HITTEST_NORMAL
	}
	ret := cd.callback((*Window)(window), (*Point)(unsafe.Pointer(area)), cd.data)
	return C.SDL_HitTestResult(ret)
}

// Provide a callback that decides if a window region has special properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest)
func (window *Window) SetHitTest(callback HitTestCallback, data any) error {
	windowHitTestCallbacks[window] = hitTestCallbackData{callback: callback, data: data}
	ret := C.SDL_SetWindowHitTest(window.cptr(), C.SDL_HitTest(C.cgoHitTestCallback), nil)
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set the shape of a transparent window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowShape)
func (window *Window) SetShape(shape *Surface) (err error) {
	ret := C.SDL_SetWindowShape(window.cptr(), shape.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Request a window to demand attention from the user.
// (https://wiki.libsdl.org/SDL3/SDL_FlashWindow)
func (window *Window) Flash(operation FlashOperation) (err error) {
	ret := C.SDL_FlashWindow(window.cptr(), C.SDL_FlashOperation(operation))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Destroy a window.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyWindow)
func (window *Window) Destroy() {
	C.SDL_DestroyWindow(window.cptr())
}

// Check whether the screensaver is currently enabled.
// (https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled)
func ScreenSaverEnabled() bool {
	return C.SDL_ScreenSaverEnabled() != 0
}

// Allow the screen to be blanked by a screen saver.
// (https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver)
func EnableScreenSaver() (err error) {
	ret := C.SDL_EnableScreenSaver()
	if ret != 0 {
		err = GetError()
	}
	return
}

// Prevent the screen from being blanked by a screen saver.
// (https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver)
func DisableScreenSaver() (err error) {
	ret := C.SDL_DisableScreenSaver()
	if ret != 0 {
		err = GetError()
	}
	return
}

//	\name OpenGL support functions
//
// /
// (https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary)
func GL_LoadLibrary(path string) (err error) {
	var cPath *C.char
	if len(path) > 0 {
		cPath = C.CString(path)
	}
	ret := C.SDL_GL_LoadLibrary(cPath)
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get an OpenGL function by name.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress)
func GL_GetProcAddress(proc string) FunctionPointer {
	return FunctionPointer(C.SDL_GL_GetProcAddress(C.CString(proc)))
}

// Get an EGL library function by name.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress)
func EGL_GetProcAddress(proc string) FunctionPointer {
	return FunctionPointer(C.SDL_EGL_GetProcAddress(C.CString(proc)))
}

// Unload the OpenGL library previously loaded by SDL_GL_LoadLibrary().
// (https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary)
func GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

// Check if an OpenGL extension is supported for the current context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported)
func GL_ExtensionSupported(extension string) bool {
	return C.SDL_GL_ExtensionSupported(C.CString(extension)) != 0
}

// Reset all previously set OpenGL context attributes to their default values.
// (https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes)
func GL_ResetAttributes() {
	C.SDL_GL_ResetAttributes()
}

// Set an OpenGL window attribute before window creation.
// (https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute)
func GL_SetAttribute(attr GLAttr, value int32) (err error) {
	ret := C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the actual value for an attribute from the current context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute)
func GL_GetAttribute(attr GLAttr) (value int32, err error) {
	ret := C.SDL_GL_GetAttribute(C.SDL_GLattr(attr), (*C.int)(&value))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Create an OpenGL context for an OpenGL window, and make it current.
// (https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext)
func (window *Window) GL_CreateContext() (ctx GLContext, err error) {
	ctx = GLContext(C.SDL_GL_CreateContext(window.cptr()))
	if ctx == nil {
		err = GetError()
	}
	return
}

// Set up an OpenGL context for rendering into an OpenGL window.
// (https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent)
func (window *Window) GL_MakeCurrent(context GLContext) (err error) {
	ret := C.SDL_GL_MakeCurrent(window.cptr(), C.SDL_GLContext(context))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the currently active OpenGL window.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow)
func GL_GetCurrentWindow() (window *Window, err error) {
	window = (*Window)(unsafe.Pointer(C.SDL_GL_GetCurrentWindow()))
	if window == nil {
		err = GetError()
	}
	return
}

// Get the currently active OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext)
func GL_GetCurrentContext() (context GLContext, err error) {
	context = GLContext(C.SDL_GL_GetCurrentContext())
	if context == nil {
		err = GetError()
	}
	return
}

// Get the currently active EGL display.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentEGLDisplay)
func EGL_GetCurrentDisplay() (display EGLDisplay, err error) {
	display = EGLDisplay(C.SDL_EGL_GetCurrentDisplay())
	if display == nil {
		err = GetError()
	}
	return
}

// Get the currently active EGL config.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig)
func EGL_GetCurrentConfig() (config EGLConfig, err error) {
	config = EGLConfig(C.SDL_EGL_GetCurrentConfig())
	if config == nil {
		err = GetError()
	}
	return
}

// Get the EGL surface associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface)
func (window *Window) EGL_GetWindowSurface() (surface EGLSurface, err error) {
	surface = EGLSurface(C.SDL_EGL_GetWindowSurface(window.cptr()))
	if surface == nil {
		err = GetError()
	}
	return
}

// Sets the callbacks for defining custom EGLAttrib arrays for EGL
// initialization.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks)
func EGL_SetEGLAttributeCallbacks(platformAttribCallback EGLAttribArrayCallback,
	surfaceAttribCallback,
	contextAttribCallback EGLIntArrayCallback) {
	C.SDL_EGL_SetAttributeCallbacks(
		C.SDL_EGLAttribArrayCallback(platformAttribCallback), C.SDL_EGLIntArrayCallback(surfaceAttribCallback),
		C.SDL_EGLIntArrayCallback(contextAttribCallback))
}

// Set the swap interval for the current OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval)
func GL_SetSwapInterval(interval int32) (err error) {
	ret := C.SDL_GL_SetSwapInterval(C.int(interval))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the swap interval for the current OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval)
func GL_GetSwapInterval() (interval int32, err error) {
	ret := C.SDL_GL_GetSwapInterval((*C.int)(&interval))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Update a window with OpenGL rendering.
// (https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow)
func (window *Window) GL_SwapWindow() (err error) {
	ret := C.SDL_GL_SwapWindow(window.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Delete an OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext)
func GL_DestroyContext(context GLContext) (err error) {
	ret := C.SDL_GL_DestroyContext(C.SDL_GLContext(context))
	if ret != 0 {
		err = GetError()
	}
	return
}
