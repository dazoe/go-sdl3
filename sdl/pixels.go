package sdl

//#include <SDL3/SDL_pixels.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// A fully opaque 8-bit alpha value
const SDL_ALPHA_OPAQUE = 255

// A fully transparent 8-bit alpha value
const SDL_ALPHA_TRANSPARENT = 0

// Pixel type.
type PixelType C.SDL_PixelType

const (
	PIXELTYPE_UNKNOWN  = PixelType(C.SDL_PIXELTYPE_UNKNOWN)
	PIXELTYPE_INDEX1   = PixelType(C.SDL_PIXELTYPE_INDEX1)
	PIXELTYPE_INDEX4   = PixelType(C.SDL_PIXELTYPE_INDEX4)
	PIXELTYPE_INDEX8   = PixelType(C.SDL_PIXELTYPE_INDEX8)
	PIXELTYPE_PACKED8  = PixelType(C.SDL_PIXELTYPE_PACKED8)
	PIXELTYPE_PACKED16 = PixelType(C.SDL_PIXELTYPE_PACKED16)
	PIXELTYPE_PACKED32 = PixelType(C.SDL_PIXELTYPE_PACKED32)
	PIXELTYPE_ARRAYU8  = PixelType(C.SDL_PIXELTYPE_ARRAYU8)
	PIXELTYPE_ARRAYU16 = PixelType(C.SDL_PIXELTYPE_ARRAYU16)
	PIXELTYPE_ARRAYU32 = PixelType(C.SDL_PIXELTYPE_ARRAYU32)
	PIXELTYPE_ARRAYF16 = PixelType(C.SDL_PIXELTYPE_ARRAYF16)
	PIXELTYPE_ARRAYF32 = PixelType(C.SDL_PIXELTYPE_ARRAYF32)
	PIXELTYPE_INDEX2   = PixelType(C.SDL_PIXELTYPE_INDEX2) /* appended at the end for compatibility with sdl2-compat:  */
)

// Bitmap pixel order, high bit -> low bit.
type BitmapOrder C.SDL_BitmapOrder

const (
	BITMAPORDER_NONE = BitmapOrder(C.SDL_BITMAPORDER_NONE)
	BITMAPORDER_4321 = BitmapOrder(C.SDL_BITMAPORDER_4321)
	BITMAPORDER_1234 = BitmapOrder(C.SDL_BITMAPORDER_1234)
)

// Packed component order, high bit -> low bit.
type PackedOrder C.SDL_PackedOrder

const (
	PACKEDORDER_NONE = PackedOrder(C.SDL_PACKEDORDER_NONE)
	PACKEDORDER_XRGB = PackedOrder(C.SDL_PACKEDORDER_XRGB)
	PACKEDORDER_RGBX = PackedOrder(C.SDL_PACKEDORDER_RGBX)
	PACKEDORDER_ARGB = PackedOrder(C.SDL_PACKEDORDER_ARGB)
	PACKEDORDER_RGBA = PackedOrder(C.SDL_PACKEDORDER_RGBA)
	PACKEDORDER_XBGR = PackedOrder(C.SDL_PACKEDORDER_XBGR)
	PACKEDORDER_BGRX = PackedOrder(C.SDL_PACKEDORDER_BGRX)
	PACKEDORDER_ABGR = PackedOrder(C.SDL_PACKEDORDER_ABGR)
	PACKEDORDER_BGRA = PackedOrder(C.SDL_PACKEDORDER_BGRA)
)

// Array component order, low byte -> high byte.
type ArrayOrder C.SDL_ArrayOrder

const (
	ARRAYORDER_NONE = ArrayOrder(C.SDL_ARRAYORDER_NONE)
	ARRAYORDER_RGB  = ArrayOrder(C.SDL_ARRAYORDER_RGB)
	ARRAYORDER_RGBA = ArrayOrder(C.SDL_ARRAYORDER_RGBA)
	ARRAYORDER_ARGB = ArrayOrder(C.SDL_ARRAYORDER_ARGB)
	ARRAYORDER_BGR  = ArrayOrder(C.SDL_ARRAYORDER_BGR)
	ARRAYORDER_BGRA = ArrayOrder(C.SDL_ARRAYORDER_BGRA)
	ARRAYORDER_ABGR = ArrayOrder(C.SDL_ARRAYORDER_ABGR)
)

