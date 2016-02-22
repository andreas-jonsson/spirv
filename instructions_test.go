// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

type InstructionTest struct {
	in   []uint32
	want Instruction
	err  error
}

func testInstruction(t *testing.T, st InstructionTest) {
	dec := NewDecoder(testWordReader(st.in))
	have, err := dec.DecodeInstruction()

	if err != nil {
		if !reflect.DeepEqual(err, st.err) {
			t.Fatalf("decode error mismatch: %v\nHave: %v\nWant: %v",
				st.in, err, st.err)
		}

		// We got an error as we expected, so no further processing
		// is needed.
		return
	}

	err = have.Verify()
	if err != nil {
		if !reflect.DeepEqual(err, st.err) {
			t.Fatalf("decode error mismatch: %v\nHave: %v\nWant: %v",
				st.in, err, st.err)
		}

		// We got an error as we expected, so no further processing
		// is needed.
		return
	}

	if !reflect.DeepEqual(have, st.want) {
		t.Fatalf("decode value mismatch: %v\nHave: %T(%+v)\nWant: %T(%+v)",
			st.in, have, have, st.want, st.want)
	}

	// Encode the instruction again and compare the output to the
	// original input. We can only encode to bytes, so we need to
	// jump through some hoops to compare to the []uint32 input slice.

	var outa bytes.Buffer
	enc := NewEncoder(&outa)
	err = enc.EncodeInstruction(have)
	if err != nil {
		t.Fatal(err)
	}

	outb := testWordReader(st.in)
	if !bytes.Equal(outa.Bytes(), outb.Bytes()) {
		t.Fatalf("encode mismatch: %T(%v)\nHave: %v\nWant: %v",
			have, have, outa.Bytes(), outb.Bytes())
	}
}

// Test the decoding behaviour of a few edge cases.
func TestFaultyInstructions(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in:   nil,
			want: nil,
			err:  io.EOF,
		},
		{
			in:  []uint32{0x00000001},
			err: ErrInvalidInstructionSize,
		},
		{
			in:  []uint32{0x00010001},
			err: ErrMissingInstructionArgs,
		},
		{
			in:  []uint32{0x00010000},
			err: ErrUnacceptable,
		},
		{
			in:  []uint32{0x0001ffff},
			err: fmt.Errorf("unknown instruction: 0000ffff"),
		},
	} {
		testInstruction(t, st)
	}
}

// testWordReader returns a type which reads the given set of words
// as a sequence of bytes.
func testWordReader(data []uint32) *bytes.Buffer {
	var buf bytes.Buffer

	for _, word := range data {
		buf.WriteByte(byte(word))
		buf.WriteByte(byte(word >> 8))
		buf.WriteByte(byte(word >> 16))
		buf.WriteByte(byte(word >> 24))
	}

	return &buf
}
