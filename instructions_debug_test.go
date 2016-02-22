// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

func TestDebug(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x0030001, SourceLanguageGLSL, 450},
			want: &OpSource{
				SourceLanguage: SourceLanguageGLSL,
				Version:        450,
			},
		},
		{
			in: []uint32{
				0x00070002,
				0x74736574,
				0x756f7320,
				0x20656372,
				0x65747865,
				0x6f69736e,
				0x0000006e,
			},
			want: &OpSourceExtension{
				Extension: "test source extension",
			},
		},
		{
			in: []uint32{0x00050036, 1, 0x74736574, 0x6d616e5f, 0x00000065},
			want: &OpName{
				Target: 1,
				Name:   "test_name",
			},
		},
		{
			in: []uint32{0x00060037, 1, 2, 0x74736574, 0x6d616e5f, 0x00000065},
			want: &OpMemberName{
				Type:   1,
				Member: 2,
				Name:   "test_name",
			},
		},
		{
			in: []uint32{0x00050038, 1, 0x74736574, 0x72747320, 0x676e69},
			want: &OpString{
				ResultId: 1,
				String:   "test string",
			},
		},
		{
			in: []uint32{0x00050039, 1, 2, 3, 4},
			want: &OpLine{
				Target: 1,
				File:   2,
				Line:   3,
				Column: 4,
			},
		},
	} {
		testInstruction(t, st)
	}
}
