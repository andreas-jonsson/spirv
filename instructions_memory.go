// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpVariable allocates an object in memory, resulting in a pointer
// to it, which can be used with OpLoad and OpStore.
type OpVariable struct {
	// Result Type is a type from OpTypePointer, where the type pointed to
	// is the type of object in memory.
	ResultType Id
	ResultId   Id

	// Storage Class is the kind of memory holding the object.
	StorageClass StorageClass

	// Initializer is optional. If Initializer is present, it will be the
	// initial value of the variable’s memory content. Initializer must
	// be an <id> from a constant instruction. Initializer must have the same
	// type as the type pointed to by Result Type.
	Initializer Id `spirv:"optional"`
}

func (c *OpVariable) Opcode() uint32 { return opcodeVariable }
func (c *OpVariable) Optional() bool { return false }
func (c *OpVariable) Verify() error  { return nil }

// OpVariableArray allocates N objects sequentially in memory,
// resulting in a pointer to the first such object.
//
// This is not the same thing as allocating a single object that is an array.
type OpVariableArray struct {
	// Result Type is a type from OpTypePointer whose type pointed to is
	// the type of one of the N objects allocated in memory
	ResultType Id

	ResultId Id

	// Storage Class is the kind of memory holding the object.
	StorageClass StorageClass

	// N is the number of objects to allocate.
	N Id
}

func (c *OpVariableArray) Opcode() uint32 { return opcodeVariableArray }
func (c *OpVariableArray) Optional() bool { return false }
func (c *OpVariableArray) Verify() error  { return nil }

// OpLoad loads data through a pointer.
type OpLoad struct {
	// Result Type is a type from OpTypePointer whose type pointed to is
	// the type of one of the N objects allocated in memory
	ResultType Id

	ResultId Id

	// Pointer is the pointer to load through. It must have a type of
	// OpTypePointer whose operand is the same as Result Type.
	Pointer Id

	// MemoryAccess must be one or more Memory Access literals.
	// FIXME: Incorrect type in specification.
	MemoryAccess []MemoryAccess
}

func (c *OpLoad) Opcode() uint32 { return opcodeLoad }
func (c *OpLoad) Optional() bool { return false }
func (c *OpLoad) Verify() error  { return nil }

// OpStore stores data through a pointer.
type OpStore struct {
	// Pointer is the pointer to store through. It must have a type
	// of OpTypePointer whose operand is the same as the type of Object.
	Pointer Id

	// Object is the object to store.
	Object Id

	// MemoryAccess must be one or more Memory Access literals.
	// FIXME: Incorrect type in specification.
	MemoryAccess []MemoryAccess
}

func (c *OpStore) Opcode() uint32 { return opcodeStore }
func (c *OpStore) Optional() bool { return false }
func (c *OpStore) Verify() error  { return nil }

// OpCopyMemory copies from the memory pointed to by Source to the
// memory pointed to by Target.
//
// Both operands must be non-void pointers of the same type.
// Matching storage class is not required. The amount of memory copied is
// the size of the type pointed to.
type OpCopyMemory struct {
	// The target address.
	Target Id

	// The source address.
	Source Id

	// MemoryAccess must be one or more Memory Access literals.
	// FIXME: Incorrect type in specification.
	MemoryAccess []MemoryAccess
}

func (c *OpCopyMemory) Opcode() uint32 { return opcodeCopyMemory }
func (c *OpCopyMemory) Optional() bool { return false }
func (c *OpCopyMemory) Verify() error  { return nil }

// OpCopyMemorySized copies from the memory pointed to by Source to the
// memory pointed to by Target.
//
// Both operands must be non-void pointers of the same type.
// Matching storage class is not required.
type OpCopyMemorySized struct {
	// The target address.
	Target Id

	// The source address.
	Source Id

	// Size is the number of bytes to copy.
	Size Id

	// MemoryAccess must be one or more Memory Access literals.
	// FIXME: Incorrect type in specification.
	MemoryAccess []MemoryAccess
}

