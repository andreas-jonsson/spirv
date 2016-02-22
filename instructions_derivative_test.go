// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestDerivative(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x000400af, 1, 2, 3},
			want: &OpDPdx{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b0, 1, 2, 3},
			want: &OpDPdy{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b1, 1, 2, 3},
			want: &OpFwidth{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b2, 1, 2, 3},
			want: &OpDPdxFine{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b3, 1, 2, 3},
			want: &OpDPdyFine{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b4, 1, 2, 3},
			want: &OpFwidthFine{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b5, 1, 2, 3},
			want: &OpDPdxCoarse{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b6, 1, 2, 3},
			want: &OpDPdyCoarse{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400b7, 1, 2, 3},
			want: &OpFwidthCoarse{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
	} {
		testInstruction(t, st)
	}
}
