// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

// OpSampler creates a sampler containing both a filter and texture.
type OpSampler struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Filter     Id
}

func (c *OpSampler) Opcode() uint32 { return opcodeSampler }
func (c *OpSampler) Optional() bool { return false }
func (c *OpSampler) Verify() error  { return nil }

// OpTextureSample samples a texture with an implicit level of detail.
type OpTextureSample struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Bias       Id `spirv:"optional"`
}

func (c *OpTextureSample) Opcode() uint32 { return opcodeTextureSample }
func (c *OpTextureSample) Optional() bool { return false }
func (c *OpTextureSample) Verify() error  { return nil }

// OpTextureSampleDref samples a cube-map-array texture with depth
// comparison using an implicit level of detail.
type OpTextureSampleDref struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Dref       Id
}

func (c *OpTextureSampleDref) Opcode() uint32 { return opcodeTextureSampleDref }
func (c *OpTextureSampleDref) Optional() bool { return false }
func (c *OpTextureSampleDref) Verify() error  { return nil }

// OpTextureSampleLod samples a texture using an explicit level of detail.
type OpTextureSampleLod struct {
	ResultType    Id
	ResultId      Id
	Sampler       Id
	Coordinate    Id
	LevelofDetail Id
}

func (c *OpTextureSampleLod) Opcode() uint32 { return opcodeTextureSampleLod }
func (c *OpTextureSampleLod) Optional() bool { return false }
func (c *OpTextureSampleLod) Verify() error  { return nil }

// OpTextureSampleProj sample a texture with a projective coordinate
// using an implicit level of detail.
type OpTextureSampleProj struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Bias       Id `spirv:"optional"`
}

func (c *OpTextureSampleProj) Opcode() uint32 { return opcodeTextureSampleProj }
func (c *OpTextureSampleProj) Optional() bool { return false }
func (c *OpTextureSampleProj) Verify() error  { return nil }

// OpTextureSampleGrad samples a texture with an explicit gradient.
type OpTextureSampleGrad struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Dx         Id
	Dy         Id
}

func (c *OpTextureSampleGrad) Opcode() uint32 { return opcodeTextureSampleGrad }
func (c *OpTextureSampleGrad) Optional() bool { return false }
func (c *OpTextureSampleGrad) Verify() error  { return nil }

// OpTextureSampleOffset samples a texture with an offset from a
// coordinate using an implicit level of detail.
type OpTextureSampleOffset struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Offset     Id
	Bias       Id `spirv:"optional"`
}

func (c *OpTextureSampleOffset) Opcode() uint32 { return opcodeTextureSampleOffset }
func (c *OpTextureSampleOffset) Optional() bool { return false }
func (c *OpTextureSampleOffset) Verify() error  { return nil }

// OpTextureSampleProjLod samples a texture with a projective
// coordinate using an explicit level of detail.
type OpTextureSampleProjLod struct {
	ResultType    Id
	ResultId      Id
	Sampler       Id
	Coordinate    Id
	LevelofDetail Id
}

func (c *OpTextureSampleProjLod) Opcode() uint32 { return opcodeTextureSampleProjLod }
func (c *OpTextureSampleProjLod) Optional() bool { return false }
func (c *OpTextureSampleProjLod) Verify() error  { return nil }

// OpTextureSampleProjGrad sample a texture with a projective
// coordinate using an explicit gradient.
type OpTextureSampleProjGrad struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Dx         Id
	Dy         Id
}

func (c *OpTextureSampleProjGrad) Opcode() uint32 { return opcodeTextureSampleProjGrad }
func (c *OpTextureSampleProjGrad) Optional() bool { return false }
func (c *OpTextureSampleProjGrad) Verify() error  { return nil }

// OpTextureSampleLodOffset samples a texture with explicit level of
// detail using an offset from a coordinate.
type OpTextureSampleLodOffset struct {
	ResultType    Id
	ResultId      Id
	Sampler       Id
	Coordinate    Id
	LevelofDetail Id
	Offset        Id
}

func (c *OpTextureSampleLodOffset) Opcode() uint32 { return opcodeTextureSampleLodOffset }
func (c *OpTextureSampleLodOffset) Optional() bool { return false }
func (c *OpTextureSampleLodOffset) Verify() error  { return nil }

// OpTextureSampleProjOffset samples a texture with an offset from a
// projective coordinate using an implicit level of detail.
type OpTextureSampleProjOffset struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Offset     Id
	Bias       Id `spirv:"optional"`
}

