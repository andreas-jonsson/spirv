// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"fmt"
	"reflect"
)

// Instruction defines a generic instruction.
type Instruction interface {
	Verifiable

	// Opcode returns the opcode for this instruction.
	// It is used by the encoder to find the correct codec in the
	// instruction set library.
	Opcode() uint32

	// Optional returns true if the instruction has no semantic meaning.
	// Its presence is mostly for debugging purposes.
	Optional() bool
}

// instructionResultId returns the value of the instruction's result Id,
// provided it defines one.
func instructionResultId(i Instruction) (Id, bool) {
	rv := reflect.ValueOf(i)
	rv = reflect.Indirect(rv)

	field := rv.FieldByName("ResultId")
	if field.Kind() == reflect.Invalid {
		return 0, false
	}

	id := field.Interface().(Id)
	return id, true
}

// instructionName returns the name for the given instruction.
// This is the type name, minus some package cruft.
func instructionName(i Instruction) string {
	name := fmt.Sprintf("%T", i)
	return name[len("*spirv."):]
}

// List of known opcodes.
const (
	opcodeNop                                     = 0
	opcodeSource                                  = 1
	opcodeSourceExtension                         = 2
	opcodeExtension                               = 3
	opcodeExtInstImport                           = 4
	opcodeMemoryModel                             = 5
	opcodeEntryPoint                              = 6
	opcodeExecutionMode                           = 7
	opcodeTypeVoid                                = 8
	opcodeTypeBool                                = 9
	opcodeTypeInt                                 = 10
	opcodeTypeFloat                               = 11
	opcodeTypeVector                              = 12
	opcodeTypeMatrix                              = 13
	opcodeTypeSampler                             = 14
	opcodeTypeFilter                              = 15
	opcodeTypeArray                               = 16
	opcodeTypeRuntimeArray                        = 17
	opcodeTypeStruct                              = 18
	opcodeTypeOpaque                              = 19
	opcodeTypePointer                             = 20
	opcodeTypeFunction                            = 21
	opcodeTypeEvent                               = 22
	opcodeTypeDeviceEvent                         = 23
	opcodeTypeReserveId                           = 24
	opcodeTypeQueue                               = 25
	opcodeTypePipe                                = 26
	opcodeConstantTrue                            = 27
	opcodeConstantFalse                           = 28
	opcodeConstant                                = 29
	opcodeConstantComposite                       = 30
	opcodeConstantSampler                         = 31
	opcodeConstantNullPointer                     = 32
	opcodeConstantNullObject                      = 33
	opcodeSpecConstantTrue                        = 34
	opcodeSpecConstantFalse                       = 35
	opcodeSpecConstant                            = 36
	opcodeSpecConstantComposite                   = 37
	opcodeVariable                                = 38
	opcodeVariableArray                           = 39
	opcodeFunction                                = 40
	opcodeFunctionParameter                       = 41
	opcodeFunctionEnd                             = 42
	opcodeFunctionCall                            = 43
	opcodeExtInst                                 = 44
	opcodeUndef                                   = 45
	opcodeLoad                                    = 46
	opcodeStore                                   = 47
	opcodePhi                                     = 48
	opcodeDecorationGroup                         = 49
	opcodeDecorate                                = 50
	opcodeMemberDecorate                          = 51
	opcodeGroupDecorate                           = 52
	opcodeGroupMemberDecorate                     = 53
	opcodeName                                    = 54
	opcodeMemberName                              = 55
	opcodeString                                  = 56
	opcodeLine                                    = 57
	opcodeVectorExtractDynamic                    = 58
	opcodeVectorInsertDynamic                     = 59
	opcodeVectorShuffle                           = 60
	opcodeCompositeConstruct                      = 61
	opcodeCompositeExtract                        = 62
	opcodeCompositeInsert                         = 63
	opcodeCopyObject                              = 64
	opcodeCopyMemory                              = 65
	opcodeCopyMemorySized                         = 66
	opcodeSampler                                 = 67
	opcodeTextureSample                           = 68
	opcodeTextureSampleDref                       = 69
	opcodeTextureSampleLod                        = 70
	opcodeTextureSampleProj                       = 71
	opcodeTextureSampleGrad                       = 72
	opcodeTextureSampleOffset                     = 73
	opcodeTextureSampleProjLod                    = 74
	opcodeTextureSampleProjGrad                   = 75
	opcodeTextureSampleLodOffset                  = 76
	opcodeTextureSampleProjOffset                 = 77
	opcodeTextureSampleGradOffset                 = 78
	opcodeTextureSampleProjLodOffset              = 79
	opcodeTextureSampleProjGradOffset             = 80
	opcodeTextureFetchTexel                       = 81
	opcodeTextureFetchTexelOffset                 = 82
	opcodeTextureFetchSample                      = 83
	opcodeTextureFetchBuffer                      = 84
	opcodeTextureGather                           = 85
	opcodeTextureGatherOffset                     = 86
	opcodeTextureGatherOffsets                    = 87
	opcodeTextureQuerySizeLod                     = 88
	opcodeTextureQuerySize                        = 89
	opcodeTextureQueryLod                         = 90
	opcodeTextureQueryLevels                      = 91
	opcodeTextureQuerySamples                     = 92
	opcodeAccessChain                             = 93
	opcodeInboundsAccessChain                     = 94
	opcodeSNegate                                 = 95
	opcodeFNegate                                 = 96
	opcodeNot                                     = 97
	opcodeAny                                     = 98
	opcodeAll                                     = 99
	opcodeConvertFToU                             = 100
	opcodeConvertFToS                             = 101
	opcodeConvertSToF                             = 102
	opcodeConvertUToF                             = 103
	opcodeUConvert                                = 104
	opcodeSConvert                                = 105
	opcodeFConvert                                = 106
	opcodeConvertPtrToU                           = 107
	opcodeConvertUToPtr                           = 108
	opcodePtrCastToGeneric                        = 109
	opcodeGenericCastToPtr                        = 110
	opcodeBitcast                                 = 111
	opcodeTranspose                               = 112
	opcodeIsNan                                   = 113
	opcodeIsInf                                   = 114
	opcodeIsFinite                                = 115
	opcodeIsNormal                                = 116
	opcodeSignBitSet                              = 117
	opcodeLessOrGreater                           = 118
	opcodeOrdered                                 = 119
	opcodeUnordered                               = 120
	opcodeArraylength                             = 121
	opcodeIAdd                                    = 122
	opcodeFAdd                                    = 123
	opcodeISub                                    = 124
	opcodeFSub                                    = 125
	opcodeIMul                                    = 126
	opcodeFMul                                    = 127
	opcodeUDiv                                    = 128
	opcodeSDiv                                    = 129
	opcodeFDiv                                    = 130
	opcodeUMod                                    = 131
	opcodeSRem                                    = 132
	opcodeSMod                                    = 133
	opcodeFRem                                    = 134
	opcodeFMod                                    = 135
	opcodeVectorTimesScalar                       = 136
	opcodeMatrixTimesScalar                       = 137
	opcodeVectorTimesMatrix                       = 138
	opcodeMatrixTimesVector                       = 139
	opcodeMatrixTimesMatrix                       = 140
	opcodeOuterProduct                            = 141
	opcodeDot                                     = 142
	opcodeShiftRightLogical                       = 143
	opcodeShiftRightArithmetic                    = 144
	opcodeShiftLeftLogical                        = 145
	opcodeLogicalOr                               = 146
	opcodeLogicalXor                              = 147
	opcodeLogicalAnd                              = 148
	opcodeBitwiseOr                               = 149
	opcodeBitwiseXor                              = 150
	opcodeBitwiseAnd                              = 151
	opcodeSelect                                  = 152
	opcodeIEqual                                  = 153
	opcodeFOrdEqual                               = 154
	opcodeFUnordEqual                             = 155
	opcodeINotEqual                               = 156
	opcodeFOrdNotEqual                            = 157
	opcodeFUnordNotEqual                          = 158
	opcodeULessThan                               = 159
	opcodeSLessThan                               = 160
	opcodeFOrdLessThan                            = 161
	opcodeFUnordLessThan                          = 162
	opcodeUGreaterThan                            = 163
	opcodeSGreaterThan                            = 164
	opcodeFOrdGreaterThan                         = 165
	opcodeFUnordGreaterThan                       = 166
	opcodeULessThanEqual                          = 167
	opcodeSLessThanEqual                          = 168
	opcodeFOrdLessThanEqual                       = 169
	opcodeFUnordLessThanEqual                     = 170
	opcodeUGreaterThanEqual                       = 171
	opcodeSGreaterThanEqual                       = 172
	opcodeFOrdGreaterThanEqual                    = 173
	opcodeFUnordGreaterThanEqual                  = 174
	opcodeDPdx                                    = 175
	opcodeDPdy                                    = 176
	opcodeFwidth                                  = 177
	opcodeDPdxFine                                = 178
	opcodeDPdyFine                                = 179
	opcodeFwidthFine                              = 180
	opcodeDPdxCoarse                              = 181
	opcodeDPdyCoarse                              = 182
	opcodeFwidthCoarse                            = 183
	opcodeEmitVertex                              = 184
	opcodeEndPrimitive                            = 185
	opcodeEmitStreamVertex                        = 186
	opcodeEndStreamPrimitive                      = 187
	opcodeControlBarrier                          = 188
	opcodeMemoryBarrier                           = 189
	opcodeImagePointer                            = 190
	opcodeAtomicInit                              = 191
	opcodeAtomicLoad                              = 192
	opcodeAtomicStore                             = 193
	opcodeAtomicExchange                          = 194
	opcodeAtomicCompareExchange                   = 195
	opcodeAtomicCompareExchangeWeak               = 196
	opcodeAtomicIIncrement                        = 197
	opcodeAtomicIDecrement                        = 198
	opcodeAtomicIAdd                              = 199
	opcodeAtomicISub                              = 200
	opcodeAtomicUMin                              = 201
	opcodeAtomicUMax                              = 202
	opcodeAtomicAnd                               = 203
	opcodeAtomicOr                                = 204
	opcodeAtomicXor                               = 205
	opcodeLoopMerge                               = 206
	opcodeSelectionMerge                          = 207
	opcodeLabel                                   = 208
	opcodeBranch                                  = 209
	opcodeBranchConditional                       = 210
	opcodeSwitch                                  = 211
	opcodeKill                                    = 212
	opcodeReturn                                  = 213
	opcodeReturnValue                             = 214
	opcodeUnreachable                             = 215
	opcodeLifetimeStart                           = 216
	opcodeLifetimeStop                            = 217
	opcodeCompileFlag                             = 218
	opcodeAsyncGroupCopy                          = 219
	opcodeWaitGroupEvents                         = 220
	opcodeGroupAll                                = 221
	opcodeGroupAny                                = 222
	opcodeGroupBroadcast                          = 223
	opcodeGroupIAdd                               = 224
	opcodeGroupFAdd                               = 225
	opcodeGroupFMin                               = 226
	opcodeGroupUMin                               = 227
	opcodeGroupSMin                               = 228
	opcodeGroupFMax                               = 229
	opcodeGroupUMax                               = 230
	opcodeGroupSMax                               = 231
	opcodeGenericCastToPtrExplicit                = 232
	opcodeGenericPtrMemSemantics                  = 233
	opcodeReadPipe                                = 234
	opcodeWritePipe                               = 235
	opcodeReservedReadPipe                        = 236
	opcodeReservedWritePipe                       = 237
	opcodeReserveReadPipePackets                  = 238
	opcodeReserveWritePipePackets                 = 239
	opcodeCommitReadPipe                          = 240
	opcodeCommitWritePipe                         = 241
	opcodeIsValidReserveId                        = 242
	opcodeGetNumPipePackets                       = 243
	opcodeGetMaxPipePackets                       = 244
	opcodeGroupReserveReadPipePackets             = 245
	opcodeGroupReserveWritePipePackets            = 246
	opcodeGroupCommitReadPipe                     = 247
	opcodeGroupCommitWritePipe                    = 248
	opcodeEnqueueMarker                           = 249
	opcodeEnqueueKernel                           = 250
	opcodeGetKernelNDrangeSubGroupCount           = 251
	opcodeGetKernelNDrangeMaxSubGroupSize         = 252
	opcodeGetKernelWorkGroupSize                  = 253
	opcodeGetKernelPreferredWorkGroupSizeMultiple = 254
	opcodeRetainEvent                             = 255
	opcodeReleaseEvent                            = 256
	opcodeCreateUserEvent                         = 257
	opcodeIsValidEvent                            = 258
	opcodeSetUserEventStatus                      = 259
	opcodeCaptureEventProfilingInfo               = 260
	opcodeGetDefaultQueue                         = 261
	opcodeBuildNDRange                            = 262
)
