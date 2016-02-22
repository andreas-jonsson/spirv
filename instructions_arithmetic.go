// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSNegate performs signed-integer subtract of Operand from zero.
type OpSNegate struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpSNegate) Opcode() uint32 { return opcodeSNegate }
func (c *OpSNegate) Optional() bool { return false }
func (c *OpSNegate) Verify() error  { return nil }

// OpFNegate performs Floating-point subtract of Operand from zero.
type OpFNegate struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpFNegate) Opcode() uint32 { return opcodeFNegate }
func (c *OpFNegate) Optional() bool { return false }
func (c *OpFNegate) Verify() error  { return nil }

// OpNot complements the bits of Operand.
type OpNot struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpNot) Opcode() uint32 { return opcodeNot }
func (c *OpNot) Optional() bool { return false }
func (c *OpNot) Verify() error  { return nil }

// OpIAdd performs Integer addition of Operand 1 and Operand 2.
type OpIAdd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpIAdd) Opcode() uint32 { return opcodeIAdd }
func (c *OpIAdd) Optional() bool { return false }
func (c *OpIAdd) Verify() error  { return nil }

// OpFAdd performs Floating-point addition of Operand 1 and Operand 2.
type OpFAdd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFAdd) Opcode() uint32 { return opcodeFAdd }
func (c *OpFAdd) Optional() bool { return false }
func (c *OpFAdd) Verify() error  { return nil }

// OpISub performs Integer subtraction of Operand 2 from Operand 1.
type OpISub struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpISub) Opcode() uint32 { return opcodeISub }
func (c *OpISub) Optional() bool { return false }
func (c *OpISub) Verify() error  { return nil }

// OpFSub performs Floating-point subtraction of Operand 2 from Operand 1.
type OpFSub struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFSub) Opcode() uint32 { return opcodeFSub }
func (c *OpFSub) Optional() bool { return false }
func (c *OpFSub) Verify() error  { return nil }

// OpIMul performs Integer multiplication of Operand 1 and Operand 2.
type OpIMul struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpIMul) Opcode() uint32 { return opcodeIMul }
func (c *OpIMul) Optional() bool { return false }
func (c *OpIMul) Verify() error  { return nil }

// OpFMul performs Floating-point multiplication of Operand 1 and Operand 2.
type OpFMul struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFMul) Opcode() uint32 { return opcodeFMul }
func (c *OpFMul) Optional() bool { return false }
func (c *OpFMul) Verify() error  { return nil }

// OpUDiv performs Unsigned-integer division of Operand 1 divided by Operand 2.
type OpUDiv struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpUDiv) Opcode() uint32 { return opcodeUDiv }
func (c *OpUDiv) Optional() bool { return false }
func (c *OpUDiv) Verify() error  { return nil }

// OpSDiv performs Signed-integer division of Operand 1 divided by Operand 2.
type OpSDiv struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpSDiv) Opcode() uint32 { return opcodeSDiv }
func (c *OpSDiv) Optional() bool { return false }
func (c *OpSDiv) Verify() error  { return nil }

// OpFDiv performs Floating-point division of Operand 1 divided by Operand 2.
type OpFDiv struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFDiv) Opcode() uint32 { return opcodeFDiv }
func (c *OpFDiv) Optional() bool { return false }
func (c *OpFDiv) Verify() error  { return nil }

// OpUMod performs Unsigned modulo operation of Operand 1 modulo Operand 2.
type OpUMod struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpUMod) Opcode() uint32 { return opcodeUMod }
func (c *OpUMod) Optional() bool { return false }
func (c *OpUMod) Verify() error  { return nil }

// OpSRem performs Signed remainder operation of Operand 1 divided by Operand 2.
type OpSRem struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpSRem) Opcode() uint32 { return opcodeSRem }
func (c *OpSRem) Optional() bool { return false }
func (c *OpSRem) Verify() error  { return nil }

