// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"io"
	"reflect"
)

// Decoder defines a decoder for the SPIR-V format.
// It reads binary data from a stream and yields sequences
// of 32-bit words.
type Decoder struct {
	r      io.Reader
	ubuf   []uint32 // Scratch buffer for instruction decoding.
	bbuf   [4]byte  // Scratch buffer for the word reader.
	endian Endian
}

// NewDecoder creates a new decoder for the given stream and instruction set.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:      r,
		endian: LittleEndian,
		ubuf:   make([]uint32, 16),
	}
}

// DecodeHeader reads a module header from the underlying stream.
//
// The magic value's byte order will be used to determine the byte order
// for the entire stream.
//
// Returns an error if the magic value is invalid.
func (d *Decoder) DecodeHeader() (Header, error) {
	var hdr Header

	// Read the magic value. This is the one value we read as separate bytes.
	// The order will tell us the byte order of the rest of the stream.
	_, err := io.ReadFull(d.r, d.bbuf[:])
	if err != nil {
		if err == io.EOF {
			return hdr, ErrUnexpectedEOF
		}
		return hdr, err
	}

	hdr.Magic = uint32(d.bbuf[0]) | uint32(d.bbuf[1])<<8 |
		uint32(d.bbuf[2])<<16 | uint32(d.bbuf[3])<<24

	// Make sure it's a valid number. The order of the magic bytes lets us
	// determine the endianness of the stream's remaining data.
	switch hdr.Magic {
	case MagicLE:
		d.endian = LittleEndian
	case MagicBE:
		d.endian = BigEndian
	default:
		return hdr, ErrInvalidMagicValue
	}

	// Read remaining header.
	err = d.read(d.ubuf[:4])
	if err != nil {
		return hdr, err
	}

	hdr.Version = d.ubuf[0]
	hdr.GeneratorMagic = d.ubuf[1]
	hdr.Bound = d.ubuf[2]
	hdr.Reserved = d.ubuf[3]
	return hdr, nil
}

// DecodeInstructionWords decodes the next instruction from the underlying
// stream. The returned slice of words contains all data for the entire
// instruction.
//
// The data remains valid until the next call to any of the decoder methods.
func (d *Decoder) DecodeInstructionWords() ([]uint32, error) {
	// Read the first word: word count + opcode.
	err := d.read(d.ubuf[:1])
	if err != nil {
		if err == ErrUnexpectedEOF {
			return nil, io.EOF // Not fatal in this one case -- just end of stream.
		}
		return nil, err
	}

	words := int(d.ubuf[0] >> 16)
	if words < 1 {
		return nil, ErrInvalidInstructionSize
	}

	if words > 1 {
		// Resize read buffer if necessary.
		if words >= len(d.ubuf) {
			tmp := d.ubuf[0]
			d.ubuf = make([]uint32, words)
			d.ubuf[0] = tmp
		}

		// words defines the number of words for the entire instruction.
		// This includes the first one we just read. So remaining number of
		// operands are words - 1.
		err = d.read(d.ubuf[1:words])
		if err != nil {
			return nil, err
		}
	}

	return d.ubuf[:words:words], nil
}

// DecodeInstruction decodes the next instruction from the underlying stream.
// The returned structure defines a copy of the stream data and will
// remain valid as long as you need it to be.
//
// Returns an error if there is no matching instruction or the
// decoding failed.
func (d *Decoder) DecodeInstruction() (Instruction, error) {
	words, err := d.DecodeInstructionWords()
	if err != nil {
		return nil, err
	}

	return DecodeInstruction(words)
}

