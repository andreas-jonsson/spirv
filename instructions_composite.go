// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVectorExtractDynamic reads a single, dynamically selected, component of
// a vector.
type OpVectorExtractDynamic struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Index      Id
}

func (c *OpVectorExtractDynamic) Opcode() uint32 { return opcodeVectorExtractDynamic }
func (c *OpVectorExtractDynamic) Optional() bool { return false }
func (c *OpVectorExtractDynamic) Verify() error  { return nil }

// OpVectorInsertDynamic writes a single, variably selected, component
// into a vector.
type OpVectorInsertDynamic struct {
	ResultType Id
	ResultId   Id
	Vector     Id
	Component  Id
	Index      Id
}

func (c *OpVectorInsertDynamic) Opcode() uint32 { return opcodeVectorInsertDynamic }
func (c *OpVectorInsertDynamic) Optional() bool { return false }
func (c *OpVectorInsertDynamic) Verify() error  { return nil }

// OpVectorShuffle selects arbitrary components from two vectors to make
// a new vector.
type OpVectorShuffle struct {
	ResultType Id
	ResultId   Id
	Vector1    Id
	Vector2    Id
	Components []uint32
}

func (c *OpVectorShuffle) Opcode() uint32 { return opcodeVectorShuffle }
func (c *OpVectorShuffle) Optional() bool { return false }
func (c *OpVectorShuffle) Verify() error  { return nil }

// OpCompositeConstruct constructs a new composite object from a set of
// constituent objects that will fully form it
type OpCompositeConstruct struct {
	ResultType   Id
	ResultId     Id
	Constituents []Id
}

func (c *OpCompositeConstruct) Opcode() uint32 { return opcodeCompositeConstruct }
func (c *OpCompositeConstruct) Optional() bool { return false }
func (c *OpCompositeConstruct) Verify() error  { return nil }

// OpCompositeExtract extracts a part of a composite object.
type OpCompositeExtract struct {
	ResultType Id
	ResultId   Id
	Composite  Id
	Indices    []uint32
}

func (c *OpCompositeExtract) Opcode() uint32 { return opcodeCompositeExtract }
func (c *OpCompositeExtract) Optional() bool { return false }
func (c *OpCompositeExtract) Verify() error  { return nil }

// OpCompositeInsert inserts into a composite object.
type OpCompositeInsert struct {
	ResultType Id
	ResultId   Id
	Object     Id
	Composite  Id
	Indices    []uint32
}

func (c *OpCompositeInsert) Opcode() uint32 { return opcodeCompositeInsert }
func (c *OpCompositeInsert) Optional() bool { return false }
func (c *OpCompositeInsert) Verify() error  { return nil }

// OpCopyObject makes a copy of Operand.
// There are no dereferences involved.
type OpCopyObject struct {
	ResultType Id
	ResultId   Id
	Operand    Id
}

func (c *OpCopyObject) Opcode() uint32 { return opcodeCopyObject }
func (c *OpCopyObject) Optional() bool { return false }
func (c *OpCopyObject) Verify() error  { return nil }

// OpTranspose transposes a matrix.
type OpTranspose struct {
	ResultType Id
	ResultId   Id
	Matrix     Id
}

func (c *OpTranspose) Opcode() uint32 { return opcodeTranspose }
func (c *OpTranspose) Optional() bool { return false }
func (c *OpTranspose) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpVectorExtractDynamic{} })
	bind(func() Instruction { return &OpVectorInsertDynamic{} })
	bind(func() Instruction { return &OpVectorShuffle{} })
	bind(func() Instruction { return &OpCompositeConstruct{} })
	bind(func() Instruction { return &OpCompositeExtract{} })
	bind(func() Instruction { return &OpCompositeInsert{} })
	bind(func() Instruction { return &OpCopyObject{} })
	bind(func() Instruction { return &OpTranspose{} })
}
