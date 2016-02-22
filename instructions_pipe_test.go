// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestPipe(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x000500ea, 1, 2, 3, 4},
			want: &OpReadPipe{
				ResultType: 1,
				ResultId:   2,
				P:          3,
				Ptr:        4,
			},
		},
		{
			in: []uint32{0x000500eb, 1, 2, 3, 4},
			want: &OpWritePipe{
				ResultType: 1,
				ResultId:   2,
				P:          3,
				Ptr:        4,
			},
		},
		{
			in: []uint32{0x000700ec, 1, 2, 3, 4, 5, 6},
			want: &OpReservedReadPipe{
				ResultType: 1,
				ResultId:   2,
				P:          3,
				ReserveId:  4,
				Index:      5,
				Ptr:        6,
			},
		},
		{
			in: []uint32{0x000700ed, 1, 2, 3, 4, 5, 6},
			want: &OpReservedWritePipe{
				ResultType: 1,
				ResultId:   2,
				P:          3,
				ReserveId:  4,
				Index:      5,
				Ptr:        6,
			},
		},
		{
			in: []uint32{0x000500ee, 1, 2, 3, 4},
			want: &OpReserveReadPipePackets{
				ResultType: 1,
				ResultId:   2,
				P:          3,
				NumPackets: 4,
			},
		},
		{
			in: []uint32{0x000500ef, 1, 2, 3, 4},
			want: &OpReserveWritePipePackets{
				ResultType: 1,
				ResultId:   2,
				P:          3,
				NumPackets: 4,
			},
		},
		{
			in: []uint32{0x000300f0, 1, 2},
			want: &OpCommitReadPipe{
				P:         1,
				ReserveId: 2,
			},
		},
		{
			in: []uint32{0x000300f1, 1, 2},
			want: &OpCommitWritePipe{
				P:         1,
				ReserveId: 2,
			},
		},
		{
			in: []uint32{0x000400f2, 1, 2, 3},
			want: &OpIsValidReserveId{
				ResultType: 1,
				ResultId:   2,
				ReserveId:  3,
			},
		},
		{
			in: []uint32{0x000400f3, 1, 2, 3},
			want: &OpGetNumPipePackets{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000400f4, 1, 2, 3},
			want: &OpGetMaxPipePackets{
				ResultType: 1,
				ResultId:   2,
				P:          3,
			},
		},
		{
			in: []uint32{0x000600f5, 1, 2, 3, 4, 5},
			want: &OpGroupReserveReadPipePackets{
				ResultType: 1,
				ResultId:   2,
				Scope:      3,
				P:          4,
				NumPackets: 5,
			},
		},
		{
			in: []uint32{0x000600f6, 1, 2, 3, 4, 5},
			want: &OpGroupReserveWritePipePackets{
				ResultType: 1,
				ResultId:   2,
				Scope:      3,
				P:          4,
				NumPackets: 5,
			},
		},
		{
			in: []uint32{0x000400f7, 1, 2, 3},
			want: &OpGroupCommitReadPipe{
				Scope:     1,
				P:         2,
				ReserveId: 3,
			},
		},
		{
			in: []uint32{0x000400f8, 1, 2, 3},
			want: &OpGroupCommitWritePipe{
				Scope:     1,
				P:         2,
				ReserveId: 3,
			},
		},
	} {
		testInstruction(t, st)
	}
}
