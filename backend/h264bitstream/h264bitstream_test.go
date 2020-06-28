package h264bitstream

import (
	"testing"
)

func TestFindNal(t *testing.T) {
	stream, err := NewStream("samples/x264_test.264")
	if err != nil {
		panic(err)
	}
	defer stream.Release()

	stream.ReadNextNalUnit()
	stream.ReadNextNalUnit()
}
