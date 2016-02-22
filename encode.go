// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

// Encoder defines an encoder for the SPIR-V format.
// It writes SPIR-V sequences of words into a binary stream.
type Encoder struct {
	w      io.Writer
	buf    []uint32
	endian Endian
}

// NewEncoder creates a new encoder for the given stream and
// instruction set.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:      w,
		endian: LittleEndian,
		buf:    make([]uint32, 12),
	}
}

// EncodeHeader writes the SPIR-V encoding of header h to the
// underlying stream. The magic value (first element in the set), determines
// the byte order for all remaining data being written in this- and subsequent
// calls to the Encoder.
//
// This assumes the header has been validated and is correct.
func (e *Encoder) EncodeHeader(h Header) error {
	// The magic value should be written byte-for-byte, regardless
	// of the endianess. Infact, its byte order defines the byte order
	// for the remaining stream.
	if h.Magic == MagicLE {
		e.endian = LittleEndian
	} else {
		e.endian = BigEndian
	}

	_, err := e.w.Write([]byte{
		byte(h.Magic),
		byte(h.Magic >> 8),
		byte(h.Magic >> 16),
		byte(h.Magic >> 24),
	})

	if err != nil {
		return err
	}

	return e.write([]uint32{
		h.Version,
		h.GeneratorMagic,
		h.Bound,
		h.Reserved,
	})
}

// EncodeInstructionWords writes the SPIR-V encoding of the given instruction
// to the underlying stream. The first word in the set defines the opcode
// and word count. The word count must be <= len(data).
func (e *Encoder) EncodeInstructionWords(data []uint32) error {
	if len(data) == 0 {
		return nil
	}

	// Make sure that the size of the instruction, as defined in
	// its first word, actually matches the data we are receiving.
	size := int(data[0] >> 16)
	if len(data) != size {
		return ErrInvalidInstructionSize
	}

	return e.write(data[:size])
}

// Encode encodes the given instruction into a list of words and
// writes them to the underlying stream. It calls Instruction.Verify
// first to ensure the data is within specifications.
//
// This assumes the instruction has been validated and is correct.
func (e *Encoder) EncodeInstruction(i Instruction) error {
	// Make sure the scratch buffer has sufficient space.
	size := EncodedLen(i)
	if size > len(e.buf) {
		e.buf = make([]uint32, size)
	}

	// Encode the instruction arguments.
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)

	argc, err := encodeValue(rv, e.buf[1:])
	if err != nil {
		return err
	}

	argc++

	// Set the first instruction word.
	e.buf[0] = EncodeOpcode(argc, i.Opcode())

	// Write the words to the underlying stream.
	return e.EncodeInstructionWords(e.buf[:argc])
}

// Write writes exactly len(p) words to the underlying stream.
// It returns an error if this failed.
func (e *Encoder) write(p []uint32) error {
	var buf [4]byte

	for _, word := range p {
		if e.endian == LittleEndian {
			buf[0] = byte(word)
			buf[1] = byte(word >> 8)
			buf[2] = byte(word >> 16)
			buf[3] = byte(word >> 24)
		} else {
			buf[0] = byte(word >> 24)
			buf[1] = byte(word >> 16)
			buf[2] = byte(word >> 8)
			buf[3] = byte(word)
		}

		_, err := e.w.Write(buf[:])
		if err != nil {
			return err
		}
	}

	return nil
}

// EncodedLen returns the number of words the given instruction
// will occupy once encoded.
func EncodedLen(i Instruction) int {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)
	return encodedValueLen(rv) + 1
}

func encodedValueLen(rv reflect.Value) int {
	switch rv.Kind() {
	case reflect.Struct:
		return encodedStructLen(rv)
	case reflect.Slice, reflect.Array:
		return encodedSliceLen(rv)
	case reflect.Uint32:
		return 1
	case reflect.String:
		return int(String(rv.String()).EncodedLen())
	}
	return 0
}

func encodedStructLen(rv reflect.Value) int {
	var len int

	for i := 0; i < rv.NumField(); i++ {
		len += encodedValueLen(rv.Field(i))
	}

	return len
}

func encodedSliceLen(rv reflect.Value) int {
	var len int

	for i := 0; i < rv.Len(); i++ {
		len += encodedValueLen(rv.Index(i))
	}

	return len
}

func encodeValue(rv reflect.Value, out []uint32) (uint32, error) {
	switch rv.Kind() {
	case reflect.Struct:
		return encodeStruct(rv, out)

	case reflect.Slice, reflect.Array:
		return encodeSlice(rv, out)

	case reflect.Uint32:
		out[0] = uint32(rv.Uint())
		return 1, nil

	case reflect.String:
		str := String(rv.String())
		size := str.EncodedLen()
		str.Encode(out)
		return size, nil
	}

	return 0, fmt.Errorf("unsupported type: %v", rv.Kind())
}

// encodeStruct encodes the given struct.
func encodeStruct(rv reflect.Value, out []uint32) (uint32, error) {
	var index uint32

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		tag := ft.Tag.Get("spirv")
		if hasFieldOption(tag, "optional") && valueIsNil(fv) {
			continue
		}

		argc, err := encodeValue(fv, out[index:])
		if err != nil {
			return 0, err
		}

		index += argc
	}

	return index, nil
}

// valueIsNil returns true if the given value is considered empty.
// This depends on the underlying type.
func valueIsNil(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		return rv.Len() == 0

	case reflect.Uint32:
		return rv.Uint() == 0
	}

	return false
}

// hasFieldOption returns true if the given struct field tag
// contains the specified option.
func hasFieldOption(tag, option string) bool {
	fields := strings.Split(tag, ",")

	for _, fld := range fields {
		if len(fld) > 0 && strings.EqualFold(fld, option) {
			return true
		}
	}

	return false
}

// encodeSlice encodes the given slice
func encodeSlice(rv reflect.Value, out []uint32) (uint32, error) {
	var index uint32

	for i := 0; i < rv.Len(); i++ {
		fld := rv.Index(i)

		argc, err := encodeValue(fld, out[index:])
		if err != nil {
			return 0, err
		}

		index += argc
	}

	return index, nil
}
