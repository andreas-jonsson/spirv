// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// OpMemoryModel represents the OpMemoryModel instruction.
//
// It sets addressing model and memory model for the entire module.
type OpMemoryModel struct {
	AddressingModel AddressingModel
	MemoryModel     MemoryModel
}

func (c *OpMemoryModel) Opcode() uint32 { return opcodeMemoryModel }
func (c *OpMemoryModel) Optional() bool { return false }
func (c *OpMemoryModel) Verify() error  { return nil }

// OpEntryPoint represents the OpEntryPoint instruction.
// It declares an entry point and its execution model.
type OpEntryPoint struct {
	ExecutionModel ExecutionModel
	ResultId       Id
}

func (c *OpEntryPoint) Opcode() uint32 { return opcodeEntryPoint }
func (c *OpEntryPoint) Optional() bool { return false }
func (c *OpEntryPoint) Verify() error  { return nil }

// OpExecutionMode represents the OpExecutionMode instruction.
// It declares an execution mode for an entry point.
type OpExecutionMode struct {
	EntryPoint Id
	Mode       ExecutionMode
	Argv       []uint32
}

func (c *OpExecutionMode) Opcode() uint32 { return opcodeExecutionMode }
func (c *OpExecutionMode) Optional() bool { return false }
func (c *OpExecutionMode) Verify() error {
	argc := len(c.Argv)

	switch c.Mode {
	case ExecutionModeInvocations,
		ExecutionModeOutputVertices,
		ExecutionModeVecTypeHint:
		if argc != 1 {
			return fmt.Errorf("OpExecutionMode: ExecutionMode(%d) must have 1 argument", c.Mode)
		}

		return nil

	case ExecutionModeLocalSize,
		ExecutionModeLocalSizeHint:
		if argc != 3 {
			return fmt.Errorf("OpExecutionMode: ExecutionMode(%d) must have 3 arguments", c.Mode)
		}

		return nil
	}

	if argc > 0 {
		return fmt.Errorf("OpExecutionMode: extraneous arguments for ExecutionMode(%d)",
			c.Mode)
	}

	return nil
}

// OpCompileFlag represents the OpCompileFlag instruction.
type OpCompileFlag struct {
	Flag String
}

func (c *OpCompileFlag) Opcode() uint32 { return opcodeCompileFlag }
func (c *OpCompileFlag) Optional() bool { return false }
func (c *OpCompileFlag) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpMemoryModel{} })
	bind(func() Instruction { return &OpEntryPoint{} })
	bind(func() Instruction { return &OpExecutionMode{} })
	bind(func() Instruction { return &OpCompileFlag{} })
}
