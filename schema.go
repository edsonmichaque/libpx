package libpx

type SchemaType uint

const (
	SchemaInt SchemaType = iota
	SchameInt32
	SchemaInt64
)

type Schema struct {
	Type        SchemaType
	Description string
	Required    bool
}
