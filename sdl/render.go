package sdl

/*
#include <SDL3/SDL_render.h>
#include "stdlib.h"
*/
import "C"
import (
	"unsafe"
)

// # CategoryRender
// (https://wiki.libsdl.org/SDL3/CategoryRender)

// The name of the software renderer.
const SOFTWARE_RENDERER = "software"

// Vertex structure.
// (https://wiki.libsdl.org/SDL3/SDL_Vertex)
type Vertex struct {
	Position FPoint
	Color    FColor
	TexCoord FPoint
}

// The access pattern allowed for a texture.
// (https://wiki.libsdl.org/SDL3/SDL_TextureAccess)
type TextureAccess C.SDL_TextureAccess

func (access TextureAccess) c() C.SDL_TextureAccess {
	return C.SDL_TextureAccess(access)
}

const (
	TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC    /**< Changes rarely, not lockable */
	TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING /**< Changes frequently, lockable */
	TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET    /**< Texture can be used as a render target */
)

// How the logical size is mapped to the output.
// (https://wiki.libsdl.org/SDL3/SDL_RendererLogicalPresentation)
type RendererLogicalPresentation C.SDL_RendererLogicalPresentation

func (presentation RendererLogicalPresentation) c() C.SDL_RendererLogicalPresentation {
	return C.SDL_RendererLogicalPresentation(presentation)
}
func (presentation *RendererLogicalPresentation) cptr() *C.SDL_RendererLogicalPresentation {
	return (*C.SDL_RendererLogicalPresentation)(presentation)
}

const (
	LOGICAL_PRESENTATION_DISABLED      = C.SDL_LOGICAL_PRESENTATION_DISABLED      /**< There is no logical size in effect */
	LOGICAL_PRESENTATION_STRETCH       = C.SDL_LOGICAL_PRESENTATION_STRETCH       /**< The rendered content is stretched to the output resolution */
	LOGICAL_PRESENTATION_LETTERBOX     = C.SDL_LOGICAL_PRESENTATION_LETTERBOX     /**< The rendered content is fit to the largest dimension and the other dimension is letterboxed with black bars */
	LOGICAL_PRESENTATION_OVERSCAN      = C.SDL_LOGICAL_PRESENTATION_OVERSCAN      /**< The rendered content is fit to the smallest dimension and the other dimension extends beyond the output bounds */
	LOGICAL_PRESENTATION_INTEGER_SCALE = C.SDL_LOGICAL_PRESENTATION_INTEGER_SCALE /**< The rendered content is scaled up by integer multiples to fit the output resolution */
)

// A structure representing rendering state
// (https://wiki.libsdl.org/SDL3/SDL_Renderer)
type Renderer C.SDL_Renderer

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(r)
}

// An efficient driver-specific representation of pixel data
// (https://wiki.libsdl.org/SDL3/SDL_Texture)
type Texture C.SDL_Texture

func (t *Texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(t)
}

/* Function prototypes */

// Get the number of 2D rendering drivers available for the current display.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers)
func GetNumRenderDrivers() (int32, error) {
	ret := C.SDL_GetNumRenderDrivers()
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Use this function to get the name of a built in 2D rendering driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver)
func GetRenderDriver(index int32) string {
	return C.GoString(C.SDL_GetRenderDriver(C.int(index)))
}

// Create a window and default renderer.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer)
func CreateWindowAndRenderer(title string, width, height int32, flags WindowFlags) (*Window, *Renderer, error) {
	var cTitle *C.char = C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	var cWindow *C.SDL_Window
	var cRenderer *C.SDL_Renderer
	ret := C.SDL_CreateWindowAndRenderer(cTitle, C.int(width), C.int(height), flags.c(), &cWindow, &cRenderer)
	if ret < 0 {
		return nil, nil, GetError()
	}
	return (*Window)(unsafe.Pointer(cWindow)), (*Renderer)(unsafe.Pointer(cRenderer)), nil
}

// Create a 2D rendering context for a window.
// (https://wiki.libsdl.org/SDL3/SDL_CreateRenderer)
func CreateRenderer(window *Window, name string) (*Renderer, error) {
	var cName *C.char
	if name != "" {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}
	cRenderer := C.SDL_CreateRenderer(window.cptr(), cName)
	if cRenderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(cRenderer)), nil
}

// Create a 2D rendering context for a window, with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties)
func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	cRenderer := C.SDL_CreateRendererWithProperties(props.c())
	if cRenderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(cRenderer)), nil
}

