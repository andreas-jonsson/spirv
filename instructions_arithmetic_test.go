// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestArithmetic(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x0004005f, 1, 2, 3},
			want: &OpSNegate{
				ResultType: 1,
				ResultId:   2,
				Operand:    3,
			},
		},
		{
			in: []uint32{0x00040060, 1, 2, 3},
			want: &OpFNegate{
				ResultType: 1,
				ResultId:   2,
				Operand:    3,
			},
		},
		{
			in: []uint32{0x00040061, 1, 2, 3},
			want: &OpNot{
				ResultType: 1,
				ResultId:   2,
				Operand:    3,
			},
		},
		{
			in: []uint32{0x0005007a, 1, 2, 3, 4},
			want: &OpIAdd{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0005007b, 1, 2, 3, 4},
			want: &OpFAdd{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0005007c, 1, 2, 3, 4},
			want: &OpISub{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0005007d, 1, 2, 3, 4},
			want: &OpFSub{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0005007e, 1, 2, 3, 4},
			want: &OpIMul{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x0005007f, 1, 2, 3, 4},
			want: &OpFMul{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050080, 1, 2, 3, 4},
			want: &OpUDiv{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050081, 1, 2, 3, 4},
			want: &OpSDiv{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050082, 1, 2, 3, 4},
			want: &OpFDiv{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050083, 1, 2, 3, 4},
			want: &OpUMod{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050084, 1, 2, 3, 4},
			want: &OpSRem{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050085, 1, 2, 3, 4},
			want: &OpSMod{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050086, 1, 2, 3, 4},
			want: &OpFRem{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050087, 1, 2, 3, 4},
			want: &OpFMod{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050088, 1, 2, 3, 4},
			want: &OpVectorTimesScalar{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
				Scalar:     4,
			},
		},
		{
			in: []uint32{0x00050089, 1, 2, 3, 4},
			want: &OpMatrixTimesScalar{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
				Scalar:     4,
			},
		},
		{
			in: []uint32{0x0005008a, 1, 2, 3, 4},
			want: &OpVectorTimesMatrix{
				ResultType: 1,
				ResultId:   2,
				Vector:     3,
				Matrix:     4,
			},
		},
		{
			in: []uint32{0x0005008b, 1, 2, 3, 4},
			want: &OpMatrixTimesVector{
				ResultType: 1,
				ResultId:   2,
				Matrix:     3,
				Vector:     4,
			},
		},
		{
			in: []uint32{0x0005008c, 1, 2, 3, 4},
			want: &OpMatrixTimesMatrix{
				ResultType: 1,
				ResultId:   2,
				Left:       3,
				Right:      4,
			},
		},
		{
			in: []uint32{0x0005008d, 1, 2, 3, 4},
			want: &OpOuterProduct{
				ResultType: 1,
				ResultId:   2,
				Vector1:    3,
				Vector2:    4,
			},
		},
		{
			in: []uint32{0x0005008e, 1, 2, 3, 4},
			want: &OpDot{
				ResultType: 1,
				ResultId:   2,
				Vector1:    3,
				Vector2:    4,
			},
		},
		{
			in: []uint32{0x0005008f, 1, 2, 3, 4},
			want: &OpShiftRightLogical{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050090, 1, 2, 3, 4},
			want: &OpShiftRightArithmetic{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050091, 1, 2, 3, 4},
			want: &OpShiftLeftLogical{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050095, 1, 2, 3, 4},
			want: &OpBitwiseOr{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050096, 1, 2, 3, 4},
			want: &OpBitwiseXor{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
		{
			in: []uint32{0x00050097, 1, 2, 3, 4},
			want: &OpBitwiseAnd{
				ResultType: 1,
				ResultId:   2,
				Operand1:   3,
				Operand2:   4,
			},
		},
	} {
		testInstruction(t, st)
	}
}
