// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"reflect"
	"testing"
)

type StringTest struct {
	in      String
	want    []uint32
	wantLen uint32
}

func TestString(t *testing.T) {
	for i, st := range []StringTest{
		{"", []uint32{0}, 1},
		{"foo", []uint32{0x006f6f66}, 1},
		{"foof", []uint32{0x666f6f66, 0}, 2},
		{"foo bar", []uint32{0x206f6f66, 0x00726162}, 2},
		{"foo bars", []uint32{0x206f6f66, 0x73726162, 0}, 3},
		{"fo\no bar", []uint32{0x6f0a6f66, 0x72616220, 0}, 3},
		{"世界", []uint32{0xe796b8e4, 0x00008c95}, 2},
		{"\nsomenewline", []uint32{0x6d6f730a, 0x77656e65, 0x656e696c, 0}, 4},
	} {
		haveLen := st.in.EncodedLen()
		if st.wantLen != haveLen {
			t.Fatalf("case %d: string length mismatch: %q\nHave: %d\nWant: %d",
				i, st.in, st.wantLen, haveLen)
		}

		have := make([]uint32, haveLen)
		n := st.in.Encode(have)
		if n != len(have) {
			t.Fatalf("case %d: string length mismatch: %q\nHave: %d\nWant: %d",
				i, st.in, n, len(have))
		}

		if !reflect.DeepEqual(have, st.want) {
			t.Fatalf("case %d: fromString mismatch: %q.\nHave: %08x\nWant: %08x",
				i, st.in, have, st.want)
		}

		haveStr := DecodeString(have)
		if haveStr != st.in {
			t.Fatalf("case %d: toString mismatch: %v.\nHave: %q\nWant: %q",
				i, have, haveStr, st.in)
		}
	}
}
