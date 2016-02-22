// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpConvertFToU converts (value preserving) Float Value from floating
// point to unsigned integer, with rounding toward 0
type OpConvertFToU struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertFToU) Opcode() uint32 { return opcodeConvertFToU }
func (c *OpConvertFToU) Optional() bool { return false }
func (c *OpConvertFToU) Verify() error  { return nil }

// OpConvertFToS converts (value preserving) Float Value from floating
// point to signed integer, with round toward 0
type OpConvertFToS struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertFToS) Opcode() uint32 { return opcodeConvertFToS }
func (c *OpConvertFToS) Optional() bool { return false }
func (c *OpConvertFToS) Verify() error  { return nil }

// OpConvertSToF converts (value preserving) Signed Value from signed integer
// to floating point.
type OpConvertSToF struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertSToF) Opcode() uint32 { return opcodeConvertSToF }
func (c *OpConvertSToF) Optional() bool { return false }
func (c *OpConvertSToF) Verify() error  { return nil }

// OpConvertUToF converts (value preserving) Unsigned value from unsigned
// integer to floating point
type OpConvertUToF struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertUToF) Opcode() uint32 { return opcodeConvertUToF }
func (c *OpConvertUToF) Optional() bool { return false }
func (c *OpConvertUToF) Verify() error  { return nil }

// OpUConvert converts (value preserving) the width of Unsigned value.
// This is either a truncate or a zero extend.
type OpUConvert struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpUConvert) Opcode() uint32 { return opcodeUConvert }
func (c *OpUConvert) Optional() bool { return false }
func (c *OpUConvert) Verify() error  { return nil }

// OpSConvert converts (value preserving) the width of Signed Value.
// This is either a truncate or a sign extend.
type OpSConvert struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpSConvert) Opcode() uint32 { return opcodeSConvert }
func (c *OpSConvert) Optional() bool { return false }
func (c *OpSConvert) Verify() error  { return nil }

// OpFConvert converts (value preserving) the width of Float Value.
//
// Results are computed per component. The operand’s type and Result Type must
// have the same number of components. The widths of the components of the
// operand and the Result Type must be different.
type OpFConvert struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpFConvert) Opcode() uint32 { return opcodeFConvert }
func (c *OpFConvert) Optional() bool { return false }
func (c *OpFConvert) Verify() error  { return nil }

// OpConvertPtrToU converts Pointer to an unsigned integer type. A Result Type
// width larger than the width of Pointer will zero extend.
type OpConvertPtrToU struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertPtrToU) Opcode() uint32 { return opcodeConvertPtrToU }
func (c *OpConvertPtrToU) Optional() bool { return false }
func (c *OpConvertPtrToU) Verify() error  { return nil }

// OpConvertUToPtr converts Integer value to a pointer. A Result Type width
// smaller than the width of Integer value pointer will truncate.
type OpConvertUToPtr struct {
	ResultType Id
	ResultId   Id
	Value      Id
}

func (c *OpConvertUToPtr) Opcode() uint32 { return opcodeConvertUToPtr }
func (c *OpConvertUToPtr) Optional() bool { return false }
func (c *OpConvertUToPtr) Verify() error  { return nil }

// OpPtrCastToGeneric converts Source pointer to a pointer value pointing to
// storage class Generic. Source pointer must point to storage class
// WorkgroupLocal, WorkgroupGlobal or Private.
type OpPtrCastToGeneric struct {
	ResultType Id
	ResultId   Id
	Source     Id
}

func (c *OpPtrCastToGeneric) Opcode() uint32 { return opcodePtrCastToGeneric }
func (c *OpPtrCastToGeneric) Optional() bool { return false }
func (c *OpPtrCastToGeneric) Verify() error  { return nil }

// OpGenericCastToPtr converts Source pointer to a non-Generic storage-class
// pointer value. Source pointer must point to Generic.
type OpGenericCastToPtr struct {
	ResultType Id
	ResultId   Id
	Source     Id
}

func (c *OpGenericCastToPtr) Opcode() uint32 { return opcodeGenericCastToPtr }
func (c *OpGenericCastToPtr) Optional() bool { return false }
func (c *OpGenericCastToPtr) Verify() error  { return nil }

// OpBitcast defines a Bit-pattern preserving type conversion for
// Numerical-type or pointer-type vectors and scalars.
type OpBitcast struct {
	ResultType Id
	ResultId   Id
	Operand    Id // Operand is the bit pattern whose type will change
}

func (c *OpBitcast) Opcode() uint32 { return opcodeBitcast }
func (c *OpBitcast) Optional() bool { return false }
func (c *OpBitcast) Verify() error  { return nil }

// OpGenericCastToPtrExplicit attempts to explicitly convert Source pointer
// to storage storage-class pointer value.
type OpGenericCastToPtrExplicit struct {
	ResultType Id
	ResultId   Id
	SourcePtr  Id
	Storage    StorageClass
}

func (c *OpGenericCastToPtrExplicit) Opcode() uint32 { return opcodeGenericCastToPtrExplicit }
func (c *OpGenericCastToPtrExplicit) Optional() bool { return false }
func (c *OpGenericCastToPtrExplicit) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpConvertFToU{} })
	bind(func() Instruction { return &OpConvertFToS{} })
	bind(func() Instruction { return &OpConvertSToF{} })
	bind(func() Instruction { return &OpConvertUToF{} })
	bind(func() Instruction { return &OpUConvert{} })
	bind(func() Instruction { return &OpSConvert{} })
	bind(func() Instruction { return &OpFConvert{} })
	bind(func() Instruction { return &OpConvertPtrToU{} })
	bind(func() Instruction { return &OpConvertUToPtr{} })
	bind(func() Instruction { return &OpPtrCastToGeneric{} })
	bind(func() Instruction { return &OpGenericCastToPtr{} })
	bind(func() Instruction { return &OpBitcast{} })
	bind(func() Instruction { return &OpGenericCastToPtrExplicit{} })
}