func (c *OpTextureSampleProjOffset) Opcode() uint32 { return opcodeTextureSampleProjOffset }
func (c *OpTextureSampleProjOffset) Optional() bool { return false }
func (c *OpTextureSampleProjOffset) Verify() error  { return nil }

// OpTextureSampleGradOffset samples a texture with an offset
// coordinate and an explicit gradient.
type OpTextureSampleGradOffset struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Dx         Id
	Dy         Id
	Offset     Id
}

func (c *OpTextureSampleGradOffset) Opcode() uint32 { return opcodeTextureSampleGradOffset }
func (c *OpTextureSampleGradOffset) Optional() bool { return false }
func (c *OpTextureSampleGradOffset) Verify() error  { return nil }

// OpTextureSampleProjLodOffset samples a texture with an offset from
// a projective coordinate and an explicit level of detail.
type OpTextureSampleProjLodOffset struct {
	ResultType    Id
	ResultId      Id
	Sampler       Id
	Coordinate    Id
	LevelofDetail Id
	Offset        Id
}

func (c *OpTextureSampleProjLodOffset) Opcode() uint32 { return opcodeTextureSampleProjLodOffset }
func (c *OpTextureSampleProjLodOffset) Optional() bool { return false }
func (c *OpTextureSampleProjLodOffset) Verify() error  { return nil }

// OpTextureSampleProjGradOffset samples a texture with an offset from
// a projective coordinate and an explicit gradient.
type OpTextureSampleProjGradOffset struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Dx         Id
	Dy         Id
	Offset     Id
}

func (c *OpTextureSampleProjGradOffset) Opcode() uint32 { return opcodeTextureSampleProjGradOffset }
func (c *OpTextureSampleProjGradOffset) Optional() bool { return false }
func (c *OpTextureSampleProjGradOffset) Verify() error  { return nil }

// OpTextureFetchTexel fetches a single texel from a texture.
type OpTextureFetchTexel struct {
	ResultType    Id
	ResultId      Id
	Sampler       Id
	Coordinate    Id
	LevelofDetail Id
}

func (c *OpTextureFetchTexel) Opcode() uint32 { return opcodeTextureFetchTexel }
func (c *OpTextureFetchTexel) Optional() bool { return false }
func (c *OpTextureFetchTexel) Verify() error  { return nil }

// OpTextureFetchTexelOffset fetches a single offset texel from a texture.
type OpTextureFetchTexelOffset struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Offset     Id
}

func (c *OpTextureFetchTexelOffset) Opcode() uint32 { return opcodeTextureFetchTexelOffset }
func (c *OpTextureFetchTexelOffset) Optional() bool { return false }
func (c *OpTextureFetchTexelOffset) Verify() error  { return nil }

// OpTextureFetchSample fetches a single sample from a multi-sample texture.
type OpTextureFetchSample struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Sample     Id
}

func (c *OpTextureFetchSample) Opcode() uint32 { return opcodeTextureFetchSample }
func (c *OpTextureFetchSample) Optional() bool { return false }
func (c *OpTextureFetchSample) Verify() error  { return nil }

// OpTextureFetchBuffer fetches an element out of a buffer texture.
type OpTextureFetchBuffer struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Element    Id
}

func (c *OpTextureFetchBuffer) Opcode() uint32 { return opcodeTextureFetchBuffer }
func (c *OpTextureFetchBuffer) Optional() bool { return false }
func (c *OpTextureFetchBuffer) Verify() error  { return nil }

// OpTextureGather gathers the requested component from four sampled texels.
type OpTextureGather struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Component  Id
}

func (c *OpTextureGather) Opcode() uint32 { return opcodeTextureGather }
func (c *OpTextureGather) Optional() bool { return false }
func (c *OpTextureGather) Verify() error  { return nil }

// OpTextureGatherOffset gathers the requested component from four
// offset sampled texels.
type OpTextureGatherOffset struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Component  Id
	Offset     Id
}

func (c *OpTextureGatherOffset) Opcode() uint32 { return opcodeTextureGatherOffset }
func (c *OpTextureGatherOffset) Optional() bool { return false }
func (c *OpTextureGatherOffset) Verify() error  { return nil }

// OpTextureGatherOffsets gathers the requested component from four
// offset sampled texels.
type OpTextureGatherOffsets struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
	Component  Id
	Offsets    Id
}

