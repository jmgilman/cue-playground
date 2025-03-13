package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	ctx := cuecontext.New()
	insts := load.Instances([]string{"."}, &load.Config{
		Package: "a",
	})

	v := ctx.BuildInstance(insts[0])
	if v.Err() != nil {
		fmt.Println(v.Err())
	}
}
