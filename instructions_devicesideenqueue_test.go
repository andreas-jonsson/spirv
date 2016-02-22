// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestDeviceSideEnqueue(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x000700f9, 1, 2, 3, 4, 5, 6},
			want: &OpEnqueueMarker{
				ResultType: 1,
				ResultId:   2,
				Q:          3,
				NumEvents:  4,
				WaitEvents: 5,
				RetEvent:   6,
			},
		},
		{
			in: []uint32{0x000d00fa, 1, 2, 3, uint32(KernelEnqueueFlagNoWait), 5, 6, 7, 8, 9, 10, 11, 12},
			want: &OpEnqueueKernel{
				ResultType: 1,
				ResultId:   2,
				Q:          3,
				Flags:      KernelEnqueueFlagNoWait,
				NDRange:    5,
				NumEvents:  6,
				WaitEvents: 7,
				RetEvent:   8,
				Invoke:     9,
				Param:      10,
				ParamSize:  11,
				ParamAlign: 12,
				LocalSize:  []Id{}, // optional
			},
		},
		{
			in: []uint32{0x000e00fa, 1, 2, 3, uint32(KernelEnqueueFlagWaitWorkGroup), 5, 6, 7, 8, 9, 10, 11, 12, 13},
			want: &OpEnqueueKernel{
				ResultType: 1,
				ResultId:   2,
				Q:          3,
				Flags:      KernelEnqueueFlagWaitWorkGroup,
				NDRange:    5,
				NumEvents:  6,
				WaitEvents: 7,
				RetEvent:   8,
				Invoke:     9,
				Param:      10,
				ParamSize:  11,
				ParamAlign: 12,
				LocalSize:  []Id{13}, // optional
			},
		},
		{
			in: []uint32{0x000500fb, 1, 2, 3, 4},
			want: &OpGetKernelNDrangeSubGroupCount{
				ResultType: 1,
				ResultId:   2,
				NDRange:    3,
				Invoke:     4,
			},
		},
		{
			in: []uint32{0x000500fc, 1, 2, 3, 4},
			want: &OpGetKernelNDrangeMaxSubGroupSize{
				ResultType: 1,
				ResultId:   2,
				NDRange:    3,
				Invoke:     4,
			},
		},
		{
			in: []uint32{0x000400fd, 1, 2, 3},
			want: &OpGetKernelWorkGroupSize{
				ResultType: 1,
				ResultId:   2,
				Invoke:     3,
			},
		},
		{
			in: []uint32{0x000400fe, 1, 2, 3},
			want: &OpGetKernelPreferredWorkGroupSizeMultiple{
				ResultType: 1,
				ResultId:   2,
				Invoke:     3,
			},
		},
		{
			in: []uint32{0x000200ff, 1},
			want: &OpRetainEvent{
				Event: 1,
			},
		},
		{
			in: []uint32{0x00020100, 1},
			want: &OpReleaseEvent{
				Event: 1,
			},
		},
		{
			in: []uint32{0x00030101, 1, 2},
			want: &OpCreateUserEvent{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x00040102, 1, 2, 3},
			want: &OpIsValidEvent{
				ResultType: 1,
				ResultId:   2,
				Event:      3,
			},
		},
		{
			in: []uint32{0x00030103, 1, 2},
			want: &OpSetUserEventStatus{
				Event:  1,
				Status: 2,
			},
		},
		{
			in: []uint32{0x00040104, 1, 2, 3},
			want: &OpCaptureEventProfilingInfo{
				Event: 1,
				Info:  2,
				Value: 3,
			},
		},
		{
			in: []uint32{0x00030105, 1, 2},
			want: &OpGetDefaultQueue{
				ResultType: 1,
				ResultId:   2,
			},
		},
		{
			in: []uint32{0x00060106, 1, 2, 3, 4, 5},
			want: &OpBuildNDRange{
				ResultType:       1,
				ResultId:         2,
				GlobalWorkSize:   3,
				LocalWorkSize:    4,
				GlobalWorkOffset: 5,
			},
		},
	} {
		testInstruction(t, st)
	}
}
