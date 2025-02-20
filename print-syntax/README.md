# Print Syntax

An example showing the result of printing a partial CUE value.

Run with:

```
go run main.go
```

Produces:

```cue
{
        _#def
        _#def: {
                [string]: {
                        value: string
                }
        } & {
                test: {
                        value: "test"
                }
        }
}
```