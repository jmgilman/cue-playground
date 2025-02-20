package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/load"
)

func main() {
	// Load the module
	ctx := cuecontext.New()
	insts := load.Instances(nil, &load.Config{
		Package: "main",
	})
	v := ctx.BuildInstance(insts[0])

	// Dump the envs as CUE source
	envs := v.LookupPath(cue.ParsePath("test.foo"))
	src, err := format.Node(envs.Syntax())
	if err != nil {
		fmt.Printf("failed to format envs: %v\n", err)
	}

	fmt.Printf("%s\n", src)
}
