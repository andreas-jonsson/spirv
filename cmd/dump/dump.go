// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/andreas-jonsson/spirv"
)

// dump prints a human-readable dump of the given module.
func dump(mod *spirv.Module) {
	dumpHeader(mod.Header)

	fmt.Printf("Instructions (%d):\n", len(mod.Code))
	for _, instr := range mod.Code {
		dumpInstruction(instr)
	}
}

// dumpHeader prints a human-readable dump of the given header.
func dumpHeader(hdr spirv.Header) {
	fmt.Printf("Version:      %d\n", hdr.Version)
	fmt.Printf("Generator:    0x%08x\n", hdr.GeneratorMagic)
	fmt.Printf("Id Bound:     %d\n", hdr.Bound)
}

// dumpInstruction prints a human-readable dump of the given instruction.
func dumpInstruction(i spirv.Instruction) {
	name := fmt.Sprintf("%T", i)
	name = name[len("*spirv.Op"):]
	fields := instructionFields(i)

	fmt.Println(" ", name)

	for _, fld := range fields {
		fmt.Printf("    %-25s: %s\n", fld[0], fld[1])
	}
}

// instructionFields returns a list of string versions for the given
// instruction's fields.
func instructionFields(i spirv.Instruction) [][]string {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)
	rt := rv.Type()

	out := make([][]string, rv.NumField())

	for i := range out {
		fv := rv.Field(i)
		ft := rt.Field(i)

		v := fv.Interface()
		k := fmt.Sprintf("%T", v)
		k = strings.Replace(k, "spirv.", "", -1)

		out[i] = []string{
			ft.Name,
			fmt.Sprintf("%s(%v)", k, v),
		}
	}

	return out
}
