// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

type OpcodeTest struct {
	a, b uint32
	want uint32
}

func TestOpcode(t *testing.T) {
	for i, st := range []OpcodeTest{
		{0, 0, 0},
		{1, 0, 0x00010000},
		{0, 1, 0x00000001},
		{0xffff, 1, 0xffff0001},
		{0x10000, 1, 0x00000001},
		{0x10000, 0x10000, 0},
	} {
		have := EncodeOpcode(st.a, st.b)
		if have != st.want {
			t.Fatalf("case %d: encode mismatch: %d/%d\nHave: %08x\nWant: %08x",
				i, st.a, st.b, have, st.want)
		}

		haveA, haveB := DecodeOpcode(have)
		if haveA != st.a&0xffff || haveB != st.b&0xffff {
			t.Fatalf("case %d: decode mismatch: %08x\nHave: %d/%d\nWant: %d/%d",
				i, have, haveA, haveB, st.a, st.b)
		}
	}
}
