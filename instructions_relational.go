// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpAny result is true if any component of Vector is true,
// otherwise result is false.
type OpAny struct {
	ResultType Id
	ResultId   Id
	Vector     Id
}

func (c *OpAny) Opcode() uint32 { return opcodeAny }
func (c *OpAny) Optional() bool { return false }
func (c *OpAny) Verify() error  { return nil }

// OpAll result is true if all components of Vector are true,
// otherwise result is false
type OpAll struct {
	ResultType Id
	ResultId   Id
	Vector     Id
}

func (c *OpAll) Opcode() uint32 { return opcodeAll }
func (c *OpAll) Optional() bool { return false }
func (c *OpAll) Verify() error  { return nil }

// OpIsNan result is true if x is an IEEE NaN,
// otherwise result is false.
type OpIsNan struct {
	ResultType Id
	ResultId   Id
	X          Id
}

func (c *OpIsNan) Opcode() uint32 { return opcodeIsNan }
func (c *OpIsNan) Optional() bool { return false }
func (c *OpIsNan) Verify() error  { return nil }

// OpIsInf result is true if x is an IEEE Inf,
// otherwise result is false.
type OpIsInf struct {
	ResultType Id
	ResultId   Id
	X          Id
}

func (c *OpIsInf) Opcode() uint32 { return opcodeIsInf }
func (c *OpIsInf) Optional() bool { return false }
func (c *OpIsInf) Verify() error  { return nil }

// OpIsFinite result is true if x is an IEEE finite number,
// otherwise result is false..
type OpIsFinite struct {
	ResultType Id
	ResultId   Id
	X          Id
}

func (c *OpIsFinite) Opcode() uint32 { return opcodeIsFinite }
func (c *OpIsFinite) Optional() bool { return false }
func (c *OpIsFinite) Verify() error  { return nil }

// OpIsNormal result is true if x is an IEEE normal number,
// otherwise result is false.
type OpIsNormal struct {
	ResultType Id
	ResultId   Id
	X          Id
}

func (c *OpIsNormal) Opcode() uint32 { return opcodeIsNormal }
func (c *OpIsNormal) Optional() bool { return false }
func (c *OpIsNormal) Verify() error  { return nil }

// OpSignBitSet result is true if x has its sign bit set,
// otherwise result is false.
type OpSignBitSet struct {
	ResultType Id
	ResultId   Id
	X          Id
}

func (c *OpSignBitSet) Opcode() uint32 { return opcodeSignBitSet }
func (c *OpSignBitSet) Optional() bool { return false }
func (c *OpSignBitSet) Verify() error  { return nil }

// OpLessOrGreater result is true if x < y or x > y,
// where IEEE comparisons are used, otherwise result is false.
type OpLessOrGreater struct {
	ResultType Id
	ResultId   Id
	X          Id
}

func (c *OpLessOrGreater) Opcode() uint32 { return opcodeLessOrGreater }
func (c *OpLessOrGreater) Optional() bool { return false }
func (c *OpLessOrGreater) Verify() error  { return nil }

// OpOrdered result is true if both x == x and y == y are true,
// where IEEE comparison is used, otherwise result is false.
type OpOrdered struct {
	ResultType Id
	ResultId   Id
	X          Id
	Y          Id
}

func (c *OpOrdered) Opcode() uint32 { return opcodeOrdered }
func (c *OpOrdered) Optional() bool { return false }
func (c *OpOrdered) Verify() error  { return nil }

// OpUnordered result is true if either x or y is an IEEE NaN,
// otherwise result is false.
type OpUnordered struct {
	ResultType Id
	ResultId   Id
	X          Id
	Y          Id
}

func (c *OpUnordered) Opcode() uint32 { return opcodeUnordered }
func (c *OpUnordered) Optional() bool { return false }
func (c *OpUnordered) Verify() error  { return nil }