func (c *OpTextureGatherOffsets) Opcode() uint32 { return opcodeTextureGatherOffsets }
func (c *OpTextureGatherOffsets) Optional() bool { return false }
func (c *OpTextureGatherOffsets) Verify() error  { return nil }

// OpTextureQuerySizeLod queries the dimensions of the texture for
// Sampler for mipmap level for Level of Detail.
type OpTextureQuerySizeLod struct {
	ResultType    Id
	ResultId      Id
	Sampler       Id
	LevelofDetail Id
}

func (c *OpTextureQuerySizeLod) Opcode() uint32 { return opcodeTextureQuerySizeLod }
func (c *OpTextureQuerySizeLod) Optional() bool { return false }
func (c *OpTextureQuerySizeLod) Verify() error  { return nil }

// OpTextureQuerySize queries the dimensions of the texture for
// Sampler, with no level of detail.
type OpTextureQuerySize struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
}

func (c *OpTextureQuerySize) Opcode() uint32 { return opcodeTextureQuerySize }
func (c *OpTextureQuerySize) Optional() bool { return false }
func (c *OpTextureQuerySize) Verify() error  { return nil }

// OpTextureQueryLod queries the mipmap level and the level of detail
// for a hypothetical sampling of Sampler at Coordinate using an
// implicit level of detail.
type OpTextureQueryLod struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
	Coordinate Id
}

func (c *OpTextureQueryLod) Opcode() uint32 { return opcodeTextureQueryLod }
func (c *OpTextureQueryLod) Optional() bool { return false }
func (c *OpTextureQueryLod) Verify() error  { return nil }

// OpTextureQueryLevels queries the number of mipmap levels accessible
// through Sampler.
type OpTextureQueryLevels struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
}

func (c *OpTextureQueryLevels) Opcode() uint32 { return opcodeTextureQueryLevels }
func (c *OpTextureQueryLevels) Optional() bool { return false }
func (c *OpTextureQueryLevels) Verify() error  { return nil }

// OpTextureQuerySamples queries the number of samples available per
// texel fetch in a multisample texture.
type OpTextureQuerySamples struct {
	ResultType Id
	ResultId   Id
	Sampler    Id
}

func (c *OpTextureQuerySamples) Opcode() uint32 { return opcodeTextureQuerySamples }
func (c *OpTextureQuerySamples) Optional() bool { return false }
func (c *OpTextureQuerySamples) Verify() error  { return nil }

func init() {
	bind(func() Instruction { return &OpSampler{} })
	bind(func() Instruction { return &OpTextureSample{Bias: 0} })
	bind(func() Instruction { return &OpTextureSampleDref{} })
	bind(func() Instruction { return &OpTextureSampleLod{} })
	bind(func() Instruction { return &OpTextureSampleProj{Bias: 0} })
	bind(func() Instruction { return &OpTextureSampleGrad{} })
	bind(func() Instruction { return &OpTextureSampleOffset{Bias: 0} })
	bind(func() Instruction { return &OpTextureSampleProjLod{} })
	bind(func() Instruction { return &OpTextureSampleProjGrad{} })
	bind(func() Instruction { return &OpTextureSampleLodOffset{} })
	bind(func() Instruction { return &OpTextureSampleProjOffset{Bias: 0} })
	bind(func() Instruction { return &OpTextureSampleGradOffset{} })
	bind(func() Instruction { return &OpTextureSampleProjLodOffset{} })
	bind(func() Instruction { return &OpTextureSampleProjGradOffset{} })
	bind(func() Instruction { return &OpTextureFetchTexel{} })
	bind(func() Instruction { return &OpTextureFetchTexelOffset{} })
	bind(func() Instruction { return &OpTextureFetchSample{} })
	bind(func() Instruction { return &OpTextureFetchBuffer{} })
	bind(func() Instruction { return &OpTextureGather{} })
	bind(func() Instruction { return &OpTextureGatherOffset{} })
	bind(func() Instruction { return &OpTextureGatherOffsets{} })
	bind(func() Instruction { return &OpTextureQuerySizeLod{} })
	bind(func() Instruction { return &OpTextureQuerySize{} })
	bind(func() Instruction { return &OpTextureQueryLod{} })
	bind(func() Instruction { return &OpTextureQueryLevels{} })
	bind(func() Instruction { return &OpTextureQuerySamples{} })
}
