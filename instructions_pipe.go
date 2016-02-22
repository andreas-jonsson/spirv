// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpReadPipe reads a packet from P into Ptr.
type OpReadPipe struct {
	ResultType Id
	ResultId   Id
	P          Id
	Ptr        Id
}

func (c *OpReadPipe) Opcode() uint32 { return opcodeReadPipe }
func (c *OpReadPipe) Optional() bool { return false }
func (c *OpReadPipe) Verify() error  { return nil }

// OpWritePipe writes a packet from Ptr to P.
type OpWritePipe struct {
	ResultType Id
	ResultId   Id
	P          Id
	Ptr        Id
}

func (c *OpWritePipe) Opcode() uint32 { return opcodeWritePipe }
func (c *OpWritePipe) Optional() bool { return false }
func (c *OpWritePipe) Verify() error  { return nil }

// OpReservedReadPipe reads a packet from the reserved area specified by
// ReserveId and Index from P into Ptr.
type OpReservedReadPipe struct {
	ResultType Id
	ResultId   Id
	P          Id
	ReserveId  Id
	Index      Id
	Ptr        Id
}

func (c *OpReservedReadPipe) Opcode() uint32 { return opcodeReservedReadPipe }
func (c *OpReservedReadPipe) Optional() bool { return false }
func (c *OpReservedReadPipe) Verify() error  { return nil }

// OpReservedWritePipe writes a packet from Ptr into the reserved area
// specified by ReserveId and Index into P.
type OpReservedWritePipe struct {
	ResultType Id
	ResultId   Id
	P          Id
	ReserveId  Id
	Index      Id
	Ptr        Id
}

func (c *OpReservedWritePipe) Opcode() uint32 { return opcodeReservedWritePipe }
func (c *OpReservedWritePipe) Optional() bool { return false }
func (c *OpReservedWritePipe) Verify() error  { return nil }

// OpReserveReadPipePackets reserves NumPackets entries for reading from P.
type OpReserveReadPipePackets struct {
	ResultType Id
	ResultId   Id
	P          Id
	NumPackets Id
}

func (c *OpReserveReadPipePackets) Opcode() uint32 { return opcodeReserveReadPipePackets }
func (c *OpReserveReadPipePackets) Optional() bool { return false }
func (c *OpReserveReadPipePackets) Verify() error  { return nil }

// OpReserveWritePipePackets reserves NumPackets entries for writing to P.
type OpReserveWritePipePackets struct {
	ResultType Id
	ResultId   Id
	P          Id
	NumPackets Id
}

func (c *OpReserveWritePipePackets) Opcode() uint32 { return opcodeReserveWritePipePackets }
func (c *OpReserveWritePipePackets) Optional() bool { return false }
func (c *OpReserveWritePipePackets) Verify() error  { return nil }

// OpCommitReadPipe indicates that all reads to NumPackets associated with
// ReserveId in P are completed.
type OpCommitReadPipe struct {
	P         Id
	ReserveId Id
}

func (c *OpCommitReadPipe) Opcode() uint32 { return opcodeCommitReadPipe }
func (c *OpCommitReadPipe) Optional() bool { return false }
func (c *OpCommitReadPipe) Verify() error  { return nil }

// OpCommitWritePipe indicates that all writes to NumPackets associated
// with ReserveId in P are completed.
type OpCommitWritePipe struct {
	P         Id
	ReserveId Id
}

func (c *OpCommitWritePipe) Opcode() uint32 { return opcodeCommitWritePipe }
func (c *OpCommitWritePipe) Optional() bool { return false }
func (c *OpCommitWritePipe) Verify() error  { return nil }

// OpIsValidReserveId returns true if ReserveId is a valid reservation ID
// and false otherwise.
type OpIsValidReserveId struct {
	ResultType Id
	ResultId   Id
	ReserveId  Id
}

func (c *OpIsValidReserveId) Opcode() uint32 { return opcodeIsValidReserveId }
func (c *OpIsValidReserveId) Optional() bool { return false }
func (c *OpIsValidReserveId) Verify() error  { return nil }

