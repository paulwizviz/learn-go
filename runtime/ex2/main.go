//go:build darwin && cgo

package main

/*
#include <sys/sysctl.h>
#include <stdlib.h> // Include stdlib.h for size_t
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func getPhysicalCoreCount() int {
	var coreCount C.int
	var size C.size_t = C.size_t(unsafe.Sizeof(coreCount)) // Correct type
	C.sysctlbyname(C.CString("hw.physicalcpu"), unsafe.Pointer(&coreCount), &size, nil, 0)
	return int(coreCount)
}

func main() {
	physicalCores := getPhysicalCoreCount()
	fmt.Println("Physical Cores:", physicalCores)
}
