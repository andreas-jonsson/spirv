// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpAtomicInit initializes atomic memory to Value.
// This is not done atomically with respect to anything.
type OpAtomicInit struct {
	Pointer Id
	Value   Id
}

func (c *OpAtomicInit) Opcode() uint32 { return opcodeAtomicInit }
func (c *OpAtomicInit) Optional() bool { return false }
func (c *OpAtomicInit) Verify() error  { return nil }

// OpAtomicLoad atomically loads through Pointer using the given Semantics.
//
// All subparts of the value that is loaded will be read atomically with
// respect to all other atomic accesses to it within Scope.
type OpAtomicLoad struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
}

func (c *OpAtomicLoad) Opcode() uint32 { return opcodeAtomicLoad }
func (c *OpAtomicLoad) Optional() bool { return false }
func (c *OpAtomicLoad) Verify() error  { return nil }

// OpAtomicStore atomically stores through Pointer using the given Semantics.
//
// All subparts of Value will be written atomically with respect to all
// other atomic accesses to it within Scope.
type OpAtomicStore struct {
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicStore) Opcode() uint32 { return opcodeAtomicStore }
func (c *OpAtomicStore) Optional() bool { return false }
func (c *OpAtomicStore) Verify() error  { return nil }

// OpAtomicExchange performs the following steps atomically with respect to any
// other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value from copying Value, and
//   3) store the New Value back through Pointer.
//
// The instruction’s result is the Original Value.
type OpAtomicExchange struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicExchange) Opcode() uint32 { return opcodeAtomicExchange }
func (c *OpAtomicExchange) Optional() bool { return false }
func (c *OpAtomicExchange) Verify() error  { return nil }

// OpAtomicCompareExchange performs the following steps atomically with respect
// to any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by selecting Value if Original Value equals Comparator
//      or selecting Original Value otherwise, and
//   3) store the New Value back through Pointer.
//
// The instruction’s result is the Original Value.
type OpAtomicCompareExchange struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
	Comparator     Id
}

func (c *OpAtomicCompareExchange) Opcode() uint32 { return opcodeAtomicCompareExchange }
func (c *OpAtomicCompareExchange) Optional() bool { return false }
func (c *OpAtomicCompareExchange) Verify() error  { return nil }

// OpAtomicCompareExchangeWeak performs the following steps atomically
// with respect to any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by selecting Value if Original Value equals Comparator
//      or selecting Original Value otherwise, and
//   3) store the New Value back through Pointer.
//
type OpAtomicCompareExchangeWeak struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
	Comparator     Id
}

func (c *OpAtomicCompareExchangeWeak) Opcode() uint32 { return opcodeAtomicCompareExchangeWeak }
func (c *OpAtomicCompareExchangeWeak) Optional() bool { return false }
func (c *OpAtomicCompareExchangeWeak) Verify() error  { return nil }

// OpAtomicIIncrement performs the following steps atomically with respect
// to any other atomic accesses within Scope to the same location
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value through integer addition of 1 to Original Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicIIncrement struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
}

func (c *OpAtomicIIncrement) Opcode() uint32 { return opcodeAtomicIIncrement }
func (c *OpAtomicIIncrement) Optional() bool { return false }
func (c *OpAtomicIIncrement) Verify() error  { return nil }

// OpAtomicIDecrement performs the following steps atomically with respect
// to any other atomic accesses within Scope to the same location
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value through integer subtraction of 1 from Original Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicIDecrement struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
}

func (c *OpAtomicIDecrement) Opcode() uint32 { return opcodeAtomicIDecrement }
func (c *OpAtomicIDecrement) Optional() bool { return false }
func (c *OpAtomicIDecrement) Verify() error  { return nil }

// OpAtomicIAdd performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by integer addition of Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicIAdd struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicIAdd) Opcode() uint32 { return opcodeAtomicIAdd }
func (c *OpAtomicIAdd) Optional() bool { return false }
func (c *OpAtomicIAdd) Verify() error  { return nil }

// OpAtomicISub performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by integer subtraction of Value from Original Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicISub struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicISub) Opcode() uint32 { return opcodeAtomicISub }
func (c *OpAtomicISub) Optional() bool { return false }
func (c *OpAtomicISub) Verify() error  { return nil }

// OpAtomicUMin performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by finding the smallest unsigned integer of
//      Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicUMin struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicUMin) Opcode() uint32 { return opcodeAtomicUMin }
func (c *OpAtomicUMin) Optional() bool { return false }
func (c *OpAtomicUMin) Verify() error  { return nil }

// OpAtomicUMax performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by finding the largest unsigned integer of Original
//      Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicUMax struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicUMax) Opcode() uint32 { return opcodeAtomicUMax }
func (c *OpAtomicUMax) Optional() bool { return false }
func (c *OpAtomicUMax) Verify() error  { return nil }

// OpAtomicAnd performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by the bitwise AND of Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicAnd struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicAnd) Opcode() uint32 { return opcodeAtomicAnd }
func (c *OpAtomicAnd) Optional() bool { return false }
func (c *OpAtomicAnd) Verify() error  { return nil }

// OpAtomicOr performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by the bitwise OR of Original Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicOr struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicOr) Opcode() uint32 { return opcodeAtomicOr }
func (c *OpAtomicOr) Optional() bool { return false }
func (c *OpAtomicOr) Verify() error  { return nil }

// OpAtomicXor performs the following steps atomically with respect to
// any other atomic accesses within Scope to the same location:
//
//   1) load through Pointer to get an Original Value,
//   2) get a New Value by the bitwise exclusive OR of Original
//      Value and Value, and
//   3) store the New Value back through Pointer.
//
type OpAtomicXor struct {
	ResultType     Id
	ResultId       Id
	Pointer        Id
	ExecutionScope ExecutionScope
	MemorySemantic MemorySemantic
	Value          Id
}

func (c *OpAtomicXor) Opcode() uint32 { return opcodeAtomicXor }
func (c *OpAtomicXor) Optional() bool { return false }
func (c *OpAtomicXor) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpAtomicInit{} })
	bind(func() Instruction { return &OpAtomicLoad{} })
	bind(func() Instruction { return &OpAtomicStore{} })
	bind(func() Instruction { return &OpAtomicExchange{} })
	bind(func() Instruction { return &OpAtomicCompareExchange{} })
	bind(func() Instruction { return &OpAtomicCompareExchangeWeak{} })
	bind(func() Instruction { return &OpAtomicIIncrement{} })
	bind(func() Instruction { return &OpAtomicIDecrement{} })
	bind(func() Instruction { return &OpAtomicIAdd{} })
	bind(func() Instruction { return &OpAtomicISub{} })
	bind(func() Instruction { return &OpAtomicUMin{} })
	bind(func() Instruction { return &OpAtomicUMax{} })
	bind(func() Instruction { return &OpAtomicAnd{} })
	bind(func() Instruction { return &OpAtomicOr{} })
	bind(func() Instruction { return &OpAtomicXor{} })
}
