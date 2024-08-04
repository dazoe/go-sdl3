package sdl

//#include <SDL3/SDL_joystick.h>
import "C"
import "unsafe"

// # CategoryJoystick
// * SDL joystick support.
// (https://wiki.libsdl.org/SDL3/CategoryJoystick)

// The joystick structure used to identify an SDL joystick.
// (https://wiki.libsdl.org/SDL3/SDL_Joystick)
type Joystick C.SDL_Joystick

func (j *Joystick) cptr() *C.SDL_Joystick {
	return (*C.SDL_Joystick)(j)
}

// This is a unique ID for a joystick for the time it is connected to the
// (https://wiki.libsdl.org/SDL3/SDL_JoystickID)
type JoystickID C.SDL_JoystickID

func (id JoystickID) c() C.SDL_JoystickID {
	return C.SDL_JoystickID(id)
}

// An enum of some common joystick types.
// (https://wiki.libsdl.org/SDL3/SDL_JoystickType)
type JoystickType C.SDL_JoystickType

const (
	JOYSTICK_TYPE_UNKNOWN      = JoystickType(C.SDL_JOYSTICK_TYPE_UNKNOWN)
	JOYSTICK_TYPE_GAMEPAD      = JoystickType(C.SDL_JOYSTICK_TYPE_GAMEPAD)
	JOYSTICK_TYPE_WHEEL        = JoystickType(C.SDL_JOYSTICK_TYPE_WHEEL)
	JOYSTICK_TYPE_ARCADE_STICK = JoystickType(C.SDL_JOYSTICK_TYPE_ARCADE_STICK)
	JOYSTICK_TYPE_FLIGHT_STICK = JoystickType(C.SDL_JOYSTICK_TYPE_FLIGHT_STICK)
	JOYSTICK_TYPE_DANCE_PAD    = JoystickType(C.SDL_JOYSTICK_TYPE_DANCE_PAD)
	JOYSTICK_TYPE_GUITAR       = JoystickType(C.SDL_JOYSTICK_TYPE_GUITAR)
	JOYSTICK_TYPE_DRUM_KIT     = JoystickType(C.SDL_JOYSTICK_TYPE_DRUM_KIT)
	JOYSTICK_TYPE_ARCADE_PAD   = JoystickType(C.SDL_JOYSTICK_TYPE_ARCADE_PAD)
	JOYSTICK_TYPE_THROTTLE     = JoystickType(C.SDL_JOYSTICK_TYPE_THROTTLE)
)

// Possible connection states for a joystick device.
// (https://wiki.libsdl.org/SDL3/SDL_JoystickConnectionState)
type JoystickConnectionState C.SDL_JoystickConnectionState

const (
	JOYSTICK_CONNECTION_INVALID  = JoystickConnectionState(C.SDL_JOYSTICK_CONNECTION_INVALID)
	JOYSTICK_CONNECTION_UNKNOWN  = JoystickConnectionState(C.SDL_JOYSTICK_CONNECTION_UNKNOWN)
	JOYSTICK_CONNECTION_WIRED    = JoystickConnectionState(C.SDL_JOYSTICK_CONNECTION_WIRED)
	JOYSTICK_CONNECTION_WIRELESS = JoystickConnectionState(C.SDL_JOYSTICK_CONNECTION_WIRELESS)
)

// The largest value an SDL_Joystick's axis can report.
// (https://wiki.libsdl.org/SDL3/SDL_JOYSTICK_AXIS_MAX)
const JOYSTICK_AXIS_MAX = 32767

// The smallest value an SDL_Joystick's axis can report.
// (https://wiki.libsdl.org/SDL3/SDL_JOYSTICK_AXIS_MIN)
const JOYSTICK_AXIS_MIN = -32768

// Set max recognized G-force from accelerometer
// See src/joystick/uikit/SDL_sysjoystick.m for notes on why this is needed
const SDL_IPHONE_MAX_GFORCE = 5.0

/* Function prototypes */

// Locking for atomic access to the joystick API.
// (https://wiki.libsdl.org/SDL3/SDL_LockJoysticks)
func LockJoysticks() {
	C.SDL_LockJoysticks()
}

// Unlocking for atomic access to the joystick API.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks)
func UnlockJoysticks() {
	C.SDL_UnlockJoysticks()
}

