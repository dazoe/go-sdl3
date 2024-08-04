package sdl

//#include <SDL3/SDL_rect.h>
import "C"
import (
	"math"
	"unsafe"
)

/**
 * # CategoryRect
 *
 * Some helper functions for managing rectangles and 2D points, in both
 * interger and floating point versions.
 *
 * (https://wiki.libsdl.org/SDL3/CategoryRect)
 */

// The structure that defines a point (using integers).
// (https://wiki.libsdl.org/SDL3/SDL_Point)
type Point struct {
	X, Y int32
}

func (p *Point) cptr() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

// The structure that defines a point (using floating point values).
// (https://wiki.libsdl.org/SDL3/SDL_FPoint)
type FPoint struct {
	X, Y float32
}

func (p *FPoint) cptr() *C.SDL_FPoint {
	return (*C.SDL_FPoint)(unsafe.Pointer(p))
}

// A rectangle, with the origin at the upper left (using integers).
// (https://wiki.libsdl.org/SDL3/SDL_Rect)
type Rect struct {
	X, Y, W, H int32
}

func (r *Rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(r))
}

// A rectangle, with the origin at the upper left (using floating point
// values).
// (https://wiki.libsdl.org/SDL3/SDL_FRect)
type FRect struct {
	X, Y, W, H float32
}

func (r *FRect) cptr() *C.SDL_FRect {
	return (*C.SDL_FRect)(unsafe.Pointer(r))
}

// Determine whether a point resides inside a rectangle.
// (https://wiki.libsdl.org/SDL3/SDL_PointInRect)
func (p *Point) InRect(r *Rect) bool {
	return (p != nil && r != nil &&
		p.X >= r.X && p.X < (r.X+r.W) &&
		p.Y >= r.Y && p.Y < (r.Y+r.H))
}

// Determine whether a rectangle has no area.
// (https://wiki.libsdl.org/SDL3/SDL_RectEmpty)
func (r *Rect) Empty() bool {
	return r == nil || r.W <= 0 || r.H <= 0
}

// Determine whether two rectangles are equal.
// (https://wiki.libsdl.org/SDL3/SDL_RectsEqual)
func (a *Rect) Equal(b *Rect) bool {
	return a != nil && b != nil && *a == *b
}

// Using a local functions is much faster and can be inlined

// Determine whether two rectangles intersect.
// (https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection)
func (a *Rect) HasIntersection(b *Rect) bool {
	if a == nil || a.W <= 0 || a.H <= 0 {
		return false
	}
	if b == nil || b.W <= 0 || b.H <= 0 {
		return false
	}
	if a.X >= b.X+b.W || a.X+a.W <= b.X || a.Y >= b.Y+b.H || a.Y+a.H <= b.Y {
		return false
	}
	return true
}

