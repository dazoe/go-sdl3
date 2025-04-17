package sdl

// TODO: left off here

/*
#include <SDL3/SDL_gamepad.h>
#include "stdlib.h"
*/
import "C"
import (
	"encoding/binary"
	"unsafe"
)

// # CategoryGamepad
// (https://wiki.libsdl.org/SDL3/CategoryGamepad)

// The structure used to identify an SDL gamepad
// (https://wiki.libsdl.org/SDL3/SDL_Gamepad)
type Gamepad C.SDL_Gamepad

func (g *Gamepad) cptr() *C.SDL_Gamepad {
	return (*C.SDL_Gamepad)(unsafe.Pointer(g))
}

// Standard gamepad types.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadType)
type GamepadType C.SDL_GamepadType

func (t GamepadType) c() C.SDL_GamepadType {
	return C.SDL_GamepadType(t)
}

func (t GamepadType) String() string {
	return GetGamepadStringForType(t)
}

const (
	GAMEPAD_TYPE_UNKNOWN                      = GamepadType(C.SDL_GAMEPAD_TYPE_UNKNOWN)
	GAMEPAD_TYPE_STANDARD                     = GamepadType(C.SDL_GAMEPAD_TYPE_STANDARD)
	GAMEPAD_TYPE_XBOX360                      = GamepadType(C.SDL_GAMEPAD_TYPE_XBOX360)
	GAMEPAD_TYPE_XBOXONE                      = GamepadType(C.SDL_GAMEPAD_TYPE_XBOXONE)
	GAMEPAD_TYPE_PS3                          = GamepadType(C.SDL_GAMEPAD_TYPE_PS3)
	GAMEPAD_TYPE_PS4                          = GamepadType(C.SDL_GAMEPAD_TYPE_PS4)
	GAMEPAD_TYPE_PS5                          = GamepadType(C.SDL_GAMEPAD_TYPE_PS5)
	GAMEPAD_TYPE_NINTENDO_SWITCH_PRO          = GamepadType(C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_PRO)
	GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_LEFT  = GamepadType(C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_LEFT)
	GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_RIGHT = GamepadType(C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_RIGHT)
	GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_PAIR  = GamepadType(C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_PAIR)
	GAMEPAD_TYPE_COUNT                        = GamepadType(C.SDL_GAMEPAD_TYPE_COUNT)
)

// The list of buttons available on a gamepad
// (https://wiki.libsdl.org/SDL3/SDL_GamepadButton)
type GamepadButton C.SDL_GamepadButton

func (b GamepadButton) c() C.SDL_GamepadButton {
	return C.SDL_GamepadButton(b)
}

func (b GamepadButton) String() string {
	return GetGamepadStringForButton(b)
}