func (c *OpCopyMemorySized) Opcode() uint32 { return opcodeCopyMemorySized }
func (c *OpCopyMemorySized) Optional() bool { return false }
func (c *OpCopyMemorySized) Verify() error  { return nil }

// OpAccessChain creates a pointer into a composite object that can be
// used with OpLoad and OpStore.
//
// The storage class of the pointer created will be the same as the storage
// class of the base operand.
type OpAccessChain struct {
	ResultType Id
	ResultId   Id

	// Base must be a pointer type, pointing to the base of the object.
	Base Id

	// Indices walk the type hierarchy to the desired depth, potentially
	// down to scalar granularity. The type of the pointer created will be to
	// the type reached by walking the type hierarchy down to the last
	// provided index.
	Indices []Id
}

func (c *OpAccessChain) Opcode() uint32 { return opcodeAccessChain }
func (c *OpAccessChain) Optional() bool { return false }
func (c *OpAccessChain) Verify() error  { return nil }

// OpInboundsAccessChain has the same semantics as OpAccessChain, with the
// addition that the resulting pointer is known to point within the base object.
type OpInboundsAccessChain struct {
	ResultType Id
	ResultId   Id

	// Base must be a pointer type, pointing to the base of the object.
	Base Id

	// Indices walk the type hierarchy to the desired depth, potentially
	// down to scalar granularity. The type of the pointer created will be to
	// the type reached by walking the type hierarchy down to the last
	// provided index.
	Indices []Id
}

func (c *OpInboundsAccessChain) Opcode() uint32 { return opcodeInboundsAccessChain }
func (c *OpInboundsAccessChain) Optional() bool { return false }
func (c *OpInboundsAccessChain) Verify() error  { return nil }

// OpArraylength results in the length of a run-time array.
type OpArraylength struct {
	ResultType Id
	ResultId   Id

	// Structure must be an object of type OpTypeStruct that contains
	// a member that is a run-time array.
	Structure Id

	// Array member is a member number of Structure that must have a
	// type from OpTypeRuntimeArray.
	Member uint32
}

func (c *OpArraylength) Opcode() uint32 { return opcodeArraylength }
func (c *OpArraylength) Optional() bool { return false }
func (c *OpArraylength) Verify() error  { return nil }

// OpImagePointer forms a pointer to a texel of an image.
// Use of such a pointer is limited to atomic operations.
//
// TODO: This requires an Image storage class to be added.
type OpImagePointer struct {
	ResultType Id
	ResultId   Id

	// Image is a pointer to a variable of type of OpTypeSampler.
	Image Id

	// Coordinate and Sample specify which texel and sample within
	// the image to form an address of.
	Coordinate Id
	Sample     Id
}

func (c *OpImagePointer) Opcode() uint32 { return opcodeImagePointer }
func (c *OpImagePointer) Optional() bool { return false }
func (c *OpImagePointer) Verify() error  { return nil }

// OpGenericPtrMemSemantics returns a valid Memory Semantics
// value for ptr.
type OpGenericPtrMemSemantics struct {
	ResultType Id // Result Type must be a 32-bits wide OpTypeInt value
	ResultId   Id
	Ptr        Id
}

func (c *OpGenericPtrMemSemantics) Opcode() uint32 { return opcodeGenericPtrMemSemantics }
func (c *OpGenericPtrMemSemantics) Optional() bool { return false }
func (c *OpGenericPtrMemSemantics) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpVariable{Initializer: 0} })
	bind(func() Instruction { return &OpVariableArray{} })
	bind(func() Instruction { return &OpLoad{} })
	bind(func() Instruction { return &OpStore{} })
	bind(func() Instruction { return &OpCopyMemory{} })
	bind(func() Instruction { return &OpCopyMemorySized{} })
	bind(func() Instruction { return &OpAccessChain{} })
	bind(func() Instruction { return &OpInboundsAccessChain{} })
	bind(func() Instruction { return &OpArraylength{} })
	bind(func() Instruction { return &OpImagePointer{} })
	bind(func() Instruction { return &OpGenericPtrMemSemantics{} })
}
