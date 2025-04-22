package sdl

/*
#include <SDL3/SDL_gpu.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// # CategoryGPU
// (https://wiki.libsdl.org/SDL3/CategoryGPU)

// Type Declarations

// An opaque handle representing the SDL_GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_GPUDevice)
type GPUDevice C.SDL_GPUDevice

// An opaque handle representing a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBuffer)
type GPUBuffer C.SDL_GPUBuffer

// An opaque handle representing a transfer buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTransferBuffer)
type GPUTransferBuffer C.SDL_GPUTransferBuffer

// An opaque handle representing a texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTexture)
type GPUTexture C.SDL_GPUTexture

// An opaque handle representing a sampler.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSampler)
type GPUSampler C.SDL_GPUSampler

// An opaque handle representing a compiled shader object.
// (https://wiki.libsdl.org/SDL3/SDL_GPUShader)
type GPUShader C.SDL_GPUShader

// An opaque handle representing a compute pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUComputePipeline)
type GPUComputePipeline C.SDL_GPUComputePipeline

// An opaque handle representing a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipeline)
type GPUGraphicsPipeline C.SDL_GPUGraphicsPipeline

// An opaque handle representing a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUCommandBuffer)
type GPUCommandBuffer C.SDL_GPUCommandBuffer

// An opaque handle representing a render pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPURenderPass)
type GPURenderPass C.SDL_GPURenderPass

// An opaque handle representing a compute pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUComputePass)
type GPUComputePass C.SDL_GPUComputePass

// An opaque handle representing a copy pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUCopyPass)
type GPUCopyPass C.SDL_GPUCopyPass

// An opaque handle representing a fence.
// (https://wiki.libsdl.org/SDL3/SDL_GPUFence)
type GPUFence C.SDL_GPUFence

// Specifies the primitive topology of a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUPrimitiveType)
type GPUPrimitiveType C.SDL_GPUPrimitiveType

const (
	GPU_PRIMITIVETYPE_TRIANGLELIST  = GPUPrimitiveType(C.SDL_GPU_PRIMITIVETYPE_TRIANGLELIST)  /**< A series of separate triangles. */
	GPU_PRIMITIVETYPE_TRIANGLESTRIP = GPUPrimitiveType(C.SDL_GPU_PRIMITIVETYPE_TRIANGLESTRIP) /**< A series of connected triangles. */
	GPU_PRIMITIVETYPE_LINELIST      = GPUPrimitiveType(C.SDL_GPU_PRIMITIVETYPE_LINELIST)      /**< A series of separate lines. */
	GPU_PRIMITIVETYPE_LINESTRIP     = GPUPrimitiveType(C.SDL_GPU_PRIMITIVETYPE_LINESTRIP)     /**< A series of connected lines. */
	GPU_PRIMITIVETYPE_POINTLIST     = GPUPrimitiveType(C.SDL_GPU_PRIMITIVETYPE_POINTLIST)     /**< A series of separate points. */
)

// Specifies how the contents of a texture attached to a render pass are
// treated at the beginning of the render pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPULoadOp)
type GPULoadOp C.SDL_GPULoadOp

const (
	GPU_LOADOP_LOAD      = GPULoadOp(C.SDL_GPU_LOADOP_LOAD)      /**< The previous contents of the texture will be preserved. */
	GPU_LOADOP_CLEAR     = GPULoadOp(C.SDL_GPU_LOADOP_CLEAR)     /**< The contents of the texture will be cleared to a color. */
	GPU_LOADOP_DONT_CARE = GPULoadOp(C.SDL_GPU_LOADOP_DONT_CARE) /**< The previous contents of the texture need not be preserved. The contents will be undefined. */
)

// Specifies how the contents of a texture attached to a render pass are
// treated at the end of the render pass.
// (http://wiki.libsdl.org/SDL3/SDL_GPUStoreOp)
type GPUStoreOp C.SDL_GPUStoreOp

const (
	GPU_STOREOP_STORE             = GPUStoreOp(C.SDL_GPU_STOREOP_STORE)             /**< The contents generated during the render pass will be written to memory. */
	GPU_STOREOP_DONT_CARE         = GPUStoreOp(C.SDL_GPU_STOREOP_DONT_CARE)         /**< The contents generated during the render pass are not needed and may be discarded. The contents will be undefined. */
	GPU_STOREOP_RESOLVE           = GPUStoreOp(C.SDL_GPU_STOREOP_RESOLVE)           /**< The multisample contents generated during the render pass will be resolved to a non-multisample texture. The contents in the multisample texture may then be discarded and will be undefined. */
	GPU_STOREOP_RESOLVE_AND_STORE = GPUStoreOp(C.SDL_GPU_STOREOP_RESOLVE_AND_STORE) /**< The multisample contents generated during the render pass will be resolved to a non-multisample texture. The contents in the multisample texture will be written to memory. */
)

// Specifies the size of elements in an index buffer.
// (http://wiki.libsdl.org/SDL3/SDL_GPUIndexElementSize)
type GPUIndexElementSize C.SDL_GPUIndexElementSize

const (
	GPU_INDEXELEMENTSIZE_16BIT = GPUIndexElementSize(C.SDL_GPU_INDEXELEMENTSIZE_16BIT) /**< The index elements are 16-bit. */
	GPU_INDEXELEMENTSIZE_32BIT = GPUIndexElementSize(C.SDL_GPU_INDEXELEMENTSIZE_32BIT) /**< The index elements are 32-bit. */
)

// Specifies the pixel format of a texture.
// (http://wiki.libsdl.org/SDL3/SDL_GPUTextureFormat)
type GPUTextureFormat C.SDL_GPUTextureFormat

const (
	GPU_TEXTUREFORMAT_INVALID = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_INVALID)

	/* Unsigned Normalized Float Color Formats */
	GPU_TEXTUREFORMAT_A8_UNORM           = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_A8_UNORM)
	GPU_TEXTUREFORMAT_R8_UNORM           = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8_UNORM)
	GPU_TEXTUREFORMAT_R8G8_UNORM         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8_UNORM)
	GPU_TEXTUREFORMAT_R8G8B8A8_UNORM     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_UNORM)
	GPU_TEXTUREFORMAT_R16_UNORM          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16_UNORM)
	GPU_TEXTUREFORMAT_R16G16_UNORM       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16_UNORM)
	GPU_TEXTUREFORMAT_R16G16B16A16_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_UNORM)
	GPU_TEXTUREFORMAT_R10G10B10A2_UNORM  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R10G10B10A2_UNORM)
	GPU_TEXTUREFORMAT_B5G6R5_UNORM       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_B5G6R5_UNORM)
	GPU_TEXTUREFORMAT_B5G5R5A1_UNORM     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_B5G5R5A1_UNORM)
	GPU_TEXTUREFORMAT_B4G4R4A4_UNORM     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_B4G4R4A4_UNORM)
	GPU_TEXTUREFORMAT_B8G8R8A8_UNORM     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_B8G8R8A8_UNORM)
	/* Compressed Unsigned Normalized Float Color Formats */
	GPU_TEXTUREFORMAT_BC1_RGBA_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC1_RGBA_UNORM)
	GPU_TEXTUREFORMAT_BC2_RGBA_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC2_RGBA_UNORM)
	GPU_TEXTUREFORMAT_BC3_RGBA_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC3_RGBA_UNORM)
	GPU_TEXTUREFORMAT_BC4_R_UNORM    = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC4_R_UNORM)
	GPU_TEXTUREFORMAT_BC5_RG_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC5_RG_UNORM)
	GPU_TEXTUREFORMAT_BC7_RGBA_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC7_RGBA_UNORM)
	/* Compressed Signed Float Color Formats */
	GPU_TEXTUREFORMAT_BC6H_RGB_FLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC6H_RGB_FLOAT)
	/* Compressed Unsigned Float Color Formats */
	GPU_TEXTUREFORMAT_BC6H_RGB_UFLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC6H_RGB_UFLOAT)
	/* Signed Normalized Float Color Formats  */
	GPU_TEXTUREFORMAT_R8_SNORM           = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8_SNORM)
	GPU_TEXTUREFORMAT_R8G8_SNORM         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8_SNORM)
	GPU_TEXTUREFORMAT_R8G8B8A8_SNORM     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_SNORM)
	GPU_TEXTUREFORMAT_R16_SNORM          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16_SNORM)
	GPU_TEXTUREFORMAT_R16G16_SNORM       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16_SNORM)
	GPU_TEXTUREFORMAT_R16G16B16A16_SNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_SNORM)
	/* Signed Float Color Formats */
	GPU_TEXTUREFORMAT_R16_FLOAT          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16_FLOAT)
	GPU_TEXTUREFORMAT_R16G16_FLOAT       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16_FLOAT)
	GPU_TEXTUREFORMAT_R16G16B16A16_FLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_FLOAT)
	GPU_TEXTUREFORMAT_R32_FLOAT          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32_FLOAT)
	GPU_TEXTUREFORMAT_R32G32_FLOAT       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32G32_FLOAT)
	GPU_TEXTUREFORMAT_R32G32B32A32_FLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32G32B32A32_FLOAT)
	/* Unsigned Float Color Formats */
	GPU_TEXTUREFORMAT_R11G11B10_UFLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R11G11B10_UFLOAT)
	/* Unsigned Integer Color Formats */
	GPU_TEXTUREFORMAT_R8_UINT           = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8_UINT)
	GPU_TEXTUREFORMAT_R8G8_UINT         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8_UINT)
	GPU_TEXTUREFORMAT_R8G8B8A8_UINT     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_UINT)
	GPU_TEXTUREFORMAT_R16_UINT          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16_UINT)
	GPU_TEXTUREFORMAT_R16G16_UINT       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16_UINT)
	GPU_TEXTUREFORMAT_R16G16B16A16_UINT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_UINT)
	GPU_TEXTUREFORMAT_R32_UINT          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32_UINT)
	GPU_TEXTUREFORMAT_R32G32_UINT       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32G32_UINT)
	GPU_TEXTUREFORMAT_R32G32B32A32_UINT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32G32B32A32_UINT)
	/* Signed Integer Color Formats */
	GPU_TEXTUREFORMAT_R8_INT           = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8_INT)
	GPU_TEXTUREFORMAT_R8G8_INT         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8_INT)
	GPU_TEXTUREFORMAT_R8G8B8A8_INT     = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_INT)
	GPU_TEXTUREFORMAT_R16_INT          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16_INT)
	GPU_TEXTUREFORMAT_R16G16_INT       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16_INT)
	GPU_TEXTUREFORMAT_R16G16B16A16_INT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_INT)
	GPU_TEXTUREFORMAT_R32_INT          = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32_INT)
	GPU_TEXTUREFORMAT_R32G32_INT       = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32G32_INT)
	GPU_TEXTUREFORMAT_R32G32B32A32_INT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R32G32B32A32_INT)
	/* SRGB Unsigned Normalized Color Formats */
	GPU_TEXTUREFORMAT_R8G8B8A8_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_UNORM_SRGB)
	GPU_TEXTUREFORMAT_B8G8R8A8_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_B8G8R8A8_UNORM_SRGB)
	/* Compressed SRGB Unsigned Normalized Color Formats */
	GPU_TEXTUREFORMAT_BC1_RGBA_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC1_RGBA_UNORM_SRGB)
	GPU_TEXTUREFORMAT_BC2_RGBA_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC2_RGBA_UNORM_SRGB)
	GPU_TEXTUREFORMAT_BC3_RGBA_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC3_RGBA_UNORM_SRGB)
	GPU_TEXTUREFORMAT_BC7_RGBA_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_BC7_RGBA_UNORM_SRGB)
	/* Depth Formats */
	GPU_TEXTUREFORMAT_D16_UNORM         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_D16_UNORM)
	GPU_TEXTUREFORMAT_D24_UNORM         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_D24_UNORM)
	GPU_TEXTUREFORMAT_D32_FLOAT         = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_D32_FLOAT)
	GPU_TEXTUREFORMAT_D24_UNORM_S8_UINT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_D24_UNORM_S8_UINT)
	GPU_TEXTUREFORMAT_D32_FLOAT_S8_UINT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_D32_FLOAT_S8_UINT)
	/* Compressed ASTC Normalized Float Color Formats*/
	GPU_TEXTUREFORMAT_ASTC_4x4_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_4x4_UNORM)
	GPU_TEXTUREFORMAT_ASTC_5x4_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_5x4_UNORM)
	GPU_TEXTUREFORMAT_ASTC_5x5_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_5x5_UNORM)
	GPU_TEXTUREFORMAT_ASTC_6x5_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_6x5_UNORM)
	GPU_TEXTUREFORMAT_ASTC_6x6_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_6x6_UNORM)
	GPU_TEXTUREFORMAT_ASTC_8x5_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x5_UNORM)
	GPU_TEXTUREFORMAT_ASTC_8x6_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x6_UNORM)
	GPU_TEXTUREFORMAT_ASTC_8x8_UNORM   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x8_UNORM)
	GPU_TEXTUREFORMAT_ASTC_10x5_UNORM  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x5_UNORM)
	GPU_TEXTUREFORMAT_ASTC_10x6_UNORM  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x6_UNORM)
	GPU_TEXTUREFORMAT_ASTC_10x8_UNORM  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x8_UNORM)
	GPU_TEXTUREFORMAT_ASTC_10x10_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x10_UNORM)
	GPU_TEXTUREFORMAT_ASTC_12x10_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_12x10_UNORM)
	GPU_TEXTUREFORMAT_ASTC_12x12_UNORM = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_12x12_UNORM)
	/* Compressed SRGB ASTC Normalized Float Color Formats*/
	GPU_TEXTUREFORMAT_ASTC_4x4_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_4x4_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_5x4_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_5x4_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_5x5_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_5x5_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_6x5_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_6x5_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_6x6_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_6x6_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_8x5_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x5_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_8x6_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x6_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_8x8_UNORM_SRGB   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x8_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_10x5_UNORM_SRGB  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x5_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_10x6_UNORM_SRGB  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x6_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_10x8_UNORM_SRGB  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x8_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_10x10_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x10_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_12x10_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_12x10_UNORM_SRGB)
	GPU_TEXTUREFORMAT_ASTC_12x12_UNORM_SRGB = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_12x12_UNORM_SRGB)
	/* Compressed ASTC Signed Float Color Formats*/
	GPU_TEXTUREFORMAT_ASTC_4x4_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_4x4_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_5x4_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_5x4_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_5x5_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_5x5_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_6x5_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_6x5_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_6x6_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_6x6_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_8x5_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x5_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_8x6_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x6_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_8x8_FLOAT   = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_8x8_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_10x5_FLOAT  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x5_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_10x6_FLOAT  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x6_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_10x8_FLOAT  = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x8_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_10x10_FLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_10x10_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_12x10_FLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_12x10_FLOAT)
	GPU_TEXTUREFORMAT_ASTC_12x12_FLOAT = GPUTextureFormat(C.SDL_GPU_TEXTUREFORMAT_ASTC_12x12_FLOAT)
)

