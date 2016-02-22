// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"testing"
)

func TestAnnotations(t *testing.T) {
	for _, st := range []InstructionTest{
		{
			in: []uint32{0x00020031, 0xff},
			want: &OpDecorationGroup{
				ResultId: 0xff,
			},
		},
		{
			in:  []uint32{0x00060032, 1, DecorationNoStaticUse, 2, 3, 4},
			err: errors.New("OpDecorate: extraneous arguments for Decoration(26)"),
		},
		{
			in:  []uint32{0x00040032, 1, DecorationLinkageType, 2},
			err: errors.New("OpDecorate: invalid LinkageType value"),
		},
		{
			in:  []uint32{0x00040032, 1, DecorationLinkageType, LinkageTypeExport},
			err: errors.New("OpDecorate: invalid LinkageType value"),
			want: &OpDecorate{
				Target:     1,
				Decoration: DecorationLinkageType,
				Argv:       []uint32{LinkageTypeExport},
			},
		},
		{
			in:  []uint32{0x00060033, 1, 2, DecorationNoStaticUse, 3, 4},
			err: errors.New("OpMemberDecorate: extraneous arguments for Decoration(26)"),
		},
		{
			in:  []uint32{0x00050033, 1, 2, DecorationLinkageType, 3},
			err: errors.New("OpMemberDecorate: invalid LinkageType value"),
		},
		{
			in: []uint32{0x00050033, 1, 2, DecorationLinkageType, LinkageTypeExport},
			want: &OpMemberDecorate{
				StructType: 1,
				Member:     2,
				Decoration: DecorationLinkageType,
				Argv:       []uint32{LinkageTypeExport},
			},
		},
		{
			in: []uint32{0x00060034, 1, 2, 3, 4, 5},
			want: &OpGroupDecorate{
				Group:   1,
				Targets: []Id{2, 3, 4, 5},
			},
		},
		{
			in: []uint32{0x00060035, 1, 2, 3, 4, 5},
			want: &OpGroupMemberDecorate{
				Group:   1,
				Targets: []Id{2, 3, 4, 5},
			},
		},
	} {
		testInstruction(t, st)
	}
}
