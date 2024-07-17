package schema

//go:generate cue get go --local .
//go:generate mv schema_go_gen.cue ../../cue/schema/schema_go_gen.cue

type Schema struct {
	Field1 string    `json:"field1"`
	Field2 int       `json:"field2"`
	Field3 bool      `json:"field3"`
	Sub    SubSchema `json:"sub,omitempty"`
}

type SubSchema struct {
	Field1 string `json:"field1"`
}
