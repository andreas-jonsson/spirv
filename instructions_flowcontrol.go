// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "fmt"

// TODO: Context-aware validation. Most instructions here require context
// (i.e., module) to properly validate.

// OpPhi is the SSA phi function.
type OpPhi struct {
	ResultType Id
	ResultId   Id
	Operands   []Id
}

func (c *OpPhi) Opcode() uint32 { return opcodePhi }
func (c *OpPhi) Optional() bool { return false }
func (c *OpPhi) Verify() error {
	if len(c.Operands) == 0 {
		return fmt.Errorf("OpPhi: expected operands")
	}

	if len(c.Operands)%2 != 0 {
		return fmt.Errorf("OpPhi: expected array of (Variable, ParentBlock) operand pairs")
	}

	return nil
}

// OpLoopMerge declares and controls a structured control-flow loop construct.
type OpLoopMerge struct {
	Label       Id
	LoopControl LoopControl
}

func (c *OpLoopMerge) Opcode() uint32 { return opcodeLoopMerge }
func (c *OpLoopMerge) Optional() bool { return false }
func (c *OpLoopMerge) Verify() error  { return nil }

// OpSelectionMerge declares and controls a structured control-flow selection
// construct, used with OpBranchConditional or OpSwitch.
type OpSelectionMerge struct {
	Label            Id
	SelectionControl SelectionControl
}

func (c *OpSelectionMerge) Opcode() uint32 { return opcodeSelectionMerge }
func (c *OpSelectionMerge) Optional() bool { return false }
func (c *OpSelectionMerge) Verify() error  { return nil }

// OpLabel defines a block label.
type OpLabel struct {
	ResultId Id
}

func (c *OpLabel) Opcode() uint32 { return opcodeLabel }
func (c *OpLabel) Optional() bool { return false }
func (c *OpLabel) Verify() error  { return nil }

// OpBranch is an unconditional branch to TargetLabel.
type OpBranch struct {
	TargetLabel Id
}

func (c *OpBranch) Opcode() uint32 { return opcodeBranch }
func (c *OpBranch) Optional() bool { return false }
func (c *OpBranch) Verify() error  { return nil }

// OpBranchConditional branches to TrueLabel if Condition is true, or to
// False Label if Condition is false.
type OpBranchConditional struct {
	Condition     Id
	TrueLabel     Id
	FalseLabel    Id
	BranchWeights []uint32 `spirv:"optional"`
}

func (c *OpBranchConditional) Opcode() uint32 { return opcodeBranchConditional }
func (c *OpBranchConditional) Optional() bool { return false }
func (c *OpBranchConditional) Verify() error {
	if len(c.BranchWeights) != 0 && len(c.BranchWeights) != 2 {
		return fmt.Errorf("OpBranchConditional: BranchWeights expects 0 or 2 elements")
	}
	return nil
}

// OpSwitch branches to a matching operand label.
type OpSwitch struct {
	Selector Id
	Default  Id
	// TODO: struct for pair (Value, LabelId)
	Target []uint32 `spirv:"optional"`
}

func (c *OpSwitch) Opcode() uint32 { return opcodeSwitch }
func (c *OpSwitch) Optional() bool { return false }
func (c *OpSwitch) Verify() error {
	if len(c.Target)%2 != 0 {
		return fmt.Errorf("OpSwitch: Target expects array of (LiteralNumber, Label) pairs")
	}

	for j := 0; j < len(c.Target); j += 2 {
		for k := j + 2; k < len(c.Target); k += 2 {
			if c.Target[j] == c.Target[k] {
				return fmt.Errorf("OpSwitch: Target literals must be unique")
			}
		}
	}

	return nil
}

// OpKill discards the fragment shader.
type OpKill struct{}

func (c *OpKill) Opcode() uint32 { return opcodeKill }
func (c *OpKill) Optional() bool { return false }
func (c *OpKill) Verify() error  { return nil }

func init() {
	bind(func() Instruction {
		return &OpKill{}
	})
}

// OpReturn returns with no value from a function with void return type.
type OpReturn struct{}

func (c *OpReturn) Opcode() uint32 { return opcodeReturn }
func (c *OpReturn) Optional() bool { return false }
func (c *OpReturn) Verify() error  { return nil }

// OpReturnValue returns with a value from a function.
type OpReturnValue struct {
	Value Id
}

func (c *OpReturnValue) Opcode() uint32 { return opcodeReturnValue }
func (c *OpReturnValue) Optional() bool { return false }
func (c *OpReturnValue) Verify() error  { return nil }

// OpUnreachable declares that this block is not reachable in the Control
// Flow Graph.
type OpUnreachable struct{}

func (c *OpUnreachable) Opcode() uint32 { return opcodeUnreachable }
func (c *OpUnreachable) Optional() bool { return false }
func (c *OpUnreachable) Verify() error  { return nil }

// OpLifetimeStart declares that the content of the object pointed to was
// not defined before this instruction.
type OpLifetimeStart struct {
	Object       Id
	MemoryAmount uint32
}

func (c *OpLifetimeStart) Opcode() uint32 { return opcodeLifetimeStart }
func (c *OpLifetimeStart) Optional() bool { return false }
func (c *OpLifetimeStart) Verify() error  { return nil }

// OpLifetimeStop declares that the content of the object pointed to is
// dead after this instruction.
type OpLifetimeStop struct {
	Object       Id
	MemoryAmount uint32
}

func (c *OpLifetimeStop) Opcode() uint32 { return opcodeLifetimeStop }
func (c *OpLifetimeStop) Optional() bool { return false }
func (c *OpLifetimeStop) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpPhi{} })
	bind(func() Instruction { return &OpLoopMerge{} })
	bind(func() Instruction { return &OpSelectionMerge{} })
	bind(func() Instruction { return &OpLabel{} })
	bind(func() Instruction { return &OpBranch{} })
	bind(func() Instruction { return &OpBranchConditional{} })
	bind(func() Instruction { return &OpSwitch{} })
	bind(func() Instruction { return &OpReturn{} })
	bind(func() Instruction { return &OpReturnValue{} })
	bind(func() Instruction { return &OpUnreachable{} })
	bind(func() Instruction { return &OpLifetimeStart{} })
	bind(func() Instruction { return &OpLifetimeStop{} })
}
