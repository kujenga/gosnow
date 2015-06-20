package gosnow

/*
#cgo CFLAGS: -I./drafter/src/ -I./drafter/ext/snowcrash/src/
#cgo LDFLAGS: -ldrafter
#include <stdlib.h>
#include <stdio.h>
#include "cdrafter.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Parse parses the inputted string and passes it to the drafter library for parsing.
// The JSON result is returned from the function
func Parse(source string) (string, error) {
	cs := C.CString(source)
	var result string
	cr := C.CString(result)
	ret := int(C.drafter_c_parse(cs, 0, &cr))
	if ret != 0 {
		return "", fmt.Errorf("drafter_c_parse faile with code: %v", ret)
	}

	res := C.GoString(cr)
	C.free(unsafe.Pointer(cr)) /* we MUST release allocted memory for result */

	return res, nil
}
