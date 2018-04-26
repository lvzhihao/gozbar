package gozbar

import (
	"fmt"
	"image/jpeg"
	"os"
	"strings"
	"testing"
)

func TestZbarFourcc(t *testing.T) {
	y800 := ZbarFourcc('Y', '8', '0', '0')
	if strings.Compare(fmt.Sprintf("0x%x", y800), "0x30303859") != 0 {
		t.Error("test zbar_fourcc error")
	}
}

func TestQrImage(t *testing.T) {
	fp, err := os.Open("testdata/qr_normal.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	img, err := jpeg.Decode(fp)
	if err != nil {
		t.Error(err)
		return
	}
	zbarImg := ImageCreate(img)

	scan := ImageScannerCreate()
	scan.SetConfig(ZBAR_NONE, ZBAR_CFG_ENABLE, 0)   //disable all symbol
	scan.SetConfig(ZBAR_QRCODE, ZBAR_CFG_ENABLE, 1) //enable qrcode only
	ret := scan.ScanImage(zbarImg.Image())
	if ret <= 0 {
		t.Error("no result")
		return
	}
	t.Log("result count:", ret)
	symbol := zbarImg.FirstSymbol()
	symbol.EachData(func(info string) {
		t.Logf("%+v\n", info)
	})
}

func TestQrExtImage(t *testing.T) {
	fp, err := os.Open("testdata/qrext.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	img, err := jpeg.Decode(fp)
	if err != nil {
		t.Error(err)
		return
	}
	zbarImg := ImageCreate(img)

	scan := ImageScannerCreate()
	ok := scan.SetConfig(ZBAR_NONE, ZBAR_CFG_ENABLE, 1)
	if !ok {
		t.Error("image scanner set config error")
		return
	}
	ret := scan.ScanImage(zbarImg.Image())
	if ret <= 0 {
		t.Error("no result")
	}
	t.Log("result count:", ret)
	symbol := zbarImg.FirstSymbol()
	symbol.Each(func(info *SymbolInfo) {
		t.Logf("%+v\n", info)
	})
}
