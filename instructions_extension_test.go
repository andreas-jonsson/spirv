// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestExtensions(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{
				0x00070003,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: &OpExtension{
				Name: "test source extension",
			},
		},
		{
			in: []uint32{
				0x00080004,
				0x00000023,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: &OpExtInstImport{
				ResultId: 0x23,
				Name:     "test source extension",
			},
		},
		{
			in: []uint32{0x0008002c, 1, 2, 3, 4, 5, 6, 7},
			want: &OpExtInst{
				ResultType:  1,
				ResultId:    2,
				Set:         3,
				Instruction: 4,
				Operands:    []Id{5, 6, 7},
			},
		},
	} {
		testInstruction(t, st)
	}
}
