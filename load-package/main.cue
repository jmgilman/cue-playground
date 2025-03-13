package main

import (
	"github.com/jmgilman/cue-playground/load-package/a"
	"github.com/jmgilman/cue-playground/load-package/b"
)

#Main: {
	av:  a.#A
	bv:  b.#B
	foo: string | *"bar"
}
