package sdl

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
#include <SDL3/SDL_events.h>
#include "events.h"
*/
import "C"

// # CategoryEvents
// (https://wiki.libsdl.org/SDL3/CategoryEvents)

// The types of events that can be delivered.
// (https://wiki.libsdl.org/SDL3/SDL_EventType)
type EventType C.SDL_EventType

func (t EventType) c() C.Uint32 {
	return C.Uint32(t)
}

const (
	EVENT_FIRST = EventType(C.SDL_EVENT_FIRST) /**< Unused (do not remove) */

	/* Application events */
	EVENT_QUIT = EventType(C.SDL_EVENT_QUIT) /**< User-requested quit */

	/* These application events have special meaning on iOS, see README-ios.md for details */
	EVENT_TERMINATING = EventType(C.SDL_EVENT_TERMINATING) /**< The application is being terminated by the OS
	  Called on iOS in applicationWillTerminate()
	  Called on Android in onDestroy()
	*/
	EVENT_LOW_MEMORY = EventType(C.SDL_EVENT_LOW_MEMORY) /**< The application is low on memory, free memory if possible.
	  Called on iOS in applicationDidReceiveMemoryWarning()
	  Called on Android in onLowMemory()
	*/
	EVENT_WILL_ENTER_BACKGROUND = EventType(C.SDL_EVENT_WILL_ENTER_BACKGROUND) /**< The application is about to enter the background
	  Called on iOS in applicationWillResignActive()
	  Called on Android in onPause()
	*/
	EVENT_DID_ENTER_BACKGROUND = EventType(C.SDL_EVENT_DID_ENTER_BACKGROUND) /**< The application did enter the background and may not get CPU for some time
	  Called on iOS in applicationDidEnterBackground()
	  Called on Android in onPause()
	*/
	EVENT_WILL_ENTER_FOREGROUND = EventType(C.SDL_EVENT_WILL_ENTER_FOREGROUND) /**< The application is about to enter the foreground
	  Called on iOS in applicationWillEnterForeground()
	  Called on Android in onResume()
	*/
	EVENT_DID_ENTER_FOREGROUND = EventType(C.SDL_EVENT_DID_ENTER_FOREGROUND) /**< The application is now interactive
	  Called on iOS in applicationDidBecomeActive()
	  Called on Android in onResume()
	*/

	EVENT_LOCALE_CHANGED = EventType(C.SDL_EVENT_LOCALE_CHANGED) /**< The user's locale preferences have changed. */

	EVENT_SYSTEM_THEME_CHANGED = EventType(C.SDL_EVENT_SYSTEM_THEME_CHANGED) /**< The system theme changed */

	/* Display events */
	/* 0x150 was SDL_DISPLAYEVENT, reserve the number for sdl2-compat */
	EVENT_DISPLAY_ORIENTATION           = EventType(C.SDL_EVENT_DISPLAY_ORIENTATION)           /**< Display orientation has changed to data1 */
	EVENT_DISPLAY_ADDED                 = EventType(C.SDL_EVENT_DISPLAY_ADDED)                 /**< Display has been added to the system */
	EVENT_DISPLAY_REMOVED               = EventType(C.SDL_EVENT_DISPLAY_REMOVED)               /**< Display has been removed from the system */
	EVENT_DISPLAY_MOVED                 = EventType(C.SDL_EVENT_DISPLAY_MOVED)                 /**< Display has changed position */
	EVENT_DISPLAY_DESKTOP_MODE_CHANGED  = EventType(C.SDL_EVENT_DISPLAY_DESKTOP_MODE_CHANGED)  /**< Display has changed desktop mode */
	EVENT_DISPLAY_CURRENT_MODE_CHANGED  = EventType(C.SDL_EVENT_DISPLAY_CURRENT_MODE_CHANGED)  /**< Display has changed current mode */
	EVENT_DISPLAY_CONTENT_SCALE_CHANGED = EventType(C.SDL_EVENT_DISPLAY_CONTENT_SCALE_CHANGED) /**< Display has changed content scale */
	EVENT_DISPLAY_FIRST                 = EventType(C.SDL_EVENT_DISPLAY_FIRST)
	EVENT_DISPLAY_LAST                  = EventType(C.SDL_EVENT_DISPLAY_LAST)

	/* Window events */
	/* 0x200 was SDL_WINDOWEVENT, reserve the number for sdl2-compat */
	/* 0x201 was SDL_EVENT_SYSWM, reserve the number for sdl2-compat */
	EVENT_WINDOW_SHOWN                 = EventType(C.SDL_EVENT_WINDOW_SHOWN)                 /**< Window has been shown */
	EVENT_WINDOW_HIDDEN                = EventType(C.SDL_EVENT_WINDOW_HIDDEN)                /**< Window has been hidden */
	EVENT_WINDOW_EXPOSED               = EventType(C.SDL_EVENT_WINDOW_EXPOSED)               /**< Window has been exposed and should be redrawn, and can be redrawn directly from event watchers for this event */
	EVENT_WINDOW_MOVED                 = EventType(C.SDL_EVENT_WINDOW_MOVED)                 /**< Window has been moved to data1, data2 */
	EVENT_WINDOW_RESIZED               = EventType(C.SDL_EVENT_WINDOW_RESIZED)               /**< Window has been resized to data1xdata2 */
	EVENT_WINDOW_PIXEL_SIZE_CHANGED    = EventType(C.SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED)    /**< The pixel size of the window has changed to data1xdata2 */
	EVENT_WINDOW_METAL_VIEW_RESIZED    = EventType(C.SDL_EVENT_WINDOW_METAL_VIEW_RESIZED)    /**< The pixel size of a Metal view associated with the window has changed */
	EVENT_WINDOW_MINIMIZED             = EventType(C.SDL_EVENT_WINDOW_MINIMIZED)             /**< Window has been minimized */
	EVENT_WINDOW_MAXIMIZED             = EventType(C.SDL_EVENT_WINDOW_MAXIMIZED)             /**< Window has been maximized */
	EVENT_WINDOW_RESTORED              = EventType(C.SDL_EVENT_WINDOW_RESTORED)              /**< Window has been restored to normal size and position */
	EVENT_WINDOW_MOUSE_ENTER           = EventType(C.SDL_EVENT_WINDOW_MOUSE_ENTER)           /**< Window has gained mouse focus */
	EVENT_WINDOW_MOUSE_LEAVE           = EventType(C.SDL_EVENT_WINDOW_MOUSE_LEAVE)           /**< Window has lost mouse focus */
	EVENT_WINDOW_FOCUS_GAINED          = EventType(C.SDL_EVENT_WINDOW_FOCUS_GAINED)          /**< Window has gained keyboard focus */
	EVENT_WINDOW_FOCUS_LOST            = EventType(C.SDL_EVENT_WINDOW_FOCUS_LOST)            /**< Window has lost keyboard focus */
	EVENT_WINDOW_CLOSE_REQUESTED       = EventType(C.SDL_EVENT_WINDOW_CLOSE_REQUESTED)       /**< The window manager requests that the window be closed */
	EVENT_WINDOW_HIT_TEST              = EventType(C.SDL_EVENT_WINDOW_HIT_TEST)              /**< Window had a hit test that wasn't SDL_HITTEST_NORMAL */
	EVENT_WINDOW_ICCPROF_CHANGED       = EventType(C.SDL_EVENT_WINDOW_ICCPROF_CHANGED)       /**< The ICC profile of the window's display has changed */
	EVENT_WINDOW_DISPLAY_CHANGED       = EventType(C.SDL_EVENT_WINDOW_DISPLAY_CHANGED)       /**< Window has been moved to display data1 */
	EVENT_WINDOW_DISPLAY_SCALE_CHANGED = EventType(C.SDL_EVENT_WINDOW_DISPLAY_SCALE_CHANGED) /**< Window display scale has been changed */
	EVENT_WINDOW_SAFE_AREA_CHANGED     = EventType(C.SDL_EVENT_WINDOW_SAFE_AREA_CHANGED)     /**< The window safe area has been changed */
	EVENT_WINDOW_OCCLUDED              = EventType(C.SDL_EVENT_WINDOW_OCCLUDED)              /**< The window has been occluded */
	EVENT_WINDOW_ENTER_FULLSCREEN      = EventType(C.SDL_EVENT_WINDOW_ENTER_FULLSCREEN)      /**< The window has entered fullscreen mode */
	EVENT_WINDOW_LEAVE_FULLSCREEN      = EventType(C.SDL_EVENT_WINDOW_LEAVE_FULLSCREEN)      /**< The window has left fullscreen mode */
	EVENT_WINDOW_DESTROYED             = EventType(C.SDL_EVENT_WINDOW_DESTROYED)             /**< The window with the associated ID is being or has been destroyed.
	  If this message is being handled in an event watcher, the window handle is still valid
	  and can still be used to retrieve any userdata associated with the window. Otherwise,
	  the handle has already been destroyed and all resources associated with it are invalid */
	// EVENT_WINDOW_PEN_ENTER         = EventType(C.SDL_EVENT_WINDOW_PEN_ENTER)         /**< Window has gained focus of the pressure-sensitive pen with ID "data1" */
	// EVENT_WINDOW_PEN_LEAVE         = EventType(C.SDL_EVENT_WINDOW_PEN_LEAVE)         /**< Window has lost focus of the pressure-sensitive pen with ID "data1" */
	EVENT_WINDOW_HDR_STATE_CHANGED = EventType(C.SDL_EVENT_WINDOW_HDR_STATE_CHANGED) /**< Window HDR properties have changed */
	EVENT_WINDOW_FIRST             = EventType(C.SDL_EVENT_WINDOW_FIRST)
	EVENT_WINDOW_LAST              = EventType(C.SDL_EVENT_WINDOW_LAST)

	/* Keyboard events */
	EVENT_KEY_DOWN                = EventType(C.SDL_EVENT_KEY_DOWN)                /**< Key pressed */
	EVENT_KEY_UP                  = EventType(C.SDL_EVENT_KEY_UP)                  /**< Key released */
	EVENT_TEXT_EDITING            = EventType(C.SDL_EVENT_TEXT_EDITING)            /**< Keyboard text editing (composition) */
	EVENT_TEXT_INPUT              = EventType(C.SDL_EVENT_TEXT_INPUT)              /**< Keyboard text input */
	EVENT_KEYMAP_CHANGED          = EventType(C.SDL_EVENT_KEYMAP_CHANGED)          /**< Keymap changed due to a system event such as an input language or keyboard layout change. */
	EVENT_KEYBOARD_ADDED          = EventType(C.SDL_EVENT_KEYBOARD_ADDED)          /**< A new keyboard has been inserted into the system */
	EVENT_KEYBOARD_REMOVED        = EventType(C.SDL_EVENT_KEYBOARD_REMOVED)        /**< A keyboard has been removed */
	EVENT_TEXT_EDITING_CANDIDATES = EventType(C.SDL_EVENT_TEXT_EDITING_CANDIDATES) /**< Keyboard text editing candidates */

	/* Mouse events */
	EVENT_MOUSE_MOTION      = EventType(C.SDL_EVENT_MOUSE_MOTION)      /**< Mouse moved */
	EVENT_MOUSE_BUTTON_DOWN = EventType(C.SDL_EVENT_MOUSE_BUTTON_DOWN) /**< Mouse button pressed */
	EVENT_MOUSE_BUTTON_UP   = EventType(C.SDL_EVENT_MOUSE_BUTTON_UP)   /**< Mouse button released */
	EVENT_MOUSE_WHEEL       = EventType(C.SDL_EVENT_MOUSE_WHEEL)       /**< Mouse wheel motion */
	EVENT_MOUSE_ADDED       = EventType(C.SDL_EVENT_MOUSE_ADDED)       /**< A new mouse has been inserted into the system */
	EVENT_MOUSE_REMOVED     = EventType(C.SDL_EVENT_MOUSE_REMOVED)     /**< A mouse has been removed */

	/* Joystick events */
	EVENT_JOYSTICK_AXIS_MOTION     = EventType(C.SDL_EVENT_JOYSTICK_AXIS_MOTION)     /**< Joystick axis motion */
	EVENT_JOYSTICK_BALL_MOTION     = EventType(C.SDL_EVENT_JOYSTICK_BALL_MOTION)     /**< Joystick trackball motion */
	EVENT_JOYSTICK_HAT_MOTION      = EventType(C.SDL_EVENT_JOYSTICK_HAT_MOTION)      /**< Joystick hat position change */
	EVENT_JOYSTICK_BUTTON_DOWN     = EventType(C.SDL_EVENT_JOYSTICK_BUTTON_DOWN)     /**< Joystick button pressed */
	EVENT_JOYSTICK_BUTTON_UP       = EventType(C.SDL_EVENT_JOYSTICK_BUTTON_UP)       /**< Joystick button released */
	EVENT_JOYSTICK_ADDED           = EventType(C.SDL_EVENT_JOYSTICK_ADDED)           /**< A new joystick has been inserted into the system */
	EVENT_JOYSTICK_REMOVED         = EventType(C.SDL_EVENT_JOYSTICK_REMOVED)         /**< An opened joystick has been removed */
	EVENT_JOYSTICK_BATTERY_UPDATED = EventType(C.SDL_EVENT_JOYSTICK_BATTERY_UPDATED) /**< Joystick battery level change */
	EVENT_JOYSTICK_UPDATE_COMPLETE = EventType(C.SDL_EVENT_JOYSTICK_UPDATE_COMPLETE) /**< Joystick update is complete */

	/* Gamepad events */
	EVENT_GAMEPAD_AXIS_MOTION          = EventType(C.SDL_EVENT_GAMEPAD_AXIS_MOTION)          /**< Gamepad axis motion */
	EVENT_GAMEPAD_BUTTON_DOWN          = EventType(C.SDL_EVENT_GAMEPAD_BUTTON_DOWN)          /**< Gamepad button pressed */
	EVENT_GAMEPAD_BUTTON_UP            = EventType(C.SDL_EVENT_GAMEPAD_BUTTON_UP)            /**< Gamepad button released */
	EVENT_GAMEPAD_ADDED                = EventType(C.SDL_EVENT_GAMEPAD_ADDED)                /**< A new gamepad has been inserted into the system */
	EVENT_GAMEPAD_REMOVED              = EventType(C.SDL_EVENT_GAMEPAD_REMOVED)              /**< A gamepad has been removed */
	EVENT_GAMEPAD_REMAPPED             = EventType(C.SDL_EVENT_GAMEPAD_REMAPPED)             /**< The gamepad mapping was updated */
	EVENT_GAMEPAD_TOUCHPAD_DOWN        = EventType(C.SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN)        /**< Gamepad touchpad was touched */
	EVENT_GAMEPAD_TOUCHPAD_MOTION      = EventType(C.SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION)      /**< Gamepad touchpad finger was moved */
	EVENT_GAMEPAD_TOUCHPAD_UP          = EventType(C.SDL_EVENT_GAMEPAD_TOUCHPAD_UP)          /**< Gamepad touchpad finger was lifted */
	EVENT_GAMEPAD_SENSOR_UPDATE        = EventType(C.SDL_EVENT_GAMEPAD_SENSOR_UPDATE)        /**< Gamepad sensor was updated */
	EVENT_GAMEPAD_UPDATE_COMPLETE      = EventType(C.SDL_EVENT_GAMEPAD_UPDATE_COMPLETE)      /**< Gamepad update is complete */
	EVENT_GAMEPAD_STEAM_HANDLE_UPDATED = EventType(C.SDL_EVENT_GAMEPAD_STEAM_HANDLE_UPDATED) /**< Gamepad Steam handle has changed */

	/* Touch events */
	EVENT_FINGER_DOWN   = EventType(C.SDL_EVENT_FINGER_DOWN)
	EVENT_FINGER_UP     = EventType(C.SDL_EVENT_FINGER_UP)
	EVENT_FINGER_MOTION = EventType(C.SDL_EVENT_FINGER_MOTION)

	/* 0x800, 0x801, and 0x802 were the Gesture events from SDL2. Do not reuse these values! sdl2-compat needs them! */

	/* Clipboard events */
	EVENT_CLIPBOARD_UPDATE = EventType(C.SDL_EVENT_CLIPBOARD_UPDATE) /**< The clipboard or primary selection changed */

	/* Drag and drop events */
	EVENT_DROP_FILE     = EventType(C.SDL_EVENT_DROP_FILE)     /**< The system requests a file open */
	EVENT_DROP_TEXT     = EventType(C.SDL_EVENT_DROP_TEXT)     /**< text/plain drag-and-drop event */
	EVENT_DROP_BEGIN    = EventType(C.SDL_EVENT_DROP_BEGIN)    /**< A new set of drops is beginning (NULL filename) */
	EVENT_DROP_COMPLETE = EventType(C.SDL_EVENT_DROP_COMPLETE) /**< Current set of drops is now complete (NULL filename) */
	EVENT_DROP_POSITION = EventType(C.SDL_EVENT_DROP_POSITION) /**< Position while moving over the window */

	/* Audio hotplug events */
	EVENT_AUDIO_DEVICE_ADDED          = EventType(C.SDL_EVENT_AUDIO_DEVICE_ADDED)          /**< A new audio device is available */
	EVENT_AUDIO_DEVICE_REMOVED        = EventType(C.SDL_EVENT_AUDIO_DEVICE_REMOVED)        /**< An audio device has been removed. */
	EVENT_AUDIO_DEVICE_FORMAT_CHANGED = EventType(C.SDL_EVENT_AUDIO_DEVICE_FORMAT_CHANGED) /**< An audio device's format has been changed by the system. */

	/* Sensor events */
	EVENT_SENSOR_UPDATE = EventType(C.SDL_EVENT_SENSOR_UPDATE) /**< A sensor was updated */

	/* Pressure-sensitive pen events */
	EVENT_PEN_PROXIMITY_IN  = EventType(C.SDL_EVENT_PEN_PROXIMITY_IN)  /**< Pressure-sensitive pen has become available */
	EVENT_PEN_PROXIMITY_OUT = EventType(C.SDL_EVENT_PEN_PROXIMITY_OUT) /**< Pressure-sensitive pen has become unavailable */
	EVENT_PEN_DOWN          = EventType(C.SDL_EVENT_PEN_DOWN)          /**< Pressure-sensitive pen touched drawing surface */
	EVENT_PEN_UP            = EventType(C.SDL_EVENT_PEN_UP)            /**< Pressure-sensitive pen stopped touching drawing surface */
	EVENT_PEN_BUTTON_DOWN   = EventType(C.SDL_EVENT_PEN_BUTTON_DOWN)   /**< Pressure-sensitive pen button pressed */
	EVENT_PEN_BUTTON_UP     = EventType(C.SDL_EVENT_PEN_BUTTON_UP)     /**< Pressure-sensitive pen button released */
	EVENT_PEN_MOTION        = EventType(C.SDL_EVENT_PEN_MOTION)        /**< Pressure-sensitive pen is moving on the tablet */
	EVENT_PEN_AXIS          = EventType(C.SDL_EVENT_PEN_AXIS)          /**< Pressure-sensitive pen angle/pressure/etc changed */

	/* Camera hotplug events */
	EVENT_CAMERA_DEVICE_ADDED    = EventType(C.SDL_EVENT_CAMERA_DEVICE_ADDED)    /**< A new camera device is available */
	EVENT_CAMERA_DEVICE_REMOVED  = EventType(C.SDL_EVENT_CAMERA_DEVICE_REMOVED)  /**< A camera device has been removed. */
	EVENT_CAMERA_DEVICE_APPROVED = EventType(C.SDL_EVENT_CAMERA_DEVICE_APPROVED) /**< A camera device has been approved for use by the user. */
	EVENT_CAMERA_DEVICE_DENIED   = EventType(C.SDL_EVENT_CAMERA_DEVICE_DENIED)   /**< A camera device has been denied for use by the user. */

	/* Render events */
	EVENT_RENDER_TARGETS_RESET = EventType(C.SDL_EVENT_RENDER_TARGETS_RESET) /**< The render targets have been reset and their contents need to be updated */
	EVENT_RENDER_DEVICE_RESET  = EventType(C.SDL_EVENT_RENDER_DEVICE_RESET)  /**< The device has been reset and all textures need to be recreated */
	EVENT_RENDER_DEVICE_LOST   = EventType(C.SDL_EVENT_RENDER_DEVICE_LOST)   /**< The device has been lost and can't be recovered. */

	/* Internal events */
	EVENT_POLL_SENTINEL = EventType(C.SDL_EVENT_POLL_SENTINEL) /**< Signals the end of an event poll cycle */

	/** Events SDL_EVENT_USER through SDL_EVENT_LAST are for your use, and should be allocated with SDL_RegisterEvents()*/
	EVENT_USER = EventType(C.SDL_EVENT_USER)

	/** This last event is only for bounding internal arrays */
	EVENT_LAST = EventType(C.SDL_EVENT_LAST)

	/* This just makes sure the enum is the size of Uint32 */
	// EVENT_ENUM_PADDING = EventType(C.SDL_EVENT_ENUM_PADDING)
)