// OpLogicalOr result is true if either Operand 1 or Operand 2 is true.
// Result is false if both Operand 1 and Operand 2 are false.
type OpLogicalOr struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpLogicalOr) Opcode() uint32 { return opcodeLogicalOr }
func (c *OpLogicalOr) Optional() bool { return false }
func (c *OpLogicalOr) Verify() error  { return nil }

// OpLogicalXor result is true if exactly one of Operand 1 or
// Operand 2 is true. Result is false if Operand 1 and Operand 2
// have the same value.
type OpLogicalXor struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpLogicalXor) Opcode() uint32 { return opcodeLogicalXor }
func (c *OpLogicalXor) Optional() bool { return false }
func (c *OpLogicalXor) Verify() error  { return nil }

// OpLogicalAnd result is true if both Operand 1 and Operand 2 are true.
// Result is false if either Operand 1 or Operand 2 are false
type OpLogicalAnd struct {
	ResultType Id
	ResultId   Id
	Operand1   Id
	Operand2   Id
}

func (c *OpLogicalAnd) Opcode() uint32 { return opcodeLogicalAnd }
func (c *OpLogicalAnd) Optional() bool { return false }
func (c *OpLogicalAnd) Verify() error  { return nil }

// OpSelect select between two objects.
// Results are computed per component
type OpSelect struct {
	ResultType Id
	ResultId   Id
	Condition  Id
	Object1    Id
	Object2    Id
}

func (c *OpSelect) Opcode() uint32 { return opcodeSelect }
func (c *OpSelect) Optional() bool { return false }
func (c *OpSelect) Verify() error  { return nil }

// OpIEqual performs Integer comparison for equality.
type OpIEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpIEqual) Opcode() uint32 { return opcodeIEqual }
func (c *OpIEqual) Optional() bool { return false }
func (c *OpIEqual) Verify() error  { return nil }

// OpFOrdEqual performs Floating-point comparison for being
// ordered and equal.
type OpFOrdEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFOrdEqual) Opcode() uint32 { return opcodeFOrdEqual }
func (c *OpFOrdEqual) Optional() bool { return false }
func (c *OpFOrdEqual) Verify() error  { return nil }

// OpFUnordEqual performs Floating-point comparison for being
// unordered or equal.
type OpFUnordEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFUnordEqual) Opcode() uint32 { return opcodeFUnordEqual }
func (c *OpFUnordEqual) Optional() bool { return false }
func (c *OpFUnordEqual) Verify() error  { return nil }

// OpINotEqual performs Integer comparison for inequality
type OpINotEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpINotEqual) Opcode() uint32 { return opcodeINotEqual }
func (c *OpINotEqual) Optional() bool { return false }
func (c *OpINotEqual) Verify() error  { return nil }

// OpFOrdNotEqual performs Floating-point comparison for being
// ordered and not equal.
type OpFOrdNotEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFOrdNotEqual) Opcode() uint32 { return opcodeFOrdNotEqual }
func (c *OpFOrdNotEqual) Optional() bool { return false }
func (c *OpFOrdNotEqual) Verify() error  { return nil }

// OpFUnordNotEqual performs Floating-point comparison for
// being unordered or not equal.
type OpFUnordNotEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFUnordNotEqual) Opcode() uint32 { return opcodeFUnordNotEqual }
func (c *OpFUnordNotEqual) Optional() bool { return false }
func (c *OpFUnordNotEqual) Verify() error  { return nil }

// OpULessThan performs Unsigned-integer comparison if Operand 1
// is less than Operand 2.
type OpULessThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpULessThan) Opcode() uint32 { return opcodeULessThan }
func (c *OpULessThan) Optional() bool { return false }
func (c *OpULessThan) Verify() error  { return nil }

// OpSLessThan performs Signed-integer comparison if Operand 1
// is less than Operand 2.
type OpSLessThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpSLessThan) Opcode() uint32 { return opcodeSLessThan }
func (c *OpSLessThan) Optional() bool { return false }
func (c *OpSLessThan) Verify() error  { return nil }

