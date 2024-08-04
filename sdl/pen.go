package sdl

//#include <SDL3/SDL_pen.h>
import "C"
import "unsafe"

// # CategoryPen
// (https://wiki.libsdl.org/SDL3/CategoryPen)

type PenID C.SDL_PenID

func (id PenID) c() C.SDL_PenID {
	return C.SDL_PenID(id)
}

const PEN_INVALID = PenID(C.SDL_PEN_INVALID)   /**< Reserved invalid SDL_PenID is valid */
const PEN_MOUSEID = MouseID(C.SDL_PEN_MOUSEID) /**< Device ID for mouse events triggered by pen events */
const PEN_INFO_UNKNOWN = PenID(0xffffffff)     /**< Marks unknown information when querying the pen */

// Pen axis indices
// (https://wiki.libsdl.org/SDL3/SDL_PenAxis)
type PenAxis C.SDL_PenAxis

const (
	PEN_AXIS_PRESSURE = PenAxis(C.SDL_PEN_AXIS_PRESSURE) /**< Pen pressure.  Unidirectional: 0..1.0 */
	PEN_AXIS_XTILT    = PenAxis(C.SDL_PEN_AXIS_XTILT)    /**< Pen horizontal tilt angle.  Bidirectional: -90.0..90.0 (left-to-right).
	  The physical max/min tilt may be smaller than -90.0 / 90.0, cf. SDL_PenCapabilityInfo */
	PEN_AXIS_YTILT = PenAxis(C.SDL_PEN_AXIS_YTILT) /**< Pen vertical tilt angle.  Bidirectional: -90.0..90.0 (top-to-down).
	  The physical max/min tilt may be smaller than -90.0 / 90.0, cf. SDL_PenCapabilityInfo */
	PEN_AXIS_DISTANCE = PenAxis(C.SDL_PEN_AXIS_DISTANCE) /**< Pen distance to drawing surface.  Unidirectional: 0.0..1.0 */
	PEN_AXIS_ROTATION = PenAxis(C.SDL_PEN_AXIS_ROTATION) /**< Pen barrel rotation.  Bidirectional: -180..179.9 (clockwise, 0 is facing up, -180.0 is facing down). */
	PEN_AXIS_SLIDER   = PenAxis(C.SDL_PEN_AXIS_SLIDER)   /**< Pen finger wheel or slider (e.g., Airbrush Pen).  Unidirectional: 0..1.0 */
	PEN_NUM_AXES      = PenAxis(C.SDL_PEN_NUM_AXES)      /**< Last valid axis index */
	PEN_AXIS_LAST     = PenAxis(C.SDL_PEN_AXIS_LAST)     /**< Last axis index plus 1 */
)

const (
	/* Pen flags.  These share a bitmask space with SDL_BUTTON_LEFT and friends. */
	PEN_FLAG_DOWN_BIT_INDEX   = MouseButtonFlags(C.SDL_PEN_FLAG_DOWN_BIT_INDEX)   /* Bit for storing that pen is touching the surface */
	PEN_FLAG_INK_BIT_INDEX    = MouseButtonFlags(C.SDL_PEN_FLAG_INK_BIT_INDEX)    /* Bit for storing has-non-eraser-capability status */
	PEN_FLAG_ERASER_BIT_INDEX = MouseButtonFlags(C.SDL_PEN_FLAG_ERASER_BIT_INDEX) /* Bit for storing is-eraser or has-eraser-capability property */
	PEN_FLAG_AXIS_BIT_OFFSET  = MouseButtonFlags(C.SDL_PEN_FLAG_AXIS_BIT_OFFSET)  /* Bit for storing has-axis-0 property */
	/* Pen tips */
	PEN_TIP_INK    = MouseButtonFlags(C.SDL_PEN_TIP_INK)    /**< Regular pen tip (for drawing) touched the surface */
	PEN_TIP_ERASER = MouseButtonFlags(C.SDL_PEN_TIP_ERASER) /**< Eraser pen tip touched the surface */
)