const (
	GAMEPAD_BUTTON_INVALID        = GamepadButton(C.SDL_GAMEPAD_BUTTON_INVALID)
	GAMEPAD_BUTTON_SOUTH          = GamepadButton(C.SDL_GAMEPAD_BUTTON_SOUTH) /* Bottom face button (e.g. Xbox A button) */
	GAMEPAD_BUTTON_EAST           = GamepadButton(C.SDL_GAMEPAD_BUTTON_EAST)  /* Right face button (e.g. Xbox B button) */
	GAMEPAD_BUTTON_WEST           = GamepadButton(C.SDL_GAMEPAD_BUTTON_WEST)  /* Left face button (e.g. Xbox X button) */
	GAMEPAD_BUTTON_NORTH          = GamepadButton(C.SDL_GAMEPAD_BUTTON_NORTH) /* Top face button (e.g. Xbox Y button) */
	GAMEPAD_BUTTON_BACK           = GamepadButton(C.SDL_GAMEPAD_BUTTON_BACK)
	GAMEPAD_BUTTON_GUIDE          = GamepadButton(C.SDL_GAMEPAD_BUTTON_GUIDE)
	GAMEPAD_BUTTON_START          = GamepadButton(C.SDL_GAMEPAD_BUTTON_START)
	GAMEPAD_BUTTON_LEFT_STICK     = GamepadButton(C.SDL_GAMEPAD_BUTTON_LEFT_STICK)
	GAMEPAD_BUTTON_RIGHT_STICK    = GamepadButton(C.SDL_GAMEPAD_BUTTON_RIGHT_STICK)
	GAMEPAD_BUTTON_LEFT_SHOULDER  = GamepadButton(C.SDL_GAMEPAD_BUTTON_LEFT_SHOULDER)
	GAMEPAD_BUTTON_RIGHT_SHOULDER = GamepadButton(C.SDL_GAMEPAD_BUTTON_RIGHT_SHOULDER)
	GAMEPAD_BUTTON_DPAD_UP        = GamepadButton(C.SDL_GAMEPAD_BUTTON_DPAD_UP)
	GAMEPAD_BUTTON_DPAD_DOWN      = GamepadButton(C.SDL_GAMEPAD_BUTTON_DPAD_DOWN)
	GAMEPAD_BUTTON_DPAD_LEFT      = GamepadButton(C.SDL_GAMEPAD_BUTTON_DPAD_LEFT)
	GAMEPAD_BUTTON_DPAD_RIGHT     = GamepadButton(C.SDL_GAMEPAD_BUTTON_DPAD_RIGHT)
	GAMEPAD_BUTTON_MISC1          = GamepadButton(C.SDL_GAMEPAD_BUTTON_MISC1)         /* Additional button (e.g. Xbox Series X share button, PS5 microphone button, Nintendo Switch Pro capture button, Amazon Luna microphone button, Google Stadia capture button) */
	GAMEPAD_BUTTON_RIGHT_PADDLE1  = GamepadButton(C.SDL_GAMEPAD_BUTTON_RIGHT_PADDLE1) /* Upper or primary paddle, under your right hand (e.g. Xbox Elite paddle P1) */
	GAMEPAD_BUTTON_LEFT_PADDLE1   = GamepadButton(C.SDL_GAMEPAD_BUTTON_LEFT_PADDLE1)  /* Upper or primary paddle, under your left hand (e.g. Xbox Elite paddle P3) */
	GAMEPAD_BUTTON_RIGHT_PADDLE2  = GamepadButton(C.SDL_GAMEPAD_BUTTON_RIGHT_PADDLE2) /* Lower or secondary paddle, under your right hand (e.g. Xbox Elite paddle P2) */
	GAMEPAD_BUTTON_LEFT_PADDLE2   = GamepadButton(C.SDL_GAMEPAD_BUTTON_LEFT_PADDLE2)  /* Lower or secondary paddle, under your left hand (e.g. Xbox Elite paddle P4) */
	GAMEPAD_BUTTON_TOUCHPAD       = GamepadButton(C.SDL_GAMEPAD_BUTTON_TOUCHPAD)      /* PS4/PS5 touchpad button */
	GAMEPAD_BUTTON_MISC2          = GamepadButton(C.SDL_GAMEPAD_BUTTON_MISC2)         /* Additional button */
	GAMEPAD_BUTTON_MISC3          = GamepadButton(C.SDL_GAMEPAD_BUTTON_MISC3)         /* Additional button */
	GAMEPAD_BUTTON_MISC4          = GamepadButton(C.SDL_GAMEPAD_BUTTON_MISC4)         /* Additional button */
	GAMEPAD_BUTTON_MISC5          = GamepadButton(C.SDL_GAMEPAD_BUTTON_MISC5)         /* Additional button */
	GAMEPAD_BUTTON_MISC6          = GamepadButton(C.SDL_GAMEPAD_BUTTON_MISC6)         /* Additional button */
	GAMEPAD_BUTTON_COUNT          = GamepadButton(C.SDL_GAMEPAD_BUTTON_COUNT)
)

// The set of gamepad button labels
// (https://wiki.libsdl.org/SDL3/SDL_GamepadButtonLabel)
type GamepadButtonLabel C.SDL_GamepadButtonLabel