// OpFOrdLessThan performs Floating-point comparison if operands are
// ordered and Operand 1 is less than Operand 2.
type OpFOrdLessThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFOrdLessThan) Opcode() uint32 { return opcodeFOrdLessThan }
func (c *OpFOrdLessThan) Optional() bool { return false }
func (c *OpFOrdLessThan) Verify() error  { return nil }

// OpFUnordLessThan performs Floating-point comparison if operands
// are unordered or Operand 1 is less than Operand 2.
type OpFUnordLessThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFUnordLessThan) Opcode() uint32 { return opcodeFUnordLessThan }
func (c *OpFUnordLessThan) Optional() bool { return false }
func (c *OpFUnordLessThan) Verify() error  { return nil }

// OpUGreaterThan performs Unsigned-integer comparison if Operand 1
// is greater than Operand 2.
type OpUGreaterThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpUGreaterThan) Opcode() uint32 { return opcodeUGreaterThan }
func (c *OpUGreaterThan) Optional() bool { return false }
func (c *OpUGreaterThan) Verify() error  { return nil }

// OpSGreaterThan performs Signed-integer comparison if Operand 1
// is greater than Operand 2.
type OpSGreaterThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpSGreaterThan) Opcode() uint32 { return opcodeSGreaterThan }
func (c *OpSGreaterThan) Optional() bool { return false }
func (c *OpSGreaterThan) Verify() error  { return nil }

// OpFOrdGreaterThan performs Floating-point comparison if operands
// are ordered and Operand 1 is greater than Operand 2.
type OpFOrdGreaterThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFOrdGreaterThan) Opcode() uint32 { return opcodeFOrdGreaterThan }
func (c *OpFOrdGreaterThan) Optional() bool { return false }
func (c *OpFOrdGreaterThan) Verify() error  { return nil }

// OpFUnordGreaterThan performs Floating-point comparison if
// operands are unordered or Operand 1 is greater than Operand 2.
type OpFUnordGreaterThan struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFUnordGreaterThan) Opcode() uint32 { return opcodeFUnordGreaterThan }
func (c *OpFUnordGreaterThan) Optional() bool { return false }
func (c *OpFUnordGreaterThan) Verify() error  { return nil }

// OpULessThanEqual performs Unsigned-integer comparison if Operand 1 is
// less than or equal to Operand 2.
type OpULessThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpULessThanEqual) Opcode() uint32 { return opcodeULessThanEqual }
func (c *OpULessThanEqual) Optional() bool { return false }
func (c *OpULessThanEqual) Verify() error  { return nil }

// OpSLessThanEqual performs Signed-integer comparison if Operand 1 is
// less than or equal to Operand 2.
type OpSLessThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpSLessThanEqual) Opcode() uint32 { return opcodeSLessThanEqual }
func (c *OpSLessThanEqual) Optional() bool { return false }
func (c *OpSLessThanEqual) Verify() error  { return nil }

// OpFOrdLessThanEqual performs Floating-point comparison if operands
// are ordered and Operand 1 is less than or equal to Operand 2.
type OpFOrdLessThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFOrdLessThanEqual) Opcode() uint32 { return opcodeFOrdLessThanEqual }
func (c *OpFOrdLessThanEqual) Optional() bool { return false }
func (c *OpFOrdLessThanEqual) Verify() error  { return nil }

// OpFUnordLessThanEqual performs Floating-point comparison if
// operands are unordered or Operand 1 is less than or equal to Operand 2.
type OpFUnordLessThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFUnordLessThanEqual) Opcode() uint32 { return opcodeFUnordLessThanEqual }
func (c *OpFUnordLessThanEqual) Optional() bool { return false }
func (c *OpFUnordLessThanEqual) Verify() error  { return nil }