// OpGetNumPipePackets returns the number of available entries in P.
type OpGetNumPipePackets struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpGetNumPipePackets) Opcode() uint32 { return opcodeGetNumPipePackets }
func (c *OpGetNumPipePackets) Optional() bool { return false }
func (c *OpGetNumPipePackets) Verify() error  { return nil }

// OpGetMaxPipePackets returns the maximum number of packets specified when
// P was created.
type OpGetMaxPipePackets struct {
	ResultType Id
	ResultId   Id
	P          Id
}

func (c *OpGetMaxPipePackets) Opcode() uint32 { return opcodeGetMaxPipePackets }
func (c *OpGetMaxPipePackets) Optional() bool { return false }
func (c *OpGetMaxPipePackets) Verify() error  { return nil }

// OpGroupReserveReadPipePackets reserves NumPackets entries for reading
// from P at group level.
type OpGroupReserveReadPipePackets struct {
	ResultType Id
	ResultId   Id
	Scope      ExecutionScope
	P          Id
	NumPackets Id
}

func (c *OpGroupReserveReadPipePackets) Opcode() uint32 { return opcodeGroupReserveReadPipePackets }
func (c *OpGroupReserveReadPipePackets) Optional() bool { return false }
func (c *OpGroupReserveReadPipePackets) Verify() error  { return nil }

// OpGroupReserveWritePipePackets reserves NumPackets entries for writing
// to P at group level.
type OpGroupReserveWritePipePackets struct {
	ResultType Id
	ResultId   Id
	Scope      ExecutionScope
	P          Id
	NumPackets Id
}

func (c *OpGroupReserveWritePipePackets) Opcode() uint32 { return opcodeGroupReserveWritePipePackets }
func (c *OpGroupReserveWritePipePackets) Optional() bool { return false }
func (c *OpGroupReserveWritePipePackets) Verify() error  { return nil }

// OpGroupCommitReadPipe is a group level indication that all reads to
// NumPackets associated with ReserveId to P are completed.
type OpGroupCommitReadPipe struct {
	Scope     ExecutionScope
	P         Id
	ReserveId Id
}

func (c *OpGroupCommitReadPipe) Opcode() uint32 { return opcodeGroupCommitReadPipe }
func (c *OpGroupCommitReadPipe) Optional() bool { return false }
func (c *OpGroupCommitReadPipe) Verify() error  { return nil }

// OpGroupCommitWritePipe is a group level indication that all writes to
// NumPackets associated with ReserveId to P are completed.
type OpGroupCommitWritePipe struct {
	Scope     ExecutionScope
	P         Id
	ReserveId Id
}

func (c *OpGroupCommitWritePipe) Opcode() uint32 { return opcodeGroupCommitWritePipe }
func (c *OpGroupCommitWritePipe) Optional() bool { return false }
func (c *OpGroupCommitWritePipe) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpReadPipe{} })
	bind(func() Instruction { return &OpWritePipe{} })
	bind(func() Instruction { return &OpReservedReadPipe{} })
	bind(func() Instruction { return &OpReservedWritePipe{} })
	bind(func() Instruction { return &OpReserveReadPipePackets{} })
	bind(func() Instruction { return &OpReserveWritePipePackets{} })
	bind(func() Instruction { return &OpCommitReadPipe{} })
	bind(func() Instruction { return &OpCommitWritePipe{} })
	bind(func() Instruction { return &OpIsValidReserveId{} })
	bind(func() Instruction { return &OpGetNumPipePackets{} })
	bind(func() Instruction { return &OpGetMaxPipePackets{} })
	bind(func() Instruction { return &OpGroupReserveReadPipePackets{} })
	bind(func() Instruction { return &OpGroupReserveWritePipePackets{} })
	bind(func() Instruction { return &OpGroupCommitReadPipe{} })
	bind(func() Instruction { return &OpGroupCommitWritePipe{} })
}
