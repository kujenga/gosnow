package gosnow

/*
#cgo CFLAGS: -I./drafter/src/ -I./drafter/ext/snowcrash/src/
#cgo LDFLAGS: -ldrafter -L"./drafter/build/out/Release/" -L"./drafter/build/out/Release/lib.target/"
#include <stdlib.h>
#include <stdio.h>
#include "cdrafter.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	// ScRenderDescriptionsOptionKey causes the parser to render markdown in description
	ScRenderDescriptionsOptionKey = 1 << 0
	// RequireBlueprintNameOptionKey causes the parser to treat missing blueprint name as error
	RequireBlueprintNameOptionKey = 1 << 1
	// ExportSourcemapOptionKey causes the parser to export source maps AST
	ExportSourcemapOptionKey = 1 << 2
)

// OptionParse parses the inputted string and passes it to the drafter library for parsing.
// The JSON result is returned from the function
func OptionParse(source string, flags int) (string, error) {
	cs := C.CString(source)
	var result string
	cr := C.CString(result)
	ret := int(C.drafter_c_parse(cs, C.sc_blueprint_parser_options(flags), &cr))
	if ret != 0 {
		return "", fmt.Errorf("drafter_c_parse failed with code: %v", ret)
	}

	res := C.GoString(cr)
	C.free(unsafe.Pointer(cr)) /* we MUST release allocted memory for result */

	return res, nil
}

// Parse parses the inputted string and passes it to the drafter library for parsing.
// The JSON result is returned from the function
func Parse(source string) (string, error) {
	return OptionParse(source, 0)
}
