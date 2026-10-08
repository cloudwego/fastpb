package fastpb

import "testing"

// FuzzConsumeBytes feeds arbitrary wire bytes and requires ConsumeBytes to
// either return a value or a clean negative length -- never panic.
func FuzzConsumeBytes(f *testing.F) {
	for _, s := range [][]byte{
		{0x03, 'a', 'b', 'c'},
		{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01},
		{0x80, 0x80, 0x80, 0x80},
		{},
	} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, b []byte) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("panic on %x: %v", b, r)
			}
		}()
		_, _ = ConsumeBytes(b)
	})
}