// Packed component layout.
type PackedLayout C.SDL_PackedLayout

const (
	PACKEDLAYOUT_NONE    = PackedLayout(C.SDL_PACKEDLAYOUT_NONE)
	PACKEDLAYOUT_332     = PackedLayout(C.SDL_PACKEDLAYOUT_332)
	PACKEDLAYOUT_4444    = PackedLayout(C.SDL_PACKEDLAYOUT_4444)
	PACKEDLAYOUT_1555    = PackedLayout(C.SDL_PACKEDLAYOUT_1555)
	PACKEDLAYOUT_5551    = PackedLayout(C.SDL_PACKEDLAYOUT_5551)
	PACKEDLAYOUT_565     = PackedLayout(C.SDL_PACKEDLAYOUT_565)
	PACKEDLAYOUT_8888    = PackedLayout(C.SDL_PACKEDLAYOUT_8888)
	PACKEDLAYOUT_2101010 = PackedLayout(C.SDL_PACKEDLAYOUT_2101010)
	PACKEDLAYOUT_1010102 = PackedLayout(C.SDL_PACKEDLAYOUT_1010102)
)

//TODO: incomplete: macros skipped

// Pixel format
// (https://wiki.libsdl.org/SDL3/SDL_PixelFormat)
type PixelFormat C.SDL_PixelFormat

func (format PixelFormat) c() C.SDL_PixelFormat {
	return (C.SDL_PixelFormat)(format)
}

