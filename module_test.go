// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"bytes"
	"reflect"
	"testing"
)

var mod = NewModule()

func TestModuleVerifyLogicalLayout1(t *testing.T) {
	// CompileFlag instruction is too early.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpSource{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	if err == nil {
		t.Fatalf("expected failure")
	}
}

func TestModuleVerifyLogicalLayout2(t *testing.T) {
	// Valid module.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionParameter{},
		&OpLabel{},
		&OpIAdd{},
		&OpBranch{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleVerifyLogicalLayout3(t *testing.T) {
	// Faulty module with missing OpFunctionEnd.
	mod.Code = []Instruction{
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},
		&OpFunction{},
	}

	err := mod.verifyLogicalLayout()
	if err == nil {
		t.Fatalf("expected failure")
	}
}

func TestModuleVerifyLogicalLayout4(t *testing.T) {
	mod.Code = []Instruction{}
	err := mod.verifyLogicalLayout()
	want := ErrMemoryModel

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalLayout5(t *testing.T) {
	// Module with global variable. it has an unacceptable storage class.
	mod.Code = []Instruction{
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpVariable{
			StorageClass: StorageClassFunction,
		},

		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	want := NewLayoutError(3, "global variable: storage class can not be StorageClassFunction")

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalLayout6(t *testing.T) {
	// Missing Entrypoint
	mod.Code = []Instruction{
		&OpMemoryModel{},
		&OpExecutionMode{},
		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	want := ErrEntrypoint

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalLayout7(t *testing.T) {
	// Missing Execution Mode
	mod.Code = []Instruction{
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	want := ErrExecutionMode

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalLayout8(t *testing.T) {
	// Variables inside functions must have StorageClassFunction
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionParameter{},
		&OpLabel{},
		&OpVariable{
			StorageClass: StorageClassAtomicCounter,
		},
		&OpBranch{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	want := NewLayoutError(7, "local variable: storage class must be StorageClassFunction")

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalLayout9(t *testing.T) {
	// Variables inside functions must be contained in the first block.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionParameter{},

		&OpLabel{},
		&OpBranch{},

		&OpLabel{},
		&OpVariable{
			StorageClass: StorageClassFunction,
		},
		&OpBranch{},

		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	want := NewLayoutError(9, "variable definition may only appear in the first block")

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalLayout10(t *testing.T) {
	// Variables inside functions must be contained in the first block and
	// not be preceeded by any other instruction.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionParameter{},

		&OpLabel{},
		&OpIAdd{},
		&OpVariable{
			StorageClass: StorageClassFunction,
		},
		&OpBranch{},

		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalLayout()
	want := NewLayoutError(7, "variable definitions must preceed all other instructions in this block")

	if !reflect.DeepEqual(err, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, err)
	}
}

func TestModuleVerifyLogicalAddressing1(t *testing.T) {
	// Faulty module: variable allocates pointer type while
	// memory model is Logical.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{
			AddressingModel: AddressingModeLogical,
		},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpVariable{
			StorageClass: StorageClassFunction,
		},

		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyLogicalAddressing()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleVerifySSA1(t *testing.T) {
	// Faulty module: 2 identical result IDs.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{},
		&OpExecutionMode{},

		&OpFunction{ResultId: 1},
		&OpFunctionEnd{},

		&OpFunction{ResultId: 1},
		&OpFunctionEnd{},
	}

	want := NewLayoutError(6, "duplicate ResultId(%d); previous definition at: $%08x", 1, 4)
	have := mod.verifySSA()

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, have)
	}
}

func TestModuleVerifyEntrypoints1(t *testing.T) {
	// Faulty module: missing OpEntryPoint and no LinkageType decoration
	// used to offset its absence.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpExecutionMode{},

		&OpFunction{},
		&OpFunctionEnd{},
	}

	want := NewLayoutError(0, "unless the Linkage capabilities are used, we require at least 1 OpEntryPoint")
	have := mod.verifyEntrypoints()

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, have)
	}
}

func TestModuleVerifyEntrypoints2(t *testing.T) {
	// Good module: missing OpEntryPoint but a LinkageType decoration
	// is present.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpExecutionMode{},

		&OpDecorate{
			Decoration: DecorationLinkageType,
		},
		&OpFunction{},
		&OpFunctionEnd{},
	}

	err := mod.verifyEntrypoints()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleVerifyEntrypoints3(t *testing.T) {
	// Bad module: function is targeted by both OpEntryPoint and
	// OpFunctionCall. This is akin to a C/Go/etc program calling
	// main() from somewhere in the program. This is not allowed in SPIR-V.
	mod.Code = []Instruction{
		&OpCompileFlag{},
		&OpMemoryModel{},
		&OpEntryPoint{
			ExecutionModel: ExecutionModelFragment,
			ResultId:       1,
		},
		&OpExecutionMode{},

		&OpFunction{ResultId: 1},
		&OpLabel{},
		&OpFunctionCall{
			ResultId: 2,
			Function: 1,
		}, // Recursive call to ourselves.
		&OpBranch{},
		&OpFunctionEnd{},
	}

	want := NewLayoutError(6, "call to function previously defined as entrypoint at $%08x", 2)
	have := mod.verifyEntrypoints()

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("error mismatch:\nWant: %v\nHave: %v", want, have)
	}
}

func TestModuleStrip(t *testing.T) {
	mod.Code = []Instruction{
		&OpSource{SourceLanguageGLSL, 450},
		&OpCompileFlag{
			Flag: "test",
		},
		&OpMemoryModel{
			AddressingModel: AddressingModeLogical,
			MemoryModel:     MemoryModelGLSL450,
		},
	}

	mod.Strip()

	want := 2
	have := len(mod.Code)

	if have != want {
		t.Fatalf("Strip error: Expected %d remaining instructions; have: %d", want, have)
	}
}

func TestModuleRoundtrip(t *testing.T) {
	var out bytes.Buffer
	err := mod.Save(&out)
	if err != nil {
		t.Fatalf("save: %v", err)
	}

	in := bytes.NewBuffer(out.Bytes())
	modb, err := Load(in)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	if !reflect.DeepEqual(mod, modb) {
		t.Fatalf("rountrip failure:\nHave: %v\nWant: %v", mod, modb)
	}
}