// Specifies how a texture is intended to be used by the client.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureUsageFlags)
type GPUTextureUsageFlags C.SDL_GPUTextureUsageFlags

const (
	GPU_TEXTUREUSAGE_SAMPLER                                 = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_SAMPLER)                                 /**< Texture supports sampling. */
	GPU_TEXTUREUSAGE_COLOR_TARGET                            = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_COLOR_TARGET)                            /**< Texture is a color render target. */
	GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET                    = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET)                    /**< Texture is a depth stencil target. */
	GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ                   = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ)                   /**< Texture supports storage reads in graphics stages. */
	GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ                    = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ)                    /**< Texture supports storage reads in the compute stage. */
	GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE                   = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE)                   /**< Texture supports storage writes in the compute stage. */
	GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE = GPUTextureUsageFlags(C.SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE) /**< Texture supports reads and writes in the same compute shader. This is NOT equivalent to READ | WRITE. */
)

// Specifies the type of a texture.
// (http://wiki.libsdl.org/SDL3/SDL_GPUTextureType)
type GPUTextureType C.SDL_GPUTextureType

const (
	GPU_TEXTURETYPE_2D         = GPUTextureType(C.SDL_GPU_TEXTURETYPE_2D)         /**< The texture is a 2-dimensional image. */
	GPU_TEXTURETYPE_2D_ARRAY   = GPUTextureType(C.SDL_GPU_TEXTURETYPE_2D_ARRAY)   /**< The texture is a 2-dimensional array image. */
	GPU_TEXTURETYPE_3D         = GPUTextureType(C.SDL_GPU_TEXTURETYPE_3D)         /**< The texture is a 3-dimensional image. */
	GPU_TEXTURETYPE_CUBE       = GPUTextureType(C.SDL_GPU_TEXTURETYPE_CUBE)       /**< The texture is a cube image. */
	GPU_TEXTURETYPE_CUBE_ARRAY = GPUTextureType(C.SDL_GPU_TEXTURETYPE_CUBE_ARRAY) /**< The texture is a cube array image. */
)

// Specifies the sample count of a texture.
// (http://wiki.libsdl.org/SDL3/SDL_GPUSampleCount)
type GPUSampleCount C.SDL_GPUSampleCount

const (
	GPU_SAMPLECOUNT_1 = GPUSampleCount(C.SDL_GPU_SAMPLECOUNT_1) /**< No multisampling. */
	GPU_SAMPLECOUNT_2 = GPUSampleCount(C.SDL_GPU_SAMPLECOUNT_2) /**< MSAA 2x */
	GPU_SAMPLECOUNT_4 = GPUSampleCount(C.SDL_GPU_SAMPLECOUNT_4) /**< MSAA 4x */
	GPU_SAMPLECOUNT_8 = GPUSampleCount(C.SDL_GPU_SAMPLECOUNT_8) /**< MSAA 8x */
)

// Specifies the face of a cube map.
// (http://wiki.libsdl.org/SDL3/SDL_GPUTextureType)
type GPUCubeMapFace C.SDL_GPUCubeMapFace

const (
	GPU_CUBEMAPFACE_POSITIVEX = GPUCubeMapFace(C.SDL_GPU_CUBEMAPFACE_POSITIVEX)
	GPU_CUBEMAPFACE_NEGATIVEX = GPUCubeMapFace(C.SDL_GPU_CUBEMAPFACE_NEGATIVEX)
	GPU_CUBEMAPFACE_POSITIVEY = GPUCubeMapFace(C.SDL_GPU_CUBEMAPFACE_POSITIVEY)
	GPU_CUBEMAPFACE_NEGATIVEY = GPUCubeMapFace(C.SDL_GPU_CUBEMAPFACE_NEGATIVEY)
	GPU_CUBEMAPFACE_POSITIVEZ = GPUCubeMapFace(C.SDL_GPU_CUBEMAPFACE_POSITIVEZ)
	GPU_CUBEMAPFACE_NEGATIVEZ = GPUCubeMapFace(C.SDL_GPU_CUBEMAPFACE_NEGATIVEZ)
)

// Specifies how a buffer is intended to be used by the client.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBufferUsageFlags)
type GPUBufferUsageFlags C.SDL_GPUBufferUsageFlags

const (
	GPU_BUFFERUSAGE_VERTEX                = GPUBufferUsageFlags(C.SDL_GPU_BUFFERUSAGE_VERTEX)                /**< Buffer is a vertex buffer. */
	GPU_BUFFERUSAGE_INDEX                 = GPUBufferUsageFlags(C.SDL_GPU_BUFFERUSAGE_INDEX)                 /**< Buffer is an index buffer. */
	GPU_BUFFERUSAGE_INDIRECT              = GPUBufferUsageFlags(C.SDL_GPU_BUFFERUSAGE_INDIRECT)              /**< Buffer is an indirect buffer. */
	GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ = GPUBufferUsageFlags(C.SDL_GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ) /**< Buffer supports storage reads in graphics stages. */
	GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ  = GPUBufferUsageFlags(C.SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ)  /**< Buffer supports storage reads in the compute stage. */
	GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE = GPUBufferUsageFlags(C.SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE) /**< Buffer supports storage writes in the compute stage. */
)

// Specifies how a transfer buffer is intended to be used by the client.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferUsage)
type GPUTransferBufferUsage C.SDL_GPUTransferBufferUsage

const (
	GPU_TRANSFERBUFFERUSAGE_UPLOAD   = GPUTransferBufferUsage(C.SDL_GPU_TRANSFERBUFFERUSAGE_UPLOAD)
	GPU_TRANSFERBUFFERUSAGE_DOWNLOAD = GPUTransferBufferUsage(C.SDL_GPU_TRANSFERBUFFERUSAGE_DOWNLOAD)
)

// Specifies which stage a shader program corresponds to.
// (https://wiki.libsdl.org/SDL3/SDL_GPUShaderStage)
type GPUShaderStage C.SDL_GPUShaderStage

const (
	GPU_SHADERSTAGE_VERTEX   = GPUShaderStage(C.SDL_GPU_SHADERSTAGE_VERTEX)
	GPU_SHADERSTAGE_FRAGMENT = GPUShaderStage(C.SDL_GPU_SHADERSTAGE_FRAGMENT)
)

// Specifies the format of shader code.
// (https://wiki.libsdl.org/SDL3/SDL_GPUShaderFormat)
type GPUShaderFormat C.SDL_GPUShaderFormat

const (
	GPU_SHADERFORMAT_INVALID  = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_INVALID)
	GPU_SHADERFORMAT_PRIVATE  = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_PRIVATE)  /**< Shaders for NDA'd platforms. */
	GPU_SHADERFORMAT_SPIRV    = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_SPIRV)    /**< SPIR-V shaders for Vulkan. */
	GPU_SHADERFORMAT_DXBC     = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_DXBC)     /**< DXBC SM5_0 shaders for D3D11. */
	GPU_SHADERFORMAT_DXIL     = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_DXIL)     /**< DXIL shaders for D3D12. */
	GPU_SHADERFORMAT_MSL      = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_MSL)      /**< MSL shaders for Metal. */
	GPU_SHADERFORMAT_METALLIB = GPUShaderFormat(C.SDL_GPU_SHADERFORMAT_METALLIB) /**< Precompiled metallib shaders for Metal. */
)

// Specifies the format of a vertex attribute.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexElementFormat)
type GPUVertexElementFormat C.SDL_GPUVertexElementFormat

const (
	GPU_VERTEXELEMENTFORMAT_INVALID = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_INVALID)

	/* 32-bit Signed Integers */
	GPU_VERTEXELEMENTFORMAT_INT  = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_INT)
	GPU_VERTEXELEMENTFORMAT_INT2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_INT2)
	GPU_VERTEXELEMENTFORMAT_INT3 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_INT3)
	GPU_VERTEXELEMENTFORMAT_INT4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_INT4)

	/* 32-bit Unsigned Integers */
	GPU_VERTEXELEMENTFORMAT_UINT  = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UINT)
	GPU_VERTEXELEMENTFORMAT_UINT2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UINT2)
	GPU_VERTEXELEMENTFORMAT_UINT3 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UINT3)
	GPU_VERTEXELEMENTFORMAT_UINT4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UINT4)

	/* 32-bit Floats */
	GPU_VERTEXELEMENTFORMAT_FLOAT  = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT)
	GPU_VERTEXELEMENTFORMAT_FLOAT2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT2)
	GPU_VERTEXELEMENTFORMAT_FLOAT3 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT3)
	GPU_VERTEXELEMENTFORMAT_FLOAT4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT4)

	/* 8-bit Signed Integers */
	GPU_VERTEXELEMENTFORMAT_BYTE2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE2)
	GPU_VERTEXELEMENTFORMAT_BYTE4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE4)

	/* 8-bit Unsigned Integers */
	GPU_VERTEXELEMENTFORMAT_UBYTE2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE2)
	GPU_VERTEXELEMENTFORMAT_UBYTE4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE4)

	/* 8-bit Signed Normalized */
	GPU_VERTEXELEMENTFORMAT_BYTE2_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE2_NORM)
	GPU_VERTEXELEMENTFORMAT_BYTE4_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE4_NORM)

	/* 8-bit Unsigned Normalized */
	GPU_VERTEXELEMENTFORMAT_UBYTE2_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE2_NORM)
	GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM)

	/* 16-bit Signed Integers */
	GPU_VERTEXELEMENTFORMAT_SHORT2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT2)
	GPU_VERTEXELEMENTFORMAT_SHORT4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT4)

	/* 16-bit Unsigned Integers */
	GPU_VERTEXELEMENTFORMAT_USHORT2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT2)
	GPU_VERTEXELEMENTFORMAT_USHORT4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT4)

	/* 16-bit Signed Normalized */
	GPU_VERTEXELEMENTFORMAT_SHORT2_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT2_NORM)
	GPU_VERTEXELEMENTFORMAT_SHORT4_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT4_NORM)

	/* 16-bit Unsigned Normalized */
	GPU_VERTEXELEMENTFORMAT_USHORT2_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT2_NORM)
	GPU_VERTEXELEMENTFORMAT_USHORT4_NORM = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT4_NORM)

	/* 16-bit Floats */
	GPU_VERTEXELEMENTFORMAT_HALF2 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_HALF2)
	GPU_VERTEXELEMENTFORMAT_HALF4 = GPUVertexElementFormat(C.SDL_GPU_VERTEXELEMENTFORMAT_HALF4)
)

