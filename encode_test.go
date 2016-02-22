// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"reflect"
	"testing"
)

type EncodeInstructionTest struct {
	in   []uint32
	want []byte
	err  error
}

func TestEncodeWords(t *testing.T) {
	for i, st := range []EncodeInstructionTest{
		{
			in:   nil,
			want: nil,
		},
		{
			in:   []uint32{0x00020001},
			want: nil,
			err:  ErrInvalidInstructionSize,
		},
		{
			in:   []uint32{0x00020001},
			want: nil,
			err:  ErrInvalidInstructionSize,
		},
		{
			in: []uint32{
				0x00020001,
				0x00000001,
			},
			want: []byte{
				0x01, 0x00, 0x02, 0x00,
				0x01, 0x00, 0x00, 0x00,
			},
		},
		{
			in:   []uint32{0x00010001},
			want: []byte{0x01, 0x00, 0x01, 0x00},
		},
	} {
		var have bytes.Buffer
		enc := NewEncoder(&have)
		err := enc.EncodeInstructionWords(st.in)

		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: error mismatch:\nHave: %v\nWant: %v",
					i, err, st.err)
			}
			continue // Expected error -- Just move on.
		}

		if !bytes.Equal(have.Bytes(), st.want) {
			t.Fatalf("case %d: value mismatch:\nHave: %v\nWant: %v",
				i, have.Bytes(), st.want)
		}
	}
}

type EncodeHeaderTest struct {
	in   Header
	want []byte
	err  error
}

func TestEncodeHeader(t *testing.T) {
	for i, st := range []EncodeHeaderTest{
		{
			in:   Header{},
			want: nil,
			err:  ErrInvalidMagicValue,
		},
		{
			in:   Header{123, 99, 1, 255, 0},
			want: nil,
			err:  ErrInvalidMagicValue,
		},
		{
			in:   Header{MagicLE, 100, 1, 255, 0},
			want: nil,
			err:  ErrInvalidVersion,
		},
		{
			in: Header{MagicLE, 99, 1, 255, 0},
			want: []byte{
				0x03, 0x02, 0x23, 0x07,
				0x63, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0xff, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			in: Header{MagicBE, 99, 1, 255, 0},
			want: []byte{
				0x07, 0x23, 0x02, 0x03,
				0x00, 0x00, 0x00, 0x63,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0xff,
				0x00, 0x00, 0x00, 0x00,
			},
		},
	} {
		err := st.in.Verify()
		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: error mismatch:\nHave: %v\nWant: %v",
					i, err, st.err)
			}
			continue // Expected error -- Just move on.
		}

		var have bytes.Buffer
		enc := NewEncoder(&have)
		err = enc.EncodeHeader(st.in)
		if err != nil {
			if !reflect.DeepEqual(err, st.err) {
				t.Fatalf("case %d: error mismatch:\nHave: %v\nWant: %v",
					i, err, st.err)
			}
			continue // Expected error -- Just move on.
		}

		if !bytes.Equal(have.Bytes(), st.want) {
			t.Fatalf("case %d: value mismatch:\nHave: %v\nWant: %v",
				i, have.Bytes(), st.want)
		}
	}
}