const (
	PIXELFORMAT_UNKNOWN       = PixelFormat(C.SDL_PIXELFORMAT_UNKNOWN)
	PIXELFORMAT_INDEX1LSB     = PixelFormat(C.SDL_PIXELFORMAT_INDEX1LSB)
	PIXELFORMAT_INDEX1MSB     = PixelFormat(C.SDL_PIXELFORMAT_INDEX1MSB)
	PIXELFORMAT_INDEX2LSB     = PixelFormat(C.SDL_PIXELFORMAT_INDEX2LSB)
	PIXELFORMAT_INDEX2MSB     = PixelFormat(C.SDL_PIXELFORMAT_INDEX2MSB)
	PIXELFORMAT_INDEX4LSB     = PixelFormat(C.SDL_PIXELFORMAT_INDEX4LSB)
	PIXELFORMAT_INDEX4MSB     = PixelFormat(C.SDL_PIXELFORMAT_INDEX4MSB)
	PIXELFORMAT_INDEX8        = PixelFormat(C.SDL_PIXELFORMAT_INDEX8)
	PIXELFORMAT_RGB332        = PixelFormat(C.SDL_PIXELFORMAT_RGB332)
	PIXELFORMAT_XRGB4444      = PixelFormat(C.SDL_PIXELFORMAT_XRGB4444)
	PIXELFORMAT_XBGR4444      = PixelFormat(C.SDL_PIXELFORMAT_XBGR4444)
	PIXELFORMAT_XRGB1555      = PixelFormat(C.SDL_PIXELFORMAT_XRGB1555)
	PIXELFORMAT_XBGR1555      = PixelFormat(C.SDL_PIXELFORMAT_XBGR1555)
	PIXELFORMAT_ARGB4444      = PixelFormat(C.SDL_PIXELFORMAT_ARGB4444)
	PIXELFORMAT_RGBA4444      = PixelFormat(C.SDL_PIXELFORMAT_RGBA4444)
	PIXELFORMAT_ABGR4444      = PixelFormat(C.SDL_PIXELFORMAT_ABGR4444)
	PIXELFORMAT_BGRA4444      = PixelFormat(C.SDL_PIXELFORMAT_BGRA4444)
	PIXELFORMAT_ARGB1555      = PixelFormat(C.SDL_PIXELFORMAT_ARGB1555)
	PIXELFORMAT_RGBA5551      = PixelFormat(C.SDL_PIXELFORMAT_RGBA5551)
	PIXELFORMAT_ABGR1555      = PixelFormat(C.SDL_PIXELFORMAT_ABGR1555)
	PIXELFORMAT_BGRA5551      = PixelFormat(C.SDL_PIXELFORMAT_BGRA5551)
	PIXELFORMAT_RGB565        = PixelFormat(C.SDL_PIXELFORMAT_RGB565)
	PIXELFORMAT_BGR565        = PixelFormat(C.SDL_PIXELFORMAT_BGR565)
	PIXELFORMAT_RGB24         = PixelFormat(C.SDL_PIXELFORMAT_RGB24)
	PIXELFORMAT_BGR24         = PixelFormat(C.SDL_PIXELFORMAT_BGR24)
	PIXELFORMAT_XRGB8888      = PixelFormat(C.SDL_PIXELFORMAT_XRGB8888)
	PIXELFORMAT_RGBX8888      = PixelFormat(C.SDL_PIXELFORMAT_RGBX8888)
	PIXELFORMAT_XBGR8888      = PixelFormat(C.SDL_PIXELFORMAT_XBGR8888)
	PIXELFORMAT_BGRX8888      = PixelFormat(C.SDL_PIXELFORMAT_BGRX8888)
	PIXELFORMAT_ARGB8888      = PixelFormat(C.SDL_PIXELFORMAT_ARGB8888)
	PIXELFORMAT_RGBA8888      = PixelFormat(C.SDL_PIXELFORMAT_RGBA8888)
	PIXELFORMAT_ABGR8888      = PixelFormat(C.SDL_PIXELFORMAT_ABGR8888)
	PIXELFORMAT_BGRA8888      = PixelFormat(C.SDL_PIXELFORMAT_BGRA8888)
	PIXELFORMAT_XRGB2101010   = PixelFormat(C.SDL_PIXELFORMAT_XRGB2101010)
	PIXELFORMAT_XBGR2101010   = PixelFormat(C.SDL_PIXELFORMAT_XBGR2101010)
	PIXELFORMAT_ARGB2101010   = PixelFormat(C.SDL_PIXELFORMAT_ARGB2101010)
	PIXELFORMAT_ABGR2101010   = PixelFormat(C.SDL_PIXELFORMAT_ABGR2101010)
	PIXELFORMAT_RGB48         = PixelFormat(C.SDL_PIXELFORMAT_RGB48)
	PIXELFORMAT_BGR48         = PixelFormat(C.SDL_PIXELFORMAT_BGR48)
	PIXELFORMAT_RGBA64        = PixelFormat(C.SDL_PIXELFORMAT_RGBA64)
	PIXELFORMAT_ARGB64        = PixelFormat(C.SDL_PIXELFORMAT_ARGB64)
	PIXELFORMAT_BGRA64        = PixelFormat(C.SDL_PIXELFORMAT_BGRA64)
	PIXELFORMAT_ABGR64        = PixelFormat(C.SDL_PIXELFORMAT_ABGR64)
	PIXELFORMAT_RGB48_FLOAT   = PixelFormat(C.SDL_PIXELFORMAT_RGB48_FLOAT)
	PIXELFORMAT_BGR48_FLOAT   = PixelFormat(C.SDL_PIXELFORMAT_BGR48_FLOAT)
	PIXELFORMAT_RGBA64_FLOAT  = PixelFormat(C.SDL_PIXELFORMAT_RGBA64_FLOAT)
	PIXELFORMAT_ARGB64_FLOAT  = PixelFormat(C.SDL_PIXELFORMAT_ARGB64_FLOAT)
	PIXELFORMAT_BGRA64_FLOAT  = PixelFormat(C.SDL_PIXELFORMAT_BGRA64_FLOAT)
	PIXELFORMAT_ABGR64_FLOAT  = PixelFormat(C.SDL_PIXELFORMAT_ABGR64_FLOAT)
	PIXELFORMAT_RGB96_FLOAT   = PixelFormat(C.SDL_PIXELFORMAT_RGB96_FLOAT)
	PIXELFORMAT_BGR96_FLOAT   = PixelFormat(C.SDL_PIXELFORMAT_BGR96_FLOAT)
	PIXELFORMAT_RGBA128_FLOAT = PixelFormat(C.SDL_PIXELFORMAT_RGBA128_FLOAT)
	PIXELFORMAT_ARGB128_FLOAT = PixelFormat(C.SDL_PIXELFORMAT_ARGB128_FLOAT)
	PIXELFORMAT_BGRA128_FLOAT = PixelFormat(C.SDL_PIXELFORMAT_BGRA128_FLOAT)
	PIXELFORMAT_ABGR128_FLOAT = PixelFormat(C.SDL_PIXELFORMAT_ABGR128_FLOAT)
	PIXELFORMAT_YV12          = PixelFormat(C.SDL_PIXELFORMAT_YV12)         /**< Planar mode: Y + V + U  (3 planes) */
	PIXELFORMAT_IYUV          = PixelFormat(C.SDL_PIXELFORMAT_IYUV)         /**< Planar mode: Y + U + V  (3 planes) */
	PIXELFORMAT_YUY2          = PixelFormat(C.SDL_PIXELFORMAT_YUY2)         /**< Packed mode: Y0+U0+Y1+V0 (1 plane) */
	PIXELFORMAT_UYVY          = PixelFormat(C.SDL_PIXELFORMAT_UYVY)         /**< Packed mode: U0+Y0+V0+Y1 (1 plane) */
	PIXELFORMAT_YVYU          = PixelFormat(C.SDL_PIXELFORMAT_YVYU)         /**< Packed mode: Y0+V0+Y1+U0 (1 plane) */
	PIXELFORMAT_NV12          = PixelFormat(C.SDL_PIXELFORMAT_NV12)         /**< Planar mode: Y + U/V interleaved  (2 planes) */
	PIXELFORMAT_NV21          = PixelFormat(C.SDL_PIXELFORMAT_NV21)         /**< Planar mode: Y + V/U interleaved  (2 planes) */
	PIXELFORMAT_P010          = PixelFormat(C.SDL_PIXELFORMAT_P010)         /**< Planar mode: Y + U/V interleaved  (2 planes) */
	PIXELFORMAT_EXTERNAL_OES  = PixelFormat(C.SDL_PIXELFORMAT_EXTERNAL_OES) /**< Android video texture format */
)

