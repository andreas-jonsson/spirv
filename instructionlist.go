// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// InstructionList defines a list of instructions with some
// utility functions.
type InstructionList []Instruction

// Functions returns all instruction subsets defining a function.
func (set InstructionList) Functions() []InstructionList {
	var out []InstructionList

	start := set.FilterIndex(opcodeFunction, 0)
	end := set.FilterIndex(opcodeFunctionEnd, 0)

	if len(start) != len(end) {
		return nil
	}

	for i, s := range start {
		e := end[i]
		out = append(out, set[s:e+1])
	}

	return out
}

// Blocks returns all function blocks. This assumes the given set
// is itself just one function. Otherwise it will return blocks for
// multiple- or all functions.
func (set InstructionList) Blocks() []InstructionList {
	var out []InstructionList

	start := set.FilterIndex(opcodeLabel, 0)
	end := set.FilterIndex(opcodeBranch, 0)

	if len(start) != len(end) {
		return nil
	}

	for i, s := range start {
		e := end[i]
		out = append(out, set[s:e+1])
	}

	return out
}

// First returns the first instruction with the given opcode.
// Returns nil if it could not be found.
func (set InstructionList) First(opcode uint32) Instruction {
	idx := set.Index(opcode)
	if idx > -1 {
		return set[idx]
	}
	return nil
}

// Index returns the index of the first instance of the given opcode.
// Returns -1 if none was found.
func (set InstructionList) Index(opcode uint32) int {
	for i, v := range set {
		if v.Opcode() == opcode {
			return i
		}
	}

	return -1
}

// Filter returns all instructions with the given opcode.
func (set InstructionList) Filter(opcode uint32) InstructionList {
	out := make(InstructionList, 0, len(set))

	for _, v := range set {
		if v.Opcode() == opcode {
			out = append(out, v)
		}
	}

	return out
}

// FilterIndex returns all instruction indices with the given opcode.
//
// The offset value is optionally added to a resulting index. This can be useful
// if you want indices relative to some other point than the start of the
// provided set.
func (set InstructionList) FilterIndex(opcode uint32, offset int) []int {
	out := make([]int, 0, len(set))

	for i, v := range set {
		if v.Opcode() == opcode {
			out = append(out, i+offset)
		}
	}

	return out
}

// Count returns the number of instances of the given instruction in the set.
func (set InstructionList) Count(opcode uint32) int {
	var count int

	for _, v := range set {
		if v.Opcode() == opcode {
			count++
		}
	}

	return count
}

// localVariables returns all local variables defined in all functions.
func (set InstructionList) localVariables() []int {
	start := set.FilterIndex(opcodeFunction, 0)
	end := set.FilterIndex(opcodeFunctionEnd, 0)

	if len(start) != len(end) {
		return nil
	}

	var out []int

	for i, s := range start {
		e := end[i]
		v := set[s:e].FilterIndex(opcodeVariable, s)
		out = append(out, v...)
	}

	return out
}

// globalVariables returns all global variables defined in the set
func (set InstructionList) globalVariables() []int {
	funcIndex := set.Index(opcodeFunction)
	return set[:funcIndex].FilterIndex(opcodeVariable, 0)
}
