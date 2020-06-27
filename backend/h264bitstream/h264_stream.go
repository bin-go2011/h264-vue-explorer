package h264bitstream

//#include "h264_stream.h"
import "C"

type h264Stream struct {
	handle *C.h264_stream_t
}

func NewStream() *h264Stream {
	return &h264Stream{
		handle: C.h264_new(),
	}
}

func (s *h264Stream) Release() {
	C.h264_free(s.handle)
}