// OpSMod performs Signed modulo operation of Operand 1 modulo Operand 2.
type OpSMod struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpSMod) Opcode() uint32 { return opcodeSMod }
func (c *OpSMod) Optional() bool { return false }
func (c *OpSMod) Verify() error  { return nil }

// OpFRem performs Floating-point remainder operation of Operand 1 divided by Operand 2.
type OpFRem struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFRem) Opcode() uint32 { return opcodeFRem }
func (c *OpFRem) Optional() bool { return false }
func (c *OpFRem) Verify() error  { return nil }

// OpFMod performs Floating-point modulo operation of Operand 1 modulo Operand 2.
type OpFMod struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpFMod) Opcode() uint32 { return opcodeFMod }
func (c *OpFMod) Optional() bool { return false }
func (c *OpFMod) Verify() error  { return nil }

// OpVectorTimesScalar scales a floating-point vector.
type OpVectorTimesScalar struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Scalar     Id
}

func (c *OpVectorTimesScalar) Opcode() uint32 { return opcodeVectorTimesScalar }
func (c *OpVectorTimesScalar) Optional() bool { return false }
func (c *OpVectorTimesScalar) Verify() error  { return nil }

// OpMatrixTimesScalar scales a floating-point matrix.
type OpMatrixTimesScalar struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Scalar     Id
}

func (c *OpMatrixTimesScalar) Opcode() uint32 { return opcodeMatrixTimesScalar }
func (c *OpMatrixTimesScalar) Optional() bool { return false }
func (c *OpMatrixTimesScalar) Verify() error  { return nil }

// OpVectorTimesMatrix performs Linear-algebraic Vector X Matrix.
type OpVectorTimesMatrix struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Matrix     Id
}

func (c *OpVectorTimesMatrix) Opcode() uint32 { return opcodeVectorTimesMatrix }
func (c *OpVectorTimesMatrix) Optional() bool { return false }
func (c *OpVectorTimesMatrix) Verify() error  { return nil }

// OpMatrixTimesVector performs Linear-algebraic Vector X Matrix.
type OpMatrixTimesVector struct {
	ResultType Id
	ResultId   Id
	Matrix     Id
	Vector     Id
}

func (c *OpMatrixTimesVector) Opcode() uint32 { return opcodeMatrixTimesVector }
func (c *OpMatrixTimesVector) Optional() bool { return false }
func (c *OpMatrixTimesVector) Verify() error  { return nil }

// OpMatrixTimesMatrix performs Linear-algebraic multiply of Left X Right.
type OpMatrixTimesMatrix struct {
	ResultType Id
	ResultId   Id
	Left       Id
	Right      Id
}

func (c *OpMatrixTimesMatrix) Opcode() uint32 { return opcodeMatrixTimesMatrix }
func (c *OpMatrixTimesMatrix) Optional() bool { return false }
func (c *OpMatrixTimesMatrix) Verify() error  { return nil }

// OpOuterProduct performs Linear-algebraic outer product of Vector 1 and Vector 2.
type OpOuterProduct struct {
	ResultType Id
	ResultId   Id
	Vector1    Id
	Vector2    Id
}

func (c *OpOuterProduct) Opcode() uint32 { return opcodeOuterProduct }
func (c *OpOuterProduct) Optional() bool { return false }
func (c *OpOuterProduct) Verify() error  { return nil }

// OpDot performs Dot product of Vector 1 and Vector 2
type OpDot struct {
	ResultType Id
	ResultId   Id
	Vector1    Id
	Vector2    Id
}

func (c *OpDot) Opcode() uint32 { return opcodeDot }
func (c *OpDot) Optional() bool { return false }
func (c *OpDot) Verify() error  { return nil }

// OpShiftRightLogical shifts the bits in Operand 1 right by the number
// of bits specified in Operand 2
type OpShiftRightLogical struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpShiftRightLogical) Opcode() uint32 { return opcodeShiftRightLogical }
func (c *OpShiftRightLogical) Optional() bool { return false }
func (c *OpShiftRightLogical) Verify() error  { return nil }