const (
	GAMEPAD_BUTTON_LABEL_UNKNOWN  = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_UNKNOWN)
	GAMEPAD_BUTTON_LABEL_A        = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_A)
	GAMEPAD_BUTTON_LABEL_B        = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_B)
	GAMEPAD_BUTTON_LABEL_X        = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_X)
	GAMEPAD_BUTTON_LABEL_Y        = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_Y)
	GAMEPAD_BUTTON_LABEL_CROSS    = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_CROSS)
	GAMEPAD_BUTTON_LABEL_CIRCLE   = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_CIRCLE)
	GAMEPAD_BUTTON_LABEL_SQUARE   = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_SQUARE)
	GAMEPAD_BUTTON_LABEL_TRIANGLE = GamepadButtonLabel(C.SDL_GAMEPAD_BUTTON_LABEL_TRIANGLE)
)

// The list of axes available on a gamepad
// (https://wiki.libsdl.org/SDL3/SDL_GamepadAxis)
type GamepadAxis C.SDL_GamepadAxis

func (a GamepadAxis) c() C.SDL_GamepadAxis {
	return C.SDL_GamepadAxis(a)
}

func (a GamepadAxis) String() string {
	return GetGamepadStringForAxis(a)
}

const (
	GAMEPAD_AXIS_INVALID       = GamepadAxis(C.SDL_GAMEPAD_AXIS_INVALID)
	GAMEPAD_AXIS_LEFTX         = GamepadAxis(C.SDL_GAMEPAD_AXIS_LEFTX)
	GAMEPAD_AXIS_LEFTY         = GamepadAxis(C.SDL_GAMEPAD_AXIS_LEFTY)
	GAMEPAD_AXIS_RIGHTX        = GamepadAxis(C.SDL_GAMEPAD_AXIS_RIGHTX)
	GAMEPAD_AXIS_RIGHTY        = GamepadAxis(C.SDL_GAMEPAD_AXIS_RIGHTY)
	GAMEPAD_AXIS_LEFT_TRIGGER  = GamepadAxis(C.SDL_GAMEPAD_AXIS_LEFT_TRIGGER)
	GAMEPAD_AXIS_RIGHT_TRIGGER = GamepadAxis(C.SDL_GAMEPAD_AXIS_RIGHT_TRIGGER)
	GAMEPAD_AXIS_COUNT         = GamepadAxis(C.SDL_GAMEPAD_AXIS_COUNT)
)

// Types of gamepad control bindings.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadBindingType)
type GamepadBindingType C.SDL_GamepadBindingType

const (
	GAMEPAD_BINDTYPE_NONE   = GamepadBindingType(C.SDL_GAMEPAD_BINDTYPE_NONE)
	GAMEPAD_BINDTYPE_BUTTON = GamepadBindingType(C.SDL_GAMEPAD_BINDTYPE_BUTTON)
	GAMEPAD_BINDTYPE_AXIS   = GamepadBindingType(C.SDL_GAMEPAD_BINDTYPE_AXIS)
	GAMEPAD_BINDTYPE_HAT    = GamepadBindingType(C.SDL_GAMEPAD_BINDTYPE_HAT)
)

// TODO: make better?
// A mapping between one joystick input to a gamepad control.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadBinding)
type GamepadBinding C.SDL_GamepadBinding

