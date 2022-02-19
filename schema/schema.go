package schema

type Type uint

const (
	Int Type = iota
	Int32
	Int64

	Float
	Float32
	Float64

	String

	Bool
)

type Schema struct {
	Name        string
	Type        Type
	Description string
	Required    bool
}
