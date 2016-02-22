// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestAtomic(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00300bf, 1, 2},
			want: &OpAtomicInit{
				Pointer: 1,
				Value:   2,
			},
		},
		{
			in: []uint32{
				0x00600c0, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
			},
			want: &OpAtomicLoad{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
			},
		},
		{
			in: []uint32{
				0x00500c1, 1,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				2,
			},
			want: &OpAtomicStore{
				Pointer:        1,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          2,
			},
		},
		{
			in: []uint32{
				0x00700c2, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicExchange{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00800c3, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
				5,
			},
			want: &OpAtomicCompareExchange{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
				Comparator:     5,
			},
		},
		{
			in: []uint32{
				0x00800c4, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
				5,
			},
			want: &OpAtomicCompareExchangeWeak{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
				Comparator:     5,
			},
		},
		{
			in: []uint32{
				0x00600c5, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
			},
			want: &OpAtomicIIncrement{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
			},
		},
		{
			in: []uint32{
				0x00600c6, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
			},
			want: &OpAtomicIDecrement{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
			},
		},
		{
			in: []uint32{
				0x00700c7, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicIAdd{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00700c8, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicISub{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00700c9, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicUMin{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00700ca, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicUMax{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00700cb, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicAnd{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00700cc, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicOr{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
		{
			in: []uint32{
				0x00700cd, 1, 2, 3,
				uint32(ExecutionScopeCrossDevice),
				uint32(MemorySemanticAtomicCounterMemory),
				4,
			},
			want: &OpAtomicXor{
				ResultType:     1,
				ResultId:       2,
				Pointer:        3,
				ExecutionScope: ExecutionScopeCrossDevice,
				MemorySemantic: MemorySemanticAtomicCounterMemory,
				Value:          4,
			},
		},
	} {
		testInstruction(t, st)
	}
}
