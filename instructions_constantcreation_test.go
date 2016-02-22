// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"testing"
)

func TestConstantCreation(t *testing.T) {
	for _, st := range []InstructionTest{

		{
			in: []uint32{0x003001b, 1, 2},
			want: &OpConstantTrue{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x003001c, 1, 2},
			want: &OpConstantFalse{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x003001d, 1, 2},
			want: &OpConstant{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x006001d, 1, 2, 3, 4, 5},
			want: &OpConstant{
				ResultType: 1,
				ResultId:   2,
				Value:      []uint32{3, 4, 5},
			},
		},
		{
			in: []uint32{0x003001e, 1, 2},
			want: &OpConstantComposite{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x006001e, 1, 2, 3, 4, 5},
			want: &OpConstantComposite{
				ResultType:   1,
				ResultId:     2,
				Constituents: []Id{3, 4, 5},
			},
		},
		{
			in: []uint32{0x006001f, 1, 2,
				uint32(AddressingModePhysical64), 4, uint32(SamplerFilterModeNearest)},
			err: errors.New("OpConstantSampler: Param must be 0 or 1"),
		},
		{
			in: []uint32{0x006001f, 1, 2,
				uint32(AddressingModePhysical64), 1, uint32(SamplerFilterModeNearest)},
			want: &OpConstantSampler{
				ResultType: 1,
				ResultId:   2,
				Mode:       AddressingModePhysical64,
				Param:      1,
				Filter:     SamplerFilterModeNearest,
			},
		},
		{
			in: []uint32{0x0030020, 1, 2},
			want: &OpConstantNullPointer{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x0030021, 1, 2},
			want: &OpConstantNullObject{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x0030022, 1, 2},
			want: &OpSpecConstantTrue{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x0030023, 1, 2},
			want: &OpSpecConstantFalse{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x0030024, 1, 2},
			want: &OpSpecConstant{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x0060024, 1, 2, 3, 4, 5},
			want: &OpSpecConstant{
				ResultType: 1,
				ResultId:   2,
				Value:      []uint32{3, 4, 5},
			},
		},
		{
			in: []uint32{0x0030025, 1, 2},
			want: &OpSpecConstantComposite{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x0060025, 1, 2, 3, 4, 5},
			want: &OpSpecConstantComposite{
				ResultType:   1,
				ResultId:     2,
				Constituents: []Id{3, 4, 5},
			},
		},
	} {
		testInstruction(t, st)
	}
}