// Specifies the rate at which vertex attributes are pulled from buffers.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputRate)
type GPUVertexInputRate C.SDL_GPUVertexInputRate

const (
	GPU_VERTEXINPUTRATE_VERTEX   = GPUVertexInputRate(C.SDL_GPU_VERTEXINPUTRATE_VERTEX)   /**< Attribute addressing is a function of the vertex index. */
	GPU_VERTEXINPUTRATE_INSTANCE = GPUVertexInputRate(C.SDL_GPU_VERTEXINPUTRATE_INSTANCE) /**< Attribute addressing is a function of the instance index. */
)

// Specifies the fill mode of the graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUFillMode)
type GPUFillMode C.SDL_GPUFillMode

const (
	GPU_FILLMODE_FILL = GPUFillMode(C.SDL_GPU_FILLMODE_FILL) /**< Polygons will be rendered via rasterization. */
	GPU_FILLMODE_LINE = GPUFillMode(C.SDL_GPU_FILLMODE_LINE) /**< Polygon edges will be drawn as line segments. */
)

// Specifies the facing direction in which triangle faces will be culled.
// (https://wiki.libsdl.org/SDL3/SDL_GPUCullMode)
type GPUCullMode C.SDL_GPUCullMode

const (
	GPU_CULLMODE_NONE  = GPUCullMode(C.SDL_GPU_CULLMODE_NONE)  /**< No triangles are culled. */
	GPU_CULLMODE_FRONT = GPUCullMode(C.SDL_GPU_CULLMODE_FRONT) /**< Front-facing triangles are culled. */
	GPU_CULLMODE_BACK  = GPUCullMode(C.SDL_GPU_CULLMODE_BACK)  /**< Back-facing triangles are culled. */
)

// Specifies the vertex winding that will cause a triangle to be determined to
// be front-facing.
// (https://wiki.libsdl.org/SDL3/SDL_GPUFrontFace)
type GPUFrontFace C.SDL_GPUFrontFace

const (
	GPU_FRONTFACE_COUNTER_CLOCKWISE = GPUFrontFace(C.SDL_GPU_FRONTFACE_COUNTER_CLOCKWISE) /**< A triangle with counter-clockwise vertex winding will be considered front-facing. */
	GPU_FRONTFACE_CLOCKWISE         = GPUFrontFace(C.SDL_GPU_FRONTFACE_CLOCKWISE)         /**< A triangle with clockwise vertex winding will be considered front-facing. */
)

// Specifies a comparison operator for depth, stencil and sampler operations.
// (https://wiki.libsdl.org/SDL3/SDL_GPUCompareOp)
type GPUCompareOp C.SDL_GPUCompareOp

const (
	GPU_COMPAREOP_INVALID          = GPUCompareOp(C.SDL_GPU_COMPAREOP_INVALID)
	GPU_COMPAREOP_NEVER            = GPUCompareOp(C.SDL_GPU_COMPAREOP_NEVER)            /**< The comparison always evaluates false. */
	GPU_COMPAREOP_LESS             = GPUCompareOp(C.SDL_GPU_COMPAREOP_LESS)             /**< The comparison evaluates reference < test. */
	GPU_COMPAREOP_EQUAL            = GPUCompareOp(C.SDL_GPU_COMPAREOP_EQUAL)            /**< The comparison evaluates reference == test. */
	GPU_COMPAREOP_LESS_OR_EQUAL    = GPUCompareOp(C.SDL_GPU_COMPAREOP_LESS_OR_EQUAL)    /**< The comparison evaluates reference <= test. */
	GPU_COMPAREOP_GREATER          = GPUCompareOp(C.SDL_GPU_COMPAREOP_GREATER)          /**< The comparison evaluates reference > test. */
	GPU_COMPAREOP_NOT_EQUAL        = GPUCompareOp(C.SDL_GPU_COMPAREOP_NOT_EQUAL)        /**< The comparison evaluates reference != test. */
	GPU_COMPAREOP_GREATER_OR_EQUAL = GPUCompareOp(C.SDL_GPU_COMPAREOP_GREATER_OR_EQUAL) /**< The comparison evalutes reference >= test. */
	GPU_COMPAREOP_ALWAYS           = GPUCompareOp(C.SDL_GPU_COMPAREOP_ALWAYS)           /**< The comparison always evaluates true. */
)

// Specifies what happens to a stored stencil value if stencil tests fail or
// pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUStencilOp)
type GPUStencilOp C.SDL_GPUStencilOp

const (
	GPU_STENCILOP_INVALID             = GPUStencilOp(C.SDL_GPU_STENCILOP_INVALID)
	GPU_STENCILOP_KEEP                = GPUStencilOp(C.SDL_GPU_STENCILOP_KEEP)                /**< Keeps the current value. */
	GPU_STENCILOP_ZERO                = GPUStencilOp(C.SDL_GPU_STENCILOP_ZERO)                /**< Sets the value to 0. */
	GPU_STENCILOP_REPLACE             = GPUStencilOp(C.SDL_GPU_STENCILOP_REPLACE)             /**< Sets the value to reference. */
	GPU_STENCILOP_INCREMENT_AND_CLAMP = GPUStencilOp(C.SDL_GPU_STENCILOP_INCREMENT_AND_CLAMP) /**< Increments the current value and clamps to the maximum value. */
	GPU_STENCILOP_DECREMENT_AND_CLAMP = GPUStencilOp(C.SDL_GPU_STENCILOP_DECREMENT_AND_CLAMP) /**< Decrements the current value and clamps to 0. */
	GPU_STENCILOP_INVERT              = GPUStencilOp(C.SDL_GPU_STENCILOP_INVERT)              /**< Bitwise-inverts the current value. */
	GPU_STENCILOP_INCREMENT_AND_WRAP  = GPUStencilOp(C.SDL_GPU_STENCILOP_INCREMENT_AND_WRAP)  /**< Increments the current value and wraps back to 0. */
	GPU_STENCILOP_DECREMENT_AND_WRAP  = GPUStencilOp(C.SDL_GPU_STENCILOP_DECREMENT_AND_WRAP)  /**< Decrements the current value and wraps to the maximum value. */
)

// Specifies the operator to be used when pixels in a render target are
// blended with existing pixels in the texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBlendOp)
type GPUBlendOp C.SDL_GPUBlendOp

const (
	GPU_BLENDOP_INVALID          = GPUBlendOp(C.SDL_GPU_BLENDOP_INVALID)
	GPU_BLENDOP_ADD              = GPUBlendOp(C.SDL_GPU_BLENDOP_ADD)              /**< (source * source_factor) + (destination * destination_factor) */
	GPU_BLENDOP_SUBTRACT         = GPUBlendOp(C.SDL_GPU_BLENDOP_SUBTRACT)         /**< (source * source_factor) - (destination * destination_factor) */
	GPU_BLENDOP_REVERSE_SUBTRACT = GPUBlendOp(C.SDL_GPU_BLENDOP_REVERSE_SUBTRACT) /**< (destination * destination_factor) - (source * source_factor) */
	GPU_BLENDOP_MIN              = GPUBlendOp(C.SDL_GPU_BLENDOP_MIN)              /**< min(source, destination) */
	GPU_BLENDOP_MAX              = GPUBlendOp(C.SDL_GPU_BLENDOP_MAX)              /**< max(source, destination) */
)

// Specifies a blending factor to be used when pixels in a render target are
// blended with existing pixels in the texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBlendFactor)
type GPUBlendFactor C.SDL_GPUBlendFactor

const (
	GPU_BLENDFACTOR_INVALID                  = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_INVALID)
	GPU_BLENDFACTOR_ZERO                     = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ZERO)                     /**< 0 */
	GPU_BLENDFACTOR_ONE                      = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ONE)                      /**< 1 */
	GPU_BLENDFACTOR_SRC_COLOR                = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_SRC_COLOR)                /**< source color */
	GPU_BLENDFACTOR_ONE_MINUS_SRC_COLOR      = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ONE_MINUS_SRC_COLOR)      /**< 1 - source color */
	GPU_BLENDFACTOR_DST_COLOR                = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_DST_COLOR)                /**< destination color */
	GPU_BLENDFACTOR_ONE_MINUS_DST_COLOR      = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ONE_MINUS_DST_COLOR)      /**< 1 - destination color */
	GPU_BLENDFACTOR_SRC_ALPHA                = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_SRC_ALPHA)                /**< source alpha */
	GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA      = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA)      /**< 1 - source alpha */
	GPU_BLENDFACTOR_DST_ALPHA                = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_DST_ALPHA)                /**< destination alpha */
	GPU_BLENDFACTOR_ONE_MINUS_DST_ALPHA      = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ONE_MINUS_DST_ALPHA)      /**< 1 - destination alpha */
	GPU_BLENDFACTOR_CONSTANT_COLOR           = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_CONSTANT_COLOR)           /**< blend constant */
	GPU_BLENDFACTOR_ONE_MINUS_CONSTANT_COLOR = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_ONE_MINUS_CONSTANT_COLOR) /**< 1 - blend constant */
	GPU_BLENDFACTOR_SRC_ALPHA_SATURATE       = GPUBlendFactor(C.SDL_GPU_BLENDFACTOR_SRC_ALPHA_SATURATE)       /**< min(source alpha, 1 - destination alpha) */
)

// Specifies which color components are written in a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUColorComponentFlags)
type GPUColorComponentFlags C.SDL_GPUColorComponentFlags

const (
	GPU_COLORCOMPONENT_R = GPUColorComponentFlags(C.SDL_GPU_COLORCOMPONENT_R) /**< the red component */
	GPU_COLORCOMPONENT_G = GPUColorComponentFlags(C.SDL_GPU_COLORCOMPONENT_G) /**< the green component */
	GPU_COLORCOMPONENT_B = GPUColorComponentFlags(C.SDL_GPU_COLORCOMPONENT_B) /**< the blue component */
	GPU_COLORCOMPONENT_A = GPUColorComponentFlags(C.SDL_GPU_COLORCOMPONENT_A) /**< the alpha component */
)

// Specifies a filter operation used by a sampler.
// (https://wiki.libsdl.org/SDL3/SDL_GPUFilter)
type GPUFilter C.SDL_GPUFilter

const (
	GPU_FILTER_NEAREST = GPUFilter(C.SDL_GPU_FILTER_NEAREST) /**< Point filtering. */
	GPU_FILTER_LINEAR  = GPUFilter(C.SDL_GPU_FILTER_LINEAR)  /**< Linear filtering. */
)

// Specifies a mipmap mode used by a sampler.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSamplerMipmapMode)
type GPUSamplerMipmapMode C.SDL_GPUSamplerMipmapMode

const (
	GPU_SAMPLERMIPMAPMODE_NEAREST = GPUSamplerMipmapMode(C.SDL_GPU_SAMPLERMIPMAPMODE_NEAREST) /**< Point filtering. */
	GPU_SAMPLERMIPMAPMODE_LINEAR  = GPUSamplerMipmapMode(C.SDL_GPU_SAMPLERMIPMAPMODE_LINEAR)  /**< Linear filtering. */
)

// Specifies behavior of texture sampling when the coordinates exceed the 0-1
// range.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSamplerAddressMode)
type GPUSamplerAddressMode C.SDL_GPUSamplerAddressMode

const (
	GPU_SAMPLERADDRESSMODE_REPEAT          = GPUSamplerAddressMode(C.SDL_GPU_SAMPLERADDRESSMODE_REPEAT)          /**< Specifies that the coordinates will wrap around. */
	GPU_SAMPLERADDRESSMODE_MIRRORED_REPEAT = GPUSamplerAddressMode(C.SDL_GPU_SAMPLERADDRESSMODE_MIRRORED_REPEAT) /**< Specifies that the coordinates will wrap around mirrored. */
	GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE   = GPUSamplerAddressMode(C.SDL_GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE)   /**< Specifies that the coordinates will clamp to the 0-1 range. */
)

// Specifies the timing that will be used to present swapchain textures to the
// OS.
// (https://wiki.libsdl.org/SDL3/SDL_GPUPresentMode)
type GPUPresentMode C.SDL_GPUPresentMode