const (
	PROP_RENDERER_CREATE_NAME_STRING                               = "name"
	PROP_RENDERER_CREATE_WINDOW_POINTER                            = "window"
	PROP_RENDERER_CREATE_SURFACE_POINTER                           = "surface"
	PROP_RENDERER_CREATE_OUTPUT_COLORSPACE_NUMBER                  = "output_colorspace"
	PROP_RENDERER_CREATE_PRESENT_VSYNC_NUMBER                      = "present_vsync"
	PROP_RENDERER_CREATE_VULKAN_INSTANCE_POINTER                   = "vulkan.instance"
	PROP_RENDERER_CREATE_VULKAN_SURFACE_NUMBER                     = "vulkan.surface"
	PROP_RENDERER_CREATE_VULKAN_PHYSICAL_DEVICE_POINTER            = "vulkan.physical_device"
	PROP_RENDERER_CREATE_VULKAN_DEVICE_POINTER                     = "vulkan.device"
	PROP_RENDERER_CREATE_VULKAN_GRAPHICS_QUEUE_FAMILY_INDEX_NUMBER = "vulkan.graphics_queue_family_index"
	PROP_RENDERER_CREATE_VULKAN_PRESENT_QUEUE_FAMILY_INDEX_NUMBER  = "vulkan.present_queue_family_index"
)

// Create a 2D software rendering context for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer)
func CreateSoftwareRenderer(surface *Surface) (*Renderer, error) {
	cRenderer := C.SDL_CreateSoftwareRenderer(surface.cptr())
	if cRenderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(cRenderer)), nil
}

// Get the renderer associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderer)
func GetRenderer(window *Window) (*Renderer, error) {
	cRenderer := C.SDL_GetRenderer(window.cptr())
	if cRenderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(cRenderer)), nil
}

// Get the window associated with a renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow)
func (renderer *Renderer) GetWindow() (*Window, error) {
	cWindow := C.SDL_GetRenderWindow(renderer.cptr())
	if cWindow == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(cWindow)), nil
}

// Get the name of a renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRendererName)
func (renderer *Renderer) GetName() (string, error) {
	cName := C.SDL_GetRendererName(renderer.cptr())
	if cName == nil {
		return "", GetError()
	}
	return C.GoString(cName), nil
}

