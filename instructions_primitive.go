// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpEmitVertex emits the current values of all output variables to the
// current output primitive. After execution, the values of all output
// variables are undefined.
type OpEmitVertex struct{}

func (c *OpEmitVertex) Opcode() uint32 { return opcodeEmitVertex }
func (c *OpEmitVertex) Optional() bool { return false }
func (c *OpEmitVertex) Verify() error  { return nil }

// OpEndPrimitive finishes the current primitive and starts a new one.
// No vertex is emitted.
//
// This instruction can only be used when only one stream is present.
type OpEndPrimitive struct{}

func (c *OpEndPrimitive) Opcode() uint32 { return opcodeEndPrimitive }
func (c *OpEndPrimitive) Optional() bool { return false }
func (c *OpEndPrimitive) Verify() error  { return nil }

// OpEmitStreamVertex emits the current values of all output variables to
// the current output primitive. After execution, the values of all output
// variables are undefined.
//
// This instruction can only be used when multiple streams are present
type OpEmitStreamVertex struct {
	Stream Id
}

func (c *OpEmitStreamVertex) Opcode() uint32 { return opcodeEmitStreamVertex }
func (c *OpEmitStreamVertex) Optional() bool { return false }
func (c *OpEmitStreamVertex) Verify() error  { return nil }

// OpEndStreamPrimitive finishes the current primitive and starts a new one.
// No vertex is emitted
//
// This instruction can only be used when multiple streams are present
type OpEndStreamPrimitive struct {
	Stream Id
}

func (c *OpEndStreamPrimitive) Opcode() uint32 { return opcodeEndStreamPrimitive }
func (c *OpEndStreamPrimitive) Optional() bool { return false }
func (c *OpEndStreamPrimitive) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpEmitVertex{} })
	bind(func() Instruction { return &OpEndPrimitive{} })
	bind(func() Instruction { return &OpEmitStreamVertex{} })
	bind(func() Instruction { return &OpEndStreamPrimitive{} })
}
