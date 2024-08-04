package sdl

//#include <SDL3/SDL_blendmode.h>
import "C"

// A set of blend modes used in drawing operations.
// (https://wiki.libsdl.org/SDL3/SDL_BlendMode)
type BlendMode C.SDL_BlendMode

func (mode BlendMode) c() C.SDL_BlendMode {
	return C.SDL_BlendMode(mode)
}

func (mode *BlendMode) cptr() *C.SDL_BlendMode {
	return (*C.SDL_BlendMode)(mode)
}

const (
	BLENDMODE_NONE                = BlendMode(C.SDL_BLENDMODE_NONE)                /**< no blending: dstRGBA = srcRGBA */
	BLENDMODE_BLEND               = BlendMode(C.SDL_BLENDMODE_BLEND)               /**< alpha blending: dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)), dstA = srcA + (dstA * (1-srcA)) */
	BLENDMODE_BLEND_PREMULTIPLIED = BlendMode(C.SDL_BLENDMODE_BLEND_PREMULTIPLIED) /**< pre-multiplied alpha blending: dstRGBA = srcRGBA + (dstRGBA * (1-srcA)) */
	BLENDMODE_ADD                 = BlendMode(C.SDL_BLENDMODE_ADD)                 /**< additive blending: dstRGB = (srcRGB * srcA) + dstRGB, dstA = dstA */
	BLENDMODE_ADD_PREMULTIPLIED   = BlendMode(C.SDL_BLENDMODE_ADD_PREMULTIPLIED)   /**< pre-multiplied additive blending: dstRGB = srcRGB + dstRGB, dstA = dstA */
	BLENDMODE_MOD                 = BlendMode(C.SDL_BLENDMODE_MOD)                 /**< color modulate: dstRGB = srcRGB * dstRGB, dstA = dstA */
	BLENDMODE_MUL                 = BlendMode(C.SDL_BLENDMODE_MUL)                 /**< color multiply: dstRGB = (srcRGB * dstRGB) + (dstRGB * (1-srcA)), dstA = dstA */
	BLENDMODE_INVALID             = BlendMode(C.SDL_BLENDMODE_INVALID)
)

// The blend operation used when combining source and destination pixel
// components.
// (https://wiki.libsdl.org/SDL3/SDL_BlendOperation)
type BlendOperation C.SDL_BlendOperation

const (
	BLENDOPERATION_ADD          = BlendOperation(C.SDL_BLENDOPERATION_ADD)          /**< dst + src: supported by all renderers */
	BLENDOPERATION_SUBTRACT     = BlendOperation(C.SDL_BLENDOPERATION_SUBTRACT)     /**< src - dst : supported by D3D, OpenGL, OpenGLES, and Vulkan */
	BLENDOPERATION_REV_SUBTRACT = BlendOperation(C.SDL_BLENDOPERATION_REV_SUBTRACT) /**< dst - src : supported by D3D, OpenGL, OpenGLES, and Vulkan */
	BLENDOPERATION_MINIMUM      = BlendOperation(C.SDL_BLENDOPERATION_MINIMUM)      /**< min(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan */
	BLENDOPERATION_MAXIMUM      = BlendOperation(C.SDL_BLENDOPERATION_MAXIMUM)      /**< max(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan */
)

// The normalized factor used to multiply pixel components.
// (https://wiki.libsdl.org/SDL3/SDL_BlendFactor)
type BlendFactor C.SDL_BlendFactor

const (
	BLENDFACTOR_ZERO                = BlendFactor(C.SDL_BLENDFACTOR_ZERO)                /**< 0, 0, 0, 0 */
	BLENDFACTOR_ONE                 = BlendFactor(C.SDL_BLENDFACTOR_ONE)                 /**< 1, 1, 1, 1 */
	BLENDFACTOR_SRC_COLOR           = BlendFactor(C.SDL_BLENDFACTOR_SRC_COLOR)           /**< srcR, srcG, srcB, srcA */
	BLENDFACTOR_ONE_MINUS_SRC_COLOR = BlendFactor(C.SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR) /**< 1-srcR, 1-srcG, 1-srcB, 1-srcA */
	BLENDFACTOR_SRC_ALPHA           = BlendFactor(C.SDL_BLENDFACTOR_SRC_ALPHA)           /**< srcA, srcA, srcA, srcA */
	BLENDFACTOR_ONE_MINUS_SRC_ALPHA = BlendFactor(C.SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA) /**< 1-srcA, 1-srcA, 1-srcA, 1-srcA */
	BLENDFACTOR_DST_COLOR           = BlendFactor(C.SDL_BLENDFACTOR_DST_COLOR)           /**< dstR, dstG, dstB, dstA */
	BLENDFACTOR_ONE_MINUS_DST_COLOR = BlendFactor(C.SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR) /**< 1-dstR, 1-dstG, 1-dstB, 1-dstA */
	BLENDFACTOR_DST_ALPHA           = BlendFactor(C.SDL_BLENDFACTOR_DST_ALPHA)           /**< dstA, dstA, dstA, dstA */
	BLENDFACTOR_ONE_MINUS_DST_ALPHA = BlendFactor(C.SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA) /**< 1-dstA, 1-dstA, 1-dstA, 1-dstA */
)

// Compose a custom blend mode for renderers.
// (https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode)
func ComposeCustomBlendMode(srcColorFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
	return BlendMode(C.SDL_ComposeCustomBlendMode(
		C.SDL_BlendFactor(srcColorFactor),
		C.SDL_BlendFactor(dstColorFactor),
		C.SDL_BlendOperation(colorOperation),
		C.SDL_BlendFactor(srcAlphaFactor),
		C.SDL_BlendFactor(dstAlphaFactor),
		C.SDL_BlendOperation(alphaOperation)))
}
