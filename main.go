package main

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

func main() {
	source := C.CString("# My API\n## GET /message\n + Response 200 (text/plain)\n\n        Hello World\n")
	var result string
	cresult := C.CString(result)
	ret := int(C.drafter_c_parse(source, 0, &cresult))
	if ret == 0 {
		fmt.Printf("Result: %v\n", "OK")
	} else {
		fmt.Printf("Result: %v\n", "ERROR")
	}

	fmt.Printf("Serialized JSON result:\n%v\n", C.GoString(cresult))

	C.free(unsafe.Pointer(cresult)) /* we MUST release allocted memory for result */
}