// Aliases for RGBA byte arrays of color data, for the current platform
const (
	PIXELFORMAT_RGBA32 = PixelFormat(C.SDL_PIXELFORMAT_RGBA32)
	PIXELFORMAT_ARGB32 = PixelFormat(C.SDL_PIXELFORMAT_ARGB32)
	PIXELFORMAT_BGRA32 = PixelFormat(C.SDL_PIXELFORMAT_BGRA32)
	PIXELFORMAT_ABGR32 = PixelFormat(C.SDL_PIXELFORMAT_ABGR32)
	PIXELFORMAT_RGBX32 = PixelFormat(C.SDL_PIXELFORMAT_RGBX32)
	PIXELFORMAT_XRGB32 = PixelFormat(C.SDL_PIXELFORMAT_XRGB32)
	PIXELFORMAT_BGRX32 = PixelFormat(C.SDL_PIXELFORMAT_BGRX32)
	PIXELFORMAT_XBGR32 = PixelFormat(C.SDL_PIXELFORMAT_XBGR32)
)

// Colorspace color type
// (https://wiki.libsdl.org/SDL3/SDL_ColorType)
type ColorType C.SDL_ColorType

const (
	COLOR_TYPE_UNKNOWN = ColorType(C.SDL_COLOR_TYPE_UNKNOWN)
	COLOR_TYPE_RGB     = ColorType(C.SDL_COLOR_TYPE_RGB)
	COLOR_TYPE_YCBCR   = ColorType(C.SDL_COLOR_TYPE_YCBCR)
)

// Colorspace color range, as described by
// https://www.itu.int/rec/R-REC-BT.2100-2-201807-I/en
// (https://wiki.libsdl.org/SDL3/SDL_ColorRange)
type ColorRange C.SDL_ColorRange

const (
	COLOR_RANGE_UNKNOWN = ColorRange(C.SDL_COLOR_RANGE_UNKNOWN)
	COLOR_RANGE_LIMITED = ColorRange(C.SDL_COLOR_RANGE_LIMITED) /**< Narrow range, e.g. 16-235 for 8-bit RGB and luma, and 16-240 for 8-bit chroma */
	COLOR_RANGE_FULL    = ColorRange(C.SDL_COLOR_RANGE_FULL)    /**< Full range, e.g. 0-255 for 8-bit RGB and luma, and 1-255 for 8-bit chroma */
)

