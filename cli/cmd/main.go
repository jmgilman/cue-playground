package main

import (
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/alecthomas/kong"
)

var cli struct {
	Validate validateCmd `cmd:"" help:"Validates a configuration file."`
}

type validateCmd struct {
	Config string `arg:"" help:"Path to the configuration file."`
}

func (c *validateCmd) Run() error {
	// TODO: Use the embedded schema to validate the configuration file.
	// files, err := cue.SchemaFiles.ReadDir("schema")

	ctx := cuecontext.New()
	instances := load.Instances([]string{
		"./cue/schema",
		c.Config,
	}, nil)

	schema := ctx.BuildInstance(instances[0]).LookupPath(cue.ParsePath("#Schema"))
	input := ctx.BuildInstance(instances[1])

	return input.Unify(schema).Validate()
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("cli"),
		kong.Description("A CLI for testing the CUE API"))

	err := ctx.Run()
	ctx.FatalIfErrorf(err)
	os.Exit(0)
}