// Fields shared by every event
// (https://wiki.libsdl.org/SDL3/SDL_CommonEvent)
type CommonEvent struct {
	Type      EventType /**< Event type, shared with all events, Uint32 to cover user events which are not in the SDL_EventType enumeration */
	_         uint32
	Timestamp uint64 /**< In nanoseconds, populated using SDL_GetTicksNS() */
}

func (e *CommonEvent) GetType() EventType   { return e.Type }
func (e *CommonEvent) GetTimestamp() uint64 { return e.Timestamp }
func (e *CommonEvent) cptr() *C.SDL_Event   { return (*C.SDL_Event)(unsafe.Pointer(e)) }

// Display state change event data (event.display.*)
// (https://wiki.libsdl.org/SDL3/SDL_DisplayEvent)
type DisplayEvent struct {
	CommonEvent
	DisplayID DisplayID /**< The associated display */
	Data1     int32     /**< event dependent data */
	Data2     int32     /**< event dependent data */
}

// Window state change event data (event.window.*)
// (https://wiki.libsdl.org/SDL3/SDL_WindowEvent)
type WindowEvent struct {
	CommonEvent
	WindowID WindowID /**< The associated window */
	Data1    int32    /**< event dependent data */
	Data2    int32    /**< event dependent data */
}

// Keyboard device event structure (event.kdevice.*)
// (https://wiki.libsdl.org/SDL3/SDL_KeyboardDeviceEvent)
type KeyboardDeviceEvent struct {
	CommonEvent
	Which KeyboardID /**< The keyboard instance id */
}