// Colorspace color primaries, as described by
// https://www.itu.int/rec/T-REC-H.273-201612-S/en
// (https://wiki.libsdl.org/SDL3/SDL_ColorPrimaries)
type ColorPrimaries C.SDL_ColorPrimaries

const (
	COLOR_PRIMARIES_UNKNOWN      = ColorPrimaries(C.SDL_COLOR_PRIMARIES_UNKNOWN)
	COLOR_PRIMARIES_BT709        = ColorPrimaries(C.SDL_COLOR_PRIMARIES_BT709) /**< ITU-R BT.709-6 */
	COLOR_PRIMARIES_UNSPECIFIED  = ColorPrimaries(C.SDL_COLOR_PRIMARIES_UNSPECIFIED)
	COLOR_PRIMARIES_BT470M       = ColorPrimaries(C.SDL_COLOR_PRIMARIES_BT470M)       /**< ITU-R BT.470-6 System M */
	COLOR_PRIMARIES_BT470BG      = ColorPrimaries(C.SDL_COLOR_PRIMARIES_BT470BG)      /**< ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625 */
	COLOR_PRIMARIES_BT601        = ColorPrimaries(C.SDL_COLOR_PRIMARIES_BT601)        /**< ITU-R BT.601-7 525, SMPTE 170M */
	COLOR_PRIMARIES_SMPTE240     = ColorPrimaries(C.SDL_COLOR_PRIMARIES_SMPTE240)     /**< SMPTE 240M, functionally the same as SDL_COLOR_PRIMARIES_BT601 */
	COLOR_PRIMARIES_GENERIC_FILM = ColorPrimaries(C.SDL_COLOR_PRIMARIES_GENERIC_FILM) /**< Generic film (color filters using Illuminant C) */
	COLOR_PRIMARIES_BT2020       = ColorPrimaries(C.SDL_COLOR_PRIMARIES_BT2020)       /**< ITU-R BT.2020-2 / ITU-R BT.2100-0 */
	COLOR_PRIMARIES_XYZ          = ColorPrimaries(C.SDL_COLOR_PRIMARIES_XYZ)          /**< SMPTE ST 428-1 */
	COLOR_PRIMARIES_SMPTE431     = ColorPrimaries(C.SDL_COLOR_PRIMARIES_SMPTE431)     /**< SMPTE RP 431-2 */
	COLOR_PRIMARIES_SMPTE432     = ColorPrimaries(C.SDL_COLOR_PRIMARIES_SMPTE432)     /**< SMPTE EG 432-1 / DCI P3 */
	COLOR_PRIMARIES_EBU3213      = ColorPrimaries(C.SDL_COLOR_PRIMARIES_EBU3213)      /**< EBU Tech. 3213-E */
	COLOR_PRIMARIES_CUSTOM       = ColorPrimaries(C.SDL_COLOR_PRIMARIES_CUSTOM)
)

// Colorspace transfer characteristics.
// (https://wiki.libsdl.org/SDL3/SDL_TransferCharacteristics)
type TransferCharacteristics C.SDL_TransferCharacteristics

const (
	TRANSFER_CHARACTERISTICS_UNKNOWN       = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_UNKNOWN)
	TRANSFER_CHARACTERISTICS_BT709         = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_BT709) /**< Rec. ITU-R BT.709-6 / ITU-R BT1361 */
	TRANSFER_CHARACTERISTICS_UNSPECIFIED   = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_UNSPECIFIED)
	TRANSFER_CHARACTERISTICS_GAMMA22       = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_GAMMA22)  /**< ITU-R BT.470-6 System M / ITU-R BT1700 625 PAL & SECAM */
	TRANSFER_CHARACTERISTICS_GAMMA28       = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_GAMMA28)  /**< ITU-R BT.470-6 System B, G */
	TRANSFER_CHARACTERISTICS_BT601         = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_BT601)    /**< SMPTE ST 170M / ITU-R BT.601-7 525 or 625 */
	TRANSFER_CHARACTERISTICS_SMPTE240      = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_SMPTE240) /**< SMPTE ST 240M */
	TRANSFER_CHARACTERISTICS_LINEAR        = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_LINEAR)
	TRANSFER_CHARACTERISTICS_LOG100        = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_LOG100)
	TRANSFER_CHARACTERISTICS_LOG100_SQRT10 = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_LOG100_SQRT10)
	TRANSFER_CHARACTERISTICS_IEC61966      = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_IEC61966)     /**< IEC 61966-2-4 */
	TRANSFER_CHARACTERISTICS_BT1361        = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_BT1361)       /**< ITU-R BT1361 Extended Colour Gamut */
	TRANSFER_CHARACTERISTICS_SRGB          = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_SRGB)         /**< IEC 61966-2-1 (sRGB or sYCC) */
	TRANSFER_CHARACTERISTICS_BT2020_10BIT  = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_BT2020_10BIT) /**< ITU-R BT2020 for 10-bit system */
	TRANSFER_CHARACTERISTICS_BT2020_12BIT  = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_BT2020_12BIT) /**< ITU-R BT2020 for 12-bit system */
	TRANSFER_CHARACTERISTICS_PQ            = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_PQ)           /**< SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems */
	TRANSFER_CHARACTERISTICS_SMPTE428      = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_SMPTE428)     /**< SMPTE ST 428-1 */
	TRANSFER_CHARACTERISTICS_HLG           = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_HLG)          /**< ARIB STD-B67, known as "hybrid log-gamma" (HLG) */
	TRANSFER_CHARACTERISTICS_CUSTOM        = TransferCharacteristics(C.SDL_TRANSFER_CHARACTERISTICS_CUSTOM)
)