const (
	GPU_PRESENTMODE_VSYNC     = GPUPresentMode(C.SDL_GPU_PRESENTMODE_VSYNC)
	GPU_PRESENTMODE_IMMEDIATE = GPUPresentMode(C.SDL_GPU_PRESENTMODE_IMMEDIATE)
	GPU_PRESENTMODE_MAILBOX   = GPUPresentMode(C.SDL_GPU_PRESENTMODE_MAILBOX)
)

// Specifies the texture format and colorspace of the swapchain textures.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSwapchainComposition)
type GPUSwapchainComposition C.SDL_GPUSwapchainComposition

const (
	GPU_SWAPCHAINCOMPOSITION_SDR                 = GPUSwapchainComposition(C.SDL_GPU_SWAPCHAINCOMPOSITION_SDR)
	GPU_SWAPCHAINCOMPOSITION_SDR_LINEAR          = GPUSwapchainComposition(C.SDL_GPU_SWAPCHAINCOMPOSITION_SDR_LINEAR)
	GPU_SWAPCHAINCOMPOSITION_HDR_EXTENDED_LINEAR = GPUSwapchainComposition(C.SDL_GPU_SWAPCHAINCOMPOSITION_HDR_EXTENDED_LINEAR)
	GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084        = GPUSwapchainComposition(C.SDL_GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084)
)

//  /* Structures */

// A structure specifying a viewport.
// (https://wiki.libsdl.org/SDL3/SDL_GPUViewport)
type GPUViewport struct {
	X        float32 /**< The left offset of the viewport. */
	Y        float32 /**< The top offset of the viewport. */
	W        float32 /**< The width of the viewport. */
	H        float32 /**< The height of the viewport. */
	MinDepth float32 /**< The minimum depth of the viewport. */
	MaxDepth float32 /**< The maximum depth of the viewport. */
}

// A structure specifying parameters related to transferring data to or from a
// texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureTransferInfo)
type GPUTextureTransferInfo struct {
	TransferBuffer *GPUTransferBuffer /**< The transfer buffer used in the transfer operation. */
	Offset         uint32             /**< The starting byte of the image data in the transfer buffer. */
	PixelsPerRow   uint32             /**< The number of pixels from one row to the next. */
	RowsPerLayer   uint32             /**< The number of rows from one layer/depth-slice to the next. */
}

// A structure specifying a location in a transfer buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferLocation)
type GPUTransferBufferLocation struct {
	TransferBuffer *GPUTransferBuffer /**< The transfer buffer used in the transfer operation. */
	Offset         uint32             /**< The starting byte of the buffer data in the transfer buffer. */
}

// A structure specifying a location in a texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureLocation)
type GPUTextureLocation struct {
	Texture  *GPUTexture /**< The texture used in the copy operation. */
	MipLevel uint32      /**< The mip level index of the location. */
	Layer    uint32      /**< The layer index of the location. */
	X        uint32      /**< The left offset of the location. */
	Y        uint32      /**< The top offset of the location. */
	Z        uint32      /**< The front offset of the location. */
}

// A structure specifying a region of a texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureRegion)
type GPUTextureRegion struct {
	Texture  *GPUTexture /**< The texture used in the copy operation. */
	MipLevel uint32      /**< The mip level index to transfer. */
	Layer    uint32      /**< The layer index to transfer. */
	X        uint32      /**< The left offset of the region. */
	Y        uint32      /**< The top offset of the region. */
	Z        uint32      /**< The front offset of the region. */
	W        uint32      /**< The width of the region. */
	H        uint32      /**< The height of the region. */
	D        uint32      /**< The depth of the region. */
}

// A structure specifying a region of a texture used in the blit operation.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBlitRegion)
type GPUBlitRegion struct {
	Texture           *GPUTexture /**< The texture. */
	MipLevel          uint32      /**< The mip level index of the region. */
	LayerOrDepthPlane uint32      /**< The layer index or depth plane of the region. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures. */
	X                 uint32      /**< The left offset of the region. */
	Y                 uint32      /**< The top offset of the region.  */
	W                 uint32      /**< The width of the region. */
	H                 uint32      /**< The height of the region. */
}

// A structure specifying a location in a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBufferLocation)
type GPUBufferLocation struct {
	Buffer *GPUBuffer /**< The buffer. */
	Offset uint32     /**< The starting byte within the buffer. */
}

// A structure specifying a region of a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBufferRegion)
type GPUBufferRegion struct {
	Buffer *GPUBuffer /**< The buffer. */
	Offset uint32     /**< The starting byte within the buffer. */
	Size   uint32     /**< The size in bytes of the region. */
}

// A structure specifying the parameters of an indirect draw command.
// (https://wiki.libsdl.org/SDL3/SDL_GPUIndirectDrawCommand)
type GPUIndirectDrawCommand struct {
	NumVertices   uint32 /**< The number of vertices to draw. */
	NumInstances  uint32 /**< The number of instances to draw. */
	FirstVertex   uint32 /**< The index of the first vertex to draw. */
	FirstInstance uint32 /**< The ID of the first instance to draw. */
}

// A structure specifying the parameters of an indexed indirect draw command.
// (https://wiki.libsdl.org/SDL3/SDL_GPUIndexedIndirectDrawCommand)
type GPUIndexedIndirectDrawCommand struct {
	NumIndices    uint32 /**< The number of indices to draw per instance. */
	NumInstances  uint32 /**< The number of instances to draw. */
	FirstIndex    uint32 /**< The base index within the index buffer. */
	VertexOffset  int32  /**< The value added to the vertex index before indexing into the vertex buffer. */
	FirstInstance uint32 /**< The ID of the first instance to draw. */
}

// A structure specifying the parameters of an indexed dispatch command.
// (https://wiki.libsdl.org/SDL3/SDL_GPUIndirectDispatchCommand)
type GPUIndirectDispatchCommand struct {
	GroupcountX uint32 /**< The number of local workgroups to dispatch in the X dimension. */
	GroupcountY uint32 /**< The number of local workgroups to dispatch in the Y dimension. */
	GroupcountZ uint32 /**< The number of local workgroups to dispatch in the Z dimension. */
}

// State structures

