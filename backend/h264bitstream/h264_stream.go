package h264bitstream

//#include "h264_stream.h"
import "C"
import (
	"os"
	"unsafe"
)

const BUFSIZE = 32 * 1024 * 1024

type Stream struct {
	handle    *C.h264_stream_t
	file      *os.File
	buffer    []byte
	remaining int
	offset    int
}

func NewStream(filename string) (*Stream, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Stream{
		handle: C.h264_new(),
		file:   f,
		buffer: make([]byte, BUFSIZE),
	}, nil
}

func (s *Stream) Release() {
	C.h264_free(s.handle)
	s.file.Close()
}

func (s *Stream) ReadNextNalUnit() ([]byte, error) {
	if s.remaining <= 0 {
		n, err := s.file.Read(s.buffer)
		if err != nil {
			return nil, err
		}
		s.remaining = n
	}
	var nal_start, nal_end int

	p := s.buffer[s.offset:]
	C.find_nal_unit(
		(*C.uint8_t)(unsafe.Pointer(&p[0])),
		C.int(s.remaining),
		(*C.int)(unsafe.Pointer(&nal_start)),
		(*C.int)(unsafe.Pointer(&nal_end)),
	)

	p = p[nal_start:]
	C.read_debug_nal_unit(s.handle,
		(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])))),
		C.int(nal_end-nal_start),
	)

	s.offset += nal_end
	s.remaining -= nal_end

	return s.buffer[nal_start:nal_end], nil
}
