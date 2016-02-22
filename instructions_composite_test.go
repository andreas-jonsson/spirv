// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestComposite(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x0005003a, 1, 2, 3, 4},
			want: &OpVectorExtractDynamic{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
				Index:      4,
			},
		},
		{
			in: []uint32{0x0006003b, 1, 2, 3, 4, 5},
			want: &OpVectorInsertDynamic{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
				Component:  4,
				Index:      5,
			},
		},
		{
			in: []uint32{0x0008003c, 1, 2, 3, 4, 5, 6, 7},
			want: &OpVectorShuffle{
				ResultType: 1,
				ResultId:   2,
				Vector1:    3,
				Vector2:    4,
				Components: []uint32{5, 6, 7},
			},
		},
		{
			in: []uint32{0x0006003d, 1, 2, 3, 4, 5},
			want: &OpCompositeConstruct{
				ResultType:   1,
				ResultId:     2,
				Constituents: []Id{3, 4, 5},
			},
		},
		{
			in: []uint32{0x0004003e, 1, 2, 3},
			want: &OpCompositeExtract{
				ResultType: 1,
				ResultId:   2,
				Composite:  3,
			},
		},
		{
			in: []uint32{0x0006003e, 1, 2, 3, 4, 5},
			want: &OpCompositeExtract{
				ResultType: 1,
				ResultId:   2,
				Composite:  3,
				Indices:    []uint32{4, 5},
			},
		},
		{
			in: []uint32{0x0005003f, 1, 2, 3, 4},
			want: &OpCompositeInsert{
				ResultType: 1,
				ResultId:   2,
				Object:     3,
				Composite:  4,
			},
		},
		{
			in: []uint32{0x0007003f, 1, 2, 3, 4, 5, 6},
			want: &OpCompositeInsert{
				ResultType: 1,
				ResultId:   2,
				Object:     3,
				Composite:  4,
				Indices:    []uint32{5, 6},
			},
		},
		{
			in: []uint32{0x00040040, 1, 2, 3},
			want: &OpCopyObject{
				ResultType: 1,
				ResultId:   2,
				Operand:    3,
			},
		},
		{
			in: []uint32{0x00040070, 1, 2, 3},
			want: &OpTranspose{
				ResultType: 1,
				ResultId:   2,
				Matrix:     3,
			},
		},
	} {
		testInstruction(t, st)
	}
}