// Keyboard button event structure (event.key.*)
// (https://wiki.libsdl.org/SDL3/SDL_KeyboardEvent)
type KeyboardEvent struct {
	CommonEvent
	WindowID WindowID   /**< The window with keyboard focus, if any */
	Which    KeyboardID /**< The keyboard instance id, or 0 if unknown or virtual */
	Scancode Scancode   /**< SDL physical key code */
	Key      Keycode    /**< SDL virtual key code */
	Mod      Keymod     /**< current key modifiers */
	Raw      uint16     /**< The platform dependent scancode for this event */
	Down     bool       /**< SDL_PRESSED or SDL_RELEASED */
	Repeat   bool       /**< Non-zero if this is a key repeat */
}

// Keyboard text editing event structure (event.edit.*)
// (https://wiki.libsdl.org/SDL3/SDL_TextEditingEvent)
type TextEditingEvent struct {
	CommonEvent
	WindowID WindowID /**< The window with keyboard focus, if any */
	Text     string   /**< The editing text */
	Start    int32    /**< The start cursor of selected editing text, or -1 if not set */
	Length   int32    /**< The length of selected editing text, or -1 if not set */
}
type cTextEditingEvent C.SDL_TextEditingEvent

// Keyboard IME candidates event structure (event.edit_candidates.*)
// (https://wiki.libsdl.org/SDL3/SDL_TextEditingCandidatesEvent)
type TextEditingCandidatesEvent struct {
	CommonEvent
	WindowID   WindowID /**< The window with keyboard focus, if any */
	Candidates []string /**< The list of candidates, or NULL if there are no candidates available */
	// Num_candidates     int32    /**< The number of strings in `candidates` */
	SelectedCandidate int32 /**< The index of the selected candidate, or -1 if no candidate is selected */
	Horizontal        bool  /**< SDL_TRUE if the list is horizontal, SDL_FALSE if it's vertical */
}
type cTextEditingCandidatesEvent C.SDL_TextEditingCandidatesEvent

