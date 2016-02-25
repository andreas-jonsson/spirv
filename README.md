## SPIR-V

[![Build](https://drone.io/github.com/andreas-jonsson/spirv/status.png)](https://drone.io/github.com/andreas-jonsson/spirv/latest)
[![Coverage](https://coveralls.io/repos/github/andreas-jonsson/spirv/badge.svg?branch=master)](https://coveralls.io/github/andreas-jonsson/spirv?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/andreas-jonsson/spirv)
[![Bountysource](https://api.bountysource.com/badge/team?team_id=130721&amp;style=bounties_received)](https://www.bountysource.com/teams/go-spir-v/bounties)

Package SPIR-V is a Go encoder/decoder for the Vulkan SPIR-V format.
This is based on the [preliminary specificationn][1] (pdf) and is subject to
change as the specification matures.

Additional SPIR-V information can be found [here][2] (pdf) and [here][3].
A video lecture on Vulkan and SPIR can be seen [here][4].

[1]: https://www.khronos.org/registry/spir-v/specs/1.0/SPIRV.pdf
[2]: https://www.khronos.org/registry/spir-v/
[3]: https://www.khronos.org/spir/
[4]: https://www.youtube.com/watch?v=qKbtrVEhaw8


### Usage

By itself, this package is not very useful. All it does is decode SPIR-V
binary into sets of 32-bit words data structure and vice-versa. It is intended
as a tool to facilitate the creation of SPIR-V debugging tools, compilers,
and whatever else you may require.

At the highest level, one can operate on complete modules.
They can be loaded, saved and verified to be correct:

	module, err := spirv.Load(r)
	...

	err := module.Verify()
	...

	err := module.Save(w)
	...

The Encoder and Decoder can be used directly if you wish. They offer working
with data on a per-instruction basis and if you opt out of deserialization into
typed structures, you can examine them without any allocation overhead.


### About

SPIR-V is a binary intermediate language for representing graphical-shader
stages and compute kernels for multiple Khronos APIs. Each function in a SPIR-V
module contains a control-flow graph (CFG) of basic blocks, with additional
instructions and constraints to retain source-code structured flow control.

SPIR-V has the following goals:

* Provide a simple binary intermediate language for all functionality appearing in Khronos shaders/kernels.
* Have a short, transparent, self-contained specification (sections Specification and Binary Form).
* Map easily to other IRs, including LLVM IR.
* Be the form passed by an API into a driver to set shaders/kernels.
* Can be targeted by new front ends for novel high-level languages.
* Allow the first parts of the source compilation and reflection process to be done offline.
* Be low-level enough to require a reverse-engineering step to reconstruct source code.
* Improve portability by enabling shared tools to generate or operate on it.
* Allow separation of core specification from source language-specific sets of built-in functions.
* Reduce compile time during application run time. (Eliminating most of the compile time during application run time is not a goal of this IL. Target-specific register allocation and scheduling are expected to still take significant time.)
* Allow some optimizations to be done offline.


### Install

    go get github.com/andreas-jonsson/spirv
    
    
### Acknowledgement

This library was originally written by [Jim Teeuwen](https://github.com/jteeuwen).


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.
