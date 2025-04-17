package sdl

//#include <SDL3/SDL_surface.h>
import "C"
import "unsafe"

// The flags on an SDL_Surface.
// (https://wiki.libsdl.org/SDL3/SDL_SurfaceFlags)
type SurfaceFlags C.SDL_SurfaceFlags

const (
	SURFACE_PREALLOCATED = SurfaceFlags(C.SDL_SURFACE_PREALLOCATED) /**< Surface uses preallocated pixel memory */
	SURFACE_LOCK_NEEDED  = SurfaceFlags(C.SDL_SURFACE_LOCK_NEEDED)  /**< Surface needs to be locked to access pixels */
	SURFACE_LOCKED       = SurfaceFlags(C.SDL_SURFACE_LOCKED)       /**< Surface is currently locked */
	SURFACE_SIMD_ALIGNED = SurfaceFlags(C.SDL_SURFACE_SIMD_ALIGNED) /**< Surface uses pixel memory allocated with SDL_aligned_alloc() */
)

// Evaluates to true if the surface needs to be locked before access.
// #define SDL_MUSTLOCK(S) (((S)->flags & (SDL_SURFACE_LOCK_NEEDED | SDL_SURFACE_LOCKED)) == SDL_SURFACE_LOCK_NEEDED)

// The scaling mode.
// (https://wiki.libsdl.org/SDL3/SDL_ScaleMode)
type ScaleMode C.SDL_ScaleMode

func (mode ScaleMode) c() C.SDL_ScaleMode {
	return C.SDL_ScaleMode(mode)
}
func (mode *ScaleMode) cptr() *C.SDL_ScaleMode {
	return (*C.SDL_ScaleMode)(mode)
}

const (
	SCALEMODE_NEAREST = ScaleMode(C.SDL_SCALEMODE_NEAREST) /**< nearest pixel sampling */
	SCALEMODE_LINEAR  = ScaleMode(C.SDL_SCALEMODE_LINEAR)  /**< linear filtering */
)

// The flip mode.
// (https://wiki.libsdl.org/SDL3/SDL_FlipMode)
type FlipMode C.SDL_FlipMode

func (mode FlipMode) c() C.SDL_FlipMode {
	return C.SDL_FlipMode(mode)
}

const (
	FLIP_NONE       = FlipMode(C.SDL_FLIP_NONE)       /**< Do not flip */
	FLIP_HORIZONTAL = FlipMode(C.SDL_FLIP_HORIZONTAL) /**< flip horizontally */
	FLIP_VERTICAL   = FlipMode(C.SDL_FLIP_VERTICAL)   /**< flip vertically */
)

//TODO: removed
// // Internal surface data
// type SurfaceData C.SDL_SurfaceData

// A collection of pixels used in software blitting.
// (https://wiki.libsdl.org/SDL3/SDL_Surface)
type Surface struct {
	Flags  SurfaceFlags   /**< Read-only */
	Format PixelFormat    /**< Read-only */
	W, H   int32          /**< Read-only */
	Pitch  int32          /**< Read-only */
	Pixels unsafe.Pointer /**< Read-only pointer, writable pixels if non-NULL */

	refcount C.int /**< Application reference count, used when freeing surface */

	internal unsafe.Pointer /**< Private */
}

func (surface *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

// Allocate a new surface with a specific pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurface)
func CreateSurface(width, height int32, format PixelFormat) (surface *Surface, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_CreateSurface(C.int(width), C.int(height), format.c())))
	if surface == nil {
		err = GetError()
	}
	return
}

// Allocate a new surface with a specific pixel format and existing pixel
// data.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom)
func CreateSurfaceFrom(width, height int32, format PixelFormat, pixels unsafe.Pointer, pitch int32) (surface *Surface, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_CreateSurfaceFrom(C.int(width), C.int(height), format.c(), pixels, C.int(pitch))))
	if surface == nil {
		err = GetError()
	}
	return
}

// Free a surface.
// (https://wiki.libsdl.org/SDL3/SDL_DestroySurface)
func (surface *Surface) Destroy() {
	C.SDL_DestroySurface(surface.cptr())
	surface = nil
}