// Keyboard text input event structure (event.text.*)
// (https://wiki.libsdl.org/SDL3/SDL_TextInputEvent)
type TextInputEvent struct {
	CommonEvent
	WindowID WindowID /**< The window with keyboard focus, if any */
	Text     string   /**< The input text, UTF-8 encoded */
}
type cTextInputEvent C.SDL_TextInputEvent

// Mouse device event structure (event.mdevice.*)
// (https://wiki.libsdl.org/SDL3/SDL_MouseDeviceEvent)
type MouseDeviceEvent struct {
	CommonEvent
	Which MouseID /**< The mouse instance id */
}

// Mouse motion event structure (event.motion.*)
// (https://wiki.libsdl.org/SDL3/SDL_MouseMotionEvent)
type MouseMotionEvent struct {
	CommonEvent
	WindowID WindowID         /**< The window with mouse focus, if any */
	Which    MouseID          /**< The mouse instance id, SDL_TOUCH_MOUSEID, or SDL_PEN_MOUSEID */
	State    MouseButtonFlags /**< The current button state */
	X        float32          /**< X coordinate, relative to window */
	Y        float32          /**< Y coordinate, relative to window */
	Xrel     float32          /**< The relative motion in the X direction */
	Yrel     float32          /**< The relative motion in the Y direction */
}

// Mouse button event structure (event.button.*)
// (https://wiki.libsdl.org/SDL3/SDL_MouseButtonEvent)
type MouseButtonEvent struct {
	CommonEvent
	WindowID WindowID /**< The window with mouse focus, if any */
	Which    MouseID  /**< The mouse instance id, SDL_TOUCH_MOUSEID, or SDL_PEN_MOUSEID */
	Button   uint8    /**< The mouse button index */
	Down     bool     /**< true if the button is pressed */
	Clicks   uint8    /**< 1 for single-click, 2 for double-click, etc. */
	_        uint8
	X        float32 /**< X coordinate, relative to window */
	Y        float32 /**< Y coordinate, relative to window */
}

// Mouse wheel event structure (event.wheel.*)
// (https://wiki.libsdl.org/SDL3/SDL_MouseWheelEvent)
type MouseWheelEvent struct {
	CommonEvent
	WindowID  WindowID            /**< The window with mouse focus, if any */
	Which     MouseID             /**< The mouse instance id, SDL_TOUCH_MOUSEID, or SDL_PEN_MOUSEID */
	X         float32             /**< The amount scrolled horizontally, positive to the right and negative to the left */
	Y         float32             /**< The amount scrolled vertically, positive away from the user and negative toward the user */
	Direction MouseWheelDirection /**< Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back */
	Mouse_x   float32             /**< X coordinate, relative to window */
	Mouse_y   float32             /**< Y coordinate, relative to window */
}

// Joystick axis motion event structure (event.jaxis.*)
// (https://wiki.libsdl.org/SDL3/SDL_JoyAxisEvent)
type JoyAxisEvent struct {
	CommonEvent
	Which JoystickID /**< The joystick instance id */
	Axis  uint8      /**< The joystick axis index */
	_     uint8      // padding
	_     uint8      // padding
	_     uint8      // padding
	Value int16      /**< The axis value (range: -32768 to 32767) */
	_     uint16     // padding
}

