package rtmp

// #cgo LDFLAGS: -lrtmp
// #include <librtmp/rtmp.h>
// #include <stdlib.h>
import "C"

import (
	"io"
	"unsafe"
)

type RTMP struct {
	rtmp *C.struct_RTMP
	url  *C.char
	io.ReadCloser
}

// Initialize do RTMP_Alloc and RTMP_Init at once
func Init() (*RTMP, error) {
	r := &RTMP{}
	rtmp, err := C.RTMP_Alloc()
	if err != nil {
		return nil, err
	}
	r.rtmp = rtmp

	_, err = C.RTMP_Init(r.rtmp)
	return r, err
}

func (r *RTMP) SetupURL(url string) error {
	r.url = C.CString(url)
	_, err := C.RTMP_SetupURL(r.rtmp, r.url)
	return err
}

func (r *RTMP) Connect() error {
	ret, err := C.RTMP_Connect(r.rtmp, nil)
	if ret != 0 {
		return err
	}
	return nil
}

func (r *RTMP) Read(p []byte) (n int, err error) {
	size, err := C.RTMP_Read(r.rtmp, (*C.char)(unsafe.Pointer(&p[0])), C.int(len(p)))
	if err != nil {
		return 0, err
	}
	return int(size), nil
}

// Close do RTMP_Close and RTMP_Free at once
func (r *RTMP) Close() error {
	_, err := C.RTMP_Close(r.rtmp)
	if err != nil {
		return err
	}
	_, err = C.RTMP_Free(r.rtmp)
	if err != nil {
		return err
	}
	C.free(unsafe.Pointer(r.url))
	return nil
}