// #define SDL_PEN_CAPABILITY(capbit)    (1ul << (capbit))
// #define SDL_PEN_AXIS_CAPABILITY(axis) SDL_PEN_CAPABILITY((axis) + SDL_PEN_FLAG_AXIS_BIT_OFFSET)

// Pen capabilities reported by SDL_GetPenCapabilities.
// (https://wiki.libsdl.org/SDL3/SDL_PenCapabilityFlags)
type PenCapabilityFlags C.SDL_PenCapabilityFlags

const (
	PEN_DOWN_MASK                = PenCapabilityFlags(C.SDL_PEN_DOWN_MASK)          /**< Pen tip is currently touching the drawing surface. */
	PEN_INK_MASK                 = PenCapabilityFlags(C.SDL_PEN_INK_MASK)           /**< Pen has a regular drawing tip (SDL_GetPenCapabilities).  For events (SDL_PenButtonEvent, SDL_PenMotionEvent, SDL_GetPenStatus) this flag is mutually exclusive with SDL_PEN_ERASER_MASK .  */
	PEN_ERASER_MASK              = PenCapabilityFlags(C.SDL_PEN_ERASER_MASK)        /**< Pen has an eraser tip (SDL_GetPenCapabilities) or is being used as eraser (SDL_PenButtonEvent , SDL_PenMotionEvent , SDL_GetPenStatus)  */
	PEN_AXIS_PRESSURE_MASK       = PenCapabilityFlags(C.SDL_PEN_AXIS_PRESSURE_MASK) /**< Pen provides pressure information in axis SDL_PEN_AXIS_PRESSURE */
	PEN_AXIS_XTILT_MASK          = PenCapabilityFlags(C.SDL_PEN_AXIS_XTILT_MASK)    /**< Pen provides horizontal tilt information in axis SDL_PEN_AXIS_XTILT */
	PEN_AXIS_YTILT_MASK          = PenCapabilityFlags(C.SDL_PEN_AXIS_YTILT_MASK)    /**< Pen provides vertical tilt information in axis SDL_PEN_AXIS_YTILT */
	PEN_AXIS_DISTANCE_MASK       = PenCapabilityFlags(C.SDL_PEN_AXIS_DISTANCE_MASK) /**< Pen provides distance to drawing tablet in SDL_PEN_AXIS_DISTANCE */
	PEN_AXIS_ROTATION_MASK       = PenCapabilityFlags(C.SDL_PEN_AXIS_ROTATION_MASK) /**< Pen provides barrel rotation information in axis SDL_PEN_AXIS_ROTATION */
	PEN_AXIS_SLIDER_MASK         = PenCapabilityFlags(C.SDL_PEN_AXIS_SLIDER_MASK)   /**< Pen provides slider / finger wheel or similar in axis SDL_PEN_AXIS_SLIDER */
	PEN_AXIS_BIDIRECTIONAL_MASKS = PenCapabilityFlags(C.SDL_PEN_AXIS_BIDIRECTIONAL_MASKS)
)

// Pen types
// (https://wiki.libsdl.org/SDL3/SDL_PenSubType)
type PenSubType uint32

const (
	PEN_TYPE_UNKNOWN  = PenSubType(C.SDL_PEN_TYPE_UNKNOWN)
	PEN_TYPE_ERASER   = PenSubType(C.SDL_PEN_TYPE_ERASER)   /**< Eraser */
	PEN_TYPE_PEN      = PenSubType(C.SDL_PEN_TYPE_PEN)      /**< Generic pen; this is the default. */
	PEN_TYPE_PENCIL   = PenSubType(C.SDL_PEN_TYPE_PENCIL)   /**< Pencil */
	PEN_TYPE_BRUSH    = PenSubType(C.SDL_PEN_TYPE_BRUSH)    /**< Brush-like device */
	PEN_TYPE_AIRBRUSH = PenSubType(C.SDL_PEN_TYPE_AIRBRUSH) /**< Airbrush device that "sprays" ink */
	PEN_TYPE_LAST     = PenSubType(C.SDL_PEN_TYPE_LAST)     /**< Last valid pen type */
)