// Joystick trackball motion event structure (event.jball.*)
// (https://wiki.libsdl.org/SDL3/SDL_JoyBallEvent)
type JoyBallEvent struct {
	CommonEvent
	Which JoystickID /**< The joystick instance id */
	Ball  uint8      /**< The joystick trackball index */
	_     uint8      // padding
	_     uint8      // padding
	_     uint8      // padding
	Xrel  int16      /**< The relative motion in the X direction */
	Yrel  int16      /**< The relative motion in the Y direction */
}

// Joystick hat position change event structure (event.jhat.*)
// (https://wiki.libsdl.org/SDL3/SDL_JoyHatEvent)
type JoyHatEvent struct {
	CommonEvent
	Which JoystickID /**< The joystick instance id */
	Hat   uint8      /**< The joystick hat index */
	Value uint8      /**< The hat position value. */
	_     uint8      // padding
	_     uint8      // padding
}

// Joystick button event structure (event.jbutton.*)
// (https://wiki.libsdl.org/SDL3/SDL_JoyButtonEvent)
type JoyButtonEvent struct {
	CommonEvent
	Which  JoystickID /**< The joystick instance id */
	Button uint8      /**< The joystick button index */
	Down   bool       /**< true if the button is pressed */
	_      uint8      // padding
	_      uint8      // padding
}

// Joystick device event structure (event.jdevice.*)
// (https://wiki.libsdl.org/SDL3/SDL_JoyDeviceEvent)
type JoyDeviceEvent struct {
	CommonEvent
	Which JoystickID /**< The joystick instance id */
}

// Joysick battery level change event structure (event.jbattery.*)
// (https://wiki.libsdl.org/SDL3/SDL_JoyBatteryEvent)
type JoyBatteryEvent struct {
	CommonEvent
	Which   JoystickID /**< The joystick instance id */
	State   PowerState /**< The joystick battery state */
	Percent int32      /**< The joystick battery percent charge remaining */
}

// Gamepad axis motion event structure (event.gaxis.*)
// (https://wiki.libsdl.org/SDL3/SDL_GamepadAxisEvent)
type GamepadAxisEvent struct {
	CommonEvent
	Which JoystickID /**< The joystick instance id */
	Axis  uint8      /**< The gamepad axis (SDL_GamepadAxis) */
	_     uint8      // padding
	_     uint8      // padding
	_     uint8      // padding
	Value int16      /**< The axis value (range: -32768 to 32767) */
	_     uint16     // padding
}

// Gamepad button event structure (event.gbutton.*)
// (https://wiki.libsdl.org/SDL3/SDL_GamepadButtonEvent)
type GamepadButtonEvent struct {
	CommonEvent
	Which  JoystickID /**< The joystick instance id */
	Button uint8      /**< The gamepad button (SDL_GamepadButton) */
	Down   bool       /**< true if the button is pressed */
	_      uint8      // padding
	_      uint8      // padding
}

// Gamepad device event structure (event.gdevice.*)
// (https://wiki.libsdl.org/SDL3/SDL_GamepadDeviceEvent)
type GamepadDeviceEvent struct {
	CommonEvent
	Which JoystickID /**< The joystick instance id */
}

// Gamepad touchpad event structure (event.gtouchpad.*)
// (https://wiki.libsdl.org/SDL3/SDL_GamepadTouchpadEvent)
type GamepadTouchpadEvent struct {
	CommonEvent
	Which    JoystickID /**< The joystick instance id */
	Touchpad int32      /**< The index of the touchpad */
	Finger   int32      /**< The index of the finger on the touchpad */
	X        float32    /**< Normalized in the range 0...1 with 0 being on the left */
	Y        float32    /**< Normalized in the range 0...1 with 0 being at the top */
	Pressure float32    /**< Normalized in the range 0...1 */
}

// Gamepad sensor event structure (event.gsensor.*)
// (https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEvent)
type GamepadSensorEvent struct {
	CommonEvent
	Which           JoystickID /**< The joystick instance id */
	Sensor          int32      /**< The type of the sensor, one of the values of SDL_SensorType */
	Data            [3]float32 /**< Up to 3 values from the sensor, as defined in SDL_sensor.h */
	SensorTimestamp uint64     /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
}

// Audio device event structure (event.adevice.*)
// (https://wiki.libsdl.org/SDL3/SDL_AudioDeviceEvent)
type AudioDeviceEvent struct {
	CommonEvent
	Which     AudioDeviceID /**< SDL_AudioDeviceID for the device being added or removed or changing */
	Recording bool          /**< false if a playback device, true if a recording device. */
	_         uint8         // padding
	_         uint8         // padding
	_         uint8         // padding
}

// Camera device event structure (event.cdevice.*)
// (https://wiki.libsdl.org/SDL3/SDL_CameraDeviceEvent)
type CameraDeviceEvent struct {
	CommonEvent
	Which CameraID /**< SDL_CameraID for the device being added or removed or changing */
}

// Touch finger event structure (event.tfinger.*)
// (https://wiki.libsdl.org/SDL3/SDL_TouchFingerEvent)
type TouchFingerEvent struct {
	CommonEvent
	TouchID  TouchID /**< The touch device id */
	FingerID FingerID
	X        float32  /**< Normalized in the range 0...1 */
	Y        float32  /**< Normalized in the range 0...1 */
	Dx       float32  /**< Normalized in the range -1...1 */
	Dy       float32  /**< Normalized in the range -1...1 */
	Pressure float32  /**< Normalized in the range 0...1 */
	WindowID WindowID /**< The window underneath the finger, if any */
}

// Pressure-sensitive pen proximity event structure (event.pmotion.*)
// (https://wiki.libsdl.org/SDL3/SDL_PenProximityEvent)
type PenProximityEvent struct {
	CommonEvent
	WindowID WindowID /**< The window with mouse focus, if any */
	Which    PenID    /**< The pen instance id */
}

// Pressure-sensitive pen motion event structure (event.pmotion.*)
// (https://wiki.libsdl.org/SDL3/SDL_PenMotionEvent)
type PenMotionEvent struct {
	CommonEvent
	WindowID WindowID      /**< The window with mouse focus, if any */
	Which    PenID         /**< The pen instance id */
	PenState PenInputFlags /**< Complete pen input state at time of event */
	X        float32       /**< X coordinate, relative to window */
	Y        float32       /**< Y coordinate, relative to window */
}

// Pressure-sensitive pen touched event structure (event.ptouch.*)
// (https://wiki.libsdl.org/SDL3/SDL_PenTouchEvent)
type PenTouchEvent struct {
	CommonEvent
	WindowID WindowID      /**< The window with mouse focus, if any */
	Which    PenID         /**< The pen instance id */
	PenState PenInputFlags /**< Complete pen input state at time of event */
	X        float32       /**< X coordinate, relative to window */
	Y        float32       /**< Y coordinate, relative to window */
	Eraser   bool          /**< true if eraser end is used (not all pens support this). */
	Down     bool          /**< true if the pen is touching or false if the pen is lifted off */
}

