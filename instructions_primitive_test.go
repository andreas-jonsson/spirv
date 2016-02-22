// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestPrimitive(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in:   []uint32{0x00100b8},
			want: &OpEmitVertex{},
		},
		{
			in:   []uint32{0x00100b9},
			want: &OpEndPrimitive{},
		},
		{
			in: []uint32{0x00200ba, 1},
			want: &OpEmitStreamVertex{
				Stream: 1,
			},
		},
		{
			in: []uint32{0x00200bb, 1},
			want: &OpEndStreamPrimitive{
				Stream: 1,
			},
		},
	} {
		testInstruction(t, st)
	}
}