// Get the properties associated with a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties)
func (surface *Surface) GetProperties() (id PropertiesID, err error) {
	id = PropertiesID(C.SDL_GetSurfaceProperties(surface.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

const (
	PROP_SURFACE_COLORSPACE_NUMBER       = "SDL.surface.colorspace"
	PROP_SURFACE_SDR_WHITE_POINT_FLOAT   = "SDL.surface.SDR_white_point"
	PROP_SURFACE_HDR_HEADROOM_FLOAT      = "SDL.surface.HDR_headroom"
	PROP_SURFACE_TONEMAP_OPERATOR_STRING = "SDL.surface.tonemap"
)

// Set the colorspace used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace)
func (surface *Surface) SetColorspace(colorspace Colorspace) (err error) {
	ret := C.SDL_SetSurfaceColorspace(surface.cptr(), C.SDL_Colorspace(colorspace))
	if !ret {
		err = GetError()
	}
	return
}

// Get the colorspace used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace)
func (surface *Surface) GetColorspace() Colorspace {
	return Colorspace(C.SDL_GetSurfaceColorspace(surface.cptr()))
}

// Create a palette and associate it with a surface.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette)
func (surface *Surface) CreatePalette() (palette *Palette, err error) {
	palette = (*Palette)(unsafe.Pointer(C.SDL_CreateSurfacePalette(surface.cptr())))
	if palette == nil {
		err = GetError()
	}
	return
}

// Set the palette used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette)
func (surface *Surface) SetPalette(palette *Palette) (err error) {
	ret := C.SDL_SetSurfacePalette(surface.cptr(), palette.cptr())
	if !ret {
		err = GetError()
	}
	return
}

// Get the palette used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette)
func (surface *Surface) GetPalette() *Palette {
	return (*Palette)(unsafe.Pointer(C.SDL_GetSurfacePalette(surface.cptr())))
}

// Set up a surface for directly accessing the pixels.
// (https://wiki.libsdl.org/SDL3/SDL_LockSurface)
func (surface *Surface) Lock() (err error) {
	ret := C.SDL_LockSurface(surface.cptr())
	if !ret {
		err = GetError()
	}
	return
}

// Release a surface after directly accessing the pixels.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockSurface)
func (surface *Surface) Unlock() {
	C.SDL_UnlockSurface(surface.cptr())
}

// Load a BMP image from a seekable SDL data stream.
// (https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO)
func (src *IOStream) LoadBMP(closeio bool) (surface *Surface, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_LoadBMP_IO(src.cptr(), (C.bool)(closeio))))
	if surface == nil {
		err = GetError()
	}
	return
}

// Load a BMP image from a file.
// (https://wiki.libsdl.org/SDL3/SDL_LoadBMP)
func LoadBMP(file string) (surface *Surface, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_LoadBMP(C.CString(file))))
	if surface == nil {
		err = GetError()
	}
	return
}

// Save a surface to a seekable SDL data stream in BMP format.
// (https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO)
func (surface *Surface) SaveBMP_IO(dst *IOStream, closeio bool) (err error) {
	ret := C.SDL_SaveBMP_IO(surface.cptr(), dst.cptr(), (C.bool)(closeio))
	if !ret {
		err = GetError()
	}
	return
}

// Save a surface to a file.
// (https://wiki.libsdl.org/SDL3/SDL_SaveBMP)
func (surface *Surface) SaveBMP(file string) (err error) {
	ret := C.SDL_SaveBMP(surface.cptr(), C.CString(file))
	if !ret {
		err = GetError()
	}
	return
}

// Set the RLE acceleration hint for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE)
func (surface *Surface) SetRLE(enabled bool) (err error) {
	ret := C.SDL_SetSurfaceRLE(surface.cptr(), (C.bool)(enabled))
	if !ret {
		err = GetError()
	}
	return
}

// Returns whether the surface is RLE enabled.
// (https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE)
func (surface *Surface) HasRLE() bool {
	return bool(C.SDL_SurfaceHasRLE(surface.cptr()))
}

// Set the color key (transparent pixel) in a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey)
func (surface *Surface) SetColorKey(enabled bool, key uint32) (err error) {
	ret := C.SDL_SetSurfaceColorKey(surface.cptr(), (C.bool)(enabled), C.Uint32(key))
	if !ret {
		err = GetError()
	}
	return
}

// Returns whether the surface has a color key.
// (https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey)
func (surface *Surface) HasColorKey() bool {
	return bool(C.SDL_SurfaceHasColorKey(surface.cptr()))
}

// Get the color key (transparent pixel) for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey)
func (surface *Surface) GetColorKey() (key uint32, err error) {
	ret := C.SDL_GetSurfaceColorKey(surface.cptr(), (*C.Uint32)(&key))
	if !ret {
		err = GetError()
	}
	return
}

// Set an additional color value multiplied into blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod)
func (surface *Surface) SetColorMod(r, g, b uint8) (err error) {
	ret := C.SDL_SetSurfaceColorMod(surface.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b))
	if !ret {
		err = GetError()
	}
	return
}

