// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import "testing"

type constantTest struct {
	verify func(uint32) error
	min    uint32
	max    uint32
	err    error
	even   bool
}

func TestConstant(t *testing.T) {
	for _, ct := range []constantTest{
		{
			min: AccessQualifierReadOnly,
			max: AccessQualifierReadWrite,
			err: ErrInvalidAccessQualifier,
			verify: func(w uint32) error {
				return AccessQualifier(w).Verify()
			},
		},
		{
			min: AddressingModeLogical,
			max: AddressingModePhysical64,
			err: ErrInvalidAddressingModel,
			verify: func(w uint32) error {
				return AddressingModel(w).Verify()
			},
		},
		{
			min: Dim1D,
			max: DimBuffer,
			err: ErrInvalidDimensionality,
			verify: func(w uint32) error {
				return Dimensionality(w).Verify()
			},
		},
		{
			min: ExecutionModeInvocations,
			max: ExecutionModeContractionOff,
			err: ErrInvalidExecutionMode,
			verify: func(w uint32) error {
				return ExecutionMode(w).Verify()
			},
		},
		{
			min: ExecutionModelVertex,
			max: ExecutionModelKernel,
			err: ErrInvalidExecutionModel,
			verify: func(w uint32) error {
				return ExecutionModel(w).Verify()
			},
		},
		{
			min: FPFastMathModeNotNaN,
			max: FPFastMathModeNotNaN | FPFastMathModeNotInf |
				FPFastMathModeNSZ | FPFastMathModeAllowRecip |
				FPFastMathModeFast,
			err: ErrInvalidFPFastMathMode,
			verify: func(w uint32) error {
				return FPFastMathMode(w).Verify()
			},
		},
		{
			min: FPRoundingModeRTE,
			max: FPRoundingModeRTN,
			err: ErrInvalidFPRoundingMode,
			verify: func(w uint32) error {
				return FPRoundingMode(w).Verify()
			},
		},
		{
			min: LinkageTypeExport,
			max: LinkageTypeImport,
			err: ErrInvalidLinkageType,
			verify: func(w uint32) error {
				return LinkageType(w).Verify()
			},
		},
		{
			min: MemoryModelSimple,
			max: MemoryModelOpenCL21,
			err: ErrInvalidMemoryModel,
			verify: func(w uint32) error {
				return MemoryModel(w).Verify()
			},
		},
		{
			min:  SamplerAddressingModeNone,
			max:  SamplerAddressingModeRepeatMirrored,
			err:  ErrInvalidSamplerAddressingMode,
			even: true,
			verify: func(w uint32) error {
				return SamplerAddressingMode(w).Verify()
			},
		},
		{
			min:  SamplerFilterModeNearest,
			max:  SamplerFilterModeLinear,
			err:  ErrInvalidSamplerFilterMode,
			even: true,
			verify: func(w uint32) error {
				return SamplerFilterMode(w).Verify()
			},
		},
		{
			min: SourceLanguageUnknown,
			max: SourceLanguageOpenCL,
			err: ErrInvalidSourceLanguage,
			verify: func(w uint32) error {
				return SourceLanguage(w).Verify()
			},
		},
		{
			min: StorageClassUniformConstant,
			max: StorageClassAtomicCounter,
			err: ErrInvalidStorageClass,
			verify: func(w uint32) error {
				return StorageClass(w).Verify()
			},
		},
		{
			min: FunctionParamAttrZext,
			max: FunctionParamAttrNoReadWrite,
			err: ErrInvalidFunctionParameter,
			verify: func(w uint32) error {
				return FunctionParameter(w).Verify()
			},
		},
		{
			min: DecorationPrecisionLow,
			max: DecorationSpecId,
			err: ErrInvalidDecoration,
			verify: func(w uint32) error {
				return Decoration(w).Verify()
			},
		},
		{
			min: BuiltinPosition,
			max: BuiltinSubgroupLocalInvocationId,
			err: ErrInvalidBuiltin,
			verify: func(w uint32) error {
				return Builtin(w).Verify()
			},
		},
		{
			min: SelectionControlNoControl,
			max: SelectionControlDontFlatten,
			err: ErrInvalidSelectionControl,
			verify: func(w uint32) error {
				return SelectionControl(w).Verify()
			},
		},
		{
			min: LoopControlNoControl,
			max: LoopControlDontUnroll,
			err: ErrInvalidLoopControl,
			verify: func(w uint32) error {
				return LoopControl(w).Verify()
			},
		},
		{
			min: FunctionControlMaskInLine,
			max: FunctionControlMaskInLine | FunctionControlMaskDontInline |
				FunctionControlMaskPure | FunctionControlMaskConst,
			err: ErrInvalidFunctionControlMask,
			verify: func(w uint32) error {
				return FunctionControlMask(w).Verify()
			},
		},
		{
			min: MemorySemanticRelaxed,
			max: MemorySemanticRelaxed |
				MemorySemanticSequentiallyConsistent |
				MemorySemanticAcquire |
				MemorySemanticRelease |
				MemorySemanticUniformMemory |
				MemorySemanticSubgroupMemory |
				MemorySemanticWorkgroupLocalMemory |
				MemorySemanticWorkgroupGlobalMemory |
				MemorySemanticAtomicCounterMemory |
				MemorySemanticImageMemory,
			err: ErrInvalidMemorySemantic,
			verify: func(w uint32) error {
				return MemorySemantic(w).Verify()
			},
		},
		{
			min: MemoryAccessVolatile,
			max: MemoryAccessAligned,
			err: ErrInvalidMemoryAccess,
			verify: func(w uint32) error {
				return MemoryAccess(w).Verify()
			},
		},
		{
			min: ExecutionScopeCrossDevice,
			max: ExecutionScopeSubgroup,
			err: ErrInvalidExecutionScope,
			verify: func(w uint32) error {
				return ExecutionScope(w).Verify()
			},
		},
		{
			min: GroupOperationReduce,
			max: GroupOperationExclusiveScan,
			err: ErrInvalidGroupOperation,
			verify: func(w uint32) error {
				return GroupOperation(w).Verify()
			},
		},
		{
			min: KernelEnqueueFlagNoWait,
			max: KernelEnqueueFlagWaitWorkGroup,
			err: ErrInvalidKernelEnqueueFlag,
			verify: func(w uint32) error {
				return KernelEnqueueFlag(w).Verify()
			},
		},
		{
			min: KernelProfilingInfoCmdExecTime,
			max: KernelProfilingInfoCmdExecTime,
			err: ErrInvalidKernelProfilingInfo,
			verify: func(w uint32) error {
				return KernelProfilingInfo(w).Verify()
			},
		},
	} {
		have := ct.verify(ct.min)
		if have != nil {
			t.Fatalf("Min value mismatch: %d\nHave: %v\nWant: success", ct.min, have)
		}

		have = ct.verify(ct.max)
		if have != nil {
			t.Fatalf("Max value mismatch: %d\nHave: %v\nWant: success", ct.max, have)
		}

		have = ct.verify(ct.max + 1)
		if have != ct.err {
			t.Fatalf("Invalid value mismatch: %d\nHave: %v\nWant: %v",
				ct.max, have, ct.err)
		}

		if !ct.even {
			continue
		}

		// SamplerAddressingMode and SamplerFilterMode are not exactly bit flags,
		// but they are numbered like one; sort of.
		//
		// Giving it odd numbers inside the valid range should be an error.
		have = ct.verify(ct.min + 1)
		if have != ct.err {
			t.Fatalf("Odd value mismatch: %d\nHave: %v\nWant: %v",
				ct.max, have, ct.err)
		}
	}
}

type bitTest struct {
	in   uint32
	none bool
	mask uint32
	want bool
}

func TestBitflag(t *testing.T) {
	for _, bt := range []bitTest{
		{0, true, 0, true},
		{0, false, 0, false},
		{0, true, 7, true},
		{2, true, 7, true},
		{5, true, 7, true},
	} {
		testBitflag(t, bt.in, bt.none, bt.mask, bt.want)
	}
}

func testBitflag(t *testing.T, in uint32, none bool, mask uint32, want bool) {
	have := verifyBitFlag(in, none, mask)
	if want != have {
		t.Fatalf("mismatch: in: %x, none: %v, mask: %x, want: %v, have: %v",
			in, none, mask, want, have)
	}
}