// Get the properties associated with a renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties)
func (renderer *Renderer) GetProperties() (PropertiesID, error) {
	ret := C.SDL_GetRendererProperties(renderer.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

const (
	PROP_RENDERER_NAME_STRING                               = "SDL.renderer.name"
	PROP_RENDERER_WINDOW_POINTER                            = "SDL.renderer.window"
	PROP_RENDERER_SURFACE_POINTER                           = "SDL.renderer.surface"
	PROP_RENDERER_VSYNC_NUMBER                              = "SDL.renderer.vsync"
	PROP_RENDERER_MAX_TEXTURE_SIZE_NUMBER                   = "SDL.renderer.max_texture_size"
	PROP_RENDERER_TEXTURE_FORMATS_POINTER                   = "SDL.renderer.texture_formats"
	PROP_RENDERER_OUTPUT_COLORSPACE_NUMBER                  = "SDL.renderer.output_colorspace"
	PROP_RENDERER_HDR_ENABLED_BOOLEAN                       = "SDL.renderer.HDR_enabled"
	PROP_RENDERER_SDR_WHITE_POINT_FLOAT                     = "SDL.renderer.SDR_white_point"
	PROP_RENDERER_HDR_HEADROOM_FLOAT                        = "SDL.renderer.HDR_headroom"
	PROP_RENDERER_D3D9_DEVICE_POINTER                       = "SDL.renderer.d3d9.device"
	PROP_RENDERER_D3D11_DEVICE_POINTER                      = "SDL.renderer.d3d11.device"
	PROP_RENDERER_D3D11_SWAPCHAIN_POINTER                   = "SDL.renderer.d3d11.swap_chain"
	PROP_RENDERER_D3D12_DEVICE_POINTER                      = "SDL.renderer.d3d12.device"
	PROP_RENDERER_D3D12_SWAPCHAIN_POINTER                   = "SDL.renderer.d3d12.swap_chain"
	PROP_RENDERER_D3D12_COMMAND_QUEUE_POINTER               = "SDL.renderer.d3d12.command_queue"
	PROP_RENDERER_VULKAN_INSTANCE_POINTER                   = "SDL.renderer.vulkan.instance"
	PROP_RENDERER_VULKAN_SURFACE_NUMBER                     = "SDL.renderer.vulkan.surface"
	PROP_RENDERER_VULKAN_PHYSICAL_DEVICE_POINTER            = "SDL.renderer.vulkan.physical_device"
	PROP_RENDERER_VULKAN_DEVICE_POINTER                     = "SDL.renderer.vulkan.device"
	PROP_RENDERER_VULKAN_GRAPHICS_QUEUE_FAMILY_INDEX_NUMBER = "SDL.renderer.vulkan.graphics_queue_family_index"
	PROP_RENDERER_VULKAN_PRESENT_QUEUE_FAMILY_INDEX_NUMBER  = "SDL.renderer.vulkan.present_queue_family_index"
	PROP_RENDERER_VULKAN_SWAPCHAIN_IMAGE_COUNT_NUMBER       = "SDL.renderer.vulkan.swapchain_image_count"
)

// Get the output size in pixels of a rendering context.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize)
func (renderer *Renderer) GetOutputSize() (w, h int32, err error) {
	ret := C.SDL_GetRenderOutputSize(renderer.cptr(), (*C.int)(&w), (*C.int)(&h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the current output size in pixels of a rendering context.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize)
func (renderer *Renderer) GetCurrentOutputSize() (w, h int32, err error) {
	ret := C.SDL_GetCurrentRenderOutputSize(renderer.cptr(), (*C.int)(&w), (*C.int)(&h))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Create a texture for a rendering context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTexture)
func (renderer *Renderer) CreateTexture(format PixelFormat, access TextureAccess, w, h int32) (*Texture, error) {
	cTexture := C.SDL_CreateTexture(renderer.cptr(), format.c(), access.c(), C.int(w), C.int(h))
	if cTexture == nil {
		return nil, GetError()
	}
	return (*Texture)(cTexture), nil
}

// Create a texture from an existing surface.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface)
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	cTexture := C.SDL_CreateTextureFromSurface(renderer.cptr(), surface.cptr())
	if cTexture == nil {
		return nil, GetError()
	}
	return (*Texture)(cTexture), nil
}

// Create a texture for a rendering context with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties)
func (renderer *Renderer) CreateTextureWithProperties(props PropertiesID) (*Texture, error) {
	cTexture := C.SDL_CreateTextureWithProperties(renderer.cptr(), props.c())
	if cTexture == nil {
		return nil, GetError()
	}
	return (*Texture)(cTexture), nil
}

const (
	PROP_TEXTURE_CREATE_COLORSPACE_NUMBER           = "colorspace"
	PROP_TEXTURE_CREATE_FORMAT_NUMBER               = "format"
	PROP_TEXTURE_CREATE_ACCESS_NUMBER               = "access"
	PROP_TEXTURE_CREATE_WIDTH_NUMBER                = "width"
	PROP_TEXTURE_CREATE_HEIGHT_NUMBER               = "height"
	PROP_TEXTURE_CREATE_SDR_WHITE_POINT_FLOAT       = "SDR_white_point"
	PROP_TEXTURE_CREATE_HDR_HEADROOM_FLOAT          = "HDR_headroom"
	PROP_TEXTURE_CREATE_D3D11_TEXTURE_POINTER       = "d3d11.texture"
	PROP_TEXTURE_CREATE_D3D11_TEXTURE_U_POINTER     = "d3d11.texture_u"
	PROP_TEXTURE_CREATE_D3D11_TEXTURE_V_POINTER     = "d3d11.texture_v"
	PROP_TEXTURE_CREATE_D3D12_TEXTURE_POINTER       = "d3d12.texture"
	PROP_TEXTURE_CREATE_D3D12_TEXTURE_U_POINTER     = "d3d12.texture_u"
	PROP_TEXTURE_CREATE_D3D12_TEXTURE_V_POINTER     = "d3d12.texture_v"
	PROP_TEXTURE_CREATE_METAL_PIXELBUFFER_POINTER   = "metal.pixelbuffer"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_NUMBER       = "opengl.texture"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_UV_NUMBER    = "opengl.texture_uv"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_U_NUMBER     = "opengl.texture_u"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_V_NUMBER     = "opengl.texture_v"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_NUMBER    = "opengles2.texture"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_UV_NUMBER = "opengles2.texture_uv"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_U_NUMBER  = "opengles2.texture_u"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_V_NUMBER  = "opengles2.texture_v"
	PROP_TEXTURE_CREATE_VULKAN_TEXTURE_NUMBER       = "vulkan.texture"
)

// Get the properties associated with a texture.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties)
func (texture *Texture) GetProperties() (props PropertiesID, err error) {
	ret := C.SDL_GetTextureProperties(texture.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

const (
	PROP_TEXTURE_COLORSPACE_NUMBER               = "SDL.texture.colorspace"
	PROP_TEXTURE_FORMAT_NUMBER                   = "SDL.texture.format"
	PROP_TEXTURE_ACCESS_NUMBER                   = "SDL.texture.access"
	PROP_TEXTURE_WIDTH_NUMBER                    = "SDL.texture.width"
	PROP_TEXTURE_HEIGHT_NUMBER                   = "SDL.texture.height"
	PROP_TEXTURE_SDR_WHITE_POINT_FLOAT           = "SDL.texture.SDR_white_point"
	PROP_TEXTURE_HDR_HEADROOM_FLOAT              = "SDL.texture.HDR_headroom"
	PROP_TEXTURE_D3D11_TEXTURE_POINTER           = "SDL.texture.d3d11.texture"
	PROP_TEXTURE_D3D11_TEXTURE_U_POINTER         = "SDL.texture.d3d11.texture_u"
	PROP_TEXTURE_D3D11_TEXTURE_V_POINTER         = "SDL.texture.d3d11.texture_v"
	PROP_TEXTURE_D3D12_TEXTURE_POINTER           = "SDL.texture.d3d12.texture"
	PROP_TEXTURE_D3D12_TEXTURE_U_POINTER         = "SDL.texture.d3d12.texture_u"
	PROP_TEXTURE_D3D12_TEXTURE_V_POINTER         = "SDL.texture.d3d12.texture_v"
	PROP_TEXTURE_OPENGL_TEXTURE_NUMBER           = "SDL.texture.opengl.texture"
	PROP_TEXTURE_OPENGL_TEXTURE_UV_NUMBER        = "SDL.texture.opengl.texture_uv"
	PROP_TEXTURE_OPENGL_TEXTURE_U_NUMBER         = "SDL.texture.opengl.texture_u"
	PROP_TEXTURE_OPENGL_TEXTURE_V_NUMBER         = "SDL.texture.opengl.texture_v"
	PROP_TEXTURE_OPENGL_TEXTURE_TARGET_NUMBER    = "SDL.texture.opengl.target"
	PROP_TEXTURE_OPENGL_TEX_W_FLOAT              = "SDL.texture.opengl.tex_w"
	PROP_TEXTURE_OPENGL_TEX_H_FLOAT              = "SDL.texture.opengl.tex_h"
	PROP_TEXTURE_OPENGLES2_TEXTURE_NUMBER        = "SDL.texture.opengles2.texture"
	PROP_TEXTURE_OPENGLES2_TEXTURE_UV_NUMBER     = "SDL.texture.opengles2.texture_uv"
	PROP_TEXTURE_OPENGLES2_TEXTURE_U_NUMBER      = "SDL.texture.opengles2.texture_u"
	PROP_TEXTURE_OPENGLES2_TEXTURE_V_NUMBER      = "SDL.texture.opengles2.texture_v"
	PROP_TEXTURE_OPENGLES2_TEXTURE_TARGET_NUMBER = "SDL.texture.opengles2.target"
	PROP_TEXTURE_VULKAN_TEXTURE_NUMBER           = "SDL.texture.vulkan.texture"
)

// Get the renderer that created an SDL_Texture.
// (https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture)
func (texture *Texture) GetRenderer() (*Renderer, error) {
	cRenderer := C.SDL_GetRendererFromTexture(texture.cptr())
	if cRenderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(cRenderer), nil
}

// Get the size of a texture, as floating point values.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureSize)
func (texture *Texture) GetSize() (w, h float32, err error) {
	ret := C.SDL_GetTextureSize(texture.cptr(), (*C.float)(&w), (*C.float)(&h))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set an additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod)
func (texture *Texture) SetColorMod(r, g, b uint8) error {
	ret := C.SDL_SetTextureColorMod(texture.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set an additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat)
func (texture *Texture) SetColorModFloat(r, g, b float32) error {
	ret := C.SDL_SetTextureColorModFloat(texture.cptr(), C.float(r), C.float(g), C.float(b))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod)
func (texture *Texture) GetColorMod() (r, g, b uint8, err error) {
	ret := C.SDL_GetTextureColorMod(texture.cptr(), (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get the additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat)
func (texture *Texture) GetColorModFloat() (r, g, b float32, err error) {
	ret := C.SDL_GetTextureColorModFloat(texture.cptr(), (*C.float)(&r), (*C.float)(&g), (*C.float)(&b))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set an additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod)
func (texture *Texture) SetAlphaMod(alpha uint8) error {
	ret := C.SDL_SetTextureAlphaMod(texture.cptr(), C.Uint8(alpha))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set an additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat)
func (texture *Texture) SetAlphaModFloat(alpha float32) error {
	ret := C.SDL_SetTextureAlphaModFloat(texture.cptr(), C.float(alpha))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod)
func (texture *Texture) GetAlphaMod() (alpha uint8, err error) {
	ret := C.SDL_GetTextureAlphaMod(texture.cptr(), (*C.Uint8)(&alpha))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get the additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat)
func (texture *Texture) GetAlphaModFloat() (alpha float32, err error) {
	ret := C.SDL_GetTextureAlphaModFloat(texture.cptr(), (*C.float)(&alpha))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set the blend mode for a texture, used by SDL_RenderTexture().
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode)
func (texture *Texture) SetBlendMode(blendMode BlendMode) error {
	ret := C.SDL_SetTextureBlendMode(texture.cptr(), blendMode.c())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the blend mode used for texture copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode)
func (texture *Texture) GetBlendMode() (blendMode BlendMode, err error) {
	ret := C.SDL_GetTextureBlendMode(texture.cptr(), blendMode.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set the scale mode used for texture scale operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode)
func (texture *Texture) SetScaleMode(scaleMode ScaleMode) error {
	ret := C.SDL_SetTextureScaleMode(texture.cptr(), scaleMode.c())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the scale mode used for texture scale operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode)
func (texture *Texture) GetScaleMode() (scaleMode ScaleMode, err error) {
	ret := C.SDL_GetTextureScaleMode(texture.cptr(), scaleMode.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Update the given texture rectangle with new pixel data.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateTexture)
func (texture *Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int32) error {
	ret := C.SDL_UpdateTexture(texture.cptr(), rect.cptr(), pixels, C.int(pitch))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Update a rectangle within a planar YV12 or IYUV texture with new pixel
// data.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture)
func (texture *Texture) UpdateYUV(rect *Rect, Yplane []uint8, Ypitch int32, Uplane []uint8, Upitch int32, Vplane []uint8, Vpitch int32) error {
	ret := C.SDL_UpdateYUVTexture(texture.cptr(), rect.cptr(),
		(*C.Uint8)(unsafe.SliceData(Yplane)), C.int(Ypitch),
		(*C.Uint8)(unsafe.SliceData(Uplane)), C.int(Upitch),
		(*C.Uint8)(unsafe.SliceData(Vplane)), C.int(Vpitch))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Update a rectangle within a planar NV12 or NV21 texture with new pixels.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture)
func (texture *Texture) UpdateNV(rect *Rect, Yplane []uint8, Ypitch int32, UVplane []uint8, UVpitch int32) error {
	ret := C.SDL_UpdateNVTexture(texture.cptr(), rect.cptr(),
		(*C.Uint8)(unsafe.SliceData(Yplane)), C.int(Ypitch),
		(*C.Uint8)(unsafe.SliceData(UVplane)), C.int(UVpitch))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// // Lock a portion of the texture for **write-only** pixel access.
// // (https://wiki.libsdl.org/SDL3/SDL_LockTexture)
// func LockTexture(SDL_Texture *texture,
// 										const SDL_Rect *rect,
// 										void **pixels, int *pitch) int {
// 	ret := C.SDL_LockTexture(pixels, pitch)
// }

// // Lock a portion of the texture for **write-only** pixel access, and expose
// // it as a SDL surface.
// // (https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface)
// func LockTextureToSurface(SDL_Texture *texture,
// 										const SDL_Rect *rect,
// 										SDL_Surface **surface) int {
// 	ret := C.SDL_LockTextureToSurface(surface)
// }

// // Unlock a texture, uploading the changes to video memory, if needed.
// // (https://wiki.libsdl.org/SDL3/SDL_UnlockTexture)
// func UnlockTexture(SDL_Texture *texture) void {
// 	ret := C.SDL_UnlockTexture(texture)
// }

// Set a texture as the current rendering target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget)
func (renderer *Renderer) SetTarget(texture *Texture) error {
	ret := C.SDL_SetRenderTarget(renderer.cptr(), texture.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the current render target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget)
func (renderer *Renderer) GetTarget() (texture *Texture, err error) {
	texture = (*Texture)(C.SDL_GetRenderTarget(renderer.cptr()))
	if texture == nil {
		err = GetError()
	}
	return
}

// Set a device independent resolution and presentation mode for rendering.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation)
func (renderer *Renderer) SetLogicalPresentation(w, h int32, mode RendererLogicalPresentation, scaleMode ScaleMode) error {
	ret := C.SDL_SetRenderLogicalPresentation(renderer.cptr(), C.int(w), C.int(h), mode.c(), scaleMode.c())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get device independent resolution and presentation mode for rendering.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation)
func (renderer *Renderer) GetLogicalPresentation() (w, h int32, mode RendererLogicalPresentation, scaleMode ScaleMode, err error) {
	ret := C.SDL_GetRenderLogicalPresentation(renderer.cptr(), (*C.int)(&w), (*C.int)(&h), mode.cptr(), scaleMode.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get the final presentation rectangle for rendering.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect)
func (renderer *Renderer) GetLogicalPresentationRect() (rect FRect, err error) {
	ret := C.SDL_GetRenderLogicalPresentationRect(renderer.cptr(), rect.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get a point in render coordinates when given a point in window coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow)
func (renderer *Renderer) CoordinatesFromWindow(windowX, windowY float32) (x, y float32, err error) {
	ret := C.SDL_RenderCoordinatesFromWindow(renderer.cptr(), C.float(windowX), C.float(windowY), (*C.float)(&x), (*C.float)(&y))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get a point in window coordinates when given a point in render coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow)
func (renderer *Renderer) CoordinatesToWindow(x, y float32) (windowX, windowY float32, err error) {
	ret := C.SDL_RenderCoordinatesToWindow(renderer.cptr(), C.float(x), C.float(y), (*C.float)(&windowX), (*C.float)(&windowY))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Convert the coordinates in an event to render coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates)
func (renderer *Renderer) ConvertEventCoordinates(event Event) error {
	ret := C.SDL_ConvertEventToRenderCoordinates(renderer.cptr(), event.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set the drawing area for rendering on the current target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport)
func (renderer *Renderer) SetViewport(rect *Rect) (err error) {
	ret := C.SDL_SetRenderViewport(renderer.cptr(), rect.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get the drawing area for the current target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport)
func (renderer *Renderer) GetViewport() (rect Rect, err error) {
	ret := C.SDL_GetRenderViewport(renderer.cptr(), rect.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Return whether an explicit rectangle was set as the viewport.
// (https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet)
func (renderer *Renderer) ViewportSet() bool {
	return C.SDL_RenderViewportSet(renderer.cptr()) == C.SDL_TRUE
}

// Set the clip rectangle for rendering on the specified target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect)
func (renderer *Renderer) SetClipRect(rect Rect) error {
	ret := C.SDL_SetRenderClipRect(renderer.cptr(), rect.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the clip rectangle for the current target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect)
func (renderer *Renderer) GetClipRect() (rect Rect, err error) {
	ret := C.SDL_GetRenderClipRect(renderer.cptr(), rect.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get whether clipping is enabled on the given renderer.
// (https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled)
func (renderer *Renderer) ClipEnabled() (bool, error) {
	ret := C.SDL_RenderClipEnabled(renderer.cptr())
	if ret != C.SDL_TRUE {
		return false, GetError()
	}
	return ret == C.SDL_TRUE, nil
}

// Set the drawing scale for rendering on the current target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderScale)
func (renderer *Renderer) SetScale(scaleX, scaleY float32) error {
	ret := C.SDL_SetRenderScale(renderer.cptr(), C.float(scaleX), C.float(scaleY))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the drawing scale for the current target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderScale)
func (renderer *Renderer) GetScale() (scaleX, scaleY float32, err error) {
	ret := C.SDL_GetRenderScale(renderer.cptr(), (*C.float)(&scaleX), (*C.float)(&scaleY))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set the color used for drawing operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor)
func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) error {
	ret := C.SDL_SetRenderDrawColor(renderer.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat)
func (renderer *Renderer) SetDrawColorFloat(r, g, b, a float32) error {
	ret := C.SDL_SetRenderDrawColorFloat(renderer.cptr(), C.float(r), C.float(g), C.float(b), C.float(a))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor)
func (renderer *Renderer) GetDrawColor() (r, g, b, a uint8, err error) {
	ret := C.SDL_GetRenderDrawColor(renderer.cptr(), (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Get the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat)
func (renderer *Renderer) GetDrawColorFloat() (r, g, b, a float32, err error) {
	ret := C.SDL_GetRenderDrawColorFloat(renderer.cptr(), (*C.float)(&r), (*C.float)(&g), (*C.float)(&b), (*C.float)(&a))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set the color scale used for render operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale)
func (renderer *Renderer) SetColorScale(scale float32) error {
	ret := C.SDL_SetRenderColorScale(renderer.cptr(), C.float(scale))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the color scale used for render operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale)
func (renderer *Renderer) GetColorScale() (scale float32, err error) {
	ret := C.SDL_GetRenderColorScale(renderer.cptr(), (*C.float)(&scale))
	if ret < 0 {
		err = GetError()
	}
	return
}

// Set the blend mode used for drawing operations (Fill and Line).
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode)
func (renderer *Renderer) SetDrawBlendMode(mode BlendMode) error {
	ret := C.SDL_SetRenderDrawBlendMode(renderer.cptr(), mode.c())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the blend mode used for drawing operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode)
func (renderer *Renderer) GetDrawBlendMode() (mode BlendMode, err error) {
	ret := C.SDL_GetRenderDrawBlendMode(renderer.cptr(), mode.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Clear the current rendering target with the drawing color.
// (https://wiki.libsdl.org/SDL3/SDL_RenderClear)
func (renderer *Renderer) Clear() error {
	ret := C.SDL_RenderClear(renderer.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Draw a point on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderPoint)
func (renderer *Renderer) Point(x, y float32) error {
	ret := C.SDL_RenderPoint(renderer.cptr(), C.float(x), C.float(y))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Draw multiple points on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderPoints)
func (renderer *Renderer) Points(points []Point) error {
	pointsPtr := unsafe.Pointer(unsafe.SliceData(points))
	ret := C.SDL_RenderPoints(renderer.cptr(), (*C.SDL_FPoint)(pointsPtr), C.int(len(points)))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Draw a line on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderLine)
func (renderer *Renderer) Line(x1, y1, x2, y2 float32) error {
	ret := C.SDL_RenderLine(renderer.cptr(), C.float(x1), C.float(y1), C.float(x2), C.float(y2))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Draw a series of connected lines on the current rendering target at
// subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderLines)
func (renderer *Renderer) Lines(points []Point) error {
	pointsPtr := unsafe.Pointer(unsafe.SliceData(points))
	ret := C.SDL_RenderLines(renderer.cptr(), (*C.SDL_FPoint)(pointsPtr), C.int(len(points)))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Draw a rectangle on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderRect)
func (renderer *Renderer) Rect(rect *FRect) error {
	ret := C.SDL_RenderRect(renderer.cptr(), rect.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Draw some number of rectangles on the current rendering target at subpixel
// precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderRects)
func (renderer *Renderer) Rects(rects []FRect) error {
	rectsPtr := unsafe.Pointer(unsafe.SliceData(rects))
	ret := C.SDL_RenderRects(renderer.cptr(), (*C.SDL_FRect)(rectsPtr), C.int(len(rects)))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Fill a rectangle on the current rendering target with the drawing color at
// subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderFillRect)
func (renderer *Renderer) FillRect(rect *FRect) error {
	ret := C.SDL_RenderFillRect(renderer.cptr(), rect.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Fill some number of rectangles on the current rendering target with the
// drawing color at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderFillRects)
func (renderer *Renderer) FillRects(rects []FRect) error {
	rectsPtr := unsafe.Pointer(unsafe.SliceData(rects))
	ret := C.SDL_RenderFillRects(renderer.cptr(), (*C.SDL_FRect)(rectsPtr), C.int(len(rects)))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Copy a portion of the texture to the current rendering target at subpixel
// precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTexture)
func (renderer *Renderer) Texture(texture *Texture, srcRect, dstRect *FRect) error {
	ret := C.SDL_RenderTexture(renderer.cptr(), texture.cptr(), srcRect.cptr(), dstRect.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// // Copy a portion of the source texture to the current rendering target, with
// // rotation and flipping, at subpixel precision.
// // (https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated)
func (renderer *Renderer) TextureRotated(texture *Texture, srcRect, dstRect *FRect, angle float64, center *FPoint, flip FlipMode) error {
	ret := C.SDL_RenderTextureRotated(renderer.cptr(), texture.cptr(), srcRect.cptr(), dstRect.cptr(), C.double(angle), center.cptr(), flip.c())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Tile a portion of the texture to the current rendering target at subpixel
// precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled)
func (renderer *Renderer) TextureTiled(texture *Texture, srcRect *FRect, scale float32, dstRect *FRect) error {
	ret := C.SDL_RenderTextureTiled(renderer.cptr(), texture.cptr(), srcRect.cptr(), C.float(scale), dstRect.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Perform a scaled copy using the 9-grid algorithm to the current rendering
// target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid)
func (renderer *Renderer) Texture9Grid(texture *Texture, srcRect FRect, cornerSize float32, scale float32, dstRect FRect) error {
	ret := C.SDL_RenderTexture9Grid(renderer.cptr(), texture.cptr(), srcRect.cptr(), C.float(cornerSize), C.float(scale), dstRect.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Render a list of triangles, optionally using a texture and indices into the
// vertex array Color and alpha modulation is done per vertex
// (https://wiki.libsdl.org/SDL3/SDL_RenderGeometry)
func (renderer *Renderer) Geometry(texture *Texture, vertices []Vertex, indices []int32) error {
	verticesPtr := unsafe.Pointer(unsafe.SliceData(vertices))
	indicesPtr := unsafe.Pointer(unsafe.SliceData(indices))
	ret := C.SDL_RenderGeometry(renderer.cptr(), texture.cptr(), (*C.SDL_Vertex)(verticesPtr), C.int(len(vertices)), (*C.int)(indicesPtr), C.int(len(indices)))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// TODO: my confidence in this function is low.
// Render a list of triangles, optionally using a texture and indices into the
// vertex arrays Color and alpha modulation is done per vertex
// (https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw)
func (renderer *Renderer) GeometryRaw(texture *Texture, xy []float32, xyStride int32, colors []FColor, colorStride int32, uv []float32, uvStride int32, indices []int32) error {
	xyPtr := unsafe.Pointer(unsafe.SliceData(xy))
	colorsPtr := unsafe.Pointer(unsafe.SliceData(colors))
	uvPtr := unsafe.Pointer(unsafe.SliceData(uv))
	count := len(colors)
	if len(xy)/2 != count || len(uv)/2 != count {
		return errorSDL("xy, uv and/or colors has incorrect number of elements")
	}
	indicesPtr := unsafe.Pointer(unsafe.SliceData(indices))
	idxSize := int32(4) // Only handles int32 indices
	ret := C.SDL_RenderGeometryRaw(renderer.cptr(), texture.cptr(),
		(*C.float)(xyPtr), C.int(xyStride),
		(*C.SDL_FColor)(colorsPtr), C.int(colorStride),
		(*C.float)(uvPtr), C.int(uvStride),
		C.int(count),
		indicesPtr, C.int(len(indices)), C.int(idxSize))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Read pixels from the current rendering target.
// (https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels)
func (renderer *Renderer) ReadPixels(rect Rect) (surface *Surface, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_RenderReadPixels(renderer.cptr(), rect.cptr())))
	if surface == nil {
		err = GetError()
	}
	return
}

// Update the screen with any rendering performed since the previous call.
// (https://wiki.libsdl.org/SDL3/SDL_RenderPresent)
func (renderer *Renderer) Present() error {
	ret := C.SDL_RenderPresent(renderer.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Destroy the specified texture.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyTexture)
func (texture *Texture) Destroy() {
	C.SDL_DestroyTexture(texture.cptr())
}

// Destroy the rendering context for a window and free all associated
// textures.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer)
func (renderer *Renderer) Destroy() {
	C.SDL_DestroyRenderer(renderer.cptr())
}

// Force the rendering context to flush any pending commands and state.
// (https://wiki.libsdl.org/SDL3/SDL_FlushRenderer)
func (renderer *Renderer) Flush() error {
	ret := C.SDL_FlushRenderer(renderer.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Get the CAMetalLayer associated with the given Metal renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer)
func (renderer *Renderer) GetMetalLayer() (metalLayer unsafe.Pointer) {
	metalLayer = C.SDL_GetRenderMetalLayer(renderer.cptr())
	return
}

// Get the Metal command encoder for the current frame.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder)
func (renderer *Renderer) GetMetalCommandEncoder() (commandEncoder unsafe.Pointer) {
	commandEncoder = C.SDL_GetRenderMetalCommandEncoder(renderer.cptr())
	return
}

// Add a set of synchronization semaphores for the current frame.
// (https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores)
func (renderer *Renderer) AddVulkanSemaphores(waitStageMask uint32, waitSemaphore, signalSemaphore int64) error {
	ret := C.SDL_AddVulkanRenderSemaphores(renderer.cptr(), C.Uint32(waitStageMask), C.Sint64(waitSemaphore), C.Sint64(signalSemaphore))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Toggle VSync of the given renderer.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync)
func (renderer *Renderer) SetVSync(vsync int32) error {
	ret := C.SDL_SetRenderVSync(renderer.cptr(), C.int(vsync))
	if ret < 0 {
		return GetError()
	}
	return nil
}

const (
	RENDERER_VSYNC_DISABLED = 0
	RENDERER_VSYNC_ADAPTIVE = (-1)
)

// Get VSync of the given renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync)
func (renderer *Renderer) GetVSync() (vsync int32, err error) {
	ret := int32(C.SDL_GetRenderVSync(renderer.cptr(), (*C.int)(&vsync)))
	if ret < 0 {
		err = GetError()
	}
	return
}
