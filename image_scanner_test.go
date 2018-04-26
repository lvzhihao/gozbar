package gozbar

import (
	"testing"
)

// scanner create testing
func TestImageScannerCreate(t *testing.T) {
	scan := ImageScannerCreate()
	scan.SetConfig(ZBAR_NONE, ZBAR_CFG_ENABLE, 1) // enable all symbols
	t.Log(scan)
}
