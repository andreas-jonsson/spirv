// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"reflect"
	"sort"
)

// InstructionFunc defines a constructor for an instruction.
type instructionFunc func() Instruction

// InstructionSet maps opcodes to an instruction  constructor.
type instructionSet map[uint32]instructionFunc

// Opcodes returns a sorted list of all registered opcodes.
func (set instructionSet) Opcodes() []int {
	out := make([]int, 0, len(set))

	for opcode := range set {
		out = append(out, int(opcode))
	}

	sort.Ints(out)
	return out
}

// Global, internal instruction set.
// This has instructions registered atomically during init.
var instructions = make(instructionSet)

// Bind registers the given instruction.
//
// This call panics if the instruction type defined by the constructor
// is not a pointer type. Additionally, this panics if there is already an
// entry for the instruction's opcode.
//
// All instructions are meant to be registered during package initialisation
func bind(fun instructionFunc) {
	obj := fun()
	rv := reflect.ValueOf(obj)

	if rv.Kind() != reflect.Ptr {
		panic(ErrInstructionNotPointer)
	}

	opcode := obj.Opcode()

	_, ok := instructions[opcode]
	if ok {
		panic(ErrDuplicateInstruction)
	}

	instructions[opcode] = fun
}
