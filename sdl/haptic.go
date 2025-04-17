package sdl

//#include <SDL3/SDL_haptic.h>
import "C"
import "unsafe"

// # CategoryHaptic
// (https://wiki.libsdl.org/SDL3/CategoryHaptic)

// The haptic structure used to identify an SDL haptic.
// (https://wiki.libsdl.org/SDL3/SDL_Haptic)
type Haptic C.SDL_Haptic

func (h *Haptic) cptr() *C.SDL_Haptic {
	return (*C.SDL_Haptic)(h)
}

type HapticEffectType C.Uint16
type HapticFeature C.Uint32

// Constant effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_CONSTANT)
const HAPTIC_CONSTANT = HapticEffectType(C.SDL_HAPTIC_CONSTANT)

// Sine wave effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_SINE)
const HAPTIC_SINE = HapticEffectType(C.SDL_HAPTIC_SINE)

// Square wave effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_SQUARE)
const HAPTIC_SQUARE = HapticEffectType(C.SDL_HAPTIC_SQUARE)

// Triangle wave effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_TRIANGLE)
const HAPTIC_TRIANGLE = HapticEffectType(C.SDL_HAPTIC_TRIANGLE)

// Sawtoothup wave effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_SAWTOOTHUP)
const HAPTIC_SAWTOOTHUP = HapticEffectType(C.SDL_HAPTIC_SAWTOOTHUP)

// Sawtoothdown wave effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_SAWTOOTHDOWN)
const HAPTIC_SAWTOOTHDOWN = HapticEffectType(C.SDL_HAPTIC_SAWTOOTHDOWN)

// Ramp effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_RAMP)
const HAPTIC_RAMP = HapticEffectType(C.SDL_HAPTIC_RAMP)

// Spring effect supported - uses axes position.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_SPRING)
const HAPTIC_SPRING = HapticEffectType(C.SDL_HAPTIC_SPRING)

// Damper effect supported - uses axes velocity.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_DAMPER)
const HAPTIC_DAMPER = HapticEffectType(C.SDL_HAPTIC_DAMPER)

// Inertia effect supported - uses axes acceleration.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_INERTIA)
const HAPTIC_INERTIA = HapticEffectType(C.SDL_HAPTIC_INERTIA)

// Friction effect supported - uses axes movement.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_FRICTION)
const HAPTIC_FRICTION = HapticEffectType(C.SDL_HAPTIC_FRICTION)

// Left/Right effect supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_LEFTRIGHT)
const HAPTIC_LEFTRIGHT = HapticEffectType(C.SDL_HAPTIC_LEFTRIGHT)

// Reserved for future use
const haptic_reserved1 = HapticEffectType(C.SDL_HAPTIC_RESERVED1)
const haptic_reserved2 = HapticEffectType(C.SDL_HAPTIC_RESERVED2)
const haptic_reserved3 = HapticEffectType(C.SDL_HAPTIC_RESERVED3)

// Custom effect is supported.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_CUSTOM)
const HAPTIC_CUSTOM = HapticEffectType(C.SDL_HAPTIC_CUSTOM)

// Device can set global gain.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_GAIN)
const HAPTIC_GAIN = HapticFeature(C.SDL_HAPTIC_GAIN)

// Device can set autocenter.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_AUTOCENTER)
const HAPTIC_AUTOCENTER = HapticFeature(C.SDL_HAPTIC_AUTOCENTER)

// Device can be queried for effect status.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_STATUS)
const HAPTIC_STATUS = HapticFeature(C.SDL_HAPTIC_STATUS)

// Device can be paused.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_PAUSE)
const HAPTIC_PAUSE = HapticFeature(C.SDL_HAPTIC_PAUSE)

type HapticDirectionType C.Uint8

// Uses polar coordinates for the direction.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_POLAR)
const HAPTIC_POLAR = HapticDirectionType(C.SDL_HAPTIC_POLAR)

// Uses cartesian coordinates for the direction.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_CARTESIAN)
const HAPTIC_CARTESIAN = HapticDirectionType(C.SDL_HAPTIC_CARTESIAN)

// Uses spherical coordinates for the direction.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_SPHERICAL)
const HAPTIC_SPHERICAL = HapticDirectionType(C.SDL_HAPTIC_SPHERICAL)