func (b *GamepadBinding) InputType() GamepadBindingType {
	return GamepadBindingType(b.input_type)
}
func (b *GamepadBinding) Button() int32 {
	return int32(binary.NativeEndian.Uint32(b.input[0:]))
}
func (b *GamepadBinding) Axis() GamepadAxis {
	return GamepadAxis(binary.NativeEndian.Uint32(b.input[0:]))
}
func (b *GamepadBinding) AxisMin() int32 {
	return int32(binary.NativeEndian.Uint32(b.input[4:]))
}
func (b *GamepadBinding) AxisMax() int32 {
	return int32(binary.NativeEndian.Uint32(b.input[8:]))
}
func (b *GamepadBinding) Hat() int32 {
	return int32(binary.NativeEndian.Uint32(b.input[0:]))
}
func (b *GamepadBinding) HatMask() int32 {
	return int32(binary.NativeEndian.Uint32(b.input[4:]))
}
func (b *GamepadBinding) OutputType() GamepadBindingType {
	return GamepadBindingType(b.output_type)
}
func (b *GamepadBinding) OutButton() GamepadButton {
	return GamepadButton(binary.NativeEndian.Uint32(b.output[0:]))
}
func (b *GamepadBinding) OutAxis() GamepadAxis {
	return GamepadAxis(binary.NativeEndian.Uint32(b.output[0:]))
}
func (b *GamepadBinding) OutAxisMin() int32 {
	return int32(binary.NativeEndian.Uint32(b.output[4:]))
}
func (b *GamepadBinding) OutAxisMax() int32 {
	return int32(binary.NativeEndian.Uint32(b.output[8:]))
}

// Add support for gamepads that SDL is unaware of or change the binding of an
// existing gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping)
func AddGamepadMapping(mapping string) error {
	cMapping := C.CString(mapping)
	defer C.free(unsafe.Pointer(cMapping))
	ret := C.SDL_AddGamepadMapping(cMapping)
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Load a set of gamepad mappings from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromIO)
func AddGamepadMappingsFromIO(src *IOStream, closeio bool) (n int32, err error) {
	n = int32(C.SDL_AddGamepadMappingsFromIO(src.cptr(), (C.bool)(closeio)))
	if n < 0 {
		err = GetError()
	}
	return
}

// Load a set of gamepad mappings from a file.
// (https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile)
func AddGamepadMappingsFromFile(file string) (n int32, err error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	n = int32(C.SDL_AddGamepadMappingsFromFile(cFile))
	if n < 0 {
		err = GetError()
	}
	return
}

// Reinitialize the SDL mapping database to its initial state.
// (https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings)
func ReloadGamepadMappings() error {
	ret := C.SDL_ReloadGamepadMappings()
	if !ret {
		return GetError()
	}
	return nil
}

// Get the current gamepad mappings.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings)
func GetGamepadMappings() ([]string, error) {
	var count C.int
	ret := C.SDL_GetGamepadMappings(&count)
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice(ret, count)
	mappings := make([]string, count)
	for i := 0; i < int(count); i++ {
		mappings[i] = C.GoString(cSlice[i])
	}
	return mappings, nil
}

// Get the gamepad mapping string for a given GUID.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID)
func GetGamepadMappingForGUID(guid GUID) (string, error) {
	ret := C.SDL_GetGamepadMappingForGUID(guid.c())
	if ret == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	return C.GoString(ret), nil
}

// Get the current mapping of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping)
func (g *Gamepad) GetMapping() (string, error) {
	ret := C.SDL_GetGamepadMapping(g.cptr())
	if ret == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	return C.GoString(ret), nil
}

// Set the current mapping of a joystick or gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping)
func (id JoystickID) SetMapping(mapping string) error {
	cMapping := C.CString(mapping)
	defer C.free(unsafe.Pointer(cMapping))
	ret := C.SDL_SetGamepadMapping(id.c(), cMapping)
	if !ret {
		return GetError()
	}
	return nil
}

// Return whether a gamepad is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasGamepad)
func HasGamepad() bool {
	return bool(C.SDL_HasGamepad())
}

// Get a list of currently connected gamepads.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepads)
func GetGamepads() ([]JoystickID, error) {
	var count C.int
	ret := C.SDL_GetGamepads(&count)
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice(ret, count)
	ids := make([]JoystickID, count)
	for i := 0; i < int(count); i++ {
		ids[i] = JoystickID(cSlice[i])
	}
	return ids, nil
}