// A structure specifying the parameters of a sampler.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSamplerCreateInfo)
type GPUSamplerCreateInfo struct {
	MinFilter        GPUFilter             /**< The minification filter to apply to lookups. */
	MagFilter        GPUFilter             /**< The magnification filter to apply to lookups. */
	MipmapMode       GPUSamplerMipmapMode  /**< The mipmap filter to apply to lookups. */
	AddressModeU     GPUSamplerAddressMode /**< The addressing mode for U coordinates outside [0, 1). */
	AddressModeV     GPUSamplerAddressMode /**< The addressing mode for V coordinates outside [0, 1). */
	AddressModeW     GPUSamplerAddressMode /**< The addressing mode for W coordinates outside [0, 1). */
	MipLodBias       float32               /**< The bias to be added to mipmap LOD calculation. */
	MaxAnisotropy    float32               /**< The anisotropy value clamp used by the sampler. If enable_anisotropy is false, this is ignored. */
	CompareOp        GPUCompareOp          /**< The comparison operator to apply to fetched data before filtering. */
	MinLod           float32               /**< Clamps the minimum of the computed LOD value. */
	MaxLod           float32               /**< Clamps the maximum of the computed LOD value. */
	EnableAnisotropy bool                  /**< true to enable anisotropic filtering. */
	EnableCompare    bool                  /**< true to enable comparison against a reference value during lookups. */
	_                uint8
	_                uint8
	Props            PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

// A structure specifying the parameters of vertex buffers used in a graphics
// pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexBufferDescription)
type GPUVertexBufferDescription struct {
	Slot             uint32             /**< The binding slot of the vertex buffer. */
	Pitch            uint32             /**< The byte pitch between consecutive elements of the vertex buffer. */
	InputRate        GPUVertexInputRate /**< Whether attribute addressing is a function of the vertex index or instance index. */
	InstanceStepRate uint32             /**< The number of instances to draw using the same per-instance data before advancing in the instance buffer by one element. Ignored unless input_rate is SDL_GPU_VERTEXINPUTRATE_INSTANCE */
}

// A structure specifying a vertex attribute.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexAttribute)
type GPUVertexAttribute struct {
	Location   uint32                 /**< The shader input location index. */
	BufferSlot uint32                 /**< The binding slot of the associated vertex buffer. */
	Format     GPUVertexElementFormat /**< The size and type of the attribute data. */
	Offset     uint32                 /**< The byte offset of this attribute relative to the start of the vertex element. */
}

// TODO: needs tested, not sure if pointers will pass to C or if we need to use a cGPUVertexInputState that has c types.
// A structure specifying the parameters of a graphics pipeline vertex input
// state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputState)
type GPUVertexInputState struct {
	VertexBufferDescriptions *GPUVertexBufferDescription /**< A pointer to an array of vertex buffer descriptions. */
	NumVertexBuffers         uint32                      /**< The number of vertex buffer descriptions in the above array. */
	VertexAttributes         *GPUVertexAttribute         /**< A pointer to an array of vertex attribute descriptions. */
	NumVertexAttributes      uint32                      /**< The number of vertex attribute descriptions in the above array. */
}

// A structure specifying the stencil operation state of a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUStencilOpState)
type GPUStencilOpState struct {
	FailOp      GPUStencilOp /**< The action performed on samples that fail the stencil test. */
	PassOp      GPUStencilOp /**< The action performed on samples that pass the depth and stencil tests. */
	DepthFailOp GPUStencilOp /**< The action performed on samples that pass the stencil test and fail the depth test. */
	CompareOp   GPUCompareOp /**< The comparison operator used in the stencil test. */
}

// A structure specifying the blend state of a color target.
// (https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetBlendState)
type GPUColorTargetBlendState struct {
	SrcColorBlendfactor  GPUBlendFactor         /**< The value to be multiplied by the source RGB value. */
	DstColorBlendfactor  GPUBlendFactor         /**< The value to be multiplied by the destination RGB value. */
	ColorBlendOp         GPUBlendOp             /**< The blend operation for the RGB components. */
	SrcAlphaBlendfactor  GPUBlendFactor         /**< The value to be multiplied by the source alpha. */
	DstAlphaBlendfactor  GPUBlendFactor         /**< The value to be multiplied by the destination alpha. */
	AlphaBlendOp         GPUBlendOp             /**< The blend operation for the alpha component. */
	ColorWriteMask       GPUColorComponentFlags /**< A bitmask specifying which of the RGBA components are enabled for writing. Writes to all channels if enable_color_write_mask is false. */
	EnableBlend          bool                   /**< Whether blending is enabled for the color target. */
	EnableColorWriteMask bool                   /**< Whether the color write mask is enabled. */
	_                    uint8
	_                    uint8
}

// TODO: needs tested. will need to make a cGPUShaderCreateInfo with c allocated strings.
//
//	which means we will need a free call on the created c struct etc.
//
// A structure specifying code and metadata for creating a shader object.
// (https://wiki.libsdl.org/SDL3/SDL_GPUShaderCreateInfo)
type GPUShaderCreateInfo struct {
	// CodeSize           int             /**< The size in bytes of the code pointed to. */
	Code               []byte          /**< A pointer to shader code. */
	Entrypoint         string          /**< A pointer to a null-terminated UTF-8 string specifying the entry point function name for the shader. */
	Format             GPUShaderFormat /**< The format of the shader code. */
	Stage              GPUShaderStage  /**< The stage the shader program corresponds to. */
	NumSamplers        uint32          /**< The number of samplers defined in the shader. */
	NumStorageTextures uint32          /**< The number of storage textures defined in the shader. */
	NumStorageBuffers  uint32          /**< The number of storage buffers defined in the shader. */
	NumUniformBuffers  uint32          /**< The number of uniform buffers defined in the shader. */
	Props              PropertiesID    /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

// A structure specifying the parameters of a texture.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureCreateInfo)
type GPUTextureCreateInfo struct {
	Type              GPUTextureType       /**< The base dimensionality of the texture. */
	Format            GPUTextureFormat     /**< The pixel format of the texture. */
	Usage             GPUTextureUsageFlags /**< How the texture is intended to be used by the client. */
	Width             uint32               /**< The width of the texture. */
	Height            uint32               /**< The height of the texture. */
	LayerCountOrDepth uint32               /**< The layer count or depth of the texture. This value is treated as a layer count on 2D array textures, and as a depth value on 3D textures. */
	NumLevels         uint32               /**< The number of mip levels in the texture. */
	SampleCount       GPUSampleCount       /**< The number of samples per texel. Only applies if the texture is used as a render target. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

const (
	SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_R_FLOAT       = C.SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_R_FLOAT       //"SDL.gpu.createtexture.d3d12.clear.r"
	SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_G_FLOAT       = C.SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_G_FLOAT       //"SDL.gpu.createtexture.d3d12.clear.g"
	SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_B_FLOAT       = C.SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_B_FLOAT       //"SDL.gpu.createtexture.d3d12.clear.b"
	SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_A_FLOAT       = C.SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_A_FLOAT       //"SDL.gpu.createtexture.d3d12.clear.a"
	SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_DEPTH_FLOAT   = C.SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_DEPTH_FLOAT   //"SDL.gpu.createtexture.d3d12.clear.depth"
	SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_STENCIL_UINT8 = C.SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_STENCIL_UINT8 //"SDL.gpu.createtexture.d3d12.clear.stencil"
)

// A structure specifying the parameters of a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBufferCreateInfo)
type GPUBufferCreateInfo struct {
	Usage GPUBufferUsageFlags /**< How the buffer is intended to be used by the client. */
	Size  uint32              /**< The size in bytes of the buffer. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

// A structure specifying the parameters of a transfer buffer.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferCreateInfo)
type GPUTransferBufferCreateInfo struct {
	Usage GPUTransferBufferUsage /**< How the transfer buffer is intended to be used by the client. */
	Size  uint32                 /**< The size in bytes of the transfer buffer. */

	Props PropertiesID /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

// Pipeline state structures
// A structure specifying the parameters of the graphics pipeline rasterizer
// state.
// (https://wiki.libsdl.org/SDL3/SDL_GPURasterizerState)
type GPURasterizerState struct {
	FillMode                GPUFillMode  /**< Whether polygons will be filled in or drawn as lines. */
	CullMode                GPUCullMode  /**< The facing direction in which triangles will be culled. */
	FrontFace               GPUFrontFace /**< The vertex winding that will cause a triangle to be determined as front-facing. */
	DepthBiasConstantFactor float32      /**< A scalar factor controlling the depth value added to each fragment. */
	DepthBiasClamp          float32      /**< The maximum depth bias of a fragment. */
	DepthBiasSlopeFactor    float32      /**< A scalar factor applied to a fragment's slope in depth calculations. */
	EnableDepthBias         bool         /**< true to bias fragment depth values. */
	EnableDepthClip         bool         /**< true to enable depth clip, false to enable depth clamp. */
	_                       uint8
	_                       uint8
}

// A structure specifying the parameters of the graphics pipeline multisample
// state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUMultisampleState)
type GPUMultisampleState struct {
	SampleCount GPUSampleCount /**< The number of samples to be used in rasterization. */
	SampleMask  uint32         /**< Determines which samples get updated in the render targets. Treated as 0xFFFFFFFF if enable_mask is false. */
	EnableMask  bool           /**< Enables sample masking. */
	_           uint8
	_           uint8
	_           uint8
}

// A structure specifying the parameters of the graphics pipeline depth
// stencil state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUDepthStencilState)
type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp      /**< The comparison operator used for depth testing. */
	BackStencilState  GPUStencilOpState /**< The stencil op state for back-facing triangles. */
	FrontStencilState GPUStencilOpState /**< The stencil op state for front-facing triangles. */
	CompareMask       uint8             /**< Selects the bits of the stencil values participating in the stencil test. */
	WriteMask         uint8             /**< Selects the bits of the stencil values updated by the stencil test. */
	EnableDepthTest   bool              /**< true enables the depth test. */
	EnableDepthWrite  bool              /**< true enables depth writes. Depth writes are always disabled when enable_depth_test is false. */
	EnableStencilTest bool              /**< true enables the stencil test. */
	_                 uint8
	_                 uint8
	_                 uint8
}

// A structure specifying the parameters of color targets used in a graphics
// pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetDescription)
type GPUColorTargetDescription struct {
	Format     GPUTextureFormat         /**< The pixel format of the texture to be used as a color target. */
	BlendState GPUColorTargetBlendState /**< The blend state to be used for the color target. */
}

// TODO: needs tested, struct has a pointer to a struct
// A structure specifying the descriptions of render targets used in a
// graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineTargetInfo)
type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions *GPUColorTargetDescription /**< A pointer to an array of color target descriptions. */
	NumColorTargets         uint32                     /**< The number of color target descriptions in the above array. */
	DepthStencilFormat      GPUTextureFormat           /**< The pixel format of the depth-stencil target. Ignored if has_depth_stencil_target is false. */
	HasDepthStencilTarget   bool                       /**< true specifies that the pipeline uses a depth-stencil target. */
	_                       uint8
	_                       uint8
	_                       uint8
}

// A structure specifying the parameters of a graphics pipeline state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineCreateInfo)
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader      *GPUShader                    /**< The vertex shader used by the graphics pipeline. */
	FragmentShader    *GPUShader                    /**< The fragment shader used by the graphics pipeline. */
	VertexInputState  GPUVertexInputState           /**< The vertex layout of the graphics pipeline. */
	PrimitiveType     GPUPrimitiveType              /**< The primitive topology of the graphics pipeline. */
	RasterizerState   GPURasterizerState            /**< The rasterizer state of the graphics pipeline. */
	MultisampleState  GPUMultisampleState           /**< The multisample state of the graphics pipeline. */
	DepthStencilState GPUDepthStencilState          /**< The depth-stencil state of the graphics pipeline. */
	TargetInfo        GPUGraphicsPipelineTargetInfo /**< Formats and blend modes for the render targets of the graphics pipeline. */
	Props             PropertiesID                  /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

// TODO: needs tested, struct has a pointer to a string
// A structure specifying the parameters of a compute pipeline state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUComputePipelineCreateInfo)
type GPUComputePipelineCreateInfo struct {
	// CodeSize                    int             /**< The size in bytes of the compute shader code pointed to. */
	Code                        []byte          /**< A pointer to compute shader code. */
	Entrypoint                  string          /**< A pointer to a null-terminated UTF-8 string specifying the entry point function name for the shader. */
	Format                      GPUShaderFormat /**< The format of the compute shader code. */
	NumSamplers                 uint32          /**< The number of samplers defined in the shader. */
	NumReadonlyStorageTextures  uint32          /**< The number of readonly storage textures defined in the shader. */
	NumReadonlyStorageBuffers   uint32          /**< The number of readonly storage buffers defined in the shader. */
	NumReadwriteStorageTextures uint32          /**< The number of read-write storage textures defined in the shader. */
	NumReadwriteStorageBuffers  uint32          /**< The number of read-write storage buffers defined in the shader. */
	NumUniformBuffers           uint32          /**< The number of uniform buffers defined in the shader. */
	ThreadcountX                uint32          /**< The number of threads in the X dimension. This should match the value in the shader. */
	ThreadcountY                uint32          /**< The number of threads in the Y dimension. This should match the value in the shader. */
	ThreadcountZ                uint32          /**< The number of threads in the Z dimension. This should match the value in the shader. */
	Props                       PropertiesID    /**< A properties ID for extensions. Should be 0 if no extensions are needed. */
}

// A structure specifying the parameters of a color target used by a render
// pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetInfo)
type GPUColorTargetInfo struct {
	Texture             *GPUTexture /**< The texture that will be used as a color target by a render pass. */
	MipLevel            uint32      /**< The mip level to use as a color target. */
	LayerOrDepthPlane   uint32      /**< The layer index or depth plane to use as a color target. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures. */
	ClearColor          FColor      /**< The color to clear the color target to at the start of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used. */
	LoadOp              GPULoadOp   /**< What is done with the contents of the color target at the beginning of the render pass. */
	StoreOp             GPUStoreOp  /**< What is done with the results of the render pass. */
	ResolveTexture      *GPUTexture /**< The texture that will receive the results of a multisample resolve operation. Ignored if a RESOLVE* store_op is not used. */
	ResolveMipLevel     uint32      /**< The mip level of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used. */
	ResolveLayer        uint32      /**< The layer index of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used. */
	Cycle               bool        /**< true cycles the texture if the texture is bound and load_op is not LOAD */
	CycleResolveTexture bool        /**< true cycles the resolve texture if the resolve texture is bound. Ignored if a RESOLVE* store_op is not used. */
	_                   uint8
	_                   uint8
}

// A structure specifying the parameters of a depth-stencil target used by a
// render pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUDepthStencilTargetInfo)
type GPUDepthStencilTargetInfo struct {
	Texture        *GPUTexture /**< The texture that will be used as the depth stencil target by the render pass. */
	ClearDepth     float32     /**< The value to clear the depth component to at the beginning of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used. */
	LoadOp         GPULoadOp   /**< What is done with the depth contents at the beginning of the render pass. */
	StoreOp        GPUStoreOp  /**< What is done with the depth results of the render pass. */
	StencilLoadOp  GPULoadOp   /**< What is done with the stencil contents at the beginning of the render pass. */
	StencilStoreOp GPUStoreOp  /**< What is done with the stencil results of the render pass. */
	Cycle          bool        /**< true cycles the texture if the texture is bound and any load ops are not LOAD */
	ClearStencil   uint8       /**< The value to clear the stencil component to at the beginning of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used. */
	_              uint8
	_              uint8
}

// A structure containing parameters for a blit command.
// (https://wiki.libsdl.org/SDL3/SDL_GPUBlitInfo)
type GPUBlitInfo struct {
	Source      GPUBlitRegion /**< The source region for the blit. */
	Destination GPUBlitRegion /**< The destination region for the blit. */
	LoadOp      GPULoadOp     /**< What is done with the contents of the destination before the blit. */
	ClearColor  FColor        /**< The color to clear the destination region to before the blit. Ignored if load_op is not SDL_GPU_LOADOP_CLEAR. */
	FlipMode    FlipMode      /**< The flip mode for the source region. */
	Filter      GPUFilter     /**< The filter mode used when blitting. */
	Cycle       bool          /**< true cycles the destination texture if it is already bound. */
	_           uint8
	_           uint8
	_           uint8
}

// (https://wiki.libsdl.org/SDL3/SDL_GPUBufferBinding)
type GPUBufferBinding struct {
	Buffer *GPUBuffer /**< The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_VERTEX for SDL_BindGPUVertexBuffers, or SDL_GPU_BUFFERUSAGE_INDEX for SDL_BindGPUIndexBuffers. */
	Offset uint32     /**< The starting byte of the data to bind in the buffer. */
}

// A structure specifying parameters in a sampler binding call.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureSamplerBinding)
type GPUTextureSamplerBinding struct {
	Texture *GPUTexture /**< The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER. */
	Sampler *GPUSampler /**< The sampler to bind. */
}

// A structure specifying parameters related to binding buffers in a compute
// pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUStorageBufferReadWriteBinding)
type GPUStorageBufferReadWriteBinding struct {
	Buffer *GPUBuffer /**< The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE. */
	Cycle  bool       /**< true cycles the buffer if it is already bound. */
	_      uint8
	_      uint8
	_      uint8
}

// A structure specifying parameters related to binding textures in a compute
// pass.
// (https://wiki.libsdl.org/SDL3/SDL_GPUStorageTextureReadWriteBinding)
type GPUStorageTextureReadWriteBinding struct {
	Texture  *GPUTexture /**< The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE or SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE. */
	MipLevel uint32      /**< The mip level index to bind. */
	Layer    uint32      /**< The layer index to bind. */
	Cycle    bool        /**< true cycles the texture if it is already bound. */
	_        uint8
	_        uint8
	_        uint8
}

// Functions

// Device

// Checks for GPU runtime support.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats)
func GPUSupportsShaderFormats(formatFlags GPUShaderFormat, name string) bool {
	var cstr *C.char
	if len(name) > 0 {
		cstr = C.CString(name)
		defer C.free(unsafe.Pointer(cstr))
	}
	return bool(C.SDL_GPUSupportsShaderFormats(C.SDL_GPUShaderFormat(formatFlags), C.CString(name)))
}

// Checks for GPU runtime support.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties)
func GPUSupportsProperties(props PropertiesID) bool {
	return bool(C.SDL_GPUSupportsProperties(C.SDL_PropertiesID(props)))
}

// Creates a GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice)
func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) (*GPUDevice, error) {
	var cstr *C.char
	if len(name) > 0 {
		cstr = C.CString(name)
		defer C.free(unsafe.Pointer(cstr))
	}
	ret := C.SDL_CreateGPUDevice(C.SDL_GPUShaderFormat(formatFlags), C.bool(debugMode), cstr)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUDevice)(ret), nil
}

