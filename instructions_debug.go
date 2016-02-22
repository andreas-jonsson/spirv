// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSource defines the OpSource instruction.
//
// It documents what source language a module was translated from.
// This has no semantic impact and can safely be removed from a module.
type OpSource struct {
	SourceLanguage SourceLanguage
	Version        uint32
}

func (c *OpSource) Opcode() uint32 { return opcodeSource }
func (c *OpSource) Optional() bool { return true }
func (c *OpSource) Verify() error  { return nil }

// OpSourceExtension defines optional extensions used within the source language.
//
// It documents an extension to the source language. This has no semantic
// impact and can safely be removed from a module.
type OpSourceExtension struct {
	Extension String
}

func (c *OpSourceExtension) Opcode() uint32 { return opcodeSourceExtension }
func (c *OpSourceExtension) Optional() bool { return true }
func (c *OpSourceExtension) Verify() error  { return nil }

// OpName defines the OpName instruction.
//
// It names a result ID.
// This has no semantic impact and can safely be removed from a module.
type OpName struct {
	Target Id
	Name   String
}

func (c *OpName) Opcode() uint32 { return opcodeName }
func (c *OpName) Optional() bool { return true }
func (c *OpName) Verify() error  { return nil }

// OpMemberName defines the OpMemberName instruction.
//
// It names a member of a structure type.
// This has no semantic impact and can safely be removed from a module.
type OpMemberName struct {
	Type   Id
	Member uint32
	Name   String
}

func (c *OpMemberName) Opcode() uint32 { return opcodeMemberName }
func (c *OpMemberName) Optional() bool { return true }
func (c *OpMemberName) Verify() error  { return nil }

// OpString defines the OpString instruction.
//
// It names a string for use with other debug instructions.
// This has no semantic impact and can safely be removed from a module.
type OpString struct {
	ResultId Id
	String   String
}

func (c *OpString) Opcode() uint32 { return opcodeString }
func (c *OpString) Optional() bool { return true }
func (c *OpString) Verify() error  { return nil }

// OpLine defines the OpLine instruction.
//
// It adds source-level location information.
// This has no semantic impact and can safely be removed from a module.
type OpLine struct {
	Target Id
	File   Id
	Line   uint32
	Column uint32
}

func (c *OpLine) Opcode() uint32 { return opcodeLine }
func (c *OpLine) Optional() bool { return true }
func (c *OpLine) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpSource{} })
	bind(func() Instruction { return &OpSourceExtension{} })
	bind(func() Instruction { return &OpName{} })
	bind(func() Instruction { return &OpMemberName{} })
	bind(func() Instruction { return &OpString{} })
	bind(func() Instruction { return &OpLine{} })
}