// Pressure-sensitive pen button event structure (event.pbutton.*)
// (https://wiki.libsdl.org/SDL3/SDL_PenButtonEvent)
type PenButtonEvent struct {
	CommonEvent
	WindowID WindowID      /**< The window with mouse focus, if any */
	Which    PenID         /**< The pen instance id */
	PenState PenInputFlags /**< Complete pen input state at time of event */
	X        float32       /**< X coordinate, relative to window */
	Y        float32       /**< Y coordinate, relative to window */
	Button   uint8         /**< The pen button index (first button is 1). */
	Down     bool          /**< true if the button is pressed */
}

// Pressure-sensitive pen pressure / angle event structure (event.paxis.*)
// (https://wiki.libsdl.org/SDL3/SDL_PenAxisEvent)
type PenAxisEvent struct {
	CommonEvent
	WindowID WindowID      /**< The window with mouse focus, if any */
	Which    PenID         /**< The pen instance id */
	PenState PenInputFlags /**< Complete pen input state at time of event */
	X        float32       /**< X coordinate, relative to window */
	Y        float32       /**< Y coordinate, relative to window */
	Axis     PenAxis       /**< Axis that has changed */
	Value    float32       /**< New value of axis */
}

// An event used to drop text or request a file open by the system// (event.drop.*)
// (https://wiki.libsdl.org/SDL3/SDL_DropEvent)
type DropEvent struct {
	CommonEvent
	WindowID WindowID /**< The window that was dropped on, if any */
	X        float32  /**< X coordinate, relative to window (not on begin) */
	Y        float32  /**< Y coordinate, relative to window (not on begin) */
	Source   string   /**< The source app that sent this drop event, or NULL if that isn't available */
	Data     string   /**< The text for SDL_EVENT_DROP_TEXT and the file name for SDL_EVENT_DROP_FILE, NULL for other events */
}
type cDropEvent C.SDL_DropEvent

// An event triggered when the clipboard contents have changed// (event.clipboard.*)
// (https://wiki.libsdl.org/SDL3/SDL_ClipboardEvent)
type ClipboardEvent struct {
	CommonEvent
	Owner bool /**< are we owning the clipboard (internal update) */
	// NumMimeTypes int32    /**< number of mime types */ /**< current mime types */
	MimeTypes []string /**< current mime types */
}
type cClipboardEvent C.SDL_ClipboardEvent

// Sensor event structure (event.sensor.*)
// (https://wiki.libsdl.org/SDL3/SDL_SensorEvent)
type SensorEvent struct {
	CommonEvent
	Which           SensorID   /**< The instance ID of the sensor */
	Data            [6]float32 /**< Up to 6 values from the sensor - additional values can be queried using SDL_GetSensorData() */
	SensorTimestamp uint64     /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
}

// The "quit requested" event
// (https://wiki.libsdl.org/SDL3/SDL_QuitEvent)
type QuitEvent struct {
	CommonEvent
}

// A user-defined event type (event.user.*)
// (https://wiki.libsdl.org/SDL3/SDL_UserEvent)
type UserEvent struct {
	CommonEvent
	WindowID WindowID       /**< The associated window if any */
	Code     int32          /**< User defined event code */
	Data1    unsafe.Pointer /**< User defined data pointer */
	Data2    unsafe.Pointer /**< User defined data pointer */
}

// Added event types
type AppEvent struct {
	CommonEvent
}
type SystemEvent struct {
	CommonEvent
}
type KeymapChangedEvent struct {
	CommonEvent
}
type RenderEvent struct {
	CommonEvent
}
type UnknownEvent struct {
	CommonEvent
}

type cEvent C.SDL_Event

func (e *cEvent) cptr() *C.SDL_Event {
	return (*C.SDL_Event)(e)
}

