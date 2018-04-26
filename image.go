package gozbar

/*
#include <zbar.h>
#cgo LDFLAGS: -lzbar
*/
import "C"
import (
	"encoding/binary"
	"image"
	"image/draw"
	"runtime"
	"unsafe"
)

/** zbar/conver.c
 * see zbar_fourcc
 */
func ZbarFourcc(a, b, c, d byte) uint32 {
	return binary.LittleEndian.Uint32([]byte{a, b, c, d})
}

/** zbar.h
 * zbar_image_t
 */
type Image struct {
	image *C.zbar_image_t
	gray  *image.Gray
}

/**
 * create gozbar image
 */
func ImageCreate(img image.Image) (obj *Image) {
	obj = new(Image)
	obj.Create(img)
	return obj
}

/** zbar.h
 * see zbar_image_create
 */
func (obj *Image) Create(img image.Image) {
	// zbar.h zbar_image_create
	obj.image = C.zbar_image_create()
	// load image rectangle
	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y
	// draw gray image
	obj.gray = image.NewGray(bounds)
	draw.Draw(obj.gray, bounds, img, image.ZP, draw.Over)
	// set zbar_image_t
	C.zbar_image_set_format(obj.image, C.ulong(ZbarFourcc('Y', '8', '0', '0'))) // Y800 (grayscale)
	C.zbar_image_set_size(obj.image, C.uint(w), C.uint(h))
	C.zbar_image_set_data(obj.image, unsafe.Pointer(&obj.gray.Pix[0]), C.ulong(len(obj.gray.Pix)), nil)
	// set gc func
	runtime.SetFinalizer(obj, (*Image).Destroy)
}

/**
 * load zbar_image_t
 */
func (obj *Image) Image() *C.zbar_image_t {
	return obj.image
}

/**
 * image decode successfully, get first symbol
 */
func (obj *Image) FirstSymbol() *Symbol {
	// first decoded symbol result for an image
	symbol := C.zbar_image_first_symbol(obj.image)
	// create gozbar symbol
	return SymbolCreate(symbol)
}

/**
 * image destroy
 */
func (obj *Image) Destroy() {
	// destroy zbar_image_t
	C.zbar_image_destroy(obj.image)
}
