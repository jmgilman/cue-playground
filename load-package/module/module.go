package module

import "embed"

//go:embed cue.mod/module.cue
//go:embed a/*.cue
//go:embed b/*.cue
//go:embed main.cue
var Module embed.FS
