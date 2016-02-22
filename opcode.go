// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// EncodeOpcode returns the first word of an instruction from the
// given word cound and opcode.
func EncodeOpcode(wordCount, opcode uint32) uint32 {
	return (wordCount&0xffff)<<16 | (opcode & 0xffff)
}

// DecodeOpcode extracts the word count and opcode from the given word.
func DecodeOpcode(word uint32) (uint32, uint32) {
	return (word >> 16), word & 0xffff
}
