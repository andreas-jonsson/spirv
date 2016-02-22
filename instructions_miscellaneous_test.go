// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestMiscellaneous(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in:  []uint32{0x00010000},
			err: ErrUnacceptable,
		},
		{
			in: []uint32{0x0003002d, 1, 2},
			want: &OpUndef{
				ResultType: 1,
				ResultId:   2,
			},
		},
	} {
		testInstruction(t, st)
	}
}
