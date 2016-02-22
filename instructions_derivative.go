// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpDPdx is equivalent to either OpDPdxFine or OpDPdxCoarse on P.
// Selection of which one is based on external factors.
type OpDPdx struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpDPdx) Opcode() uint32 { return opcodeDPdx }
func (c *OpDPdx) Optional() bool { return false }
func (c *OpDPdx) Verify() error  { return nil }

// OpDPdy is equivalent to either OpDPdyFine or OpDPdyCoarse on P.
// Selection of which one is based on external factors.
type OpDPdy struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpDPdy) Opcode() uint32 { return opcodeDPdy }
func (c *OpDPdy) Optional() bool { return false }
func (c *OpDPdy) Verify() error  { return nil }

// OpFwidth is equivalent to computing the sum of the absolute values of
// OpDPdx and OpDPdy on P.
type OpFwidth struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpFwidth) Opcode() uint32 { return opcodeFwidth }
func (c *OpFwidth) Optional() bool { return false }
func (c *OpFwidth) Verify() error  { return nil }

// OpDPdxFine calculates the partial derivative of P with respect to the
// window x coordinate.
type OpDPdxFine struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpDPdxFine) Opcode() uint32 { return opcodeDPdxFine }
func (c *OpDPdxFine) Optional() bool { return false }
func (c *OpDPdxFine) Verify() error  { return nil }

// OpDPdyFine calculates the partial derivative of P with respect to the
// window y coordinate.
type OpDPdyFine struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpDPdyFine) Opcode() uint32 { return opcodeDPdyFine }
func (c *OpDPdyFine) Optional() bool { return false }
func (c *OpDPdyFine) Verify() error  { return nil }

// OpFwidthFine is equivalent to computing the sum of the absolute values
// of OpDPdxFine and OpDPdyFine on P.
type OpFwidthFine struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpFwidthFine) Opcode() uint32 { return opcodeFwidthFine }
func (c *OpFwidthFine) Optional() bool { return false }
func (c *OpFwidthFine) Verify() error  { return nil }

// OpDPdxCoarse calculates the partial derivative of P with respect to the
// window x coordinate.
type OpDPdxCoarse struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpDPdxCoarse) Opcode() uint32 { return opcodeDPdxCoarse }
func (c *OpDPdxCoarse) Optional() bool { return false }
func (c *OpDPdxCoarse) Verify() error  { return nil }

// OpDPdyCoarse calculates the partial derivative of P with respect to the
// window y coordinate.
type OpDPdyCoarse struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpDPdyCoarse) Opcode() uint32 { return opcodeDPdyCoarse }
func (c *OpDPdyCoarse) Optional() bool { return false }
func (c *OpDPdyCoarse) Verify() error  { return nil }

// OpFwidthCoarse is equivalent to computing the sum of the absolute values
// of OpDPdxCoarse and OpDPdyCoarse on P.
type OpFwidthCoarse struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpFwidthCoarse) Opcode() uint32 { return opcodeFwidthCoarse }
func (c *OpFwidthCoarse) Optional() bool { return false }
func (c *OpFwidthCoarse) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpDPdx{} })
	bind(func() Instruction { return &OpDPdy{} })
	bind(func() Instruction { return &OpFwidth{} })
	bind(func() Instruction { return &OpDPdxFine{} })
	bind(func() Instruction { return &OpDPdyFine{} })
	bind(func() Instruction { return &OpFwidthFine{} })
	bind(func() Instruction { return &OpDPdxCoarse{} })
	bind(func() Instruction { return &OpDPdyCoarse{} })
	bind(func() Instruction { return &OpFwidthCoarse{} })
}