// Return whether a joystick is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasJoystick)
func HasJoystick() bool {
	return C.SDL_HasJoystick() != 0
}

// Get a list of currently connected joysticks.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoysticks)
func GetJoysticks() ([]JoystickID, error) {
	var count C.int
	ret := C.SDL_GetJoysticks(&count)
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice((*JoystickID)(ret), int(count))
	joysticks := make([]JoystickID, int(count))
	copy(joysticks, cSlice)
	return joysticks, nil
}

// Get the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID)
func (instance_id JoystickID) GetName() (name string, err error) {
	name = C.GoString(C.SDL_GetJoystickNameForID(C.SDL_JoystickID(instance_id)))
	if len(name) == 0 {
		err = GetError()
	}
	return
}

// Get the implementation dependent path of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID)
func (instance_id JoystickID) GetPath() (path string, err error) {
	path = C.GoString(C.SDL_GetJoystickPathForID(C.SDL_JoystickID(instance_id)))
	if len(path) == 0 {
		err = GetError()
	}
	return
}

// Get the player index of a joystick.
// (https://wiki.libsdl.org/SDL3/GetJoystickPlayerIndexForID)
func (instance_id JoystickID) GetPlayerIndex() int32 {
	return int32(C.SDL_GetJoystickPlayerIndexForID(C.SDL_JoystickID(instance_id)))
}

// Get the implementation-dependent GUID of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID)
func (instance_id JoystickID) GetGUID() GUID {
	return GUID(C.SDL_GetJoystickGUIDForID(C.SDL_JoystickID(instance_id)))
}

// Get the USB vendor ID of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID)
func (instance_id JoystickID) GetVendor() uint16 {
	return uint16(C.SDL_GetJoystickVendorForID(C.SDL_JoystickID(instance_id)))
}

// Get the USB product ID of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID)
func (instance_id JoystickID) GetProduct() uint16 {
	return uint16(C.SDL_GetJoystickProductForID(C.SDL_JoystickID(instance_id)))
}

// Get the product version of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID)
func (instance_id JoystickID) GetProductVersion() uint16 {
	return uint16(C.SDL_GetJoystickProductVersionForID(C.SDL_JoystickID(instance_id)))
}

// Get the type of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID)
func (instance_id JoystickID) GetType() JoystickType {
	return JoystickType(C.SDL_GetJoystickTypeForID(C.SDL_JoystickID(instance_id)))
}

// Open a joystick for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenJoystick)
func OpenJoystick(instance_id JoystickID) (joystick *Joystick, err error) {
	joystick = (*Joystick)(C.SDL_OpenJoystick(C.SDL_JoystickID(instance_id)))
	if joystick == nil {
		err = GetError()
	}
	return
}

// Get the SDL_Joystick associated with an instance ID, if it has been opened.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID)
func (instance_id JoystickID) GetJoystick() (joystick *Joystick, err error) {
	joystick = (*Joystick)(C.SDL_GetJoystickFromID(C.SDL_JoystickID(instance_id)))
	if joystick == nil {
		err = GetError()
	}
	return
}

// Get the SDL_Joystick associated with a player index.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex)
func GetJoystickFromPlayerIndex(player_index int32) (joystick *Joystick, err error) {
	joystick = (*Joystick)(C.SDL_GetJoystickFromPlayerIndex(C.int(player_index)))
	if joystick == nil {
		err = GetError()
	}
	return
}

// The structure that describes a virtual joystick touchpad.
// (https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickTouchpadDesc)
type VirtualJoystickTouchpadDesc struct {
	Nfingers uint16    /**< the number of simultaneous fingers on this touchpad */
	_        [3]uint16 // padding
}

// The structure that describes a virtual joystick sensor.
// (https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickSensorDesc)
type VirtualJoystickSensorDesc struct {
	Type SensorType /**< the type of this sensor */
	rate float32    /**< the update frequency of this sensor, may be 0.0f */
}

// The structure that describes a virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickSensorDesc)
// TODO: This needs extra work, struct has callbacks
type VirtualJoystickDesc C.SDL_VirtualJoystickDesc

