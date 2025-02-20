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

	// Dump foo as CUE source
	foo := v.LookupPath(cue.ParsePath("test.foo"))
	src, err := format.Node(foo.Syntax())
	if err != nil {
		fmt.Printf("failed to format foo: %v\n", err)
	}

	fmt.Printf("%s\n", src)
}
