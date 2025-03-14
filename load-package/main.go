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

	src, err := loadSrcFiles()
	if err != nil {
		log.Fatalf("failed to load source files: %v", err)
	}

	insts := load.Instances([]string{"./a"}, &load.Config{
		Dir:        "/",
		ModuleRoot: "/",
		Overlay:    src,
	})

	v := ctx.BuildInstance(insts[0])
	if v.Err() != nil {
		log.Fatalf("failed to build instance: %v", v.Err())
	}

	a := v.LookupPath(cue.ParsePath("#A"))
	if a.Err() != nil {
		log.Fatalf("failed to lookup path: %v", a.Err())
	}
}

func loadSrcFiles() (map[string]load.Source, error) {
	files := map[string]load.Source{}
	err := fs.WalkDir(module.Module, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		f, err := module.Module.ReadFile(path)
		if err != nil {
			return err
		}

		files["/"+path] = load.FromBytes(f)

		return nil
	})

	return files, err
}
