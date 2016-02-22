// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpFunction defines a function body. This instruction must be
// immediately followed by one OpFunctionParameter instruction per each
// formal parameter of this function.
//
// This function’s body will terminate with the next OpFunctionEnd
// instruction.
type OpFunction struct {
	ResultType   Id
	ResultId     Id
	ControlMask  FunctionControlMask
	FunctionType Id
}

func (c *OpFunction) Opcode() uint32 { return opcodeFunction }
func (c *OpFunction) Optional() bool { return false }
func (c *OpFunction) Verify() error  { return nil }

// OpFunctionParameter declares the <id> for a formal parameter belonging
// to the current function.
//
// This instruction must immediately follow an OpFunction or OpFunctionParameter
// instruction. The order of contiguous OpFunctionParameter instructions is
// the same order arguments will be listed in an OpFunctionCall instruction to
// this function. It is also the same order in which Parameter Type operands
// are listed in the OpTypeFunction of the Function Type operand for this
// function’s OpFunction instruction.
type OpFunctionParameter struct {
	ResultType Id
	ResultId   Id
}

func (c *OpFunctionParameter) Opcode() uint32 { return opcodeFunctionParameter }
func (c *OpFunctionParameter) Optional() bool { return false }
func (c *OpFunctionParameter) Verify() error  { return nil }

// OpFunctionParameter is the last instruction of a function definition.
type OpFunctionEnd struct{}

func (c *OpFunctionEnd) Opcode() uint32 { return opcodeFunctionEnd }
func (c *OpFunctionEnd) Optional() bool { return false }
func (c *OpFunctionEnd) Verify() error  { return nil }

// OpFunctionCall defines a function call.
//
// Note: A forward call is possible because there is no missing type
// information: Result Type must match the Return Type of the function, and
// the calling argument types must match the formal parameter types.
type OpFunctionCall struct {
	ResultType Id
	ResultId   Id
	Function   Id
	Argv       []Id
}

func (c *OpFunctionCall) Opcode() uint32 { return opcodeFunctionCall }
func (c *OpFunctionCall) Optional() bool { return false }
func (c *OpFunctionCall) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpFunction{} })
	bind(func() Instruction { return &OpFunctionParameter{} })
	bind(func() Instruction { return &OpFunctionEnd{} })
	bind(func() Instruction { return &OpFunctionCall{} })
}
