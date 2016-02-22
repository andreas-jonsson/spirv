// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "io"

// Module defines a complete SPIR-V module.
type Module struct {
	Header Header
	Code   InstructionList
}

// NewModule creates a new, default module.
func NewModule() *Module {
	return &Module{
		Header: Header{
			Magic:          MagicLE,
			Version:        SpecificationVersion,
			GeneratorMagic: 0,
			Bound:          0,
			Reserved:       0,
		},
	}
}

// Load loads a full module from the given input stream.
func Load(r io.Reader) (*Module, error) {
	var mod Module
	var err error
	var instr Instruction

	dec := NewDecoder(r)

	// Load the module header.
	mod.Header, err = dec.DecodeHeader()
	if err != nil {
		return nil, err
	}

	// Load all instructions.
	mod.Code = make([]Instruction, 0, 128)

	for {
		instr, err = dec.DecodeInstruction()
		if err != nil {
			if err == io.EOF {
				break // Not an error -- just end of stream.
			}

			return nil, err
		}

		mod.Code = append(mod.Code, instr)
	}

	return &mod, nil
}

// Save writes the module to the given stream.
func (m *Module) Save(w io.Writer) error {
	enc := NewEncoder(w)

	// Write the header.
	err := enc.EncodeHeader(m.Header)
	if err != nil {
		return err
	}

	// Write all instructions.
	for _, instr := range m.Code {
		err := enc.EncodeInstruction(instr)
		if err != nil {
			return err
		}
	}

	return nil
}

// Verify returns an error if the module contains invalid data.
//
// This applies a host of different levels of structural and semantic
// validation as defined in the spec chapter 2.15.
func (m *Module) Verify() error {
	// Check the header for structural validity.
	err := m.Header.Verify()
	if err != nil {
		return err
	}

	// Perform structural validity checks on each instruction, before
	// we proceed to the semantic checks.
	//
	// This uses reflection to call Verify() on all relevant struct fields
	// and then on the instruction itself. The latter is used by some
	// instructions to validate parts which can not be caught by the field
	// types themselves.
	for _, instr := range m.Code {
		err := verifyInstruction(instr)
		if err != nil {
			return err
		}
	}

	// Ensure logical layout is up to standards.
	err = m.verifyLogicalLayout()
	if err != nil {
		return err
	}

	// TODO: Non-structure types (scalars, vectors, arrays, etc.) with the same
	// operand parameterization cannot be type aliases. For nonstructures, two
	// type <id>s match iff the types match.

	// Perform some checks related to Logical addressing mode.
	err = m.verifyLogicalAddressing()
	if err != nil {
		return err
	}

	// Check the validity of result IDs.
	err = m.verifySSA()
	if err != nil {
		return err
	}

	// Check validity of entry point usage.
	err = m.verifyEntrypoints()
	if err != nil {
		return err
	}

	return nil
}

// Strip removes all instructions which have no semantic impact on the code.
// This includes debug symbols like source context and names.
func (m *Module) Strip() {
	for i := 0; i < len(m.Code); i++ {
		if !m.Code[i].Optional() {
			continue
		}

		copy(m.Code[i:], m.Code[i+1:])
		m.Code[len(m.Code)-1] = nil
		m.Code = m.Code[:len(m.Code)-1]
		i--
	}
}

// verifyEntrypoints performs some sanity checks on entrypoint definitions.
func (m *Module) verifyEntrypoints() error {
	// There must be at least 1 OpEntryPoint, unless a Link capability
	// is being used.
	if !m.hasLinkageType() && m.Code.Count(opcodeEntryPoint) == 0 {
		return NewLayoutError(0, "unless the Linkage capabilities are used, we require at least 1 OpEntryPoint")
	}

	// No function can be targeted by both an OpEntryPoint instruction and an
	// OpFunctionCall instruction.
	entries := m.Code.FilterIndex(opcodeEntryPoint, 0)
	calls := m.Code.FilterIndex(opcodeFunctionCall, 0)

	if len(entries) == 0 || len(calls) == 0 {
		return nil
	}

	for _, c := range calls {
		fc := m.Code[c].(*OpFunctionCall)
		addr := m.hasResultId(fc.Function, entries)
		if addr > -1 {
			return NewLayoutError(
				c, "call to function previously defined as entrypoint at $%08x", addr,
			)
		}
	}

	return nil
}

