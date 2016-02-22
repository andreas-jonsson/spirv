// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestBarrier(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00200bc, ExecutionScopeDevice},
			want: &OpControlBarrier{
				ExecutionScope: ExecutionScopeDevice,
			},
		},
		{
			in: []uint32{
				0x00300bd,
				ExecutionScopeDevice,
				MemorySemanticSequentiallyConsistent,
			},
			want: &OpMemoryBarrier{
				ExecutionScope: ExecutionScopeDevice,
				MemorySemantic: MemorySemanticSequentiallyConsistent,
			},
		},
	} {
		testInstruction(t, st)
	}
}
