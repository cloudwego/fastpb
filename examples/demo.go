package examples

import (
	"github.com/cloudwego/fastpb"
)

// Marshal
func Marshal(msg fastpb.Writer) []byte {
	buf := make([]byte, msg.Size())
	msg.FastWrite(buf)
	return buf
}

// Unmarshal
func Unmarshal(buf []byte, msg fastpb.Reader) error {
	_, err := fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), msg)
	return err
}
