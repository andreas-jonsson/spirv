// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andreas-jonsson/spirv"
)

func main() {
	file := parseArgs()

	fd, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer fd.Close()

	module, err := spirv.Load(fd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dump(module)
}

// parseArgs parses and validates command line arguments.
func parseArgs() string {
	flag.Usage = func() {
		fmt.Println("usage:", os.Args[0], "[options] <module file>")
		flag.PrintDefaults()
	}

	version := flag.Bool("version", false, "Display version information.")
	flag.Parse()

	if *version {
		fmt.Println(Version())
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	return flag.Arg(0)
}

func makeNew(file string) {
	fd, err := os.Create(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer fd.Close()

	mod := spirv.NewModule()
	mod.Code = []spirv.Instruction{
		&spirv.OpSource{spirv.SourceLanguageGLSL, 450},
		&spirv.OpExtInst{
			ResultType:  1,
			ResultId:    2,
			Set:         3,
			Instruction: 4,
			Operands:    []spirv.Id{5, 4, 5},
		},
		&spirv.OpFunction{
			ResultType:   0,
			ResultId:     1,
			ControlMask:  spirv.FunctionControlMaskInLine,
			FunctionType: 2,
		},
		&spirv.OpFunctionParameter{
			ResultType: 0,
			ResultId:   1,
		},
		&spirv.OpFunctionParameter{
			ResultType: 0,
			ResultId:   1,
		},
		&spirv.OpFunctionEnd{},
	}

	err = mod.Save(fd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