// Colorspace matrix coefficients.
// (https://wiki.libsdl.org/SDL3/SDL_MatrixCoefficients)
type MatrixCoefficients C.SDL_MatrixCoefficients

const (
	MATRIX_COEFFICIENTS_IDENTITY           = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_IDENTITY)
	MATRIX_COEFFICIENTS_BT709              = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_BT709) /**< ITU-R BT.709-6 */
	MATRIX_COEFFICIENTS_UNSPECIFIED        = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_UNSPECIFIED)
	MATRIX_COEFFICIENTS_FCC                = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_FCC)      /**< US FCC Title 47 */
	MATRIX_COEFFICIENTS_BT470BG            = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_BT470BG)  /**< ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625, functionally the same as SDL_MATRIX_COEFFICIENTS_BT601 */
	MATRIX_COEFFICIENTS_BT601              = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_BT601)    /**< ITU-R BT.601-7 525 */
	MATRIX_COEFFICIENTS_SMPTE240           = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_SMPTE240) /**< SMPTE 240M */
	MATRIX_COEFFICIENTS_YCGCO              = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_YCGCO)
	MATRIX_COEFFICIENTS_BT2020_NCL         = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_BT2020_NCL) /**< ITU-R BT.2020-2 non-constant luminance */
	MATRIX_COEFFICIENTS_BT2020_CL          = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_BT2020_CL)  /**< ITU-R BT.2020-2 constant luminance */
	MATRIX_COEFFICIENTS_SMPTE2085          = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_SMPTE2085)  /**< SMPTE ST 2085 */
	MATRIX_COEFFICIENTS_CHROMA_DERIVED_NCL = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_CHROMA_DERIVED_NCL)
	MATRIX_COEFFICIENTS_CHROMA_DERIVED_CL  = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_CHROMA_DERIVED_CL)
	MATRIX_COEFFICIENTS_ICTCP              = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_ICTCP) /**< ITU-R BT.2100-0 ICTCP */
	MATRIX_COEFFICIENTS_CUSTOM             = MatrixCoefficients(C.SDL_MATRIX_COEFFICIENTS_CUSTOM)
)

// Colorspace chroma sample location.
// (https://wiki.libsdl.org/SDL3/SDL_ChromaLocation)
type ChromaLocation C.SDL_ChromaLocation

const (
	CHROMA_LOCATION_NONE    = ChromaLocation(C.SDL_CHROMA_LOCATION_NONE)    /**< RGB, no chroma sampling */
	CHROMA_LOCATION_LEFT    = ChromaLocation(C.SDL_CHROMA_LOCATION_LEFT)    /**< In MPEG-2, MPEG-4, and AVC, Cb and Cr are taken on midpoint of the left-edge of the 2x2 square. In other words, they have the same horizontal location as the top-left pixel, but is shifted one-half pixel down vertically. */
	CHROMA_LOCATION_CENTER  = ChromaLocation(C.SDL_CHROMA_LOCATION_CENTER)  /**< In JPEG/JFIF, H.261, and MPEG-1, Cb and Cr are taken at the center of the 2x2 square. In other words, they are offset one-half pixel to the right and one-half pixel down compared to the top-left pixel. */
	CHROMA_LOCATION_TOPLEFT = ChromaLocation(C.SDL_CHROMA_LOCATION_TOPLEFT) /**< In HEVC for BT.2020 and BT.2100 content (in particular on Blu-rays), Cb and Cr are sampled at the same location as the group's top-left Y pixel ("co-sited", "co-located"). */
)

