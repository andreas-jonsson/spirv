// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestFunctions(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00050028, 1, 2, 3, 4},
			want: &OpFunction{
				ResultType:   1,
				ResultId:     2,
				ControlMask:  3,
				FunctionType: 4,
			},
		},
		{
			in: []uint32{0x00030029, 1, 2},
			want: &OpFunctionParameter{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in:   []uint32{0x0001002a},
			want: &OpFunctionEnd{},
		},
		{
			in: []uint32{0x0004002b, 1, 2, 3},
			want: &OpFunctionCall{
				ResultType: 1,
				ResultId:   2,
				Function:   3,
			},
		},
		{
			in: []uint32{0x0007002b, 1, 2, 3, 4, 5, 6},
			want: &OpFunctionCall{
				ResultType: 1,
				ResultId:   2,
				Function:   3,
				Argv:       []Id{4, 5, 6},
			},
		},
	} {
		testInstruction(t, st)
	}
}