// TODO: Any event that has strings or arrays in it will need to be handled differently.
func (e *cEvent) goEvent() Event {
	t := (*CommonEvent)(unsafe.Pointer(e)).GetType()
	switch {
	case t == EVENT_QUIT:
		return (*QuitEvent)(unsafe.Pointer(e))
	case t >= EVENT_TERMINATING && t <= EVENT_DID_ENTER_FOREGROUND:
		return (*AppEvent)(unsafe.Pointer(e))
	case t == EVENT_LOCALE_CHANGED || t == EVENT_SYSTEM_THEME_CHANGED:
		return (*SystemEvent)(unsafe.Pointer(e))
	case t >= EVENT_DISPLAY_FIRST && t <= EVENT_DISPLAY_LAST:
		return (*DisplayEvent)(unsafe.Pointer(e))
	case t >= EVENT_WINDOW_FIRST && t <= EVENT_WINDOW_LAST:
		return (*WindowEvent)(unsafe.Pointer(e))
	case t == EVENT_KEY_DOWN || t == EVENT_KEY_UP:
		return (*KeyboardEvent)(unsafe.Pointer(e))
	case t == EVENT_TEXT_EDITING:
		// TODO: has string in struct, need tested
		cEvt := (*cTextEditingEvent)(unsafe.Pointer(e))
		evt := new(TextEditingEvent)
		evt.CommonEvent = *(*CommonEvent)(unsafe.Pointer(e))
		evt.WindowID = WindowID(cEvt.windowID)
		evt.Text = C.GoString(cEvt.text)
		evt.Start = int32(cEvt.start)
		evt.Length = int32(cEvt.length)
		return evt
	case t == EVENT_TEXT_INPUT:
		// TODO: has string in struct, need tested
		cEvt := (*cTextInputEvent)(unsafe.Pointer(e))
		evt := new(TextInputEvent)
		evt.CommonEvent = *(*CommonEvent)(unsafe.Pointer(e))
		evt.WindowID = WindowID(cEvt.windowID)
		evt.Text = C.GoString(cEvt.text)
		return evt
	case t == EVENT_KEYMAP_CHANGED:
		return (*KeymapChangedEvent)(unsafe.Pointer(e))
	case t == EVENT_KEYBOARD_ADDED || t == EVENT_KEYBOARD_REMOVED:
		return (*KeyboardDeviceEvent)(unsafe.Pointer(e))
	case t == EVENT_TEXT_EDITING_CANDIDATES:
		// TODO: has an array of strings in struct, need tested
		cEvt := (*cTextEditingCandidatesEvent)(unsafe.Pointer(e))
		evt := new(TextEditingCandidatesEvent)
		evt.CommonEvent = *(*CommonEvent)(unsafe.Pointer(e))
		evt.WindowID = WindowID(cEvt.windowID)
		cSlice := unsafe.Slice(cEvt.candidates, cEvt.num_candidates)
		evt.Candidates = make([]string, len(cSlice))
		for i, c := range cSlice {
			evt.Candidates[i] = C.GoString(c)
		}
		evt.SelectedCandidate = int32(cEvt.selected_candidate)
		evt.Horizontal = bool(cEvt.horizontal)
		return evt
	case t == EVENT_MOUSE_MOTION:
		return (*MouseMotionEvent)(unsafe.Pointer(e))
	case t == EVENT_MOUSE_BUTTON_DOWN || t == EVENT_MOUSE_BUTTON_UP:
		return (*MouseButtonEvent)(unsafe.Pointer(e))
	case t == EVENT_MOUSE_WHEEL:
		return (*MouseWheelEvent)(unsafe.Pointer(e))
	case t == EVENT_MOUSE_ADDED || t == EVENT_MOUSE_REMOVED:
		return (*MouseDeviceEvent)(unsafe.Pointer(e))
	case t == EVENT_JOYSTICK_AXIS_MOTION:
		return (*JoyAxisEvent)(unsafe.Pointer(e))
	case t == EVENT_JOYSTICK_BALL_MOTION:
		return (*JoyBallEvent)(unsafe.Pointer(e))
	case t == EVENT_JOYSTICK_HAT_MOTION:
		return (*JoyHatEvent)(unsafe.Pointer(e))
	case t == EVENT_JOYSTICK_BUTTON_DOWN || t == EVENT_JOYSTICK_BUTTON_UP:
		return (*JoyButtonEvent)(unsafe.Pointer(e))
	case t == EVENT_JOYSTICK_ADDED || t == EVENT_JOYSTICK_REMOVED || t == EVENT_JOYSTICK_UPDATE_COMPLETE:
		return (*JoyDeviceEvent)(unsafe.Pointer(e))
	case t == EVENT_JOYSTICK_BATTERY_UPDATED:
		return (*JoyBatteryEvent)(unsafe.Pointer(e))
	case t == EVENT_GAMEPAD_AXIS_MOTION:
		return (*GamepadAxisEvent)(unsafe.Pointer(e))
	case t == EVENT_GAMEPAD_BUTTON_DOWN || t == EVENT_GAMEPAD_BUTTON_UP:
		return (*GamepadButtonEvent)(unsafe.Pointer(e))
	case t == EVENT_GAMEPAD_ADDED || t == EVENT_GAMEPAD_REMOVED || t == EVENT_GAMEPAD_REMAPPED ||
		t == EVENT_GAMEPAD_UPDATE_COMPLETE || t == EVENT_GAMEPAD_STEAM_HANDLE_UPDATED:
		return (*GamepadDeviceEvent)(unsafe.Pointer(e))
	case t == EVENT_GAMEPAD_TOUCHPAD_DOWN || t == EVENT_GAMEPAD_TOUCHPAD_MOTION || t == EVENT_GAMEPAD_TOUCHPAD_UP:
		return (*GamepadTouchpadEvent)(unsafe.Pointer(e))
	case t == EVENT_GAMEPAD_SENSOR_UPDATE:
		return (*GamepadSensorEvent)(unsafe.Pointer(e))
	case t == EVENT_FINGER_DOWN || t == EVENT_FINGER_UP || t == EVENT_FINGER_MOTION:
		return (*TouchFingerEvent)(unsafe.Pointer(e))
	case t == EVENT_CLIPBOARD_UPDATE:
		// TODO: has an array of strings in struct, need tested
		cEvt := (*cClipboardEvent)(unsafe.Pointer(e))
		evt := new(ClipboardEvent)
		evt.CommonEvent = *(*CommonEvent)(unsafe.Pointer(e))
		evt.Owner = bool(cEvt.owner)
		cSlice := unsafe.Slice(cEvt.mime_types, cEvt.num_mime_types)
		evt.MimeTypes = make([]string, len(cSlice))
		for i, c := range cSlice {
			evt.MimeTypes[i] = C.GoString(c)
		}
		return evt
	case t >= EVENT_DROP_FILE && t <= EVENT_DROP_POSITION:
		// TODO: has strings in struct, need tested
		cEvt := (*cDropEvent)(unsafe.Pointer(e))
		evt := new(DropEvent)
		evt.CommonEvent = *(*CommonEvent)(unsafe.Pointer(e))
		evt.WindowID = WindowID(cEvt.windowID)
		evt.X = float32(cEvt.x)
		evt.Y = float32(cEvt.y)
		evt.Source = C.GoString(cEvt.source)
		evt.Data = C.GoString(cEvt.data)
		return evt
	case t >= EVENT_AUDIO_DEVICE_ADDED && t <= EVENT_AUDIO_DEVICE_FORMAT_CHANGED:
		return (*AudioDeviceEvent)(unsafe.Pointer(e))
	case t == EVENT_SENSOR_UPDATE:
		return (*SensorEvent)(unsafe.Pointer(e))
	case t == EVENT_PEN_PROXIMITY_IN || t == EVENT_PEN_PROXIMITY_OUT:
		return (*PenProximityEvent)(unsafe.Pointer(e))
	case t == EVENT_PEN_DOWN || t == EVENT_PEN_UP:
		return (*PenTouchEvent)(unsafe.Pointer(e))
	case t == EVENT_PEN_BUTTON_DOWN || t == EVENT_PEN_BUTTON_UP:
		return (*PenButtonEvent)(unsafe.Pointer(e))
	case t == EVENT_PEN_MOTION:
		return (*PenMotionEvent)(unsafe.Pointer(e))
	case t == EVENT_PEN_AXIS:
		return (*PenAxisEvent)(unsafe.Pointer(e))
	case t >= EVENT_CAMERA_DEVICE_ADDED && t <= EVENT_CAMERA_DEVICE_DENIED:
		return (*CameraDeviceEvent)(unsafe.Pointer(e))
	case t == EVENT_RENDER_TARGETS_RESET || t == EVENT_RENDER_DEVICE_RESET || t == EVENT_RENDER_DEVICE_LOST:
		return (*RenderEvent)(unsafe.Pointer(e))
	case t >= EVENT_USER && t <= EVENT_LAST:
		return (*UserEvent)(unsafe.Pointer(e))
	default:
		return (*UnknownEvent)(unsafe.Pointer(e))
	}
}

type Event interface {
	GetType() EventType
	GetTimestamp() uint64
	cptr() *C.SDL_Event
}

/* Function prototypes */

// Pump the event loop, gathering events from the input devices.
// (https://wiki.libsdl.org/SDL3/SDL_PumpEvents)
func PumpEvents() {
	C.SDL_PumpEvents()
}

type EventAction C.SDL_EventAction

func (a EventAction) c() C.SDL_EventAction {
	return C.SDL_EventAction(a)
}

const (
	ADDEVENT  = EventAction(C.SDL_ADDEVENT)
	PEEKEVENT = EventAction(C.SDL_PEEKEVENT)
	GETEVENT  = EventAction(C.SDL_GETEVENT)
)

// Check the event queue for messages and optionally return them.
// (https://wiki.libsdl.org/SDL3/SDL_PeepEvents)
func PeepEvents(events []Event, action EventAction, minType EventType, maxType EventType) (n int, err error) {
	cevents := make([]cEvent, len(events))
	n = int(C.SDL_PeepEvents((*C.SDL_Event)(unsafe.SliceData(cevents)), C.int(len(cevents)), action.c(), minType.c(), maxType.c()))
	if n < 0 {
		return 0, GetError()
	}
	for i := 0; i < n; i++ {
		events[i] = cevents[i].goEvent()
	}
	return n, nil
}

