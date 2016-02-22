// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// OpDecorationGroup represents a collector of decorations from OpDecorate
// instructions.
//
// All such instructions must precede this instruction. Subsequent OpGroupDecorate
// and OpGroupMemberDecorate instructions can consume the Result <id> to apply
// multiple decorations to multiple target <id>s. Those are the only
// instructions allowed to consume the Result <id>.
type OpDecorationGroup struct {
	ResultId Id
}

func (c *OpDecorationGroup) Opcode() uint32 { return opcodeDecorationGroup }
func (c *OpDecorationGroup) Optional() bool { return false }
func (c *OpDecorationGroup) Verify() error  { return nil }

// OpDecorate represents the OpDecorate instruction.
// It adds a decoration to another <id>.
type OpDecorate struct {
	// Target is the <id> to decorate. It can potentially be any <id> that
	// is a forward reference. A set of decorations can be grouped together
	// by having multiple OpDecorate instructions target the same
	// OpDecorationGroup instruction.
	Target Id

	// The decoration type to apply.
	Decoration Decoration

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpDecorate) Opcode() uint32 { return opcodeDecorate }
func (c *OpDecorate) Optional() bool { return false }
func (c *OpDecorate) Verify() error {
	argc := len(c.Argv)

	switch c.Decoration {
	case DecorationStream, DecorationLocation, DecorationComponent,
		DecorationIndex, DecorationBinding, DecorationOffset,
		DecorationAlignment, DecorationXfbBuffer, DecorationStride,
		DecorationBuiltIn, DecorationFuncParamAttr, DecorationFPRoundingMode,
		DecorationFPFastMathMode, DecorationSpecId:
		if argc != 1 {
			return fmt.Errorf("OpDecorate: Decoration(%d) must have 1 argument", c.Decoration)
		}

		return nil

	case DecorationLinkageType:
		if argc != 1 {
			return fmt.Errorf("OpDecorate: Decoration(%d) must have 1 argument", c.Decoration)
		}

		err := LinkageType(c.Argv[0]).Verify()
		if err != nil {
			return fmt.Errorf("OpDecorate: %v", err)
		}

		return nil
	}

	if argc > 0 {
		return fmt.Errorf("OpDecorate: extraneous arguments for Decoration(%d)",
			c.Decoration)
	}

	return nil
}

// OpMemberDecorate represents the OpMemberDecorate instruction.
// It adds a decoration to a member of a structure type.
type OpMemberDecorate struct {
	// The <id> of a type from OpTypeStruct.
	StructType Id

	// Member is the number of the member to decorate in the structure.
	// The first member is member 0, the next is member 1, . . .
	Member uint32

	// The decoration type to apply.
	Decoration Decoration

	// Optional list of decoration arguments.
	//
	// Refer to the Decoration constant documentation for more details
	// on which values require which arguments.
	Argv []uint32
}

func (c *OpMemberDecorate) Opcode() uint32 { return opcodeMemberDecorate }
func (c *OpMemberDecorate) Optional() bool { return false }
func (c *OpMemberDecorate) Verify() error {
	argc := len(c.Argv)

	switch c.Decoration {
	case DecorationStream, DecorationLocation, DecorationComponent,
		DecorationIndex, DecorationBinding, DecorationOffset,
		DecorationAlignment, DecorationXfbBuffer, DecorationStride,
		DecorationBuiltIn, DecorationFuncParamAttr, DecorationFPRoundingMode,
		DecorationFPFastMathMode, DecorationSpecId:
		if argc != 1 {
			return fmt.Errorf("OpMemberDecorate: Decoration(%d) must have 1 argument", c.Decoration)
		}

		return nil

	case DecorationLinkageType:
		if argc != 1 {
			return fmt.Errorf("OpMemberDecorate: Decoration(%d) must have 1 argument", c.Decoration)
		}

		err := LinkageType(c.Argv[0]).Verify()
		if err != nil {
			return fmt.Errorf("OpMemberDecorate: %v", err)
		}

		return nil
	}

	if argc > 0 {
		return fmt.Errorf("OpMemberDecorate: extraneous arguments for Decoration(%d)",
			c.Decoration)
	}

	return nil
}

// OpGroupDecorate represents the OpGroupDecorate instruction.
// It adds a group of decorations to another <id>.
type OpGroupDecorate struct {
	// Decoration group is the <id> of an OpDecorationGroup instruction.
	Group Id

	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []Id
}

func (c *OpGroupDecorate) Opcode() uint32 { return opcodeGroupDecorate }
func (c *OpGroupDecorate) Optional() bool { return false }
func (c *OpGroupDecorate) Verify() error  { return nil }

// OpGroupMemberDecorate represents the OpGroupMemberDecorate instruction.
// It adds a decoration to a member of a structure type
type OpGroupMemberDecorate struct {
	// The <id> of a OpDecorationGroup instruction.
	Group Id

	// Targets are the target <id>s to decorate with the groups of decorations.
	Targets []Id
}

func (c *OpGroupMemberDecorate) Opcode() uint32 { return opcodeGroupMemberDecorate }
func (c *OpGroupMemberDecorate) Optional() bool { return false }
func (c *OpGroupMemberDecorate) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpDecorationGroup{} })
	bind(func() Instruction { return &OpDecorate{} })
	bind(func() Instruction { return &OpMemberDecorate{} })
	bind(func() Instruction { return &OpGroupDecorate{} })
	bind(func() Instruction { return &OpGroupMemberDecorate{} })
}
