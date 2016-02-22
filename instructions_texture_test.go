// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestTexture(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00050043, 1, 2, 3, 4},
			want: &OpSampler{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Filter:     4,
			},
		},
		{
			in: []uint32{0x00050044, 1, 2, 3, 4},
			want: &OpTextureSample{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Bias:       0, // optionl
			},
		},
		{
			in: []uint32{0x00060044, 1, 2, 3, 4, 5},
			want: &OpTextureSample{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Bias:       5, // optionl
			},
		},
		{
			in: []uint32{0x00060045, 1, 2, 3, 4, 5},
			want: &OpTextureSampleDref{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Dref:       5,
			},
		},
		{
			in: []uint32{0x00060046, 1, 2, 3, 4, 5},
			want: &OpTextureSampleLod{
				ResultType:    1,
				ResultId:      2,
				Sampler:       3,
				Coordinate:    4,
				LevelofDetail: 5,
			},
		},
		{
			in: []uint32{0x00050047, 1, 2, 3, 4},
			want: &OpTextureSampleProj{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Bias:       0, // optionl
			},
		},
		{
			in: []uint32{0x00060047, 1, 2, 3, 4, 5},
			want: &OpTextureSampleProj{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Bias:       5, // optionl
			},
		},
		{
			in: []uint32{0x00070048, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleGrad{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Dx:         5,
				Dy:         6,
			},
		},
		{
			in: []uint32{0x00060049, 1, 2, 3, 4, 5},
			want: &OpTextureSampleOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Offset:     5,
				Bias:       0, // optionl
			},
		},
		{
			in: []uint32{0x00070049, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Offset:     5,
				Bias:       6, // optionl
			},
		},
		{
			in: []uint32{0x0006004a, 1, 2, 3, 4, 5},
			want: &OpTextureSampleProjLod{
				ResultType:    1,
				ResultId:      2,
				Sampler:       3,
				Coordinate:    4,
				LevelofDetail: 5,
			},
		},
		{
			in: []uint32{0x0007004b, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleProjGrad{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Dx:         5,
				Dy:         6,
			},
		},
		{
			in: []uint32{0x0007004c, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleLodOffset{
				ResultType:    1,
				ResultId:      2,
				Sampler:       3,
				Coordinate:    4,
				LevelofDetail: 5,
				Offset:        6,
			},
		},
		{
			in: []uint32{0x0006004d, 1, 2, 3, 4, 5},
			want: &OpTextureSampleProjOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Offset:     5,
				Bias:       0, // optionl
			},
		},
		{
			in: []uint32{0x0007004d, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleProjOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Offset:     5,
				Bias:       6, // optionl
			},
		},
		{
			in: []uint32{0x0008004e, 1, 2, 3, 4, 5, 6, 7},
			want: &OpTextureSampleGradOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Dx:         5,
				Dy:         6,
				Offset:     7,
			},
		},
		{
			in: []uint32{0x0007004f, 1, 2, 3, 4, 5, 6},
			want: &OpTextureSampleProjLodOffset{
				ResultType:    1,
				ResultId:      2,
				Sampler:       3,
				Coordinate:    4,
				LevelofDetail: 5,
				Offset:        6,
			},
		},
		{
			in: []uint32{0x00080050, 1, 2, 3, 4, 5, 6, 7},
			want: &OpTextureSampleProjGradOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Dx:         5,
				Dy:         6,
				Offset:     7,
			},
		},
		{
			in: []uint32{0x00060051, 1, 2, 3, 4, 5},
			want: &OpTextureFetchTexel{
				ResultType:    1,
				ResultId:      2,
				Sampler:       3,
				Coordinate:    4,
				LevelofDetail: 5,
			},
		},
		{
			in: []uint32{0x00060052, 1, 2, 3, 4, 5},
			want: &OpTextureFetchTexelOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Offset:     5,
			},
		},
		{
			in: []uint32{0x00060053, 1, 2, 3, 4, 5},
			want: &OpTextureFetchSample{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Sample:     5,
			},
		},
		{
			in: []uint32{0x00050054, 1, 2, 3, 4},
			want: &OpTextureFetchBuffer{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Element:    4,
			},
		},
		{
			in: []uint32{0x00060055, 1, 2, 3, 4, 5},
			want: &OpTextureGather{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Component:  5,
			},
		},
		{
			in: []uint32{0x00070056, 1, 2, 3, 4, 5, 6},
			want: &OpTextureGatherOffset{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Component:  5,
				Offset:     6,
			},
		},
		{
			in: []uint32{0x00070057, 1, 2, 3, 4, 5, 6},
			want: &OpTextureGatherOffsets{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
				Component:  5,
				Offsets:    6,
			},
		},
		{
			in: []uint32{0x00050058, 1, 2, 3, 4},
			want: &OpTextureQuerySizeLod{
				ResultType:    1,
				ResultId:      2,
				Sampler:       3,
				LevelofDetail: 4,
			},
		},
		{
			in: []uint32{0x00040059, 1, 2, 3},
			want: &OpTextureQuerySize{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
			},
		},
		{
			in: []uint32{0x0005005a, 1, 2, 3, 4},
			want: &OpTextureQueryLod{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
				Coordinate: 4,
			},
		},
		{
			in: []uint32{0x0004005b, 1, 2, 3},
			want: &OpTextureQueryLevels{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
			},
		},
		{
			in: []uint32{0x0004005c, 1, 2, 3},
			want: &OpTextureQuerySamples{
				ResultType: 1,
				ResultId:   2,
				Sampler:    3,
			},
		},
	} {
		testInstruction(t, st)
	}
}