/* Function prototypes */

// Retrieves all pens that are connected to the system.
// (https://wiki.libsdl.org/SDL3/SDL_GetPens)
func GetPens() ([]PenID, error) {
	var count C.int
	ret := C.SDL_GetPens(&count)
	if ret == nil {
		return nil, GetError()
	}
	cslice := unsafe.Slice((*PenID)(ret), count)
	pens := make([]PenID, count)
	copy(pens, cslice)
	C.SDL_free(unsafe.Pointer(ret))
	return pens, nil
}

// Retrieves the pen's current status.
// (https://wiki.libsdl.org/SDL3/SDL_GetPenStatus)
func GetPenStatus(penID PenID, axes []float32) (x, y float32, state MouseButtonFlags, err error) {
	state = MouseButtonFlags(C.SDL_GetPenStatus(penID.c(), (*C.float)(&x), (*C.float)(&y), (*C.float)(unsafe.SliceData(axes)), C.size_t(len(axes))))
	if state == 0 {
		err = GetError()
	}
	return
}

// Retrieves an SDL_PenID for the given SDL_GUID.
// (https://wiki.libsdl.org/SDL3/SDL_GetPenFromGUID)
func GetPenFromGUID(guid GUID) PenID {
	return PenID(C.SDL_GetPenFromGUID(guid.c()))
}

// Retrieves the SDL_GUID for a given SDL_PenID.
// (https://wiki.libsdl.org/SDL3/SDL_GetPenGUID)
func GetPenGUID(penID PenID) GUID {
	return GUID(C.SDL_GetPenGUID(penID.c()))
}

// Checks whether a pen is still attached.
// (https://wiki.libsdl.org/SDL3/SDL_PenConnected)
func PenConnected(penID PenID) bool {
	return C.SDL_PenConnected(penID.c()) != 0
}

// Retrieves a human-readable description for a SDL_PenID.
// (https://wiki.libsdl.org/SDL3/SDL_GetPenName)
func GetPenName(penID PenID) (string, error) {
	ret := C.SDL_GetPenName(penID.c())
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Pen capabilities, as reported by SDL_GetPenCapabilities()
// (https://wiki.libsdl.org/SDL3/SDL_PenCapabilityInfo)
type PenCapabilityInfo struct {
	MaxTilt    float32 /**< Physical maximum tilt angle, for XTILT and YTILT, or SDL_PEN_INFO_UNKNOWN .  Pens cannot typically tilt all the way to 90 degrees, so this value is usually less than 90.0. */
	WacomID    uint32  /**< For Wacom devices: wacom tool type ID, otherwise 0 (useful e.g. with libwacom) */
	NumButtons int8    /**< Number of pen buttons (not counting the pen tip), or SDL_PEN_INFO_UNKNOWN */
}

func (info *PenCapabilityInfo) cptr() *C.SDL_PenCapabilityInfo {
	return (*C.SDL_PenCapabilityInfo)(unsafe.Pointer(info))
}

// Retrieves capability flags for a given SDL_PenID.
// (https://wiki.libsdl.org/SDL3/SDL_GetPenCapabilities)
func GetPenCapabilities(penID PenID) (info PenCapabilityInfo, flags PenCapabilityFlags) {
	flags = PenCapabilityFlags(C.SDL_GetPenCapabilities(penID.c(), info.cptr()))
	return
}

// Retrieves the pen type for a given SDL_PenID.
// (https://wiki.libsdl.org/SDL3/SDL_GetPenType)
func GetPenType(penID PenID) PenSubType {
	return PenSubType(C.SDL_GetPenType(penID.c()))
}