// Check for the existence of a certain event type in the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_HasEvent)
func HasEvent(evtType EventType) bool {
	return bool(C.SDL_HasEvent(evtType.c()))
}

// Check for the existence of certain event types in the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_HasEvents)
func HasEvents(minType EventType, maxType EventType) bool {
	return bool(C.SDL_HasEvents(minType.c(), maxType.c()))
}

// Clear events of a specific type from the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_FlushEvent)
func FlushEvent(evtType EventType) {
	C.SDL_FlushEvent(evtType.c())
}

// Clear events of a range of types from the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_FlushEvents)
func FlushEvents(minType EventType, maxType EventType) {
	C.SDL_FlushEvents(minType.c(), maxType.c())
}

// Poll for currently pending events.
// (https://wiki.libsdl.org/SDL3/SDL_PollEvent)
func PollEvent() Event {
	var e cEvent
	if !C.SDL_PollEvent(e.cptr()) {
		return nil
	}
	return e.goEvent()
}

// Wait indefinitely for the next available event.
// (https://wiki.libsdl.org/SDL3/SDL_WaitEvent)
func WaitEvent() (Event, error) {
	var e cEvent
	if !C.SDL_WaitEvent(e.cptr()) {
		return nil, GetError()
	}
	return e.goEvent(), nil
}

// Wait until the specified timeout (in milliseconds) for the next available
// event.
// (https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout)
func WaitEventTimeout(timeout int32) Event {
	var e cEvent
	if !C.SDL_WaitEventTimeout(e.cptr(), C.Sint32(timeout)) {
		return nil
	}
	return e.goEvent()
}

// Add an event to the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_PushEvent)
func PushEvent(evt Event) (filtered bool, err error) {
	ret := C.SDL_PushEvent(evt.cptr())
	if !ret {
		return false, GetError()
	}
	return bool(ret), nil
}

// A function pointer used for callbacks that watch the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_EventFilter)
type EventFilter func(event Event, userdata any) bool

var eventFilter EventFilter
var eventFilterUserdata any

// TODO: test me
//
//export goEventFilter
func goEventFilter(event *C.SDL_Event) C.bool {
	if eventFilter == nil {
		return (C.bool)(true) // Permit event
	}
	return (C.bool)(eventFilter((*cEvent)(event).goEvent(), eventFilterUserdata))
}

// Set up a filter to process all events before they change internal state and
// are posted to the internal event queue.
// (https://wiki.libsdl.org/SDL3/SDL_SetEventFilter)
func SetEventFilter(filter EventFilter, userdata any) {
	eventFilter = filter
	eventFilterUserdata = userdata
	C.SDL_SetEventFilter(C.SDL_EventFilter(C.cgoEventFilter), nil)
}

// Query the current event filter.
// (https://wiki.libsdl.org/SDL3/SDL_GetEventFilter)
func GetEventFilter() (filter EventFilter, userdata any) {
	var cFilter C.SDL_EventFilter
	var cUserdata unsafe.Pointer
	ret := C.SDL_GetEventFilter(&cFilter, &cUserdata)
	if !ret {
		eventFilter = nil
		eventFilterUserdata = nil
	} else if cFilter != C.SDL_EventFilter(C.cgoEventFilter) {
		panic("event filter not set by SetEventFilter")
	}
	return eventFilter, eventFilterUserdata
}

// TODO: event watches needs to be tested
type eventWatcher struct {
	filter   EventFilter
	userdata any
}

var eventWatcherCount uintptr = 0
var eventWatchers = make(map[uintptr]eventWatcher)

// TODO: test me
//
//export goEventWatcher
func goEventWatcher(userdata unsafe.Pointer, event *C.SDL_Event) C.bool {
	watcher, ok := eventWatchers[uintptr(userdata)]
	if !ok {
		fmt.Printf("eventWatcher not found: %v\n", uintptr(userdata))
		return (C.bool)(true) // Permit event
	}
	return (C.bool)(watcher.filter((*cEvent)(event).goEvent(), watcher.userdata))
}

// Add a callback to be triggered when an event is added to the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_AddEventWatch)
func AddEventWatch(filter EventFilter, userdata any) error {
	watcher := eventWatcher{filter, userdata}
	ret := C.SDL_AddEventWatch(C.SDL_EventFilter(C.cgoEventWatcher), unsafe.Pointer(eventWatcherCount))
	if !ret {
		return GetError()
	}
	eventWatchers[eventWatcherCount] = watcher
	eventWatcherCount++
	return nil
}

// Remove an event watch callback added with SDL_AddEventWatch().
// (https://wiki.libsdl.org/SDL3/SDL_DelEventWatch)
func DelEventWatch(filter EventFilter, userdata any) {
	idx := uintptr(0)
	found := false
	for k, v := range eventWatchers {
		// TODO: there is probably a better way to do this
		if reflect.ValueOf(v.filter).Pointer() == reflect.ValueOf(filter).Pointer() && v.userdata == userdata {
			idx = k
			break
		}
	}
	if !found {
		fmt.Printf("DelEventWatch: eventWatcher not found: %v\n", idx)
		return
	}
	C.SDL_RemoveEventWatch(C.SDL_EventFilter(C.cgoEventWatcher), unsafe.Pointer(eventWatcherCount))
	delete(eventWatchers, idx)
}

var eventFilterOnce EventFilter
var eventFilterOnceUserdata any

//export goEventFilterOnce
func goEventFilterOnce(event *C.SDL_Event) C.bool {
	if eventFilterOnce == nil {
		return (C.bool)(true) // Permit event
	}
	return (C.bool)(eventFilterOnce((*cEvent)(event).goEvent(), eventFilterOnceUserdata))
}

// Run a specific filter function on the current event queue, removing any
// events for which the filter returns 0.
// (https://wiki.libsdl.org/SDL3/SDL_FilterEvents)
func FilterEvents(filter EventFilter, userdata any) {
	eventFilterOnce = filter
	eventFilterOnceUserdata = userdata
	C.SDL_FilterEvents(C.SDL_EventFilter(C.cgoEventFilterOnce), nil)
	eventFilterOnce = nil
	eventFilterOnceUserdata = nil
}

// Set the state of processing events by type.
// (https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled)
func SetEventEnabled(evtType EventType, enabled bool) {
	C.SDL_SetEventEnabled(evtType.c(), (C.bool)(enabled))
}

// Query the state of processing events by type.
// (https://wiki.libsdl.org/SDL3/SDL_EventEnabled)
func EventEnabled(evtType EventType) bool {
	return bool(C.SDL_EventEnabled(evtType.c()))
}

// Allocate a set of user-defined events, and return the beginning event
// number for that set of events.
// (https://wiki.libsdl.org/SDL3/SDL_RegisterEvents)
func RegisterEvents(numevents int) EventType {
	return EventType(C.SDL_RegisterEvents(C.int(numevents)))
}