// Creates a GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties)
func CreateGPUDeviceWithProperties(props PropertiesID) (*GPUDevice, error) {
	ret := C.SDL_CreateGPUDeviceWithProperties(C.SDL_PropertiesID(props))
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUDevice)(ret), nil
}

const (
	PROP_GPU_DEVICE_CREATE_DEBUGMODE_BOOLEAN          = C.SDL_PROP_GPU_DEVICE_CREATE_DEBUGMODE_BOOLEAN          // "SDL.gpu.device.create.debugmode"
	PROP_GPU_DEVICE_CREATE_PREFERLOWPOWER_BOOLEAN     = C.SDL_PROP_GPU_DEVICE_CREATE_PREFERLOWPOWER_BOOLEAN     // "SDL.gpu.device.create.preferlowpower"
	PROP_GPU_DEVICE_CREATE_NAME_STRING                = C.SDL_PROP_GPU_DEVICE_CREATE_NAME_STRING                // "SDL.gpu.device.create.name"
	PROP_GPU_DEVICE_CREATE_SHADERS_PRIVATE_BOOLEAN    = C.SDL_PROP_GPU_DEVICE_CREATE_SHADERS_PRIVATE_BOOLEAN    // "SDL.gpu.device.create.shaders.private"
	PROP_GPU_DEVICE_CREATE_SHADERS_SPIRV_BOOLEAN      = C.SDL_PROP_GPU_DEVICE_CREATE_SHADERS_SPIRV_BOOLEAN      // "SDL.gpu.device.create.shaders.spirv"
	PROP_GPU_DEVICE_CREATE_SHADERS_DXBC_BOOLEAN       = C.SDL_PROP_GPU_DEVICE_CREATE_SHADERS_DXBC_BOOLEAN       // "SDL.gpu.device.create.shaders.dxbc"
	PROP_GPU_DEVICE_CREATE_SHADERS_DXIL_BOOLEAN       = C.SDL_PROP_GPU_DEVICE_CREATE_SHADERS_DXIL_BOOLEAN       // "SDL.gpu.device.create.shaders.dxil"
	PROP_GPU_DEVICE_CREATE_SHADERS_MSL_BOOLEAN        = C.SDL_PROP_GPU_DEVICE_CREATE_SHADERS_MSL_BOOLEAN        // "SDL.gpu.device.create.shaders.msl"
	PROP_GPU_DEVICE_CREATE_SHADERS_METALLIB_BOOLEAN   = C.SDL_PROP_GPU_DEVICE_CREATE_SHADERS_METALLIB_BOOLEAN   // "SDL.gpu.device.create.shaders.metallib"
	PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING = C.SDL_PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING // "SDL.gpu.device.create.d3d12.semantic"
)

// Destroys a GPU context previously returned by SDL_CreateGPUDevice.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice)
func (dev *GPUDevice) Destroy() {
	C.SDL_DestroyGPUDevice((*C.SDL_GPUDevice)(dev))
}

// Get the number of GPU drivers compiled into SDL.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers)
func GetNumGPUDrivers() int32 {
	return int32(C.SDL_GetNumGPUDrivers())
}

// Get the name of a built in GPU driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver)
func GetGPUDriver(index int32) string {
	return C.GoString(C.SDL_GetGPUDriver(C.int(index)))
}

// Returns the name of the backend used to create this GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver)
func (dev *GPUDevice) GetDriver() string {
	return C.GoString(C.SDL_GetGPUDeviceDriver((*C.SDL_GPUDevice)(dev)))
}

// Returns the supported shader formats for this GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats)
func (dev *GPUDevice) GetShaderFormats() GPUShaderFormat {
	return GPUShaderFormat(C.SDL_GetGPUShaderFormats((*C.SDL_GPUDevice)(dev)))
}

// State Creation

// TODO: needs tested
// Creates a pipeline object to be used in a compute workflow.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline)
func (dev *GPUDevice) CreateComputePipeline(createinfo *GPUComputePipelineCreateInfo) (*GPUComputePipeline, error) {
	var cCreateInfo *C.SDL_GPUComputePipelineCreateInfo
	if createinfo != nil {
		// TODO: can I pass the code directly like this or do i need to make a copy?
		cCreateInfo = &C.SDL_GPUComputePipelineCreateInfo{
			code_size:                      (C.size_t)(len(createinfo.Code)),
			code:                           (*C.uint8_t)(unsafe.Pointer(&createinfo.Code[0])),
			entrypoint:                     C.CString(createinfo.Entrypoint),
			format:                         C.SDL_GPUShaderFormat(createinfo.Format),
			num_samplers:                   (C.uint)(createinfo.NumSamplers),
			num_readonly_storage_textures:  (C.uint)(createinfo.NumReadonlyStorageTextures),
			num_readonly_storage_buffers:   (C.uint)(createinfo.NumReadonlyStorageBuffers),
			num_readwrite_storage_textures: (C.uint)(createinfo.NumReadwriteStorageTextures),
			num_readwrite_storage_buffers:  (C.uint)(createinfo.NumReadwriteStorageBuffers),
			num_uniform_buffers:            (C.uint)(createinfo.NumUniformBuffers),
			threadcount_x:                  (C.uint)(createinfo.ThreadcountX),
			threadcount_y:                  (C.uint)(createinfo.ThreadcountY),
			threadcount_z:                  (C.uint)(createinfo.ThreadcountZ),
			props:                          C.SDL_PropertiesID(createinfo.Props),
		}
		defer C.SDL_free(unsafe.Pointer(cCreateInfo.entrypoint))
	}
	ret := C.SDL_CreateGPUComputePipeline((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUComputePipeline)(ret), nil
}

// Creates a pipeline object to be used in a graphics workflow.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline)
func (dev *GPUDevice) CreateGraphicsPipeline(createinfo *GPUGraphicsPipelineCreateInfo) (*GPUGraphicsPipeline, error) {
	cCreateInfo := (*C.SDL_GPUGraphicsPipelineCreateInfo)(unsafe.Pointer(createinfo))
	ret := C.SDL_CreateGPUGraphicsPipeline((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUGraphicsPipeline)(ret), nil
}

// Creates a sampler object to be used when binding textures in a graphics
// workflow.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler)
func (dev *GPUDevice) CreateSampler(createinfo *GPUSamplerCreateInfo) (*GPUSampler, error) {
	cCreateInfo := (*C.SDL_GPUSamplerCreateInfo)(unsafe.Pointer(createinfo))
	ret := C.SDL_CreateGPUSampler((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUSampler)(ret), nil
}

// TODO: needs tested
// Creates a shader to be used when creating a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader)
func (dev *GPUDevice) CreateShader(createinfo *GPUShaderCreateInfo) (*GPUShader, error) {
	var cCreateInfo *C.SDL_GPUShaderCreateInfo
	if createinfo != nil {
		// TODO: can I pass the code directly like this or do i need to make a copy?
		cCreateInfo = &C.SDL_GPUShaderCreateInfo{
			code_size:            (C.size_t)(len(createinfo.Code)),
			code:                 (*C.uint8_t)(unsafe.Pointer(&createinfo.Code[0])),
			entrypoint:           C.CString(createinfo.Entrypoint),
			format:               C.SDL_GPUShaderFormat(createinfo.Format),
			stage:                C.SDL_GPUShaderStage(createinfo.Stage),
			num_samplers:         (C.uint)(createinfo.NumSamplers),
			num_storage_textures: (C.uint)(createinfo.NumStorageTextures),
			num_storage_buffers:  (C.uint)(createinfo.NumStorageBuffers),
			num_uniform_buffers:  (C.uint)(createinfo.NumUniformBuffers),
			props:                C.SDL_PropertiesID(createinfo.Props),
		}
		defer C.SDL_free(unsafe.Pointer(cCreateInfo.entrypoint))
	}
	ret := C.SDL_CreateGPUShader((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUShader)(ret), nil
}

// Creates a texture object to be used in graphics or compute workflows.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture)
func (dev *GPUDevice) CreateTexture(createinfo *GPUTextureCreateInfo) (*GPUTexture, error) {
	cCreateInfo := (*C.SDL_GPUTextureCreateInfo)(unsafe.Pointer(createinfo))
	ret := C.SDL_CreateGPUTexture((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUTexture)(ret), nil
}

// Creates a buffer object to be used in graphics or compute workflows.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer)
func (dev *GPUDevice) CreateBuffer(createinfo *GPUBufferCreateInfo) (*GPUBuffer, error) {
	cCreateInfo := (*C.SDL_GPUBufferCreateInfo)(unsafe.Pointer(createinfo))
	ret := C.SDL_CreateGPUBuffer((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUBuffer)(ret), nil
}

// Creates a transfer buffer to be used when uploading to or downloading from
// graphics resources.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer)
func (dev *GPUDevice) CreateTransferBuffer(createinfo *GPUTransferBufferCreateInfo) (*GPUTransferBuffer, error) {
	cCreateInfo := (*C.SDL_GPUTransferBufferCreateInfo)(unsafe.Pointer(createinfo))
	ret := C.SDL_CreateGPUTransferBuffer((*C.SDL_GPUDevice)(dev), cCreateInfo)
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUTransferBuffer)(ret), nil
}

// Debug Naming

// Sets an arbitrary string constant to label a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName)
func (dev *GPUDevice) SetBufferName(buffer *GPUBuffer, text string) {
	cText := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(cText))
	C.SDL_SetGPUBufferName((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUBuffer)(buffer), cText)
}

// Sets an arbitrary string constant to label a texture.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName)
func (dev *GPUDevice) SetTextureName(texture *GPUTexture, text string) {
	cText := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(cText))
	C.SDL_SetGPUTextureName((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUTexture)(texture), cText)
}

// Inserts an arbitrary string label into the command buffer callstream.
// (https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel)
func (commandBuffer *GPUCommandBuffer) InsertDebugLabel(text string) {
	cText := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(cText))
	C.SDL_InsertGPUDebugLabel((*C.SDL_GPUCommandBuffer)(commandBuffer), cText)
}

// Begins a debug group with an arbitary name.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup)
func (commandBuffer *GPUCommandBuffer) PushDebugGroup(text string) {
	cText := C.CString(text)
	defer C.SDL_free(unsafe.Pointer(cText))
	C.SDL_PushGPUDebugGroup((*C.SDL_GPUCommandBuffer)(commandBuffer), cText)
}

// Ends the most-recently pushed debug group.
// (https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup)
func (commandBuffer *GPUCommandBuffer) PopDebugGroup() {
	C.SDL_PopGPUDebugGroup((*C.SDL_GPUCommandBuffer)(commandBuffer))
}

// Disposal

// Frees the given texture as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture)
func (dev *GPUDevice) ReleaseTexture(texture *GPUTexture) {
	C.SDL_ReleaseGPUTexture((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUTexture)(texture))
}

// Frees the given sampler as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler)
func (dev *GPUDevice) ReleaseSampler(sampler *GPUSampler) {
	C.SDL_ReleaseGPUSampler((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUSampler)(sampler))
}

// Frees the given buffer as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer)
func (dev *GPUDevice) ReleaseBuffer(buffer *GPUBuffer) {
	C.SDL_ReleaseGPUBuffer((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUBuffer)(buffer))
}

// Frees the given transfer buffer as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer)
func (dev *GPUDevice) ReleaseTransferBuffer(transferBuffer *GPUTransferBuffer) {
	C.SDL_ReleaseGPUTransferBuffer((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUTransferBuffer)(transferBuffer))
}

// Frees the given compute pipeline as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline)
func (dev *GPUDevice) ReleaseComputePipeline(computePipeline *GPUComputePipeline) {
	C.SDL_ReleaseGPUComputePipeline((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUComputePipeline)(computePipeline))
}

// Frees the given shader as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader)
func (dev *GPUDevice) ReleaseShader(shader *GPUShader) {
	C.SDL_ReleaseGPUShader((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUShader)(shader))
}

// Frees the given graphics pipeline as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline)
func (dev *GPUDevice) ReleaseGraphicsPipeline(graphicsPipeline *GPUGraphicsPipeline) {
	C.SDL_ReleaseGPUGraphicsPipeline((*C.SDL_GPUDevice)(dev), (*C.SDL_GPUGraphicsPipeline)(graphicsPipeline))
}

// Acquire a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer)
func (dev *GPUDevice) AcquireCommandBuffer() (*GPUCommandBuffer, error) {
	ret := C.SDL_AcquireGPUCommandBuffer((*C.SDL_GPUDevice)(dev))
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUCommandBuffer)(ret), nil
}

// UNIFORM DATA

// Pushes data to a vertex uniform slot on the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData)
func (commandBuffer *GPUCommandBuffer) PushVertexUniformData(slotIndex uint32, data []byte) {
	C.SDL_PushGPUVertexUniformData(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		C.Uint32(slotIndex), unsafe.Pointer(&data[0]), C.Uint32(len(data)))
}

// Pushes data to a fragment uniform slot on the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData)
func (commandBuffer *GPUCommandBuffer) PushFragmentUniformData(slotIndex uint32, data []byte) {
	C.SDL_PushGPUFragmentUniformData(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		C.Uint32(slotIndex), unsafe.Pointer(&data[0]), C.Uint32(len(data)))
}

// Pushes data to a uniform slot on the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData)
func (commandBuffer *GPUCommandBuffer) PushComputeUniformData(slotIndex uint32, data []byte) {
	C.SDL_PushGPUComputeUniformData(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		C.Uint32(slotIndex), unsafe.Pointer(&data[0]), C.Uint32(len(data)))
}

// A NOTE ON CYCLING

// Graphics State

// Begins a render pass on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass)
func (commandBuffer *GPUCommandBuffer) BeginRenderPass(
	colorTargetInfos []GPUColorTargetInfo, depthStencilTargetInfo *GPUDepthStencilTargetInfo) (*GPURenderPass, error) {

	data := (*C.SDL_GPUColorTargetInfo)(unsafe.Pointer(unsafe.SliceData(colorTargetInfos)))
	dataLen := C.Uint32(len(colorTargetInfos))
	ret := C.SDL_BeginGPURenderPass(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		data, dataLen,
		(*C.SDL_GPUDepthStencilTargetInfo)(unsafe.Pointer(depthStencilTargetInfo)))
	if ret == nil {
		return nil, GetError()
	}
	return (*GPURenderPass)(ret), nil
}

// Binds a graphics pipeline on a render pass to be used in rendering.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline)
func (renderPass *GPURenderPass) BindGraphicsPipeline(graphicsPipeline *GPUGraphicsPipeline) {
	C.SDL_BindGPUGraphicsPipeline((*C.SDL_GPURenderPass)(renderPass), (*C.SDL_GPUGraphicsPipeline)(graphicsPipeline))
}

// Sets the current viewport state on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport)
func (renderPass *GPURenderPass) SetViewport(viewport *GPUViewport) {
	C.SDL_SetGPUViewport((*C.SDL_GPURenderPass)(renderPass), (*C.SDL_GPUViewport)(unsafe.Pointer(viewport)))
}

// Sets the current scissor state on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor)
func (renderPass *GPURenderPass) SetScissor(scissor *Rect) {
	C.SDL_SetGPUScissor((*C.SDL_GPURenderPass)(renderPass), (*C.SDL_Rect)(unsafe.Pointer(scissor)))
}

// TODO: is there a better way to convert from FColor to SDL_Color?
// Sets the current blend constants on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants)
func (renderPass *GPURenderPass) SetBlendConstants(blendConstants FColor) {
	C.SDL_SetGPUBlendConstants((*C.SDL_GPURenderPass)(renderPass), *((*C.SDL_FColor)(unsafe.Pointer(&blendConstants))))
}

// Sets the current stencil reference value on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference)
func (renderPass *GPURenderPass) SetStencilReference(reference uint8) {
	C.SDL_SetGPUStencilReference((*C.SDL_GPURenderPass)(renderPass), C.Uint8(reference))
}

// Binds vertex buffers on a command buffer for use with subsequent draw
// calls.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers)
func (renderPass *GPURenderPass) BindVertexBuffers(
	firstSlot uint32, bindings []GPUBufferBinding, indexElementSize GPUIndexElementSize) {

	data := (*C.SDL_GPUBufferBinding)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUVertexBuffers(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), data, dataLen)
}

