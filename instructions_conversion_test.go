// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"testing"
)

func TestConversions(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00040064, 1, 2, 3},
			want: &OpConvertFToU{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x00040065, 1, 2, 3},
			want: &OpConvertFToS{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x00040066, 1, 2, 3},
			want: &OpConvertSToF{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x00040067, 1, 2, 3},
			want: &OpConvertUToF{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x00040068, 1, 2, 3},
			want: &OpUConvert{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x00040069, 1, 2, 3},
			want: &OpSConvert{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x0004006a, 1, 2, 3},
			want: &OpFConvert{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x0004006b, 1, 2, 3},
			want: &OpConvertPtrToU{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x0004006c, 1, 2, 3},
			want: &OpConvertUToPtr{
				ResultType: 1,
				ResultId:   2,
				Value:      3,
			},
		},
		{
			in: []uint32{0x0004006d, 1, 2, 3},
			want: &OpPtrCastToGeneric{
				ResultType: 1,
				ResultId:   2,
				Source:     3,
			},
		},
		{
			in: []uint32{0x0004006e, 1, 2, 3},
			want: &OpGenericCastToPtr{
				ResultType: 1,
				ResultId:   2,
				Source:     3,
			},
		},
		{
			in: []uint32{0x0004006f, 1, 2, 3},
			want: &OpBitcast{
				ResultType: 1,
				ResultId:   2,
				Operand:    3,
			},
		},
		{
			in: []uint32{0x000500e8, 1, 2, 3, StorageClassFunction},
			want: &OpGenericCastToPtrExplicit{
				ResultType: 1,
				ResultId:   2,
				SourcePtr:  3,
				Storage:    StorageClassFunction,
			},
		},
	} {
		testInstruction(t, st)
	}
}
