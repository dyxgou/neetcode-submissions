/*
#include <stdlib.h>
*/
import "C"

import (
	"log/slog"
	"unsafe"
)

const zero int = 0

type DynamicArray struct {
	addr *DynamicArray
	ptr  unsafe.Pointer
	len  int
	cap  int
}

func NewDynamicArray(capacity int) *DynamicArray {
	if capacity <= 0 {
		panic("capacity should be greater or equal to 1")
	}

	ptr := C.calloc(C.size_t(capacity), C.size_t(unsafe.Sizeof(zero)))

	da := &DynamicArray{
		len: 0,
		ptr: ptr,
		cap: capacity,
	}

	da.addr = da
	return da
}

func (d *DynamicArray) copyCheck() {
	if d.addr != d {
		panic("ilegal use of DynamicArray: copied by value")
	}
}

func (d *DynamicArray) Get(i int) int {
	d.copyCheck()

	val := *(*int)(unsafe.Add(d.ptr, int(uintptr(i)*unsafe.Sizeof(zero))))
	slog.Info("get val", "i", i, "val", val, "newptr", int(uintptr(i)*unsafe.Sizeof(zero)))
	return val
}

func (d *DynamicArray) Set(i, val int) {
	d.copyCheck()

	ptr := (*int)(unsafe.Add(d.ptr, int(uintptr(i)*unsafe.Sizeof(zero))))

	*ptr = val
}

func (d *DynamicArray) resize() {
	if d.cap < 1024 {
		d.cap = d.cap * 2
	} else {
		// d.cap / 4
		d.cap += d.cap >> 2
	}

	d.ptr = C.reallocarray(d.ptr, C.ulong(d.cap), C.ulong(unsafe.Sizeof(zero)))
}

func (d *DynamicArray) Pushback(val int) {
	d.copyCheck()
	if d.len >= d.cap {
		d.resize()
	}

	d.Set(d.len, val)
	d.len++
}

func (d *DynamicArray) Popback() int {
	d.copyCheck()
	val := *(*int)(unsafe.Add(d.ptr, int(uintptr(d.len-1)*unsafe.Sizeof(zero))))
	d.len--

	return val
}

func (d *DynamicArray) GetSize() int {
	d.copyCheck()
	return d.len
}

func (d *DynamicArray) GetCapacity() int {
	d.copyCheck()
	return d.cap
}