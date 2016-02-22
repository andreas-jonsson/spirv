// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpExtension defines the OpExtension instruction.
//
// It declares use of an extension to SPIR-V. This allows validation of
// additional instructions, tokens, semantics, etc
type OpExtension struct {
	Name String
}

func (c *OpExtension) Opcode() uint32 { return opcodeExtension }
func (c *OpExtension) Optional() bool { return false }
func (c *OpExtension) Verify() error  { return nil }

// OpExtInstImport defines the OpExtInstImport instruction.
//
// It defines an import of an extended set of instructions.
// It can later be referenced by the Result <id>
type OpExtInstImport struct {
	ResultId Id
	Name     String
}

func (c *OpExtInstImport) Opcode() uint32 { return opcodeExtInstImport }
func (c *OpExtInstImport) Optional() bool { return false }
func (c *OpExtInstImport) Verify() error  { return nil }

// OpExtInst defines an instruction in an imported set of extended instructions.
type OpExtInst struct {
	ResultType  Id
	ResultId    Id
	Set         Id     // Result of an OpExtInstImport instruction.
	Instruction uint32 // Enumerant of the instruction to execute within the extended instruction Set.
	Operands    []Id   // Operands to the extended instruction.
}

func (c *OpExtInst) Opcode() uint32 { return opcodeExtInst }
func (c *OpExtInst) Optional() bool { return false }
func (c *OpExtInst) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpExtension{} })
	bind(func() Instruction { return &OpExtInstImport{} })
	bind(func() Instruction { return &OpExtInst{} })
}