// Use this value to play an effect on the steering wheel axis.
//
//	(https://wiki.libsdl.org/SDL3/SDL_HAPTIC_STEERING_AXIS)
const HAPTIC_STEERING_AXIS = HapticDirectionType(C.SDL_HAPTIC_STEERING_AXIS)

// Used to play a device an infinite number of times.
// (https://wiki.libsdl.org/SDL3/SDL_HAPTIC_INFINITY)
const HAPTIC_INFINITY = C.SDL_HAPTIC_INFINITY

// Structure that represents a haptic direction.
// (https://wiki.libsdl.org/SDL3/SDL_HapticDirection)
type HapticDirection struct {
	Type      HapticDirectionType /**< The type of encoding. */
	Direction [3]int32            /**< The encoded direction. */
}

type HapticEffectHeader struct {
	Type HapticEffectType /**< The type of effect. */
}

func (h *HapticEffectHeader) GetType() HapticEffectType {
	return h.Type
}
func (h *HapticEffectHeader) cptr() *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(h))
}

// A structure containing a template for a Constant effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticConstant)
type HapticConstant struct {
	/* Header */
	HapticEffectHeader                 /**< SDL_HAPTIC_CONSTANT */
	Direction          HapticDirection /**< Direction of the effect. */

	/* Replay */
	Length uint32 /**< Duration of the effect. */
	Delay  uint16 /**< Delay before starting the effect. */

	/* Trigger */
	Button   uint16 /**< Button that triggers the effect. */
	Interval uint16 /**< How soon it can be triggered again after button. */

	/* Constant */
	Level int16 /**< Strength of the constant effect. */

	/* Envelope */
	Attack_length uint16 /**< Duration of the attack. */
	Attack_level  uint16 /**< Level at the start of the attack. */
	Fade_length   uint16 /**< Duration of the fade. */
	Fade_level    uint16 /**< Level at the end of the fade. */
}

// A structure containing a template for a Periodic effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticPeriodic)
type HapticPeriodic struct {
	/* Header */
	HapticEffectHeader /**< SDL_HAPTIC_SINE, SDL_HAPTIC_SQUARE
	SDL_HAPTIC_TRIANGLE, SDL_HAPTIC_SAWTOOTHUP or
	SDL_HAPTIC_SAWTOOTHDOWN */
	Direction HapticDirection /**< Direction of the effect. */

	/* Replay */
	Length uint32 /**< Duration of the effect. */
	Delay  uint16 /**< Delay before starting the effect. */

	/* Trigger */
	Button   uint16 /**< Button that triggers the effect. */
	Interval uint16 /**< How soon it can be triggered again after button. */

	/* Periodic */
	Period    uint16 /**< Period of the wave. */
	Magnitude int16  /**< Peak value; if negative, equivalent to 180 degrees extra phase shift. */
	Offset    int16  /**< Mean value of the wave. */
	Phase     uint16 /**< Positive phase shift given by hundredth of a degree. */

	/* Envelope */
	Attack_length uint16 /**< Duration of the attack. */
	Attack_level  uint16 /**< Level at the start of the attack. */
	Fade_length   uint16 /**< Duration of the fade. */
	Fade_level    uint16 /**< Level at the end of the fade. */
}

// A structure containing a template for a Condition effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticCondition)
type HapticCondition struct {
	/* Header */
	HapticEffectHeader /**< SDL_HAPTIC_SPRING, SDL_HAPTIC_DAMPER,
	SDL_HAPTIC_INERTIA or SDL_HAPTIC_FRICTION */
	Direction HapticDirection /**< Direction of the effect - Not used ATM. */

	/* Replay */
	Length uint32 /**< Duration of the effect. */
	Delay  uint16 /**< Delay before starting the effect. */

	/* Trigger */
	Button   uint16 /**< Button that triggers the effect. */
	Interval uint16 /**< How soon it can be triggered again after button. */

	/* Condition */
	Right_sat   [3]uint16 /**< Level when joystick is to the positive side; max 0xFFFF. */
	Left_sat    [3]uint16 /**< Level when joystick is to the negative side; max 0xFFFF. */
	Right_coeff [3]int16  /**< How fast to increase the force towards the positive side. */
	Left_coeff  [3]int16  /**< How fast to increase the force towards the negative side. */
	Deadband    [3]uint16 /**< Size of the dead zone; max 0xFFFF: whole axis-range when 0-centered. */
	Center      [3]int16  /**< Position of the dead zone. */
}

