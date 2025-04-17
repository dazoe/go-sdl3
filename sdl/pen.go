package sdl

//#include <SDL3/SDL_pen.h>
import "C"

//TODO: pen is done already for rewrite for 3.1.6
// # CategoryPen
// (https://wiki.libsdl.org/SDL3/CategoryPen)

// SDL pen instance IDs.
type PenID C.SDL_PenID

// Pen input flags, as reported by various pen events' `pen_state` field.
type PenInputFlags C.SDL_PenInputFlags

const (
	PEN_INPUT_DOWN       = PenInputFlags(C.SDL_PEN_INPUT_DOWN)       /**< pen is pressed down */
	PEN_INPUT_BUTTON_1   = PenInputFlags(C.SDL_PEN_INPUT_BUTTON_1)   /**< button 1 is pressed */
	PEN_INPUT_BUTTON_2   = PenInputFlags(C.SDL_PEN_INPUT_BUTTON_2)   /**< button 2 is pressed */
	PEN_INPUT_BUTTON_3   = PenInputFlags(C.SDL_PEN_INPUT_BUTTON_3)   /**< button 3 is pressed */
	PEN_INPUT_BUTTON_4   = PenInputFlags(C.SDL_PEN_INPUT_BUTTON_4)   /**< button 4 is pressed */
	PEN_INPUT_BUTTON_5   = PenInputFlags(C.SDL_PEN_INPUT_BUTTON_5)   /**< button 5 is pressed */
	PEN_INPUT_ERASER_TIP = PenInputFlags(C.SDL_PEN_INPUT_ERASER_TIP) /**< eraser tip is used */
)

// Pen axis indices.
type PenAxis C.SDL_PenAxis

const (
	PEN_AXIS_PRESSURE            = PenAxis(C.SDL_PEN_AXIS_PRESSURE)            /**< Pen pressure.  Unidirectional: 0 to 1.0 */
	PEN_AXIS_XTILT               = PenAxis(C.SDL_PEN_AXIS_XTILT)               /**< Pen horizontal tilt angle.  Bidirectional: -90.0 to 90.0 (left-to-right). */
	PEN_AXIS_YTILT               = PenAxis(C.SDL_PEN_AXIS_YTILT)               /**< Pen vertical tilt angle.  Bidirectional: -90.0 to 90.0 (top-to-down). */
	PEN_AXIS_DISTANCE            = PenAxis(C.SDL_PEN_AXIS_DISTANCE)            /**< Pen distance to drawing surface.  Unidirectional: 0.0 to 1.0 */
	PEN_AXIS_ROTATION            = PenAxis(C.SDL_PEN_AXIS_ROTATION)            /**< Pen barrel rotation.  Bidirectional: -180 to 179.9 (clockwise, 0 is facing up, -180.0 is facing down). */
	PEN_AXIS_SLIDER              = PenAxis(C.SDL_PEN_AXIS_SLIDER)              /**< Pen finger wheel or slider (e.g., Airbrush Pen).  Unidirectional: 0 to 1.0 */
	PEN_AXIS_TANGENTIAL_PRESSURE = PenAxis(C.SDL_PEN_AXIS_TANGENTIAL_PRESSURE) /**< Pressure from squeezing the pen ("barrel pressure"). */
	PEN_AXIS_COUNT               = PenAxis(C.SDL_PEN_AXIS_COUNT)               /**< Total known pen axis types in this version of SDL. This number may grow in future releases! */
)