// Calculate the intersection of two rectangles.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection)
func (a *Rect) GetIntersection(b *Rect) (result Rect, intersects bool) {
	if a == nil || a.W <= 0 || a.H <= 0 {
		return
	}
	if b == nil || b.W <= 0 || b.H <= 0 {
		return
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin > aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin > aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return result, result.W > 0 && result.H > 0
}

// Calculate the union of two rectangles.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectUnion)
func (a *Rect) GetUnion(b *Rect) (result Rect, err error) {
	if a == nil {
		err = InvalidParamError("a")
		return
	}
	if b == nil {
		err = InvalidParamError("b")
	}

	/* Special cases for empty Rects */
	if a.W == 0 || a.H == 0 {
		if b.W == 0 || b.H == 0 {
			return
		}
		result = *b
		return
	}
	if b.W == 0 || b.H == 0 {
		result = *a
		return
	}

	/* Horizontal union */
	Amin := a.X
	Amax := Amin + a.W
	Bmin := b.X
	Bmax := Bmin + b.W
	if Bmin < Amin {
		Amin = Bmin
	}
	result.X = Amin
	if Bmax > Amax {
		Amax = Bmax
	}
	result.W = Amax - Amin

	/* Vertical union */
	Amin = a.Y
	Amax = Amin + a.H
	Bmin = b.Y
	Bmax = Bmin + b.H
	if Bmin < Amin {
		Amin = Bmin
	}
	result.Y = Amin
	if Bmax > Amax {
		Amax = Bmax
	}
	result.H = Amax - Amin
	return
}

// Calculate a minimal rectangle enclosing a set of points.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints)
func GetRectEnclosingPoints(points []Point, clip *Rect) (result Rect, err error) {
	if len(points) <= 0 {
		err = InvalidParamError("points")
		return
	}

	var minx, miny, maxx, maxy int32
	var x, y int32
	if clip != nil {
		if clip.Empty() {
			return
		}
		clip_minx := clip.X
		clip_miny := clip.Y
		clip_maxx := clip.X + clip.W - 1
		clip_maxy := clip.Y + clip.H - 1

		var added bool
		for i := 0; i < len(points); i++ {
			x = points[i].X
			y = points[i].Y

			if x < clip_minx || x > clip_maxx ||
				y < clip_miny || y > clip_maxy {
				continue
			}

			if !added {
				minx = x
				maxx = x
				miny = y
				maxy = y
				added = true
			}

			if x < minx {
				minx = x
			} else if x > maxx {
				maxx = x
			}
			if y < miny {
				miny = y
			} else if y > maxy {
				maxy = y
			}
		}
		return
	} else {
		minx = points[0].X
		maxx = minx
		miny = points[0].Y
		maxy = miny

		for i := 1; i < len(points); i++ {
			x = points[i].X
			y = points[i].Y

			if x < minx {
				minx = x
			} else if x > maxx {
				maxx = x
			}
			if y < miny {
				miny = y
			} else if y > maxy {
				maxy = y
			}
		}
	}
	result.X = minx
	result.Y = miny
	result.W = (maxx - minx) + 1
	result.H = (maxy - miny) + 1
	return
}

const (
	code_BOTTOM = 1
	code_TOP    = 2
	code_LEFT   = 4
	code_RIGHT  = 8
)

func computeOutCode(r *Rect, x, y int32) int {
	code := 0
	if y < r.Y {
		code |= code_TOP
	} else if y >= r.Y+r.H {
		code |= code_BOTTOM
	}
	if x < r.X {
		code |= code_LEFT
	} else if x >= r.X+r.W {
		code |= code_RIGHT
	}
	return code
}

// Calculate the intersection of a rectangle and line segment.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection)
func (r *Rect) GetLineIntersection(X1, Y1, X2, Y2 *int32) (intersects bool, err error) {
	var x, y int

	if r == nil {
		err = InvalidParamError("rect")
		return
	} else if X1 == nil {
		err = InvalidParamError("X1")
		return
	} else if Y1 == nil {
		err = InvalidParamError("Y1")
		return
	} else if X2 == nil {
		err = InvalidParamError("X2")
		return
	} else if Y2 == nil {
		err = InvalidParamError("Y2")
		return
	} else if r.Empty() {
		return
	}

	x1 := int(*X1)
	y1 := int(*Y1)
	x2 := int(*X2)
	y2 := int(*Y2)
	rectx1 := int(r.X)
	recty1 := int(r.Y)
	rectx2 := int(r.X + r.W - 1)
	recty2 := int(r.Y + r.H - 1)

	/* Check to see if entire line is inside rect */
	if x1 >= rectx1 && x1 <= rectx2 && x2 >= rectx1 && x2 <= rectx2 &&
		y1 >= recty1 && y1 <= recty2 && y2 >= recty1 && y2 <= recty2 {
		intersects = true
		return
	}

	/* Check to see if entire line is to one side of rect */
	if (x1 < rectx1 && x2 < rectx1) || (x1 > rectx2 && x2 > rectx2) ||
		(y1 < recty1 && y2 < recty1) || (y1 > recty2 && y2 > recty2) {
		intersects = false
		return
	}

	if y1 == y2 { /* Horizontal line, easy to clip */
		if x1 < rectx1 {
			*X1 = int32(rectx1)
		} else if x1 > rectx2 {
			*X1 = int32(rectx2)
		}
		if x2 < rectx1 {
			*X2 = int32(rectx1)
		} else if x2 > rectx2 {
			*X2 = int32(rectx2)
		}
		intersects = true
		return
	}

	if x1 == x2 { /* Vertical line, easy to clip */
		if y1 < recty1 {
			*Y1 = int32(recty1)
		} else if y1 > recty2 {
			*Y1 = int32(recty2)
		}
		if y2 < recty1 {
			*Y2 = int32(recty1)
		} else if y2 > recty2 {
			*Y2 = int32(recty2)
		}
		intersects = true
		return
	}

	/* More complicated Cohen-Sutherland algorithm */
	outcode1 := computeOutCode(r, int32(x1), int32(y1))
	outcode2 := computeOutCode(r, int32(x2), int32(y2))
	for outcode1 != 0 || outcode2 != 0 {
		if outcode1&outcode2 != 0 {
			intersects = false
			return
		}

		if outcode1 != 0 {
			if outcode1&code_TOP != 0 {
				y = recty1
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode1&code_BOTTOM != 0 {
				y = recty2
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode1&code_LEFT != 0 {
				x = rectx1
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			} else if outcode1&code_RIGHT != 0 {
				x = rectx2
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			}
			x1 = x
			y1 = y
			outcode1 = computeOutCode(r, int32(x), int32(y))
		} else {
			if outcode2&code_TOP != 0 {
				// SDL_assert(y2 != y1) /* if equal: division by zero. */
				y = recty1
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode2&code_BOTTOM != 0 {
				// SDL_assert(y2 != y1) /* if equal: division by zero. */
				y = recty2
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode2&code_LEFT != 0 {
				/* If this assertion ever fires, here's the static analysis that warned about it:
				   http://buildbot.libsdl.org/sdl-static-analysis/sdl-macosx-static-analysis/sdl-macosx-static-analysis-1101/report-b0d01a.html#EndPath */
				// SDL_assert(x2 != x1) /* if equal: division by zero. */
				x = rectx1
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			} else if outcode2&code_RIGHT != 0 {
				/* If this assertion ever fires, here's the static analysis that warned about it:
				   http://buildbot.libsdl.org/sdl-static-analysis/sdl-macosx-static-analysis/sdl-macosx-static-analysis-1101/report-39b114.html#EndPath */
				// SDL_assert(x2 != x1) /* if equal: division by zero. */
				x = rectx2
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			}
			x2 = x
			y2 = y
			outcode2 = computeOutCode(r, int32(x), int32(y))
		}
	}
	*X1 = int32(x1)
	*Y1 = int32(y1)
	*X2 = int32(x2)
	*Y2 = int32(y2)
	intersects = true
	return
}

/* SDL_FRect versions... */

// Determine whether a point resides inside a floating point rectangle.
// (https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat)
func (p *FPoint) InRect(r *FRect) bool {
	return (p != nil && r != nil &&
		p.X >= r.X && p.X < (r.X+r.W) &&
		p.Y >= r.Y && p.Y < (r.Y+r.H))
}

// Determine whether a floating point rectangle has no area.
// (https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat)
func (r *FRect) Empty() bool {
	return r == nil || r.W <= 0 || r.H <= 0
}

func fabs32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

// Determine whether two floating point rectangles are equal, within some
// given epsilon.
// (https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon)
func (a *FRect) EqualEpsilon(b *FRect, epsilon float32) bool {
	return (a != nil) && (b != nil) && ((a == b) ||
		(fabs32(a.X-b.X) <= epsilon &&
			fabs32(a.Y-b.Y) <= epsilon &&
			fabs32(a.W-b.W) <= epsilon &&
			fabs32(a.H-b.H) <= epsilon))
}

// Determine whether two floating point rectangles are equal, within a default
// epsilon.
// (https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat)
func (a *FRect) Equal(b *FRect) bool {
	return a.EqualEpsilon(b, FLT_EPSILON)
}

// Determine whether two rectangles intersect with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat)
func (a *FRect) HasIntersection(b *FRect) bool {
	if a == nil || a.W <= 0 || a.H <= 0 {
		return false
	}
	if b == nil || b.W <= 0 || b.H <= 0 {
		return false
	}
	if a.X >= b.X+b.W || a.X+a.W <= b.X || a.Y >= b.Y+b.H || a.Y+a.H <= b.Y {
		return false
	}
	return true
}

// Calculate the intersection of two rectangles with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat)
func (a *FRect) GetIntersection(b *FRect) (result FRect, intersects bool) {
	/* Special cases for empty rects */
	if a == nil || a.W <= 0 || a.H <= 0 {
		return
	}
	if b == nil || b.W <= 0 || b.H <= 0 {
		return
	}

	/* Horizontal intersection */
	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin > aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	/* Vertical intersection */
	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin > aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return result, result.W > 0 && result.H > 0
}

// Calculate the union of two rectangles with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat)
func (a *FRect) GetUnion(b *FRect) (result FRect, err error) {
	if a == nil {
		err = InvalidParamError("a")
		return
	}
	if b == nil {
		err = InvalidParamError("b")
	}

	/* Special cases for empty Rects */
	if a.W == 0 || a.H == 0 {
		if b.W == 0 || b.H == 0 {
			return
		}
		result = *b
		return
	}
	if b.W == 0 || b.H == 0 {
		result = *a
		return
	}

	/* Horizontal union */
	Amin := a.X
	Amax := Amin + a.W
	Bmin := b.X
	Bmax := Bmin + b.W
	if Bmin < Amin {
		Amin = Bmin
	}
	result.X = Amin
	if Bmax > Amax {
		Amax = Bmax
	}
	result.W = Amax - Amin

	/* Vertical union */
	Amin = a.Y
	Amax = Amin + a.H
	Bmin = b.Y
	Bmax = Bmin + b.H
	if Bmin < Amin {
		Amin = Bmin
	}
	result.Y = Amin
	if Bmax > Amax {
		Amax = Bmax
	}
	result.H = Amax - Amin
	return
}

// Calculate a minimal rectangle enclosing a set of points with float
// precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat)
func GetRectEnclosingPointsFloat(points []FPoint, clip *FRect) (result FRect, err error) {
	if len(points) <= 0 {
		err = InvalidParamError("points")
		return
	}

	var minx, miny, maxx, maxy float32
	var x, y float32
	if clip != nil {
		if clip.Empty() {
			return
		}
		clip_minx := clip.X
		clip_miny := clip.Y
		clip_maxx := clip.X + clip.W - 1
		clip_maxy := clip.Y + clip.H - 1

		var added bool
		for i := 0; i < len(points); i++ {
			x = points[i].X
			y = points[i].Y

			if x < clip_minx || x > clip_maxx ||
				y < clip_miny || y > clip_maxy {
				continue
			}

			if !added {
				minx = x
				maxx = x
				miny = y
				maxy = y
				added = true
			}

			if x < minx {
				minx = x
			} else if x > maxx {
				maxx = x
			}
			if y < miny {
				miny = y
			} else if y > maxy {
				maxy = y
			}
		}
		return
	} else {
		minx = points[0].X
		maxx = minx
		miny = points[0].Y
		maxy = miny

		for i := 1; i < len(points); i++ {
			x = points[i].X
			y = points[i].Y

			if x < minx {
				minx = x
			} else if x > maxx {
				maxx = x
			}
			if y < miny {
				miny = y
			} else if y > maxy {
				maxy = y
			}
		}
	}
	result.X = minx
	result.Y = miny
	result.W = (maxx - minx) + 1
	result.H = (maxy - miny) + 1
	return
}

func computeOutCodeF(r *FRect, x, y float32) int {
	code := 0
	if y < r.Y {
		code |= code_TOP
	} else if y >= r.Y+r.H {
		code |= code_BOTTOM
	}
	if x < r.X {
		code |= code_LEFT
	} else if x >= r.X+r.W {
		code |= code_RIGHT
	}
	return code
}

// Calculate the intersection of a rectangle and line segment with float
// precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat)
func (r *FRect) GetLineIntersection(X1, Y1, X2, Y2 *float32) (intersects bool, err error) {
	var x, y float64

	if r == nil {
		err = InvalidParamError("rect")
		return
	} else if X1 == nil {
		err = InvalidParamError("X1")
		return
	} else if Y1 == nil {
		err = InvalidParamError("Y1")
		return
	} else if X2 == nil {
		err = InvalidParamError("X2")
		return
	} else if Y2 == nil {
		err = InvalidParamError("Y2")
		return
	} else if r.Empty() {
		return
	}

	x1 := float64(*X1)
	y1 := float64(*Y1)
	x2 := float64(*X2)
	y2 := float64(*Y2)
	rectx1 := float64(r.X)
	recty1 := float64(r.Y)
	rectx2 := float64(r.X + r.W - 1)
	recty2 := float64(r.Y + r.H - 1)

	/* Check to see if entire line is inside rect */
	if x1 >= rectx1 && x1 <= rectx2 && x2 >= rectx1 && x2 <= rectx2 &&
		y1 >= recty1 && y1 <= recty2 && y2 >= recty1 && y2 <= recty2 {
		intersects = true
		return
	}

	/* Check to see if entire line is to one side of rect */
	if (x1 < rectx1 && x2 < rectx1) || (x1 > rectx2 && x2 > rectx2) ||
		(y1 < recty1 && y2 < recty1) || (y1 > recty2 && y2 > recty2) {
		intersects = false
		return
	}

	if y1 == y2 { /* Horizontal line, easy to clip */
		if x1 < rectx1 {
			*X1 = float32(rectx1)
		} else if x1 > rectx2 {
			*X1 = float32(rectx2)
		}
		if x2 < rectx1 {
			*X2 = float32(rectx1)
		} else if x2 > rectx2 {
			*X2 = float32(rectx2)
		}
		intersects = true
		return
	}

	if x1 == x2 { /* Vertical line, easy to clip */
		if y1 < recty1 {
			*Y1 = float32(recty1)
		} else if y1 > recty2 {
			*Y1 = float32(recty2)
		}
		if y2 < recty1 {
			*Y2 = float32(recty1)
		} else if y2 > recty2 {
			*Y2 = float32(recty2)
		}
		intersects = true
		return
	}

	/* More complicated Cohen-Sutherland algorithm */
	outcode1 := computeOutCodeF(r, float32(x1), float32(y1))
	outcode2 := computeOutCodeF(r, float32(x2), float32(y2))
	for outcode1 != 0 || outcode2 != 0 {
		if outcode1&outcode2 != 0 {
			intersects = false
			return
		}

		if outcode1 != 0 {
			if outcode1&code_TOP != 0 {
				y = recty1
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode1&code_BOTTOM != 0 {
				y = recty2
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode1&code_LEFT != 0 {
				x = rectx1
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			} else if outcode1&code_RIGHT != 0 {
				x = rectx2
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			}
			x1 = x
			y1 = y
			outcode1 = computeOutCodeF(r, float32(x), float32(y))
		} else {
			if outcode2&code_TOP != 0 {
				// SDL_assert(y2 != y1) /* if equal: division by zero. */
				y = recty1
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode2&code_BOTTOM != 0 {
				// SDL_assert(y2 != y1) /* if equal: division by zero. */
				y = recty2
				x = (x1 + ((x2-x1)*(y-y1))/(y2-y1))
			} else if outcode2&code_LEFT != 0 {
				/* If this assertion ever fires, here's the static analysis that warned about it:
				   http://buildbot.libsdl.org/sdl-static-analysis/sdl-macosx-static-analysis/sdl-macosx-static-analysis-1101/report-b0d01a.html#EndPath */
				// SDL_assert(x2 != x1) /* if equal: division by zero. */
				x = rectx1
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			} else if outcode2&code_RIGHT != 0 {
				/* If this assertion ever fires, here's the static analysis that warned about it:
				   http://buildbot.libsdl.org/sdl-static-analysis/sdl-macosx-static-analysis/sdl-macosx-static-analysis-1101/report-39b114.html#EndPath */
				// SDL_assert(x2 != x1) /* if equal: division by zero. */
				x = rectx2
				y = (y1 + ((y2-y1)*(x-x1))/(x2-x1))
			}
			x2 = x
			y2 = y
			outcode2 = computeOutCodeF(r, float32(x), float32(y))
		}
	}
	*X1 = float32(x1)
	*Y1 = float32(y1)
	*X2 = float32(x2)
	*Y2 = float32(y2)
	intersects = true
	return
}