// A structure containing a template for a Ramp effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticRamp)
type HapticRamp struct {
	/* Header */
	HapticEffectHeader                 /**< SDL_HAPTIC_RAMP */
	Direction          HapticDirection /**< Direction of the effect. */

	/* Replay */
	Length uint32 /**< Duration of the effect. */
	Delay  uint16 /**< Delay before starting the effect. */

	/* Trigger */
	Button   uint16 /**< Button that triggers the effect. */
	Interval uint16 /**< How soon it can be triggered again after button. */

	/* Ramp */
	Start int16 /**< Beginning strength level. */
	End   int16 /**< Ending strength level. */

	/* Envelope */
	Attack_length uint16 /**< Duration of the attack. */
	Attack_level  uint16 /**< Level at the start of the attack. */
	Fade_length   uint16 /**< Duration of the fade. */
	Fade_level    uint16 /**< Level at the end of the fade. */
}

// A structure containing a template for a Left/Right effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticLeftRight)
type HapticLeftRight struct {
	/* Header */
	HapticEffectHeader /**< SDL_HAPTIC_LEFTRIGHT */

	/* Replay */
	Length uint32 /**< Duration of the effect in milliseconds. */

	/* Rumble */
	Large_magnitude uint16 /**< Control of the large controller motor. */
	Small_magnitude uint16 /**< Control of the small controller motor. */
}

// A structure containing a template for the SDL_HAPTIC_CUSTOM effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticCustom)
type HapticCustom struct {
	/* Header */
	HapticEffectHeader                 /**< SDL_HAPTIC_CUSTOM */
	Direction          HapticDirection /**< Direction of the effect. */

	/* Replay */
	Length uint32 /**< Duration of the effect. */
	Delay  uint16 /**< Delay before starting the effect. */

	/* Trigger */
	Button   uint16 /**< Button that triggers the effect. */
	Interval uint16 /**< How soon it can be triggered again after button. */

	/* Custom */
	Channels uint8   /**< Axes to use, minimum of one. */
	Period   uint16  /**< Sample periods. */
	Samples  uint16  /**< Amount of samples. */
	Data     *uint16 /**< Should contain channels*samples items. */

	/* Envelope */
	Attack_length uint16 /**< Duration of the attack. */
	Attack_level  uint16 /**< Level at the start of the attack. */
	Fade_length   uint16 /**< Duration of the fade. */
	Fade_level    uint16 /**< Level at the end of the fade. */
}

// The generic template for any haptic effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticEffect)
type HapticEffect interface {
	// Get the type of haptic effect.
	GetType() HapticEffectType
	cptr() *C.SDL_HapticEffect
}
type cHapticEffect C.SDL_HapticEffect

//  typedef union SDL_HapticEffect
//  {
// 	 /* Common for all force feedback effects */
// 	 Uint16 type;                    /**< Effect type. */
// 	 SDL_HapticConstant constant;    /**< Constant effect. */
// 	 SDL_HapticPeriodic periodic;    /**< Periodic effect. */
// 	 SDL_HapticCondition condition;  /**< Condition effect. */
// 	 SDL_HapticRamp ramp;            /**< Ramp effect. */
// 	 SDL_HapticLeftRight leftright;  /**< Left/Right effect. */
// 	 SDL_HapticCustom custom;        /**< Custom effect. */
//  } SDL_HapticEffect;

// This is a unique ID for a haptic device for the time it is connected to the
// system, and is never reused for the lifetime of the application.
// (https://wiki.libsdl.org/SDL3/SDL_HapticID)
type HapticID C.SDL_HapticID

func (id HapticID) c() C.SDL_HapticID {
	return C.SDL_HapticID(id)
}

