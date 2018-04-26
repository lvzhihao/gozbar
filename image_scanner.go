package gozbar

/*
#include <zbar.h>
#cgo LDFLAGS: -lzbar
*/
import "C"
import (
	"runtime"
)

/** zbar.h
 * zbar_image_scanner_t
 */
type ImageScanner struct {
	scanner *C.zbar_image_scanner_t
}

/**
 * create gozbar ImageScanner
 */
func ImageScannerCreate() (obj *ImageScanner) {
	obj = new(ImageScanner)
	obj.Create()
	return
}

/** zbar.h
 * see zbar_image_scanner_create
 */
func (obj *ImageScanner) Create() {
	// zbar.h zbar_image_scanner_create
	obj.scanner = C.zbar_image_scanner_create()
	// set gc func
	runtime.SetFinalizer(obj, (*ImageScanner).Destroy)
}

/**copy from zlib.h
 * parse a configuration string of the form "[symbology.]config[=value]".
 * the config must match one of the recognized names.
 * the symbology, if present, must match one of the recognized names.
 * if symbology is unspecified, it will be set to 0.
 * if value is unspecified it will be set to 1.
 * @returns 0 if the config is parsed successfully, 1 otherwise
 */
func (obj *ImageScanner) SetConfig(symbology C.zbar_symbol_type_t, config C.zbar_config_t, value int) bool {
	ret := C.zbar_image_scanner_set_config(obj.scanner, symbology, config, C.int(value))
	if int(ret) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * scan image
 */
func (obj *ImageScanner) ScanImage(image *C.zbar_image_t) int {
	// zbar.h zbar_scan_image
	return int(C.zbar_scan_image(obj.scanner, image))
}

/**copy from zlib.h
 * parse configuration string using zbar_parse_config()
 * and apply to image scanner using zbar_image_scanner_set_config().
 * @returns 0 for success, non-0 for failure
 * @see zbar_parse_config()
 * @see zbar_image_scanner_set_config()
 * @since 0.4
 */
func (obj *ImageScanner) ParseConfig(config string) bool {
	ret := C.zbar_image_scanner_parse_config(obj.scanner, C.CString(config))
	if int(ret) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * ImageScanner destroy
 */
func (obj *ImageScanner) Destroy() {
	// destroy zbar_image_scanner_t
	C.zbar_image_scanner_destroy(obj.scanner)
}