// DecodeInstructionFromWords decodes an instruction from the given
// set of words. Additionally, it verifies that the values for each
// instruction field meet the requirements as defined in the specification.
//
// Returns an error if there is no matching instruction or the
// decoding failed.
func DecodeInstruction(words []uint32) (Instruction, error) {
	wordCount := words[0] >> 16
	opcode := words[0] & 0xffff

	if wordCount == 0 {
		return nil, ErrInvalidInstructionSize
	}

	if len(words) < int(wordCount) {
		return nil, ErrInvalidInstructionSize
	}

	constructor, ok := instructions[opcode]
	if !ok {
		return nil, fmt.Errorf("unknown instruction: %08x", opcode)
	}

	instr := constructor()

	// This instruction is illegal.
	// FIXME Remove this once instruction validation code is in.
	switch instr.(type) {
	case *OpNop:
		return nil, ErrUnacceptable
	}

	rv := reflect.ValueOf(instr)
	_, err := decodeValue(rv, words[1:])

	return instr, err
}

// Next reads exactly len(p) words from the stream.
// Returns an error if there is either not enough data or something else
// went wrong.
func (d *Decoder) read(p []uint32) error {
	for i := range p {
		_, err := io.ReadFull(d.r, d.bbuf[:])
		if err != nil {
			if err == io.EOF {
				// This particular EOF is unexpected.
				// We should be able to read at least len(p) words.
				return ErrUnexpectedEOF
			}
			return err
		}

		if d.endian == LittleEndian {
			p[i] = uint32(d.bbuf[0]) | uint32(d.bbuf[1])<<8 | uint32(d.bbuf[2])<<16 |
				uint32(d.bbuf[3])<<24
		} else {
			p[i] = uint32(d.bbuf[3]) | uint32(d.bbuf[2])<<8 | uint32(d.bbuf[1])<<16 |
				uint32(d.bbuf[0])<<24
		}
	}

	return nil
}

// decodeValue decodes the given words into the specified instruction.
func decodeValue(rv reflect.Value, argv []uint32) ([]uint32, error) {
	rv = reflect.Indirect(rv)

	switch rv.Kind() {
	case reflect.Struct:
		return decodeStruct(rv, argv)
	case reflect.Slice:
		return decodeSlice(rv, argv)
	}

	if len(argv) == 0 {
		return nil, ErrMissingInstructionArgs
	}

	switch rv.Kind() {
	case reflect.Uint32:
		rv.SetUint(uint64(argv[0]))
		return argv[1:], nil
	case reflect.String:
		return decodeString(rv, argv)
	}

	return argv, fmt.Errorf("unsupported type: %v", rv.Kind())
}

// decodeSlice decodes data into a slice field.
func decodeSlice(rv reflect.Value, argv []uint32) ([]uint32, error) {
	if len(argv) == 0 {
		return nil, nil
	}

	rt := rv.Type()
	et := rt.Elem()

	// Ensure that the slice elements are either uint32, or an alias thereof.
	if et.Kind() != reflect.Uint32 {
		return nil, fmt.Errorf("slice elements must resolve to uint32")
	}

	// if this is a straight-up uint32 slice, just copy the input.
	// Any other value is an alias and implements at least the Verifiable
	// interface. Which means it has at least 1 method attached to it.
	if rt.Elem().NumMethod() == 0 {
		copy := Copy(argv)
		rv.Set(reflect.ValueOf(copy))
		return nil, nil
	}

	// Otherwise we have to do some manualy copying magic.
	new := reflect.MakeSlice(rt, 0, len(argv))

	for _, word := range argv {
		elem := reflect.New(et)
		elem = reflect.Indirect(elem)
		elem.SetUint(uint64(word))
		new = reflect.Append(new, elem)
	}

	rv.Set(new)
	return nil, nil
}

// decodeStruct decodes input data into a struct.
func decodeStruct(rv reflect.Value, argv []uint32) ([]uint32, error) {
	var err error

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		tag := ft.Tag.Get("spirv")
		if hasFieldOption(tag, "optional") && len(argv) == 0 {
			continue
		}

		argv, err = decodeValue(fv, argv)
		if err != nil {
			return nil, err
		}
	}

	return argv, nil
}

// decodeString decodes input data into a string value.
func decodeString(rv reflect.Value, argv []uint32) ([]uint32, error) {
	str := DecodeString(argv)
	size := str.EncodedLen()
	rv.SetString(string(str))
	return argv[size:], nil
}
