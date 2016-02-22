// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestFlowControl(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00050030, 1, 2, 3, 4},
			want: &OpPhi{
				ResultType: 1,
				ResultId:   2,
				Operands:   []Id{3, 4},
			},
		},
		{
			in: []uint32{0x000300ce, 1, 2},
			want: &OpLoopMerge{
				Label:       1,
				LoopControl: 2,
			},
		},
		{
			in: []uint32{0x000300cf, 1, 2},
			want: &OpSelectionMerge{
				Label:            1,
				SelectionControl: 2,
			},
		},
		{
			in: []uint32{0x000200d0, 1},
			want: &OpLabel{
				ResultId: 1,
			},
		},
		{
			in: []uint32{0x000200d1, 1},
			want: &OpBranch{
				TargetLabel: 1,
			},
		},
		{
			in: []uint32{0x000400d2, 1, 2, 3},
			want: &OpBranchConditional{
				Condition:  1,
				TrueLabel:  2,
				FalseLabel: 3,
			},
		},
		{
			in: []uint32{0x000600d2, 1, 2, 3, 4, 5},
			want: &OpBranchConditional{
				Condition:     1,
				TrueLabel:     2,
				FalseLabel:    3,
				BranchWeights: []uint32{4, 5},
			},
		},
		{
			in: []uint32{0x000300d3, 1, 2},
			want: &OpSwitch{
				Selector: 1,
				Default:  2,
			},
		},
		{
			in: []uint32{0x000500d3, 1, 2, 3, 4},
			want: &OpSwitch{
				Selector: 1,
				Default:  2,
				Target:   []uint32{3, 4},
			},
		},
		{
			in:   []uint32{0x000100d4},
			want: &OpKill{},
		},
		{
			in:   []uint32{0x000100d5},
			want: &OpReturn{},
		},
		{
			in: []uint32{0x000200d6, 1},
			want: &OpReturnValue{
				Value: 1,
			},
		},
		{
			in:   []uint32{0x000100d7},
			want: &OpUnreachable{},
		},
		{
			in: []uint32{0x000300d8, 1, 2},
			want: &OpLifetimeStart{
				Object:       1,
				MemoryAmount: 2,
			},
		},
		{
			in: []uint32{0x000300d9, 1, 2},
			want: &OpLifetimeStop{
				Object:       1,
				MemoryAmount: 2,
			},
		},
	} {
		testInstruction(t, st)
	}
}
