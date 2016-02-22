// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"testing"
)

func TestGroup(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in:  []uint32{0x000900db, 1, 2, ExecutionScopeDevice, 3, 4, 5, 6, 7},
			err: errors.New("OpAsyncGroupCopy: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000900db, 1, 2, ExecutionScopeSubgroup, 3, 4, 5, 6, 7},
			want: &OpAsyncGroupCopy{
				ResultType:  1,
				ResultId:    2,
				Scope:       ExecutionScopeSubgroup,
				Destination: 3,
				Source:      4,
				NumElements: 5,
				Stride:      6,
				Event:       7,
			},
		},
		{
			in:  []uint32{0x000600dc, 1, 2, ExecutionScopeDevice, 3, 4},
			err: errors.New("OpWaitGroupEvents: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600dc, 1, 2, ExecutionScopeSubgroup, 3, 4},
			want: &OpWaitGroupEvents{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				NumEvents:  3,
				EventsList: 4,
			},
		},
		{
			in:  []uint32{0x000500dd, 1, 2, ExecutionScopeDevice, 3},
			err: errors.New("OpGroupAll: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000500dd, 1, 2, ExecutionScopeSubgroup, 3},
			want: &OpGroupAll{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Predicate:  3,
			},
		},
		{
			in:  []uint32{0x000500de, 1, 2, ExecutionScopeDevice, 3},
			err: errors.New("OpGroupAny: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000500de, 1, 2, ExecutionScopeSubgroup, 3},
			want: &OpGroupAny{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Predicate:  3,
			},
		},
		{
			in:  []uint32{0x000600df, 1, 2, ExecutionScopeDevice, 3, 4},
			err: errors.New("OpGroupBroadcast: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600df, 1, 2, ExecutionScopeSubgroup, 3, 4},
			want: &OpGroupBroadcast{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Value:      3,
				LocalId:    4,
			},
		},
		{
			in:  []uint32{0x000600e0, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupIAdd: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e0, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupIAdd{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e1, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupFAdd: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e1, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupFAdd{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e2, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupFMin: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e2, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupFMin{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e3, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupUMin: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e3, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupUMin{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e4, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupSMin: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e4, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupSMin{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e5, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupFMax: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e5, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupFMax{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e6, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupUMax: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e6, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupUMax{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
		{
			in:  []uint32{0x000600e7, 1, 2, ExecutionScopeDevice, GroupOperationReduce, 4},
			err: errors.New("OpGroupSMax: Scope must be Subgroup or Workgroup"),
		},
		{
			in: []uint32{0x000600e7, 1, 2, ExecutionScopeSubgroup, GroupOperationReduce, 4},
			want: &OpGroupSMax{
				ResultType: 1,
				ResultId:   2,
				Scope:      ExecutionScopeSubgroup,
				Operation:  GroupOperationReduce,
				X:          4,
			},
		},
	} {
		testInstruction(t, st)
	}
}
