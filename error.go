// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"errors"
	"fmt"
)

var (
	ErrUnexpectedEOF          = errors.New("unexpected EOF")
	ErrInvalidInstructionSize = errors.New("instruction has invalid size")
	ErrMissingInstructionArgs = errors.New("insufficient instruction arguments")
	ErrUnacceptable           = errors.New("use of this instruction is not allowed")
	ErrInstructionNotPointer  = errors.New("value from Codec.New is not a pointer type")
	ErrDuplicateInstruction   = errors.New("duplicate opcode being registered")
	ErrInvalidMagicValue      = errors.New("Header: invalid magic value")
	ErrInvalidVersion         = errors.New("Header: invalid version number")
	ErrMemoryModel            = errors.New("a module must define one and only one OpMemoryModel")
	ErrEntrypoint             = errors.New("a module must define at least one OpEntrypoint")
	ErrExecutionMode          = errors.New("a module must define at least one OpExecutionMode")
)

// LayoutError defines an error in a module's structural layout.
type LayoutError struct {
	Msg     string
	Address int
}

// NewLayoutError creates a new layout error for the given address and
// formatted message.
func NewLayoutError(addr int, msg string, argv ...interface{}) *LayoutError {
	return &LayoutError{
		Msg:     fmt.Sprintf(msg, argv...),
		Address: addr,
	}
}

func (e *LayoutError) Error() string {
	return fmt.Sprintf("at $%08x: %s", e.Address, e.Msg)
}