// OpShiftRightArithmetic shifts the bits in Operand 1 right by the number of
// bits specified in Operand 2
type OpShiftRightArithmetic struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpShiftRightArithmetic) Opcode() uint32 { return opcodeShiftRightArithmetic }
func (c *OpShiftRightArithmetic) Optional() bool { return false }
func (c *OpShiftRightArithmetic) Verify() error  { return nil }

// OpShiftLeftLogical shifts the bits in Operand 1 left by the number
// of bits specified in Operand 2
type OpShiftLeftLogical struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpShiftLeftLogical) Opcode() uint32 { return opcodeShiftLeftLogical }
func (c *OpShiftLeftLogical) Optional() bool { return false }
func (c *OpShiftLeftLogical) Verify() error  { return nil }

// OpBitwiseOr result is 1 if either Operand 1 or Operand 2 is 1.
// Result is 0 if both Operand 1 and Operand 2 are 0.
type OpBitwiseOr struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpBitwiseOr) Opcode() uint32 { return opcodeBitwiseOr }
func (c *OpBitwiseOr) Optional() bool { return false }
func (c *OpBitwiseOr) Verify() error  { return nil }

// OpBitwiseXor result is 1 if exactly one of Operand 1 or Operand 2 is 1.
// Result is 0 if Operand 1 and Operand 2 have the same value
type OpBitwiseXor struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpBitwiseXor) Opcode() uint32 { return opcodeBitwiseXor }
func (c *OpBitwiseXor) Optional() bool { return false }
func (c *OpBitwiseXor) Verify() error  { return nil }

// OpBitwiseAnd result is 1 if both Operand 1 and Operand 2 are 1.
// Result is 0 if either Operand 1 or Operand 2 are 0
type OpBitwiseAnd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpBitwiseAnd) Opcode() uint32 { return opcodeBitwiseAnd }
func (c *OpBitwiseAnd) Optional() bool { return false }
func (c *OpBitwiseAnd) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpSNegate{} })
	bind(func() Instruction { return &OpFNegate{} })
	bind(func() Instruction { return &OpNot{} })
	bind(func() Instruction { return &OpIAdd{} })
	bind(func() Instruction { return &OpFAdd{} })
	bind(func() Instruction { return &OpISub{} })
	bind(func() Instruction { return &OpFSub{} })
	bind(func() Instruction { return &OpIMul{} })
	bind(func() Instruction { return &OpFMul{} })
	bind(func() Instruction { return &OpUDiv{} })
	bind(func() Instruction { return &OpSDiv{} })
	bind(func() Instruction { return &OpFDiv{} })
	bind(func() Instruction { return &OpUMod{} })
	bind(func() Instruction { return &OpSRem{} })
	bind(func() Instruction { return &OpSMod{} })
	bind(func() Instruction { return &OpFRem{} })
	bind(func() Instruction { return &OpFMod{} })
	bind(func() Instruction { return &OpVectorTimesScalar{} })
	bind(func() Instruction { return &OpMatrixTimesScalar{} })
	bind(func() Instruction { return &OpVectorTimesMatrix{} })
	bind(func() Instruction { return &OpMatrixTimesVector{} })
	bind(func() Instruction { return &OpMatrixTimesMatrix{} })
	bind(func() Instruction { return &OpOuterProduct{} })
	bind(func() Instruction { return &OpDot{} })
	bind(func() Instruction { return &OpShiftRightLogical{} })
	bind(func() Instruction { return &OpShiftRightArithmetic{} })
	bind(func() Instruction { return &OpShiftLeftLogical{} })
	bind(func() Instruction { return &OpBitwiseOr{} })
	bind(func() Instruction { return &OpBitwiseXor{} })
	bind(func() Instruction { return &OpBitwiseAnd{} })
}