func (vjd *VirtualJoystickDesc) cptr() *C.SDL_VirtualJoystickDesc {
	return (*C.SDL_VirtualJoystickDesc)(vjd)
}

// Attach a new virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick)
func AttachVirtualJoystick(desc *VirtualJoystickDesc) (id JoystickID, err error) {
	id = JoystickID(C.SDL_AttachVirtualJoystick(desc.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Detach a virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick)
func DetachVirtualJoystick(instance_id JoystickID) (err error) {
	ret := C.SDL_DetachVirtualJoystick(C.SDL_JoystickID(instance_id))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Query whether or not a joystick is virtual.
// (https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual)
func IsJoystickVirtual(instance_id JoystickID) bool {
	return C.SDL_IsJoystickVirtual(C.SDL_JoystickID(instance_id)) != 0
}

// Set the state of an axis on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis)
func (joystick *Joystick) SetVirtualAxis(axis int, value int16) (err error) {
	ret := C.SDL_SetJoystickVirtualAxis(joystick.cptr(), C.int(axis), C.Sint16(value))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Generate ball motion on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall)
func (joystick *Joystick) SetVirtualBall(ball int, xrel int16, yrel int16) (err error) {
	ret := C.SDL_SetJoystickVirtualBall(joystick.cptr(), C.int(ball), C.Sint16(xrel), C.Sint16(yrel))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the state of a button on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton)
func (joystick *Joystick) SetVirtualButton(button int, value uint8) (err error) {
	ret := C.SDL_SetJoystickVirtualButton(joystick.cptr(), C.int(button), C.Uint8(value))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the state of a hat on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat)
func (joystick *Joystick) SetVirtualHat(hat int, value uint8) (err error) {
	ret := C.SDL_SetJoystickVirtualHat(joystick.cptr(), C.int(hat), C.Uint8(value))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set touchpad finger state on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad)
func (joystick *Joystick) SetVirtualTouchpad(touchpad int, finger int, state uint8, x float32, y float32, pressure float32) (err error) {
	ret := C.SDL_SetJoystickVirtualTouchpad(joystick.cptr(), C.int(touchpad), C.int(finger), C.Uint8(state), C.float(x), C.float(y), C.float(pressure))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Send a sensor update for an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData)
func (joystick *Joystick) SendVirtualSensorData(_type SensorType, sensor_timestamp uint64, data []float32) (err error) {
	dataPtr := unsafe.SliceData(data)
	ret := C.SDL_SendJoystickVirtualSensorData(joystick.cptr(), C.SDL_SensorType(_type), C.Uint64(sensor_timestamp), (*C.float)(dataPtr), C.int(len(data)))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the properties associated with a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties)
func (joystick *Joystick) GetProperties() (properties PropertiesID, err error) {
	properties = PropertiesID(C.SDL_GetJoystickProperties(joystick.cptr()))
	if properties == 0 {
		err = GetError()
	}
	return
}

const (
	PROP_JOYSTICK_CAP_MONO_LED_BOOLEAN       = "SDL.joystick.cap.mono_led"
	PROP_JOYSTICK_CAP_RGB_LED_BOOLEAN        = "SDL.joystick.cap.rgb_led"
	PROP_JOYSTICK_CAP_PLAYER_LED_BOOLEAN     = "SDL.joystick.cap.player_led"
	PROP_JOYSTICK_CAP_RUMBLE_BOOLEAN         = "SDL.joystick.cap.rumble"
	PROP_JOYSTICK_CAP_TRIGGER_RUMBLE_BOOLEAN = "SDL.joystick.cap.trigger_rumble"
)

// Get the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickName)
func (joystick *Joystick) GetName() (name string, err error) {
	name = C.GoString(C.SDL_GetJoystickName(joystick.cptr()))
	if len(name) == 0 {
		err = GetError()
	}
	return
}

// Get the implementation dependent path of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath)
func (joystick *Joystick) GetPath() (path string, err error) {
	path = C.GoString(C.SDL_GetJoystickPath(joystick.cptr()))
	if len(path) == 0 {
		err = GetError()
	}
	return
}

// Get the player index of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex)
func (joystick *Joystick) GetPlayerIndex() int32 {
	return int32(C.SDL_GetJoystickPlayerIndex(joystick.cptr()))
}

// Set the player index of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex)
func (joystick *Joystick) SetPlayerIndex(player_index int32) (err error) {
	ret := C.SDL_SetJoystickPlayerIndex(joystick.cptr(), C.int(player_index))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the implementation-dependent GUID for the joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID)
func (joystick *Joystick) GetGUID() (guid GUID, err error) {
	guid = GUID(C.SDL_GetJoystickGUID(joystick.cptr()))
	if guid.isZero() {
		err = GetError()
	}
	return
}

// Get the USB vendor ID of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor)
func (joystick *Joystick) GetVendor() uint16 {
	return uint16(C.SDL_GetJoystickVendor(joystick.cptr()))
}

// Get the USB product ID of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct)
func (joystick *Joystick) GetProduct() uint16 {
	return uint16(C.SDL_GetJoystickProduct(joystick.cptr()))
}

// Get the product version of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion)
func (joystick *Joystick) GetProductVersion() uint16 {
	return uint16(C.SDL_GetJoystickProductVersion(joystick.cptr()))
}

// Get the firmware version of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion)
func (joystick *Joystick) GetFirmwareVersion() uint16 {
	return uint16(C.SDL_GetJoystickFirmwareVersion(joystick.cptr()))
}

// Get the serial number of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial)
func (joystick *Joystick) GetSerial() string {
	return C.GoString(C.SDL_GetJoystickSerial(joystick.cptr()))
}

// Get the type of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickType)
func (joystick *Joystick) GetType() JoystickType {
	return JoystickType(C.SDL_GetJoystickType(joystick.cptr()))
}

// Get the device information encoded in a SDL_GUID structure.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo)
func GetJoystickGUIDInfo(guid GUID) (vendor uint16, product uint16, version uint16, crc16 uint16) {
	C.SDL_GetJoystickGUIDInfo(guid.c(), (*C.Uint16)(&vendor), (*C.Uint16)(&product), (*C.Uint16)(&version), (*C.Uint16)(&crc16))
	return
}

// Get the status of a specified joystick.
// (https://wiki.libsdl.org/SDL3/SDL_JoystickConnected)
func (joystick *Joystick) JoystickConnected() (connected bool, err error) {
	connected = C.SDL_JoystickConnected(joystick.cptr()) != 0
	if !connected {
		err = GetError()
	}
	return
}

// Get the instance ID of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickID)
func (joystick *Joystick) GetID() (id JoystickID, err error) {
	id = JoystickID(C.SDL_GetJoystickID(joystick.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Get the number of general axis controls on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes)
func (joystick *Joystick) GetNumAxes() (num int, err error) {
	num = int(C.SDL_GetNumJoystickAxes(joystick.cptr()))
	if num < 0 {
		err = GetError()
	}
	return
}

// Get the number of trackballs on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls)
func (joystick *Joystick) GetNumBalls() (num int, err error) {
	num = int(C.SDL_GetNumJoystickBalls(joystick.cptr()))
	if num < 0 {
		err = GetError()
	}
	return
}

// Get the number of POV hats on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats)
func (joystick *Joystick) GetNumHats() (num int, err error) {
	num = int(C.SDL_GetNumJoystickHats(joystick.cptr()))
	if num < 0 {
		err = GetError()
	}
	return
}

// Get the number of buttons on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons)
func (joystick *Joystick) GetNumButtons() (num int, err error) {
	num = int(C.SDL_GetNumJoystickButtons(joystick.cptr()))
	if num < 0 {
		err = GetError()
	}
	return
}

// Set the state of joystick event processing.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled)
func SetJoystickEventsEnabled(enabled bool) {
	C.SDL_SetJoystickEventsEnabled(sdlBool(enabled))
}

// Query the state of joystick event processing.
// (https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled)
func JoystickEventsEnabled() bool {
	return C.SDL_JoystickEventsEnabled() != 0
}

// Update the current state of the open joysticks.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks)
func UpdateJoysticks() {
	C.SDL_UpdateJoysticks()
}

// Get the current state of an axis control on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis)
func (joystick *Joystick) GetAxis(axis int32) (value int16, err error) {
	value = int16(C.SDL_GetJoystickAxis(joystick.cptr(), C.int(axis)))
	if value == 0 {
		err = GetError()
	}
	return
}

// Get the initial state of an axis control on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState)
func (joystick *Joystick) GetAxisInitialState(axis int32) (state int16, ok bool) {
	ok = C.SDL_GetJoystickAxisInitialState(joystick.cptr(), C.int(axis), (*C.Sint16)(&state)) != 0
	return
}

// Get the ball axis change since the last poll.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall)
func (joystick *Joystick) GetBall(ball int32) (dx, dy int32, err error) {
	ret := C.SDL_GetJoystickBall(joystick.cptr(), C.int(ball), (*C.int)(&dx), (*C.int)(&dy))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the current state of a POV hat on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat)
func (joystick *Joystick) GetHat(hat int32) HAT_DIRECTION {
	return HAT_DIRECTION(C.SDL_GetJoystickHat(joystick.cptr(), C.int(hat)))
}

type HAT_DIRECTION uint8

const (
	HAT_CENTERED  = HAT_DIRECTION(C.SDL_HAT_CENTERED)
	HAT_UP        = HAT_DIRECTION(C.SDL_HAT_UP)
	HAT_RIGHT     = HAT_DIRECTION(C.SDL_HAT_RIGHT)
	HAT_DOWN      = HAT_DIRECTION(C.SDL_HAT_DOWN)
	HAT_LEFT      = HAT_DIRECTION(C.SDL_HAT_LEFT)
	HAT_RIGHTUP   = HAT_DIRECTION(C.SDL_HAT_RIGHTUP)
	HAT_RIGHTDOWN = HAT_DIRECTION(C.SDL_HAT_RIGHTDOWN)
	HAT_LEFTUP    = HAT_DIRECTION(C.SDL_HAT_LEFTUP)
	HAT_LEFTDOWN  = HAT_DIRECTION(C.SDL_HAT_LEFTDOWN)
)

// Get the current state of a button on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton)
func (joystick *Joystick) GetButton(button int32) uint8 {
	return uint8(C.SDL_GetJoystickButton(joystick.cptr(), C.int(button)))
}

// Start a rumble effect.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick)
func (joystick *Joystick) Rumble(low_frequency_rumble, high_frequency_rumble uint16, duration_ms uint32) int32 {
	return int32(C.SDL_RumbleJoystick(joystick.cptr(), C.Uint16(low_frequency_rumble), C.Uint16(high_frequency_rumble), C.Uint32(duration_ms)))
}

// Start a rumble effect in the joystick's triggers.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers)
func (joystick *Joystick) RumbleTriggers(left_rumble, right_rumble uint16, duration_ms uint32) (err error) {
	ret := C.SDL_RumbleJoystickTriggers(joystick.cptr(), C.Uint16(left_rumble), C.Uint16(right_rumble), C.Uint32(duration_ms))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Update a joystick's LED color.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED)
func (joystick *Joystick) SetLED(red uint8, green uint8, blue uint8) (err error) {
	ret := C.SDL_SetJoystickLED(joystick.cptr(), C.Uint8(red), C.Uint8(green), C.Uint8(blue))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Send a joystick specific effect packet.
// (https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect)
func (joystick *Joystick) SendEffect(data []byte) (err error) {
	ret := C.SDL_SendJoystickEffect(joystick.cptr(), unsafe.Pointer(unsafe.SliceData(data)), C.int(len(data)))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Close a joystick previously opened with SDL_OpenJoystick().
// (https://wiki.libsdl.org/SDL3/SDL_CloseJoystick)
func (joystick *Joystick) Close() {
	C.SDL_CloseJoystick(joystick.cptr())
}

// Get the connection state of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState)
func (joystick *Joystick) GetConnectionState() (state JoystickConnectionState, err error) {
	state = JoystickConnectionState(C.SDL_GetJoystickConnectionState(joystick.cptr()))
	if state == JOYSTICK_CONNECTION_INVALID {
		err = GetError()
	}
	return
}

// Get the battery state of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo)
func (joystick *Joystick) GetPowerInfo() (percent int32, state PowerState, err error) {
	state = PowerState(C.SDL_GetJoystickPowerInfo(joystick.cptr(), (*C.int)(&percent)))
	if state == POWERSTATE_ERROR {
		err = GetError()
	}
	return
}