// Check if the given joystick is supported by the gamepad interface.
// (https://wiki.libsdl.org/SDL3/SDL_IsGamepad)
func (id JoystickID) IsGamepad() bool {
	return bool(C.SDL_IsGamepad(id.c()))
}

// Get the implementation dependent name of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID)
func (id JoystickID) GetGamepadName() (string, error) {
	ret := C.SDL_GetGamepadNameForID(id.c())
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Get the implementation dependent path of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID)
func (id JoystickID) GetGamepadPath() (string, error) {
	ret := C.SDL_GetGamepadPathForID(id.c())
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Get the player index of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID)
func (id JoystickID) GetGamepadPlayerIndex() int32 {
	return int32(C.SDL_GetGamepadPlayerIndexForID(id.c()))
}

// Get the implementation-dependent GUID of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID)
func (id JoystickID) GetGamepadGUID() GUID {
	return GUID(C.SDL_GetGamepadGUIDForID(id.c()))
}

// Get the USB vendor ID of a gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID)
func (id JoystickID) GetGamepadVendor() uint16 {
	return uint16(C.SDL_GetGamepadVendorForID(id.c()))
}

// Get the USB product ID of a gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID)
func (id JoystickID) GetGamepadProduct() uint16 {
	return uint16(C.SDL_GetGamepadProductForID(id.c()))
}

// Get the product version of a gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID)
func (id JoystickID) GetGamepadProductVersion() uint16 {
	return uint16(C.SDL_GetGamepadProductVersionForID(id.c()))
}

// Get the type of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID)
func (id JoystickID) GetGamepadType() GamepadType {
	return GamepadType(C.SDL_GetGamepadTypeForID(id.c()))
}

// Get the type of a gamepad, ignoring any mapping override.
// (https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID)
func (id JoystickID) GetRealGamepadType() GamepadType {
	return GamepadType(C.SDL_GetRealGamepadTypeForID(id.c()))
}

