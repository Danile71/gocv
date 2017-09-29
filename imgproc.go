package opencv3

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"
import "unsafe"

// CvtColor converts an image from one color space to another
func CvtColor(src Mat, dst Mat, code int) {
	C.CvtColor(src.p, dst.p, C.int(code))
}

// GaussianBlur blurs an image using a Gaussian filter.
func GaussianBlur(src Mat, dst Mat, ksize Size, sigmaX float64, sigmaY float64, borderType int) {
	pSize := C.struct_Size{
		height: C.int(ksize.Height),
		width:  C.int(ksize.Width),
	}

	C.GaussianBlur(src.p, dst.p, pSize, C.double(sigmaX), C.double(sigmaY), C.int(borderType))
}

// Rectangle draws a rectangle using to target image Mat.
func Rectangle(img Mat, r Rect, c Scalar) {
	cRect := C.struct_Rect{
		x:      C.int(r.X),
		y:      C.int(r.Y),
		width:  C.int(r.Width),
		height: C.int(r.Height),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.Val1),
		val2: C.double(c.Val2),
		val3: C.double(c.Val3),
		val4: C.double(c.Val4),
	}

	C.Rectangle(img.p, cRect, sColor)
}

// Based on the enum HersheyFonts
// Only a subset of Hershey fonts
// <http://sources.isc.org/utils/misc/hershey-font.txt> are supported
const (
	FontHersheySimplex       = 0  //!< normal size sans-serif font
	FontHersheyPlain         = 1  //!< small size sans-serif font
	FontHersheyDuplex        = 2  //!< normal size sans-serif font (more complex than FontHersheySIMPLEX)
	FontHersheyComplex       = 3  //!< normal size serif font
	FontHersheyTriplex       = 4  //!< normal size serif font (more complex than FontHersheyCOMPLEX)
	FontHersheyComplexSmall  = 5  //!< smaller version of FontHersheyCOMPLEX
	FontHersheyScriptSimplex = 6  //!< hand-writing style font
	FontHersheyScriptComplex = 7  //!< more complex variant of FontHersheySCRIPT_SIMPLEX
	FontItalic               = 16 //!< flag for italic font
)

// GetTextSize returns the size required to draw text using a specific font face,
// scale, and thickness.
func GetTextSize(text string, fontFace int, fontScale float64, thickness int) Size {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sz := C.GetTextSize(cText, C.int(fontFace), C.double(fontScale), C.int(thickness))
	return Size{Width: int(sz.width), Height: int(sz.height)}
}

// PutText renders the specified text string in the image.
func PutText(img Mat, text string, org Point, fontFace int, fontScale float64, color Scalar, thickness int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(color.Val1),
		val2: C.double(color.Val2),
		val3: C.double(color.Val3),
		val4: C.double(color.Val4),
	}

	C.PutText(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness))
	return
}
