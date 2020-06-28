package h264bitstream

//#include "h264_stream.h"
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

const BUFSIZE = 32 * 1024 * 1024

type Stream struct {
	handle *C.h264_stream_t
	file   *os.File
	buf    []byte
}

func NewStream(filename string) (*Stream, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Stream{
		handle: C.h264_new(),
		file:   f,
		buf:    make([]byte, BUFSIZE),
	}, nil
}

func (s *Stream) Release() {
	C.h264_free(s.handle)
	s.file.Close()
}

func (s *Stream) ReadNalUnit() {
	n, _ := s.file.Read(s.buf)
	var nal_start, nal_end int
	C.find_nal_unit(
		(*C.uint8_t)(unsafe.Pointer(&s.buf[0])),
		C.int(n),
		(*C.int)(unsafe.Pointer(&nal_start)),
		(*C.int)(unsafe.Pointer(&nal_end)),
	)
	fmt.Printf("nal_start=%d, nal_end=%d\n", nal_start, nal_end)
}