// verifySSA performs some sanity checks on result id values, as defined
// in chapter 2.15.1 of the specification.
func (m *Module) verifySSA() error {
	// Each <id> must appear exactly once as the result <id> of an instruction.
	//
	// Create a list of all result ids and ensure there are no duplicates.
	set := make(map[Id]int)

	for addr, instr := range m.Code {
		if _, ok := instr.(*OpEntryPoint); ok {
			// ResultId is entry point target. Not id for this
			// instruction itself -- ignore it.
			continue
		}

		id, ok := instructionResultId(instr)
		if !ok {
			continue
		}

		paddr, ok := set[id]
		if !ok {
			set[id] = addr
			continue
		}

		return NewLayoutError(
			addr, "duplicate ResultId(%d); previous definition at: $%08x",
			id, paddr,
		)
	}

	// TODO: The definition of an SSA <id> should dominate all uses of it,
	// with the following exceptions :... (see spec)

	return nil
}

// hasResultId returns the addr of a list entry, if src has a ResultId
// matching any of the Result Ids in list. Returns -1 otherwise.
func (m *Module) hasResultId(src Id, list []int) int {
	for _, addr := range list {
		dst, ok := instructionResultId(m.Code[addr])
		if ok && src == dst {
			return addr
		}
	}

	return -1
}

// hasLinkageType returns true if there is a OpDecorate instance
// on a (global) variable or function with a LinkageType defined.
func (m *Module) hasLinkageType() bool {
	decorations := m.Code.Filter(opcodeDecorate)

	for _, i := range decorations {
		v := i.(*OpDecorate)
		if v.Decoration == DecorationLinkageType {
			return true
		}
	}

	return false
}

// verifyLogicalAddressing performs a number of checks if the logical
// addressing mode is selected for this module.
func (m *Module) verifyLogicalAddressing() error {
	v := m.Code.First(opcodeMemoryModel).(*OpMemoryModel)

	if v.AddressingModel != AddressingModeLogical {
		// These rules apply only to AddressingModeLogical
		return nil
	}

	// Variables can not allocate pointer types.

	// FIXME: Ask for clarification on meaning of OpVariable rules.
	// ref: https://github.com/jteeuwen/spirv/issues/25
	return nil
}

// verifyLogicalLayout ensures the module meets the Logical Layout
// requirements as defined in the spec chapter 2.4.
func (m *Module) verifyLogicalLayout() error {
	// We must have one and only one OpmemoryModel.
	//
	// This will be caught by the regex match below, but here we can be
	// more specific with our error message.
	if m.Code.Count(opcodeMemoryModel) != 1 {
		return ErrMemoryModel
	}

	// We must have at least one OpEntryPoint.
	//
	// This will be caught by the regex match below, but here we can be
	// more specific with our error message.
	if m.Code.Count(opcodeEntryPoint) == 0 {
		return ErrEntrypoint
	}

	// We must have at least one OpExecutionMode.
	//
	// This will be caught by the regex match below, but here we can be
	// more specific with our error message.
	if m.Code.Count(opcodeExecutionMode) == 0 {
		return ErrExecutionMode
	}

	// Test instruction order.
	err := verifyLayoutPattern(m.Code)
	if err != nil {
		return err
	}

	// Some instructions have requirements beyond what can be tested
	// with a regular expression pattern.

	// Global Variables must not have StorageClassFunction.
	err = verifyGlobalVariables(m.Code)
	if err != nil {
		return err
	}

	// Local Variables must have StorageClassFunction.
	err = verifyLocalVariables(m.Code)
	if err != nil {
		return err
	}

	// All local variables must be the first instructions in the first block.
	return verifyFunctionStructure(m.Code)
}
