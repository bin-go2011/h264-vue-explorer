package h264bitstream

//#include "h264_stream.h"
import "C"
import "unsafe"

type h264Stream struct {
	handle uintptr
}

func NewStream() *h264Stream {
	return &h264Stream{
		handle: uintptr(unsafe.Pointer(C.h264_new())),
	}
}

func (s *h264Stream) Release() {
	h := (*C.h264_stream_t)(unsafe.Pointer(s.handle))
	C.h264_free(h)
}
