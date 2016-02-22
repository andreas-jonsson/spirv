// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestRelational(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x0040062, 1, 2, 3},
			want: &OpAny{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
			},
		},
		{
			in: []uint32{0x0040063, 1, 2, 3},
			want: &OpAll{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
			},
		},
		{
			in: []uint32{0x0040071, 1, 2, 3},
			want: &OpIsNan{
				ResultType: 1,
				ResultId:   2,
				X:          3,
			},
		},
		{
			in: []uint32{0x0040072, 1, 2, 3},
			want: &OpIsInf{
				ResultType: 1,
				ResultId:   2,
				X:          3,
			},
		},
		{
			in: []uint32{0x0040073, 1, 2, 3},
			want: &OpIsFinite{
				ResultType: 1,
				ResultId:   2,
				X:          3,
			},
		},
		{
			in: []uint32{0x0040074, 1, 2, 3},
			want: &OpIsNormal{
				ResultType: 1,
				ResultId:   2,
				X:          3,
			},
		},
		{
			in: []uint32{0x0040075, 1, 2, 3},
			want: &OpSignBitSet{
				ResultType: 1,
				ResultId:   2,
				X:          3,
			},
		},
		{
			in: []uint32{0x0040076, 1, 2, 3},
			want: &OpLessOrGreater{
				ResultType: 1,
				ResultId:   2,
				X:          3,
			},
		},
		{
			in: []uint32{0x0050077, 1, 2, 3, 4},
			want: &OpOrdered{
				ResultType: 1,
				ResultId:   2,
				X:          3,
				Y:          4,
			},
		},
		{
			in: []uint32{0x0050078, 1, 2, 3, 4},
			want: &OpUnordered{
				ResultType: 1,
				ResultId:   2,
				X:          3,
				Y:          4,
			},
		},
		{
			in: []uint32{0x0050092, 1, 2, 3, 4},
			want: &OpLogicalOr{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0050093, 1, 2, 3, 4},
			want: &OpLogicalXor{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0050094, 1, 2, 3, 4},
			want: &OpLogicalAnd{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0060098, 1, 2, 3, 4, 5},
			want: &OpSelect{
				ResultType: 1,
				ResultId:   2,
				Condition:  3,
				Object1:    4,
				Object2:    5,
			},
		},
		{
			in: []uint32{0x0050099, 1, 2, 3, 4},
			want: &OpIEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x005009a, 1, 2, 3, 4},
			want: &OpFOrdEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x005009b, 1, 2, 3, 4},
			want: &OpFUnordEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x005009c, 1, 2, 3, 4},
			want: &OpINotEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x005009d, 1, 2, 3, 4},
			want: &OpFOrdNotEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x005009e, 1, 2, 3, 4},
			want: &OpFUnordNotEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x005009f, 1, 2, 3, 4},
			want: &OpULessThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a0, 1, 2, 3, 4},
			want: &OpSLessThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a1, 1, 2, 3, 4},
			want: &OpFOrdLessThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a2, 1, 2, 3, 4},
			want: &OpFUnordLessThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a3, 1, 2, 3, 4},
			want: &OpUGreaterThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a4, 1, 2, 3, 4},
			want: &OpSGreaterThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a5, 1, 2, 3, 4},
			want: &OpFOrdGreaterThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a6, 1, 2, 3, 4},
			want: &OpFUnordGreaterThan{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a7, 1, 2, 3, 4},
			want: &OpULessThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a8, 1, 2, 3, 4},
			want: &OpSLessThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500a9, 1, 2, 3, 4},
			want: &OpFOrdLessThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500aa, 1, 2, 3, 4},
			want: &OpFUnordLessThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500ab, 1, 2, 3, 4},
			want: &OpUGreaterThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500ac, 1, 2, 3, 4},
			want: &OpSGreaterThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500ad, 1, 2, 3, 4},
			want: &OpFOrdGreaterThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
		{
			in: []uint32{0x00500ae, 1, 2, 3, 4},
			want: &OpFUnordGreaterThanEqual{
				ResultType: 1,
				ResultId:   2,
				Object1:    3,
				Object2:    4,
			},
		},
	} {
		testInstruction(t, st)
	}
}