//TODO: incomplete: macros skipped

// Colorspace definitions.
// (https://wiki.libsdl.org/SDL3/SDL_Colorspace)
type Colorspace C.SDL_Colorspace

const (
	COLORSPACE_UNKNOWN = Colorspace(C.SDL_COLORSPACE_UNKNOWN)
	/* sRGB is a gamma corrected colorspace, and the default colorspace for SDL rendering and 8-bit RGB surfaces */
	COLORSPACE_SRGB = Colorspace(C.SDL_COLORSPACE_SRGB) /**< Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G22_NONE_P709 */
	/* This is a linear colorspace and the default colorspace for floating point surfaces. On Windows this is the scRGB colorspace, and on Apple platforms this is kCGColorSpaceExtendedLinearSRGB for EDR content */
	COLORSPACE_SRGB_LINEAR = Colorspace(C.SDL_COLORSPACE_SRGB_LINEAR) /**< Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G10_NONE_P709  */
	/* HDR10 is a non-linear HDR colorspace and the default colorspace for 10-bit surfaces */
	COLORSPACE_HDR10          = Colorspace(C.SDL_COLORSPACE_HDR10)          /**< Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G2084_NONE_P2020  */
	COLORSPACE_JPEG           = Colorspace(C.SDL_COLORSPACE_JPEG)           /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_NONE_P709_X601 */
	COLORSPACE_BT601_LIMITED  = Colorspace(C.SDL_COLORSPACE_BT601_LIMITED)  /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601 */
	COLORSPACE_BT601_FULL     = Colorspace(C.SDL_COLORSPACE_BT601_FULL)     /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601 */
	COLORSPACE_BT709_LIMITED  = Colorspace(C.SDL_COLORSPACE_BT709_LIMITED)  /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709 */
	COLORSPACE_BT709_FULL     = Colorspace(C.SDL_COLORSPACE_BT709_FULL)     /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709 */
	COLORSPACE_BT2020_LIMITED = Colorspace(C.SDL_COLORSPACE_BT2020_LIMITED) /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P2020 */
	COLORSPACE_BT2020_FULL    = Colorspace(C.SDL_COLORSPACE_BT2020_FULL)    /**< Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_LEFT_P2020 */
)

/* The default colorspace for RGB surfaces if no colorspace is specified */
const SDL_COLORSPACE_RGB_DEFAULT = C.SDL_COLORSPACE_SRGB

/* The default colorspace for YUV surfaces if no colorspace is specified */
const SDL_COLORSPACE_YUV_DEFAULT = C.SDL_COLORSPACE_JPEG

// A structure that represents a color as RGBA components.
// (https://wiki.libsdl.org/SDL3/SDL_Color)
type Color struct {
	R, G, B, A uint8
}

func (c *Color) cptr() *C.SDL_Color {
	return (*C.SDL_Color)(unsafe.Pointer(c))
}

// implement go's image.Color interface
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

// The bits of this structure can be directly reinterpreted as a float-packed
// (https://wiki.libsdl.org/SDL3/SDL_FColor)
type FColor struct {
	R, G, B, A float32
}

func (c *FColor) cptr() *C.SDL_FColor {
	return (*C.SDL_FColor)(unsafe.Pointer(c))
}

// A set of indexed colors representing a palette.
// (https://wiki.libsdl.org/SDL3/SDL_Palette)
type Palette struct {
	NColors  int32  /**< number of elements in `colors`. */
	Colors   *Color /**< an array of colors, `ncolors` long. */
	version  uint32 /**< internal use only, do not touch. */
	refcount int32  /**< internal use only, do not touch. */
}

func (palette *Palette) cptr() *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(palette))
}

