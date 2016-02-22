// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpNop represents the OpNop instruction.
//
// Its use is not allowed, but it is explicitely defined in the specification.
// We will therefor have the decoder and encoder return an appropriate error
// when it is being used.
type OpNop struct{}

func (c *OpNop) Opcode() uint32 { return opcodeNop }
func (c *OpNop) Optional() bool { return false }
func (c *OpNop) Verify() error  { return ErrUnacceptable }

// OpUndef makes an intermediate object with no initialization.
type OpUndef struct {
	ResultType Id
	ResultId   Id
}

func (c *OpUndef) Opcode() uint32 { return opcodeUndef }
func (c *OpUndef) Optional() bool { return false }
func (c *OpUndef) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpNop{} })
	bind(func() Instruction { return &OpUndef{} })
}
