// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

type DecodeInstructionTest struct {
	in   []byte
	want []uint32
	err  error
}

func TestDecodeWords(t *testing.T) {
	for i, st := range []DecodeInstructionTest{
		{
			in:   nil,
			want: nil,
			err:  io.EOF,
		},
		{
			in: []byte{
				0x01, 0x00, 0x02, 0x00,
			},
			want: nil,
			err:  ErrUnexpectedEOF,
		},
		{
			in: []byte{
				0x01, 0x00, 0x02, 0x00,
				0x01, 0x00, 0x00, 0x00,
			},
			want: []uint32{
				0x00020001,
				0x00000001,
			},
		},
		{
			in: []byte{
				0x01, 0x00, 0x01, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			want: []uint32{
				0x00010001,
			},
		},
	} {
		dec := NewDecoder(bytes.NewBuffer(st.in))
		have, err := dec.DecodeInstructionWords()

		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: error mismatch:\nHave: %v\nWant: %v",
					i, err, st.err)
			}
			continue // Expected error -- Just move on.
		}

		if !reflect.DeepEqual(have, st.want) {
			t.Fatalf("case %d: value mismatch:\nHave: %v\nWant: %v",
				i, have, st.want)
		}
	}
}

type DecodeHeaderTest struct {
	in   []byte
	want Header
	err  error
}

func TestDecodeHeader(t *testing.T) {
	for i, st := range []DecodeHeaderTest{
		{
			in:  nil,
			err: ErrUnexpectedEOF,
		},
		{
			in:  []byte{0x01, 0x02, 0x03, 0x04},
			err: ErrInvalidMagicValue,
		},
		{
			in:  []byte{0x03, 0x02, 0x23, 0x07},
			err: ErrUnexpectedEOF,
		},
		{
			in: []byte{
				0x03, 0x02, 0x23, 0x07,
				0x63, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0xff, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			want: Header{MagicLE, 99, 1, 255, 0},
		},
		{
			in: []byte{
				0x07, 0x23, 0x02, 0x03,
				0x00, 0x00, 0x00, 0x63,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0xff,
				0x00, 0x00, 0x00, 0x00,
			},
			want: Header{MagicBE, 99, 1, 255, 0},
		},
	} {
		dec := NewDecoder(bytes.NewBuffer(st.in))
		have, err := dec.DecodeHeader()

		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: error mismatch:\nHave: %v\nWant: %v",
					i, err, st.err)
			}
			continue // Expected error -- Just move on.
		}

		if !reflect.DeepEqual(have, st.want) {
			t.Fatalf("case %d: value mismatch:\nHave: %v\nWant: %v",
				i, have, st.want)
		}
	}
}
