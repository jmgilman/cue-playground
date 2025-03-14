package main

import (
	"io/fs"
	"log"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/jmgilman/cue-playground/load-package/module"
)

func main() {
	ctx := cuecontext.New()

	// Load all source files from the embedded module
	src, err := loadSrcFiles()
	if err != nil {
		log.Fatalf("failed to load source files: %v", err)
	}

	// Load the CUE module from the embedded files using an overlay
	insts := load.Instances([]string{"./a"}, &load.Config{
		Dir:        "/",
		ModuleRoot: "/",
		Overlay:    src,
	})

	// Build the instance
	v := ctx.BuildInstance(insts[0])
	if v.Err() != nil {
		log.Fatalf("failed to build instance: %v", v.Err())
	}

	// Lookup the value at path #A
	a := v.LookupPath(cue.ParsePath("#A"))
	if a.Err() != nil {
		log.Fatalf("failed to lookup path: %v", a.Err())
	}
}

// loadSrcFiles loads all source files from the embdded CUE module.
func loadSrcFiles() (map[string]load.Source, error) {
	files := map[string]load.Source{}

	// Walk the embedded module and load all files
	err := fs.WalkDir(module.Module, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Read the file from the embedded module
		f, err := module.Module.ReadFile(path)
		if err != nil {
			return err
		}

		// The loader requires all overlay paths to be absolute
		files["/"+path] = load.FromBytes(f)

		return nil
	})

	return files, err
}
