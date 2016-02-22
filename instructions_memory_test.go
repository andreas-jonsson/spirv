// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestMemory(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x0040026, 1, 2, uint32(StorageClassPrivate)},
			want: &OpVariable{
				ResultType:   1,
				ResultId:     2,
				StorageClass: StorageClassPrivate,
				Initializer:  0,
			},
		},
		{
			in: []uint32{0x0050026, 1, 2, uint32(StorageClassPrivate), 3},
			want: &OpVariable{
				ResultType:   1,
				ResultId:     2,
				StorageClass: StorageClassPrivate,
				Initializer:  3,
			},
		},
		{
			in: []uint32{0x0050027, 1, 2, uint32(StorageClassPrivate), 10},
			want: &OpVariableArray{
				ResultType:   1,
				ResultId:     2,
				StorageClass: StorageClassPrivate,
				N:            10,
			},
		},
		{
			in: []uint32{0x004002e, 1, 2, 3},
			want: &OpLoad{
				ResultType: 1,
				ResultId:   2,
				Pointer:    3,
			},
		},
		{
			in: []uint32{0x006002e, 1, 2, 3,
				uint32(MemoryAccessAligned), uint32(MemoryAccessVolatile)},
			want: &OpLoad{
				ResultType: 1,
				ResultId:   2,
				Pointer:    3,
				MemoryAccess: []MemoryAccess{
					MemoryAccessAligned,
					MemoryAccessVolatile,
				},
			},
		},
		{
			in: []uint32{0x003002f, 1, 2},
			want: &OpStore{
				Pointer: 1,
				Object:  2,
			},
		},
		{
			in: []uint32{0x005002f, 1, 2,
				uint32(MemoryAccessAligned), uint32(MemoryAccessVolatile)},
			want: &OpStore{
				Pointer: 1,
				Object:  2,
				MemoryAccess: []MemoryAccess{
					MemoryAccessAligned,
					MemoryAccessVolatile,
				},
			},
		},
		{
			in: []uint32{0x0030041, 1, 2},
			want: &OpCopyMemory{
				Target: 1,
				Source: 2,
			},
		},
		{
			in: []uint32{0x0050041, 1, 2, uint32(MemoryAccessAligned), uint32(MemoryAccessVolatile)},
			want: &OpCopyMemory{
				Target:       1,
				Source:       2,
				MemoryAccess: []MemoryAccess{MemoryAccessAligned, MemoryAccessVolatile},
			},
		},
		{
			in: []uint32{0x0040042, 1, 2, 3},
			want: &OpCopyMemorySized{
				Target: 1,
				Source: 2,
				Size:   3,
			},
		},
		{
			in: []uint32{0x0060042, 1, 2, 3, uint32(MemoryAccessAligned), uint32(MemoryAccessVolatile)},
			want: &OpCopyMemorySized{
				Target:       1,
				Source:       2,
				Size:         3,
				MemoryAccess: []MemoryAccess{MemoryAccessAligned, MemoryAccessVolatile},
			},
		},
		{
			in: []uint32{0x004005d, 1, 2, 3},
			want: &OpAccessChain{
				ResultType: 1,
				ResultId:   2,
				Base:       3,
			},
		},
		{
			in: []uint32{0x006005d, 1, 2, 3, 4, 5},
			want: &OpAccessChain{
				ResultType: 1,
				ResultId:   2,
				Base:       3,
				Indices:    []Id{4, 5},
			},
		},
		{
			in: []uint32{0x004005e, 1, 2, 3},
			want: &OpInboundsAccessChain{
				ResultType: 1,
				ResultId:   2,
				Base:       3,
			},
		},
		{
			in: []uint32{0x006005e, 1, 2, 3, 4, 5},
			want: &OpInboundsAccessChain{
				ResultType: 1,
				ResultId:   2,
				Base:       3,
				Indices:    []Id{4, 5},
			},
		},
		{
			in: []uint32{0x0050079, 1, 2, 3, 4},
			want: &OpArraylength{
				ResultType: 1,
				ResultId:   2,
				Structure:  3,
				Member:     4,
			},
		},
		{
			in: []uint32{0x00600be, 1, 2, 3, 4, 5},
			want: &OpImagePointer{
				ResultType: 1,
				ResultId:   2,
				Image:      3,
				Coordinate: 4,
				Sample:     5,
			},
		},
		{
			in: []uint32{0x00400e9, 1, 2, 3},
			want: &OpGenericPtrMemSemantics{
				ResultType: 1,
				ResultId:   2,
				Ptr:        3,
			},
		},
	} {
		testInstruction(t, st)
	}
}