// Get the additional color value multiplied into blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod)
func (surface *Surface) GetColorMod() (r, g, b uint8, err error) {
	ret := C.SDL_GetSurfaceColorMod(surface.cptr(), (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	if !ret {
		err = GetError()
	}
	return
}

// Set an additional alpha value used in blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod)
func (surface *Surface) SetAlphaMod(alpha uint8) (err error) {
	ret := C.SDL_SetSurfaceAlphaMod(surface.cptr(), C.Uint8(alpha))
	if !ret {
		err = GetError()
	}
	return
}

// Get the additional alpha value used in blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod)
func (surface *Surface) GetAlphaMod() (alpha uint8, err error) {
	ret := C.SDL_GetSurfaceAlphaMod(surface.cptr(), (*C.Uint8)(&alpha))
	if !ret {
		err = GetError()
	}
	return
}

// Set the blend mode used for blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode)
func (surface *Surface) SetBlendMode(blendMode BlendMode) (err error) {
	ret := C.SDL_SetSurfaceBlendMode(surface.cptr(), C.SDL_BlendMode(blendMode))
	if !ret {
		err = GetError()
	}
	return
}

// Get the blend mode used for blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode)
func (surface *Surface) GetBlendMode() (blendmode BlendMode, err error) {
	ret := C.SDL_GetSurfaceBlendMode(surface.cptr(), (*C.SDL_BlendMode)(&blendmode))
	if !ret {
		err = GetError()
	}
	return
}

// Set the clipping rectangle for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect)
func (surface *Surface) SetClipRect(rect *Rect) bool {
	return bool(C.SDL_SetSurfaceClipRect(surface.cptr(), rect.cptr()))
}

// Get the clipping rectangle for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect)
func (surface *Surface) GetClipRect() (rect Rect, err error) {
	ret := C.SDL_GetSurfaceClipRect(surface.cptr(), rect.cptr())
	if !ret {
		err = GetError()
	}
	return
}

// Flip a surface vertically or horizontally.
// (https://wiki.libsdl.org/SDL3/SDL_FlipSurface)
func (surface *Surface) Flip(flip FlipMode) (err error) {
	ret := C.SDL_FlipSurface(surface.cptr(), C.SDL_FlipMode(flip))
	if !ret {
		err = GetError()
	}
	return
}

// Creates a new surface identical to the existing surface.
// (https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface)
func (surface *Surface) Duplicate() (duplicate *Surface, err error) {
	duplicate = (*Surface)(unsafe.Pointer(C.SDL_DuplicateSurface(surface.cptr())))
	if duplicate == nil {
		err = GetError()
	}
	return
}

// Copy an existing surface to a new surface of the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertSurface)
func (surface *Surface) Convert(format PixelFormat) (newSurface *Surface, err error) {
	newSurface = (*Surface)(unsafe.Pointer(C.SDL_ConvertSurface(surface.cptr(), C.SDL_PixelFormat(format))))
	if newSurface == nil {
		err = GetError()
	}
	return
}

// Copy an existing surface to a new surface of the specified format and
// colorspace.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace)
func (surface *Surface) ConvertColorspace(format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) (newSurface *Surface, err error) {
	newSurface = (*Surface)(unsafe.Pointer(C.SDL_ConvertSurfaceAndColorspace(
		surface.cptr(), C.SDL_PixelFormat(format), palette.cptr(), C.SDL_Colorspace(colorspace), C.SDL_PropertiesID(props))))
	if newSurface == nil {
		err = GetError()
	}
	return
}

// Copy a block of pixels of one format to another format.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertPixels)
func ConvertPixels(width, height int32, src_format PixelFormat, src unsafe.Pointer, src_pitch int32,
	dst_format PixelFormat, dst unsafe.Pointer, dst_pitch int32) (err error) {
	ret := C.SDL_ConvertPixels(C.int(width), C.int(height), C.SDL_PixelFormat(src_format), src,
		C.int(src_pitch), C.SDL_PixelFormat(dst_format), dst, C.int(dst_pitch))
	if !ret {
		err = GetError()
	}
	return
}

// Copy a block of pixels of one format and colorspace to another format and
// colorspace.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace)
func ConvertPixelsAndColorspace(width, height int32, src_format PixelFormat, src_colorspace Colorspace,
	src_properties PropertiesID, src unsafe.Pointer, src_pitch int32, dst_format PixelFormat, dst_colorspace Colorspace,
	dst_properties PropertiesID, dst unsafe.Pointer, dst_pitch int32) (err error) {
	ret := C.SDL_ConvertPixelsAndColorspace(C.int(width), C.int(height), C.SDL_PixelFormat(src_format),
		C.SDL_Colorspace(src_colorspace), C.SDL_PropertiesID(src_properties), src, C.int(src_pitch), C.SDL_PixelFormat(dst_format),
		C.SDL_Colorspace(dst_colorspace), C.SDL_PropertiesID(dst_properties), dst, C.int(dst_pitch))
	if !ret {
		err = GetError()
	}
	return
}

// Premultiply the alpha on a block of pixels.
// (https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha)
func PremultiplyAlpha(width, height int32, src_format PixelFormat, src unsafe.Pointer, src_pitch int32,
	dst_format PixelFormat, dst unsafe.Pointer, dst_pitch int32, linear bool) (err error) {
	ret := C.SDL_PremultiplyAlpha(C.int(width), C.int(height), C.SDL_PixelFormat(src_format), src, C.int(src_pitch),
		C.SDL_PixelFormat(dst_format), dst, C.int(dst_pitch), (C.bool)(linear))
	if !ret {
		err = GetError()
	}
	return
}

// Premultiply the alpha in a surface.
// (https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha)
func (surface *Surface) PremultiplySurfaceAlpha(linera bool) (err error) {
	ret := C.SDL_PremultiplySurfaceAlpha(surface.cptr(), (C.bool)(linera))
	if !ret {
		err = GetError()
	}
	return
}

// Perform a fast fill of a rectangle with a specific color.
// (https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect)
func (surface *Surface) FillRect(rect *Rect, color uint32) (err error) {
	ret := C.SDL_FillSurfaceRect(surface.cptr(), rect.cptr(), C.Uint32(color))
	if !ret {
		err = GetError()
	}
	return
}

// Perform a fast fill of a set of rectangles with a specific color.
// (https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects)
func (surface *Surface) FillRects(rects []Rect, color uint32) (err error) {
	ret := C.SDL_FillSurfaceRects(surface.cptr(), rects[0].cptr(), C.int(len(rects)), C.Uint32(color))
	if !ret {
		err = GetError()
	}
	return
}

// Performs a fast blit from the source surface to the destination surface.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurface)
func (surface *Surface) Blit(srcRect *Rect, dst *Surface, dstRect *Rect) (err error) {
	ret := C.SDL_BlitSurface(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr())
	if !ret {
		err = GetError()
	}
	return
}

// Perform low-level surface blitting only.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked)
func (surface *Surface) BlitUnchecked(srcRect *Rect, dst *Surface, dstRect *Rect) (err error) {
	ret := C.SDL_BlitSurfaceUnchecked(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr())
	if !ret {
		err = GetError()
	}
	return
}

// Perform stretch blit between two surfaces of the same format.
// (https://wiki.libsdl.org/SDL3/SDL_SoftStretch)
// removed: https://github.com/libsdl-org/SDL/commit/12e50d17a2c1f3f906698b797beeddb9d7af5613
// func (surface *Surface) SoftStretch(srcRect *Rect, dst *Surface, dstRect *Rect, scaleMode ScaleMode) (err error) {
// 	ret := C.SDL_SoftStretch(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr(), C.SDL_ScaleMode(scaleMode))
// 	if !ret {
// 		err = GetError()
// 	}
// 	return
// }

// Perform a scaled blit to a destination surface, which may be of a different
// format.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled)
func (surface *Surface) BlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect, scaleMode ScaleMode) (err error) {
	ret := C.SDL_BlitSurfaceScaled(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr(), C.SDL_ScaleMode(scaleMode))
	if !ret {
		err = GetError()
	}
	return
}

// Perform low-level surface scaled blitting only.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled)
func (surface *Surface) BlitUncheckedScaled(srcRect *Rect, dst *Surface, dstRect *Rect, scaleMode ScaleMode) (err error) {
	ret := C.SDL_BlitSurfaceUncheckedScaled(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr(), C.SDL_ScaleMode(scaleMode))
	if !ret {
		err = GetError()
	}
	return
}

// Map an RGB triple to an opaque pixel value for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB)
func (surface *Surface) MapRGB(r, g, b uint8) uint32 {
	return uint32(C.SDL_MapSurfaceRGB(surface.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// Map an RGBA quadruple to a pixel value for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA)
func (surface *Surface) MapRGBA(r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapSurfaceRGBA(surface.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// Retrieves a single pixel from a surface.
// (https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel)
func (surface *Surface) ReadPixel(x, y int32) (r, g, b, a uint8, err error) {
	ret := C.SDL_ReadSurfacePixel(surface.cptr(), C.int(x), C.int(y),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	if !ret {
		err = GetError()
	}
	return
}