// OpUGreaterThanEqual performs Unsigned-integer comparison if Operand 1 is
// greater than or equal to Operand 2.
type OpUGreaterThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpUGreaterThanEqual) Opcode() uint32 { return opcodeUGreaterThanEqual }
func (c *OpUGreaterThanEqual) Optional() bool { return false }
func (c *OpUGreaterThanEqual) Verify() error  { return nil }

// OpSGreaterThanEqual performs Signed-integer comparison if Operand 1 is
// greater than or equal to Operand 2.
type OpSGreaterThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpSGreaterThanEqual) Opcode() uint32 { return opcodeSGreaterThanEqual }
func (c *OpSGreaterThanEqual) Optional() bool { return false }
func (c *OpSGreaterThanEqual) Verify() error  { return nil }

// OpFOrdGreaterThanEqual performs Floating-point comparison if
// operands are ordered and Operand 1 is greater than or equal to Operand 2.
type OpFOrdGreaterThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFOrdGreaterThanEqual) Opcode() uint32 { return opcodeFOrdGreaterThanEqual }
func (c *OpFOrdGreaterThanEqual) Optional() bool { return false }
func (c *OpFOrdGreaterThanEqual) Verify() error  { return nil }

// OpFUnordGreaterThanEqual performs Floating-point comparison if
// operands are unordered or Operand 1 is greater than or equal
// to Operand 2.
type OpFUnordGreaterThanEqual struct {
	ResultType Id
	ResultId   Id
	Object1    Id
	Object2    Id
}

func (c *OpFUnordGreaterThanEqual) Opcode() uint32 { return opcodeFUnordGreaterThanEqual }
func (c *OpFUnordGreaterThanEqual) Optional() bool { return false }
func (c *OpFUnordGreaterThanEqual) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpAny{} })
	bind(func() Instruction { return &OpAll{} })
	bind(func() Instruction { return &OpIsNan{} })
	bind(func() Instruction { return &OpIsInf{} })
	bind(func() Instruction { return &OpIsFinite{} })
	bind(func() Instruction { return &OpIsNormal{} })
	bind(func() Instruction { return &OpSignBitSet{} })
	bind(func() Instruction { return &OpLessOrGreater{} })
	bind(func() Instruction { return &OpOrdered{} })
	bind(func() Instruction { return &OpUnordered{} })
	bind(func() Instruction { return &OpLogicalOr{} })
	bind(func() Instruction { return &OpLogicalXor{} })
	bind(func() Instruction { return &OpLogicalAnd{} })
	bind(func() Instruction { return &OpSelect{} })
	bind(func() Instruction { return &OpIEqual{} })
	bind(func() Instruction { return &OpFOrdEqual{} })
	bind(func() Instruction { return &OpFUnordEqual{} })
	bind(func() Instruction { return &OpINotEqual{} })
	bind(func() Instruction { return &OpFOrdNotEqual{} })
	bind(func() Instruction { return &OpFUnordNotEqual{} })
	bind(func() Instruction { return &OpULessThan{} })
	bind(func() Instruction { return &OpSLessThan{} })
	bind(func() Instruction { return &OpFOrdLessThan{} })
	bind(func() Instruction { return &OpFUnordLessThan{} })
	bind(func() Instruction { return &OpUGreaterThan{} })
	bind(func() Instruction { return &OpSGreaterThan{} })
	bind(func() Instruction { return &OpFOrdGreaterThan{} })
	bind(func() Instruction { return &OpFUnordGreaterThan{} })
	bind(func() Instruction { return &OpULessThanEqual{} })
	bind(func() Instruction { return &OpSLessThanEqual{} })
	bind(func() Instruction { return &OpFOrdLessThanEqual{} })
	bind(func() Instruction { return &OpFUnordLessThanEqual{} })
	bind(func() Instruction { return &OpUGreaterThanEqual{} })
	bind(func() Instruction { return &OpSGreaterThanEqual{} })
	bind(func() Instruction { return &OpFOrdGreaterThanEqual{} })
	bind(func() Instruction { return &OpFUnordGreaterThanEqual{} })
}