// Binds an index buffer on a command buffer for use with subsequent draw
// calls.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer)
func (renderPass *GPURenderPass) BindIndexBuffer(binding *GPUBufferBinding, indexElementSize GPUIndexElementSize) {
	C.SDL_BindGPUIndexBuffer(
		(*C.SDL_GPURenderPass)(renderPass),
		(*C.SDL_GPUBufferBinding)(unsafe.Pointer(binding)),
		C.SDL_GPUIndexElementSize(indexElementSize))
}

// Binds texture-sampler pairs for use on the vertex shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers)
func (renderPass *GPURenderPass) BindVertexSamplers(firstSlot uint32, bindings []GPUTextureSamplerBinding) {
	data := (*C.SDL_GPUTextureSamplerBinding)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUVertexSamplers(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), data, dataLen)
}

// TODO: is this right? pointer to pointer?
// Binds storage textures for use on the vertex shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures)
func (renderPass *GPURenderPass) BindVertexStorageTextures(firstSlot uint32, bindings []GPUTexture) {
	data := (*C.SDL_GPUTexture)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUVertexStorageTextures(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), &data, dataLen)
}

// TODO: is this right? pointer to pointer?
// Binds storage buffers for use on the vertex shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers)
func (renderPass *GPURenderPass) BindVertexStorageBuffers(firstSlot uint32, bindings []GPUBuffer) {
	data := (*C.SDL_GPUBuffer)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUVertexStorageBuffers(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), &data, dataLen)
}

// Binds texture-sampler pairs for use on the fragment shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers)
func (renderPass *GPURenderPass) BindFragmentSamplers(firstSlot uint32, bindings []GPUTextureSamplerBinding) {
	data := (*C.SDL_GPUTextureSamplerBinding)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUFragmentSamplers(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), data, dataLen)
}

// TODO: is this right? pointer to pointer?
// Binds storage textures for use on the fragment shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures)
func (renderPass *GPURenderPass) BindFragmentStorageTextures(firstSlot uint32, bindings []GPUTexture) {
	data := (*C.SDL_GPUTexture)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUFragmentStorageTextures(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), &data, dataLen)
}

// TODO: is this right? pointer to pointer?
// Binds storage buffers for use on the fragment shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers)
func (renderPass *GPURenderPass) BindFragmentStorageBuffers(firstSlot uint32, bindings []GPUBuffer) {
	data := (*C.SDL_GPUBuffer)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUFragmentStorageBuffers(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(firstSlot), &data, dataLen)
}

// Drawing

// Draws data using bound graphics state with an index buffer and instancing
// enabled.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives)
func (renderPass *GPURenderPass) DrawIndexedPrimitives(
	numIndices uint32, numInstances uint32, firstIndex uint32, vertexOffset int32, firstInstance uint32) {

	C.SDL_DrawGPUIndexedPrimitives(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(numIndices), C.Uint32(numInstances), C.Uint32(firstIndex),
		C.Sint32(vertexOffset), C.Uint32(firstInstance))
}

// Draws data using bound graphics state.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives)
func (renderPass *GPURenderPass) DrawPrimitives(
	numVertices uint32, numInstances uint32, firstVertex uint32, firstInstance uint32) {

	C.SDL_DrawGPUPrimitives(
		(*C.SDL_GPURenderPass)(renderPass),
		C.Uint32(numVertices), C.Uint32(numInstances), C.Uint32(firstVertex),
		C.Uint32(firstInstance))
}

// Draws data using bound graphics state and with draw parameters set from a
// buffer.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect)
func (renderPass *GPURenderPass) DrawPrimitivesIndirect(buffer *GPUBuffer, offset uint32, drawCount uint32) {
	C.SDL_DrawGPUPrimitivesIndirect(
		(*C.SDL_GPURenderPass)(renderPass),
		(*C.SDL_GPUBuffer)(unsafe.Pointer(buffer)), C.Uint32(offset), C.Uint32(drawCount))
}

// Draws data using bound graphics state with an index buffer enabled and with
// draw parameters set from a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect)
func (renderPass *GPURenderPass) DrawIndexedPrimitivesIndirect(buffer *GPUBuffer, offset uint32, drawCount uint32) {
	C.SDL_DrawGPUIndexedPrimitivesIndirect(
		(*C.SDL_GPURenderPass)(renderPass),
		(*C.SDL_GPUBuffer)(unsafe.Pointer(buffer)), C.Uint32(offset), C.Uint32(drawCount))
}

// Ends the given render pass.
// (https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass)
func (renderPass *GPURenderPass) End() {
	C.SDL_EndGPURenderPass(
		(*C.SDL_GPURenderPass)(renderPass))
}

// Compute Pass

// TODO: can storageTextureBindings and or storageBufferBindings be nil?
// Begins a compute pass on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass)
func (commandBuffer *GPUCommandBuffer) BeginComputePass(
	storageTextureBindings []GPUStorageTextureReadWriteBinding,
	storageBufferBindings []GPUStorageBufferReadWriteBinding) *GPUComputePass {

	data := (*C.SDL_GPUStorageTextureReadWriteBinding)(unsafe.Pointer(unsafe.SliceData(storageTextureBindings)))
	dataLen := C.Uint32(len(storageTextureBindings))
	data2 := (*C.SDL_GPUStorageBufferReadWriteBinding)(unsafe.Pointer(unsafe.SliceData(storageBufferBindings)))
	dataLen2 := C.Uint32(len(storageBufferBindings))
	return (*GPUComputePass)(C.SDL_BeginGPUComputePass(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		data, dataLen, data2, dataLen2))
}

// Binds a compute pipeline on a command buffer for use in compute dispatch.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline)
func (computePass *GPUComputePass) BindComputePipeline(computePipeline *GPUComputePipeline) {
	C.SDL_BindGPUComputePipeline(
		(*C.SDL_GPUComputePass)(computePass),
		(*C.SDL_GPUComputePipeline)(computePipeline))
}

// Binds texture-sampler pairs for use on the compute shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers)
func (computePass *GPUComputePass) BindComputeSamplers(firstSlot uint32, bindings []GPUTextureSamplerBinding) {
	data := (*C.SDL_GPUTextureSamplerBinding)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUComputeSamplers(
		(*C.SDL_GPUComputePass)(computePass),
		C.Uint32(firstSlot), data, dataLen)
}

// Binds storage textures as readonly for use on the compute pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures)
func (computePass *GPUComputePass) BindComputeStorageTextures(firstSlot uint32, bindings []GPUTexture) {
	data := (*C.SDL_GPUTexture)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUComputeStorageTextures(
		(*C.SDL_GPUComputePass)(computePass),
		C.Uint32(firstSlot), &data, dataLen)
}

// Binds storage buffers as readonly for use on the compute pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers)
func (computePass *GPUComputePass) BindComputeStorageBuffers(firstSlot uint32, bindings []GPUBuffer) {
	data := (*C.SDL_GPUBuffer)(unsafe.Pointer(unsafe.SliceData(bindings)))
	dataLen := C.Uint32(len(bindings))
	C.SDL_BindGPUComputeStorageBuffers(
		(*C.SDL_GPUComputePass)(computePass),
		C.Uint32(firstSlot), &data, dataLen)
}

// Dispatches compute work.
// (https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute)
func (computePass *GPUComputePass) Dispatch(groupCountX uint32, groupCountY uint32, groupCountZ uint32) {
	C.SDL_DispatchGPUCompute(
		(*C.SDL_GPUComputePass)(computePass),
		C.Uint32(groupCountX), C.Uint32(groupCountY), C.Uint32(groupCountZ))
}

// Dispatches compute work with parameters set from a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect)
func (computePass *GPUComputePass) DispatchIndirect(buffer *GPUBuffer, offset uint32) {
	C.SDL_DispatchGPUComputeIndirect(
		(*C.SDL_GPUComputePass)(computePass),
		(*C.SDL_GPUBuffer)(unsafe.Pointer(buffer)), C.Uint32(offset))
}

// Ends the current compute pass.
// (https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass)
func (computePass *GPUComputePass) End() {
	C.SDL_EndGPUComputePass(
		(*C.SDL_GPUComputePass)(computePass))
}

// TransferBuffer Data

// Maps a transfer buffer into application address space.
// (https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer)
func (device *GPUDevice) MapTransferBuffer(transferBuffer *GPUTransferBuffer, cycle bool) (unsafe.Pointer, error) {
	ret := C.SDL_MapGPUTransferBuffer(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_GPUTransferBuffer)(transferBuffer),
		C.bool(cycle))
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Pointer(ret), nil
}