// Get the mapping of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID)
func (id JoystickID) GetGamepadMapping() string {
	ret := C.SDL_GetGamepadMappingForID(id.c())
	if ret == nil {
		return ""
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	return C.GoString(ret)
}

// Open a gamepad for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenGamepad)
func OpenGamepad(id JoystickID) (*Gamepad, error) {
	ret := C.SDL_OpenGamepad(id.c())
	if ret == nil {
		return nil, GetError()
	}
	return (*Gamepad)(ret), nil
}

// Get the SDL_Gamepad associated with a joystick instance ID, if it has been
// opened.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID)
func (id JoystickID) GetGamepad() (*Gamepad, error) {
	ret := C.SDL_GetGamepadFromID(id.c())
	if ret == nil {
		return nil, GetError()
	}
	return (*Gamepad)(ret), nil
}

// Get the SDL_Gamepad associated with a player index.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex)
func GetGamepadFromPlayerIndex(playerIndex int32) *Gamepad {
	return (*Gamepad)(C.SDL_GetGamepadFromPlayerIndex(C.int(playerIndex)))
}

// Get the properties associated with an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties)
func (g *Gamepad) GetProperties() (PropertiesID, error) {
	ret := C.SDL_GetGamepadProperties(g.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

const (
	PROP_GAMEPAD_CAP_MONO_LED_BOOLEAN       = PROP_JOYSTICK_CAP_MONO_LED_BOOLEAN
	PROP_GAMEPAD_CAP_RGB_LED_BOOLEAN        = PROP_JOYSTICK_CAP_RGB_LED_BOOLEAN
	PROP_GAMEPAD_CAP_PLAYER_LED_BOOLEAN     = PROP_JOYSTICK_CAP_PLAYER_LED_BOOLEAN
	PROP_GAMEPAD_CAP_RUMBLE_BOOLEAN         = PROP_JOYSTICK_CAP_RUMBLE_BOOLEAN
	PROP_GAMEPAD_CAP_TRIGGER_RUMBLE_BOOLEAN = PROP_JOYSTICK_CAP_TRIGGER_RUMBLE_BOOLEAN
)

// Get the instance ID of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadID)
func (g *Gamepad) GetID() (JoystickID, error) {
	ret := C.SDL_GetGamepadID(g.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return JoystickID(ret), nil
}

// Get the implementation-dependent name for an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadName)
func (g *Gamepad) GetName() string {
	return C.GoString(C.SDL_GetGamepadName(g.cptr()))
}

// Get the implementation-dependent path for an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath)
func (g *Gamepad) GetPath() string {
	return C.GoString(C.SDL_GetGamepadPath(g.cptr()))
}

// Get the type of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadType)
func (g *Gamepad) GetType() GamepadType {
	return GamepadType(C.SDL_GetGamepadType(g.cptr()))
}

// Get the type of an opened gamepad, ignoring any mapping override.
// (https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType)
func (g *Gamepad) GetRealType() GamepadType {
	return GamepadType(C.SDL_GetRealGamepadType(g.cptr()))
}

// Get the player index of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex)
func (g *Gamepad) GetPlayerIndex() int32 {
	return int32(C.SDL_GetGamepadPlayerIndex(g.cptr()))
}

// Set the player index of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex)
func (g *Gamepad) SetPlayerIndex(playerIndex int32) error {
	ret := C.SDL_SetGamepadPlayerIndex(g.cptr(), C.int(playerIndex))
	if !ret {
		return GetError()
	}
	return nil
}

// Get the USB vendor ID of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor)
func (g *Gamepad) GetVendor() uint16 {
	return uint16(C.SDL_GetGamepadVendor(g.cptr()))
}

// Get the USB product ID of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct)
func (g *Gamepad) GetProduct() uint16 {
	return uint16(C.SDL_GetGamepadProduct(g.cptr()))
}

// Get the product version of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion)
func (g *Gamepad) GetProductVersion() uint16 {
	return uint16(C.SDL_GetGamepadProductVersion(g.cptr()))
}

// Get the firmware version of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion)
func (g *Gamepad) GetFirmwareVersion() uint16 {
	return uint16(C.SDL_GetGamepadFirmwareVersion(g.cptr()))
}

// Get the serial number of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial)
func (g *Gamepad) GetSerial() string {
	return C.GoString(C.SDL_GetGamepadSerial(g.cptr()))
}

// Get the Steam Input handle of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle)
func (g *Gamepad) GetSteamHandle() uint64 {
	return uint64(C.SDL_GetGamepadSteamHandle(g.cptr()))
}

// Get the connection state of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState)
func (g *Gamepad) GetConnectionState() (state JoystickConnectionState, err error) {
	state = JoystickConnectionState(C.SDL_GetGamepadConnectionState(g.cptr()))
	if state == JOYSTICK_CONNECTION_INVALID {
		err = GetError()
	}
	return
}

// Get the battery state of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo)
func (g *Gamepad) GetPowerInfo() (percent int32, state PowerState) {
	state = PowerState(C.SDL_GetGamepadPowerInfo(g.cptr(), (*C.int)(&percent)))
	return
}

// Check if a gamepad has been opened and is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadConnected)
func (g *Gamepad) Connected() bool {
	return bool(C.SDL_GamepadConnected(g.cptr()))
}

// Get the underlying joystick from a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick)
func (g *Gamepad) GetJoystick() (*Joystick, error) {
	ret := C.SDL_GetGamepadJoystick(g.cptr())
	if ret == nil {
		return nil, GetError()
	}
	return (*Joystick)(ret), nil
}

// Set the state of gamepad event processing.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled)
func (g *Gamepad) SetEventsEnabled(enabled bool) {
	C.SDL_SetGamepadEventsEnabled((C.bool)(enabled))
}

// Query the state of gamepad event processing.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled)
func (g *Gamepad) EventsEnabled() bool {
	return bool(C.SDL_GamepadEventsEnabled())
}

// TODO: needs to be tested, does the copy copy the pointers in the array or the data they point to?
// also what should be copied? are the pointers valid for the lifetime of the SDL?, or just Gamepad? or not after SDL_free?
// Get the SDL joystick layer bindings for a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings)
func (g *Gamepad) GetBindings() ([]*GamepadBinding, error) {
	var count C.int
	ret := C.SDL_GetGamepadBindings(g.cptr(), &count)
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice((**GamepadBinding)(unsafe.Pointer(ret)), count)
	bindings := make([]*GamepadBinding, count)
	copy(bindings, cSlice)
	return bindings, nil
}

// Manually pump gamepad updates if not using the loop.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads)
func UpdateGamepads() {
	C.SDL_UpdateGamepads()
}

// Convert a string into SDL_GamepadType enum.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString)
func GetGamepadTypeFromString(str string) GamepadType {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return GamepadType(C.SDL_GetGamepadTypeFromString(cstr))
}

// Convert from an SDL_GamepadType enum to a string.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType)
func GetGamepadStringForType(t GamepadType) string {
	return C.GoString(C.SDL_GetGamepadStringForType(t.c()))
}

// Convert a string into SDL_GamepadAxis enum.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString)
func GetGamepadAxisFromString(str string) GamepadAxis {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return GamepadAxis(C.SDL_GetGamepadAxisFromString(cstr))
}

// Convert from an SDL_GamepadAxis enum to a string.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis)
func GetGamepadStringForAxis(a GamepadAxis) string {
	return C.GoString(C.SDL_GetGamepadStringForAxis(a.c()))
}

// Query whether a gamepad has a given axis.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis)
func (g *Gamepad) HasAxis(a GamepadAxis) bool {
	return bool(C.SDL_GamepadHasAxis(g.cptr(), a.c()))
}

// Get the current state of an axis control on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis)
func (g *Gamepad) GetAxis(a GamepadAxis) (int16, error) {
	return int16(C.SDL_GetGamepadAxis(g.cptr(), a.c())), GetError()
}

// Convert a string into an SDL_GamepadButton enum.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString)
func GetGamepadButtonFromString(str string) GamepadButton {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return GamepadButton(C.SDL_GetGamepadButtonFromString(cstr))
}

// Convert from an SDL_GamepadButton enum to a string.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton)
func GetGamepadStringForButton(b GamepadButton) string {
	return C.GoString(C.SDL_GetGamepadStringForButton(b.c()))
}

// Query whether a gamepad has a given button.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton)
func (g *Gamepad) HasButton(b GamepadButton) bool {
	return bool(C.SDL_GamepadHasButton(g.cptr(), b.c()))
}

// Get the current state of a button on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton)
func (g *Gamepad) GetButton(b GamepadButton) bool {
	return bool(C.SDL_GetGamepadButton(g.cptr(), b.c()))
}

// Get the label of a button on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType)
func GetGamepadButtonLabelForType(t GamepadType, b GamepadButton) GamepadButtonLabel {
	return GamepadButtonLabel(C.SDL_GetGamepadButtonLabelForType(t.c(), b.c()))
}

// Get the label of a button on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel)
func (g *Gamepad) GetButtonLabel(b GamepadButton) GamepadButtonLabel {
	return GamepadButtonLabel(C.SDL_GetGamepadButtonLabel(g.cptr(), b.c()))
}

// Get the number of touchpads on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads)
func (g *Gamepad) NumTouchpads() int32 {
	return int32(C.SDL_GetNumGamepadTouchpads(g.cptr()))
}

// Get the number of supported simultaneous fingers on a touchpad on a game
// gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers)
func (g *Gamepad) NumTouchpadFingers(touchpad int) int32 {
	return int32(C.SDL_GetNumGamepadTouchpadFingers(g.cptr(), C.int(touchpad)))
}

// Get the current state of a finger on a touchpad on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger)
func (g *Gamepad) GetTouchpadFinger(touchpad int, finger int) (down bool, x, y, pressure float32, err error) {
	ret := C.SDL_GetGamepadTouchpadFinger(g.cptr(), C.int(touchpad), C.int(finger), (*C.bool)(&down), (*C.float)(&x), (*C.float)(&y), (*C.float)(&pressure))
	if !ret {
		err = GetError()
	}
	return
}

// Return whether a gamepad has a particular sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor)
func (g *Gamepad) HasSensor(t SensorType) bool {
	return bool(C.SDL_GamepadHasSensor(g.cptr(), t.c()))
}

// Set whether data reporting for a gamepad sensor is enabled.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled)
func (g *Gamepad) SetSensorEnabled(t SensorType, enabled bool) error {
	ret := C.SDL_SetGamepadSensorEnabled(g.cptr(), t.c(), (C.bool)(enabled))
	if !ret {
		return GetError()
	}
	return nil
}

// Query whether sensor data reporting is enabled for a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled)
func (g *Gamepad) SensorEnabled(t SensorType) bool {
	return bool(C.SDL_GamepadSensorEnabled(g.cptr(), t.c()))
}

// Get the data rate (number of events per second) of a gamepad sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate)
func (g *Gamepad) SensorDataRate(t SensorType) float32 {
	return float32(C.SDL_GetGamepadSensorDataRate(g.cptr(), t.c()))
}

// Get the current state of a gamepad sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData)
func (g *Gamepad) SensorData(t SensorType, data []float32) error {
	dataPtr := (*C.float)(unsafe.SliceData(data))
	num_values := C.int(len(data))
	ret := C.SDL_GetGamepadSensorData(g.cptr(), t.c(), dataPtr, num_values)
	if !ret {
		return GetError()
	}
	return nil
}

// Start a rumble effect on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad)
func (g *Gamepad) Rumble(low_frequency_rumble, high_frequency_rumble uint16, duration_ms uint32) (bool, error) {
	ret := C.SDL_RumbleGamepad(g.cptr(), C.Uint16(low_frequency_rumble), C.Uint16(high_frequency_rumble), C.Uint32(duration_ms))
	if !ret {
		return false, GetError()
	}
	return true, nil
}

// Start a rumble effect in the gamepad's triggers.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers)
func (g *Gamepad) RumbleTriggers(left_rumble, right_rumble uint16, duration_ms uint32) error {
	ret := C.SDL_RumbleGamepadTriggers(g.cptr(), C.Uint16(left_rumble), C.Uint16(right_rumble), C.Uint32(duration_ms))
	if !ret {
		return GetError()
	}
	return nil
}

// Update a gamepad's LED color.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED)
func (gp *Gamepad) SetLED(r, g, b uint8) error {
	ret := C.SDL_SetGamepadLED(gp.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b))
	if !ret {
		return GetError()
	}
	return nil
}

// Send a gamepad specific effect packet.
// (https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect)
func (g *Gamepad) SendEffect(data []byte) error {
	dataPtr := unsafe.Pointer(unsafe.SliceData(data))
	ret := C.SDL_SendGamepadEffect(g.cptr(), dataPtr, C.int(len(data)))
	if !ret {
		return GetError()
	}
	return nil
}

// Close a gamepad previously opened with SDL_OpenGamepad().
// (https://wiki.libsdl.org/SDL3/SDL_CloseGamepad)
func (g *Gamepad) Close() {
	C.SDL_CloseGamepad(g.cptr())
}

// Return the sfSymbolsName for a given button on a gamepad on Apple
// platforms.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton)
func (g *Gamepad) GetAppleSFSymbolsNameForButton(b GamepadButton) string {
	return C.GoString(C.SDL_GetGamepadAppleSFSymbolsNameForButton(g.cptr(), b.c()))
}

// Return the sfSymbolsName for a given axis on a gamepad on Apple platforms.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis)
func (g *Gamepad) GetAppleSFSymbolsNameForAxis(a GamepadAxis) string {
	return C.GoString(C.SDL_GetGamepadAppleSFSymbolsNameForAxis(g.cptr(), a.c()))
}