// Get a list of currently connected haptic devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetHaptics)
func GetHaptics() ([]HapticID, error) {
	var count C.int
	ret := C.SDL_GetHaptics(&count)
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice((*C.SDL_HapticID)(ret), int(count))
	haptics := make([]HapticID, count)
	for i := 0; i < int(count); i++ {
		haptics[i] = HapticID(cSlice[i])
	}
	return haptics, nil
}

// Get the implementation dependent name of a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID)
func (instance_id HapticID) GetHapticName() string {
	ret := C.SDL_GetHapticNameForID(instance_id.c())
	return C.GoString(ret)
}

// Open a haptic device for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenHaptic)
func (instance_id HapticID) OpenHaptic() (*Haptic, error) {
	ret := C.SDL_OpenHaptic(instance_id.c())
	if ret == nil {
		return nil, GetError()
	}
	return (*Haptic)(ret), nil
}

// Get the SDL_Haptic associated with an instance ID, if it has been opened.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID)
func (instance_id HapticID) GetHaptic() (*Haptic, error) {
	ret := C.SDL_GetHapticFromID(instance_id.c())
	if ret == nil {
		return nil, GetError()
	}
	return (*Haptic)(ret), nil
}

// Get the instance ID of an opened haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticID)
func (haptic *Haptic) GetID() (HapticID, error) {
	ret := C.SDL_GetHapticID(haptic.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return HapticID(ret), nil
}

// Get the implementation dependent name of a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticName)
func (haptic *Haptic) GetName() (name string, err error) {
	ret := C.SDL_GetHapticName(haptic.cptr())
	if ret == nil {
		return name, GetError()
	}
	return C.GoString(ret), nil
}

// Query whether or not the current mouse has haptic capabilities.
// (https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic)
func IsMouseHaptic() bool {
	return bool(C.SDL_IsMouseHaptic())
}

// Try to open a haptic device from the current mouse.
// (https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse)
func OpenHapticFromMouse() (*Haptic, error) {
	ret := C.SDL_OpenHapticFromMouse()
	if ret == nil {
		return nil, GetError()
	}
	return (*Haptic)(ret), nil
}

// Query if a joystick has haptic features.
// (https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic)
func (joystick *Joystick) IsHaptic() bool {
	return C.SDL_IsJoystickHaptic(joystick.cptr()) == true
}

// Open a haptic device for use from a joystick device.
// (https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick)
func (joystick *Joystick) OpenHaptic() (*Haptic, error) {
	ret := C.SDL_OpenHapticFromJoystick(joystick.cptr())
	if ret == nil {
		return nil, GetError()
	}
	return (*Haptic)(ret), nil
}

// Close a haptic device previously opened with SDL_OpenHaptic().
// (https://wiki.libsdl.org/SDL3/SDL_CloseHaptic)
func (haptic *Haptic) Close() {
	C.SDL_CloseHaptic(haptic.cptr())
}

// Get the number of effects a haptic device can store.
// (https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects)
func (haptic *Haptic) GetMaxEffects() (int32, error) {
	ret := int32(C.SDL_GetMaxHapticEffects(haptic.cptr()))
	if ret < 0 {
		return ret, GetError()
	}
	return ret, nil
}

// Get the number of effects a haptic device can play at the same time.
// (https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying)
func (haptic *Haptic) GetMaxEffectsPlaying() (int32, error) {
	ret := int32(C.SDL_GetMaxHapticEffectsPlaying(haptic.cptr()))
	if ret < 0 {
		return ret, GetError()
	}
	return ret, nil
}

// Get the haptic device's supported features in bitwise manner.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures)
func (haptic *Haptic) GetFeatures() (uint32, error) {
	// TODO: returns a bit mask of HapticEffectType and HapticFeature values
	ret := uint32(C.SDL_GetHapticFeatures(haptic.cptr()))
	if ret == 0 {
		return ret, GetError()
	}
	return ret, nil
}

// Get the number of haptic axes the device has.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes)
func (haptic *Haptic) GetNumAxes() (int32, error) {
	ret := int32(C.SDL_GetNumHapticAxes(haptic.cptr()))
	if ret < 0 {
		return ret, GetError()
	}
	return ret, nil
}

// Check to see if an effect is supported by a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported)
func (haptic *Haptic) EffectSupported(effect HapticEffect) bool {
	// TODO: needs tested, not sure if HapticEffect.cptr() works correctly for all types
	return C.SDL_HapticEffectSupported(haptic.cptr(), effect.cptr()) == true
}

type HapticEffectID C.int

func (id HapticEffectID) c() C.int {
	return C.int(id)
}

// Create a new haptic effect on a specified device.
// (https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect)
func (haptic *Haptic) CreateEffect(effect HapticEffect) (id HapticEffectID, err error) {
	id = HapticEffectID(C.SDL_CreateHapticEffect(haptic.cptr(), effect.cptr()))
	if id < 0 {
		err = GetError()
	}
	return
}

// Update the properties of an effect.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect)
func (haptic *Haptic) UpdateEffect(id HapticEffectID, effect HapticEffect) error {
	ret := C.SDL_UpdateHapticEffect(haptic.cptr(), id.c(), effect.cptr())
	if !ret {
		return GetError()
	}
	return nil
}

// Run the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect)
func (haptic *Haptic) RunEffect(id HapticEffectID, iterations uint32) error {
	ret := C.SDL_RunHapticEffect(haptic.cptr(), id.c(), C.Uint32(iterations))
	if !ret {
		return GetError()
	}
	return nil
}

// Stop the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect)
func (haptic *Haptic) StopEffect(id HapticEffectID) error {
	ret := C.SDL_StopHapticEffect(haptic.cptr(), id.c())
	if !ret {
		return GetError()
	}
	return nil
}

// Destroy a haptic effect on the device.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect)
func (haptic *Haptic) DestroyEffect(id HapticEffectID) {
	C.SDL_DestroyHapticEffect(haptic.cptr(), id.c())
}

// Get the status of the current effect on the specified haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus)
func (haptic *Haptic) GetEffectStatus(id HapticEffectID) (playing bool, err error) {
	ret := C.SDL_GetHapticEffectStatus(haptic.cptr(), id.c())
	if !ret {
		return false, GetError()
	}
	return bool(ret), nil
}

// Set the global gain of the specified haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_SetHapticGain)
func (haptic *Haptic) SetGain(gain int32) error {
	ret := C.SDL_SetHapticGain(haptic.cptr(), C.int(gain))
	if !ret {
		return GetError()
	}
	return nil
}

// Set the global autocenter of the device.
// (https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter)
func (haptic *Haptic) SetAutocenter(autocenter int32) error {
	ret := C.SDL_SetHapticAutocenter(haptic.cptr(), C.int(autocenter))
	if !ret {
		return GetError()
	}
	return nil
}

// Pause a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_PauseHaptic)
func (haptic *Haptic) Pause() error {
	ret := C.SDL_PauseHaptic(haptic.cptr())
	if !ret {
		return GetError()
	}
	return nil
}

// Resume a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic)
func (haptic *Haptic) Resume() error {
	ret := C.SDL_ResumeHaptic(haptic.cptr())
	if !ret {
		return GetError()
	}
	return nil
}

// Stop all the currently playing effects on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects)
func (haptic *Haptic) StopEffects() error {
	ret := C.SDL_StopHapticEffects(haptic.cptr())
	if !ret {
		return GetError()
	}
	return nil
}

// Check whether rumble is supported on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported)
func (haptic *Haptic) RumbleSupported() bool {
	return bool(C.SDL_HapticRumbleSupported(haptic.cptr()))
}

// Initialize a haptic device for simple rumble playback.
// (https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble)
func (haptic *Haptic) InitRumble() error {
	ret := C.SDL_InitHapticRumble(haptic.cptr())
	if !ret {
		return GetError()
	}
	return nil
}

// Run a simple rumble effect on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble)
func (haptic *Haptic) PlayRumble(strength float32, length uint32) error {
	ret := C.SDL_PlayHapticRumble(haptic.cptr(), C.float(strength), C.Uint32(length))
	if !ret {
		return GetError()
	}
	return nil
}

// Stop the simple rumble on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble)
func (haptic *Haptic) StopRumble() error {
	ret := C.SDL_StopHapticRumble(haptic.cptr())
	if !ret {
		return GetError()
	}
	return nil
}