// Details about the format of a pixel.
// (https://wiki.libsdl.org/SDL3/SDL_PixelFormatDetails)
type PixelFormatDetails struct {
	Format                         PixelFormat
	BitsPerPixel                   uint8
	BytesPerPixel                  uint8
	_                              [2]uint8 //padding
	RMask, GMask, BMask, AMask     uint32
	RBits, GBits, BBits, ABits     uint8
	RShift, GShift, BShift, AShift uint8
}

func (pfd *PixelFormatDetails) cptr() *C.SDL_PixelFormatDetails {
	return (*C.SDL_PixelFormatDetails)(unsafe.Pointer(pfd))
}

// Get the human readable name of a pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName)
func (format PixelFormat) GetName() string {
	return C.GoString(C.SDL_GetPixelFormatName(C.SDL_PixelFormat(format)))
}
func (format PixelFormat) String() string {
	name := format.GetName()
	if name == "" {
		return fmt.Sprintf("0x%x", format)
	}
	return name
}

// Convert one of the enumerated pixel formats to a bpp value and RGBA masks.
// (https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat)
func (format PixelFormat) GetMasks() (bpp int32, RMask, GMask, BMask, AMask uint32, err error) {
	ret := C.SDL_GetMasksForPixelFormat(C.SDL_PixelFormat(format),
		(*C.int)(&bpp), (*C.Uint32)(&RMask), (*C.Uint32)(&GMask), (*C.Uint32)(&BMask), (*C.Uint32)(&AMask))
	if !ret {
		err = GetError()
		return
	}
	return
}

// Convert a bpp value and RGBA masks to an enumerated pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks)
func GetPixelFormatForMasks(bpp int, Rmask, Gmask, Bmask, Amask uint32) PixelFormat {
	return PixelFormat(C.SDL_GetPixelFormatForMasks(
		C.int(bpp), C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask)))
}

// Create an SDL_PixelFormatDetails structure corresponding to a pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails)
func (format PixelFormat) GetDetails() (*PixelFormatDetails, error) {
	ret := (*PixelFormatDetails)(unsafe.Pointer(C.SDL_GetPixelFormatDetails(C.SDL_PixelFormat(format))))
	if ret == nil {
		return nil, GetError()
	}
	return ret, nil
}

// Create a palette structure with the specified number of color entries.
// (https://wiki.libsdl.org/SDL3/SDL_CreatePalette)
func CreatePalette(ncolors int) (*Palette, error) {
	ret := (*Palette)(unsafe.Pointer(C.SDL_CreatePalette(C.int(ncolors))))
	if ret == nil {
		return nil, GetError()
	}
	return ret, nil
}

// Set a range of colors in a palette.
// (https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors)
func (palette *Palette) SetColors(colors []Color) error {
	ret := C.SDL_SetPaletteColors(
		palette.cptr(),
		colors[0].cptr(),
		0, C.int(len(colors)))
	if !ret {
		return GetError()
	}
	return nil
}

// Free a palette created with SDL_CreatePalette().
// (https://wiki.libsdl.org/SDL3/SDL_DestroyPalette)
func (palette *Palette) Destroy() {
	C.SDL_DestroyPalette(palette.cptr())
}

// Map an RGB triple to an opaque pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGB)
func (format *PixelFormatDetails) MapRGB(palette *Palette, r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB(
		format.cptr(),
		palette.cptr(),
		C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// Map an RGBA quadruple to a pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGBA)
func (format *PixelFormatDetails) MapRGBA(palette *Palette, r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA(
		format.cptr(),
		palette.cptr(),
		C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// Get RGB values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_GetRGB)
func (format *PixelFormatDetails) GetRGB(pixel uint32, palette *Palette) (r, g, b uint8) {
	C.SDL_GetRGB(C.Uint32(pixel),
		format.cptr(),
		palette.cptr(),
		(*C.Uint8)(&r), (*C.Uint8)(&g),
		(*C.Uint8)(&b))
	return
}

// Get RGBA values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_GetRGB)
func (format *PixelFormatDetails) GetRGBA(pixel uint32, palette *Palette) (r, g, b, a uint8) {
	C.SDL_GetRGBA(C.Uint32(pixel),
		format.cptr(),
		palette.cptr(),
		(*C.Uint8)(&r), (*C.Uint8)(&g),
		(*C.Uint8)(&b), (*C.Uint8)(&a))
	return
}