// Unmaps a previously mapped transfer buffer.
// (https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer)
func (device *GPUDevice) UnmapTransferBuffer(transferBuffer *GPUTransferBuffer) {
	C.SDL_UnmapGPUTransferBuffer(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_GPUTransferBuffer)(transferBuffer))
}

// Copy Pass

// Begins a copy pass on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass)
func (commandBuffer *GPUCommandBuffer) BeginCopyPass() *GPUCopyPass {
	return (*GPUCopyPass)(C.SDL_BeginGPUCopyPass(
		(*C.SDL_GPUCommandBuffer)(commandBuffer)))
}

// Uploads data from a transfer buffer to a texture.
// (https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture)
func (copyPass *GPUCopyPass) UploadToTexture(
	source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {

	C.SDL_UploadToGPUTexture(
		(*C.SDL_GPUCopyPass)(copyPass),
		(*C.SDL_GPUTextureTransferInfo)(unsafe.Pointer(source)),
		(*C.SDL_GPUTextureRegion)(unsafe.Pointer(destination)),
		C.bool(cycle))
}

// Uploads data from a TransferBuffer to a Buffer.

// Uploads data from a transfer buffer to a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer)
func (copyPass *GPUCopyPass) UploadToBuffer(
	source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {

	C.SDL_UploadToGPUBuffer(
		(*C.SDL_GPUCopyPass)(copyPass),
		(*C.SDL_GPUTransferBufferLocation)(unsafe.Pointer(source)),
		(*C.SDL_GPUBufferRegion)(unsafe.Pointer(destination)),
		C.bool(cycle))
}

// Performs a texture-to-texture copy.
// (https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture)
func (copyPass *GPUCopyPass) CopyTextureToTexture(
	source *GPUTextureLocation, destination *GPUTextureLocation,
	width uint32, height uint32, depth uint32, cycle bool) {

	C.SDL_CopyGPUTextureToTexture(
		(*C.SDL_GPUCopyPass)(copyPass),
		(*C.SDL_GPUTextureLocation)(unsafe.Pointer(source)),
		(*C.SDL_GPUTextureLocation)(unsafe.Pointer(destination)),
		C.Uint32(width), C.Uint32(height), C.Uint32(depth), C.bool(cycle))
}

// Copies data from a buffer to a buffer.

// Performs a buffer-to-buffer copy.
// (https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer)
func (copyPass *GPUCopyPass) CopyBufferToBuffer(
	source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool) {

	C.SDL_CopyGPUBufferToBuffer(
		(*C.SDL_GPUCopyPass)(copyPass),
		(*C.SDL_GPUBufferLocation)(unsafe.Pointer(source)),
		(*C.SDL_GPUBufferLocation)(unsafe.Pointer(destination)),
		C.Uint32(size), C.bool(cycle))
}

// Copies data from a texture to a transfer buffer on the GPU timeline.
// (https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture)
func (copyPass *GPUCopyPass) DownloadFromTexture(source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
	C.SDL_DownloadFromGPUTexture(
		(*C.SDL_GPUCopyPass)(copyPass),
		(*C.SDL_GPUTextureRegion)(unsafe.Pointer(source)),
		(*C.SDL_GPUTextureTransferInfo)(unsafe.Pointer(destination)))
}

// Copies data from a buffer to a transfer buffer on the GPU timeline.
// (https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer)
func (copyPass *GPUCopyPass) DownloadFromBuffer(source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
	C.SDL_DownloadFromGPUBuffer(
		(*C.SDL_GPUCopyPass)(copyPass),
		(*C.SDL_GPUBufferRegion)(unsafe.Pointer(source)),
		(*C.SDL_GPUTransferBufferLocation)(unsafe.Pointer(destination)))
}

// Ends the current copy pass.
// (https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass)
func (copyPass *GPUCopyPass) End() {
	C.SDL_EndGPUCopyPass((*C.SDL_GPUCopyPass)(copyPass))
}

// Generates mipmaps for the given texture.
// (https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture)
func (commandBuffer *GPUCommandBuffer) GenerateMipmapsForTexture(texture *GPUTexture) {
	C.SDL_GenerateMipmapsForGPUTexture(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		(*C.SDL_GPUTexture)(texture))
}

// Blits from a source texture region to a destination texture region.
// (https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture)
func (commandBuffer *GPUCommandBuffer) BlitTexture(info *GPUBlitInfo) {
	C.SDL_BlitGPUTexture(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		(*C.SDL_GPUBlitInfo)(unsafe.Pointer(info)))
}

// Submission/Presentation

// Determines whether a swapchain composition is supported by the window.
// (https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition)
func (device *GPUDevice) WindowSupportsSwapchainComposition(
	window *Window, swapchainComposition GPUSwapchainComposition) bool {

	return bool(C.SDL_WindowSupportsGPUSwapchainComposition(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_Window)(window),
		C.SDL_GPUSwapchainComposition(swapchainComposition)))
}

// Determines whether a presentation mode is supported by the window.
// (https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode)
func (device *GPUDevice) WindowSupportsPresentMode(window *Window, presentMode GPUPresentMode) bool {
	return bool(C.SDL_WindowSupportsGPUPresentMode(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_Window)(window),
		C.SDL_GPUPresentMode(presentMode)))
}

// Claims a window, creating a swapchain structure for it.
// (https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice)
func (device *GPUDevice) ClaimWindow(window *Window) error {
	ret := C.SDL_ClaimWindowForGPUDevice((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window))
	if !ret {
		return GetError()
	}
	return nil
}

// Unclaims a window, destroying its swapchain structure.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice)
func (device *GPUDevice) ReleaseWindow(window *Window) {
	C.SDL_ReleaseWindowFromGPUDevice((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window))
}

// Changes the swapchain parameters for the given claimed window.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters)
func (device *GPUDevice) SetSwapchainParameters(
	window *Window, swapchainComposition GPUSwapchainComposition, presentMode GPUPresentMode) (bool, error) {

	ret := C.SDL_SetGPUSwapchainParameters(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_Window)(window),
		C.SDL_GPUSwapchainComposition(swapchainComposition),
		C.SDL_GPUPresentMode(presentMode))

	if !ret {
		return false, GetError()
	}
	return true, nil
}

// Obtains the texture format of the swapchain for the given window.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat)
func (device *GPUDevice) GetSwapchainTextureFormat(window *Window) GPUTextureFormat {
	return GPUTextureFormat(C.SDL_GetGPUSwapchainTextureFormat((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window)))
}

// Acquire a texture to use in presentation.
// (https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture)
func (commandBuffer *GPUCommandBuffer) AcquireSwapchainTexture(window *Window) (*GPUTexture, uint32, uint32, error) {
	var texture *GPUTexture
	var swapchainTextureWidth uint32
	var swapchainTextureHeight uint32

	ret := C.SDL_AcquireGPUSwapchainTexture(
		(*C.SDL_GPUCommandBuffer)(commandBuffer),
		(*C.SDL_Window)(window),
		(**C.SDL_GPUTexture)(unsafe.Pointer(texture)),
		(*C.Uint32)(unsafe.Pointer(&swapchainTextureWidth)),
		(*C.Uint32)(unsafe.Pointer(&swapchainTextureHeight)))
	if !ret {
		return nil, 0, 0, GetError()
	}
	return texture, swapchainTextureWidth, swapchainTextureHeight, nil
}

// Submits a command buffer so its commands can be processed on the GPU.
// (https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer)
func (commandBuffer *GPUCommandBuffer) Submit() error {
	ret := C.SDL_SubmitGPUCommandBuffer((*C.SDL_GPUCommandBuffer)(commandBuffer))
	if !ret {
		return GetError()
	}
	return nil
}

// Submits a command buffer so its commands can be processed on the GPU, and
// acquires a fence associated with the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence)
func (commandBuffer *GPUCommandBuffer) SubmitAndAcquireFence() (*GPUFence, error) {
	ret := C.SDL_SubmitGPUCommandBufferAndAcquireFence((*C.SDL_GPUCommandBuffer)(commandBuffer))
	if ret == nil {
		return nil, GetError()
	}
	return (*GPUFence)(ret), nil
}

// Cancels a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer)
func (commandBuffer *GPUCommandBuffer) Cancel() error {
	ret := C.SDL_CancelGPUCommandBuffer((*C.SDL_GPUCommandBuffer)(commandBuffer))
	if !ret {
		return GetError()
	}
	return nil
}

// Blocks the thread until the GPU is completely idle.
// (https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle)
func (device *GPUDevice) WaitForIdle() error {
	ret := C.SDL_WaitForGPUIdle((*C.SDL_GPUDevice)(device))
	if !ret {
		return GetError()
	}
	return nil
}

// TODO: is this right? pointer pointer?
// Blocks the thread until the given fences are signaled.
// (https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences)
func (device *GPUDevice) WaitForFences(waitAll bool, fences []*GPUFence) error {
	data := (*C.SDL_GPUFence)(unsafe.Pointer(&fences))
	dataLen := C.Uint32(len(fences))
	ret := C.SDL_WaitForGPUFences(
		(*C.SDL_GPUDevice)(device),
		C.bool(waitAll),
		&data,
		dataLen)
	if !ret {
		return GetError()
	}
	return nil
}

// Checks the status of a fence.
// (https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence)
func (device *GPUDevice) QueryFence(fence *GPUFence) bool {
	return bool(C.SDL_QueryGPUFence(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_GPUFence)(fence)))
}

// Releases a fence obtained from SDL_SubmitGPUCommandBufferAndAcquireFence.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence)
func (device *GPUDevice) ReleaseFence(fence *GPUFence) {
	C.SDL_ReleaseGPUFence(
		(*C.SDL_GPUDevice)(device),
		(*C.SDL_GPUFence)(fence))
}

// Format Info

// Obtains the texel block size for a texture format.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize)
func (format GPUTextureFormat) TexelBlockSize() uint32 {
	return uint32(C.SDL_GPUTextureFormatTexelBlockSize(C.SDL_GPUTextureFormat(format)))
}

// Determines whether a texture format is supported for a given type and
// usage.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat)
func (device *GPUDevice) TextureSupportsFormat(
	format GPUTextureFormat, textureType GPUTextureType, usage GPUTextureUsageFlags) bool {

	return bool(C.SDL_GPUTextureSupportsFormat(
		(*C.SDL_GPUDevice)(device),
		C.SDL_GPUTextureFormat(format),
		C.SDL_GPUTextureType(textureType),
		C.SDL_GPUTextureUsageFlags(usage)))
}

// Determines if a sample count for a texture format is supported.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount)
func (device *GPUDevice) TextureSupportsSampleCount(format GPUTextureFormat, sampleCount GPUSampleCount) bool {
	return bool(C.SDL_GPUTextureSupportsSampleCount(
		(*C.SDL_GPUDevice)(device),
		C.SDL_GPUTextureFormat(format),
		C.SDL_GPUSampleCount(sampleCount)))
}

// Calculate the size in bytes of a texture format with dimensions.
// (https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize)
func (format GPUTextureFormat) CalculateSize(width uint32, height uint32, depthOrLayerCount uint32) uint32 {
	return uint32(C.SDL_CalculateGPUTextureFormatSize(
		C.SDL_GPUTextureFormat(format),
		C.Uint32(width),
		C.Uint32(height),
		C.Uint32(depthOrLayerCount)))
}

// Windows GDK stuff?
//  #ifdef SDL_PLATFORM_GDK

//  /**
//   * Call this to suspend GPU operation on Xbox when you receive the
//   * SDL_EVENT_DID_ENTER_BACKGROUND event.
//   *
//   * Do NOT call any SDL_GPU functions after calling this function! This must
//   * also be called before calling SDL_GDKSuspendComplete.
//   *
//   * \param device a GPU context.
//   *
//   * \since This function is available since SDL 3.1.3.
//   *
//   * \sa SDL_AddEventWatch
//   */
//  extern SDL_DECLSPEC void SDLCALL SDL_GDKSuspendGPU(SDL_GPUDevice *device);

//  /**
//   * Call this to resume GPU operation on Xbox when you receive the
//   * SDL_EVENT_WILL_ENTER_FOREGROUND event.
//   *
//   * When resuming, this function MUST be called before calling any other
//   * SDL_GPU functions.
//   *
//   * \param device a GPU context.
//   *
//   * \since This function is available since SDL 3.1.3.
//   *
//   * \sa SDL_AddEventWatch
//   */
//  extern SDL_DECLSPEC void SDLCALL SDL_GDKResumeGPU(SDL_GPUDevice *device);
