// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestModeSetting(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{
				0x00030005,
				AddressingModePhysical32,
				MemoryModelGLSL450,
			},
			want: &OpMemoryModel{
				AddressingModel: AddressingModePhysical32,
				MemoryModel:     MemoryModelGLSL450,
			},
		},
		{
			in: []uint32{0x00030006, ExecutionModelFragment, 0x7f},
			want: &OpEntryPoint{
				ExecutionModel: ExecutionModelFragment,
				ResultId:       0x7f,
			},
		},
		{
			in: []uint32{0x00060007, 0x7f, ExecutionModeLocalSize, 1, 2, 3},
			want: &OpExecutionMode{
				EntryPoint: 0x7f,
				Mode:       ExecutionModeLocalSize,
				Argv:       []uint32{1, 2, 3},
			},
		},
		{
			in: []uint32{
				0x000700da,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: &OpCompileFlag{
				Flag: "test source extension",
			},
		},
	} {
		testInstruction(t, st)
	}
}
