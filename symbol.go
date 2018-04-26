package gozbar

/*
#include <zbar.h>
#cgo LDFLAGS: -lzbar
*/
import "C"
import "fmt"

/** zbar.h
 * zbar_symbol_t
 */
type Symbol struct {
	symbol *C.zbar_symbol_t
	info   *SymbolInfo
}

/**
 * symbol info
 */
type SymbolInfo struct {
	Type    int32    `json:"type"`
	Data    string   `json:"data"`
	Length  int32    `json:"length"`
	Quality int32    `json:"quality"`
	LocSize int32    `json:"loc_size"`
	LocXY   []string `json:"loc_x_y"`
}

/**
 * create gozbar Symbol
 */
func SymbolCreate(symbol *C.zbar_symbol_t) (ret *Symbol) {
	if symbol == nil {
		ret = nil
	} else {
		ret = new(Symbol)
		ret.Create(symbol)
	}
	return
}

/** zbar.h
 * see zbar_symbol_create
 */
func (obj *Symbol) Create(symbol *C.zbar_symbol_t) {
	obj.symbol = symbol
	obj.info = nil
}

/** copy from zbar.h
 * iterate the set to which this symbol belongs (there can be only one).
 * @returns the next symbol in the set, or
 * @returns NULL when no more results are available
 */
func (obj *Symbol) Next() *Symbol {
	// see zbar_symbol_next
	symbol := C.zbar_symbol_next(obj.symbol)
	// create gozbar Symbol
	return SymbolCreate(symbol)
}

/**
 * get gozbar SymbolInfo
 */
func (obj *Symbol) GetInfo() *SymbolInfo {
	if obj.info == nil {
		obj.info = &SymbolInfo{
			Type:    int32(obj.GetType()),
			Data:    obj.GetData(),
			Length:  int32(obj.GetDataLength()),
			Quality: int32(obj.GetQuality()),
			LocSize: 0,
			LocXY:   make([]string, 0),
		}
		for i := 0; i < obj.GetLocSize(); i++ {
			obj.info.LocXY = append(obj.info.LocXY, fmt.Sprintf("%d,%d", obj.GetLocX(i), obj.GetLocY(i)))
		}
		obj.info.LocSize = int32(len(obj.info.LocXY))
	}
	return obj.info
}

/**
 * see zbar_symbol_get_data_length
 */
func (obj *Symbol) GetDataLength() int {
	return int(C.zbar_symbol_get_data_length(obj.symbol))
}

/**
 * see zbar_symbol_get_data
 */
func (obj *Symbol) GetData() string {
	retLen := C.zbar_symbol_get_data_length(obj.symbol)
	if int(retLen) > 0 {
		return C.GoString(C.zbar_symbol_get_data(obj.symbol))
	} else {
		return ""
	}
}

/**
 * see zbar_symbol_get_quality
 */
func (obj *Symbol) GetQuality() int {
	return int(C.zbar_symbol_get_quality(obj.symbol))
}

/**
 * see zbar_symbol_get_loc_size
 */
func (obj *Symbol) GetLocSize() int {
	return int(C.zbar_symbol_get_loc_size(obj.symbol))
}

/**
 * see zbar_symbol_get_loc_x
 */
func (obj *Symbol) GetLocX(index int) int {
	return int(C.zbar_symbol_get_loc_x(obj.symbol, C.uint(index)))
}

/**
 * see zbar_symbol_get_loc_y
 */
func (obj *Symbol) GetLocY(index int) int {
	return int(C.zbar_symbol_get_loc_y(obj.symbol, C.uint(index)))
}

/**
 * see zbar_symbol_get_loc_type
 */
func (obj *Symbol) GetType() C.zbar_symbol_type_t {
	return C.zbar_symbol_get_type(obj.symbol)
}

/**
 * foreach by current symbol
 */
func (obj *Symbol) Each(handler func(*SymbolInfo)) {
	inst := obj
	for {
		handler(inst.GetInfo())
		inst = inst.Next()
		if inst == nil {
			break
		}
	}
}

/**
 * foreach by current symbol only data string
 */
func (obj *Symbol) EachData(handler func(string)) {
	inst := obj
	for {
		handler(inst.GetData())
		inst = inst.Next()
		if inst == nil {
			break
		}
	}
}
