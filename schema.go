package libpx

type SchemaType uint

const (
	SchemaTypeInt SchemaType = iota
	SchameTypeInt32
	SchemaTypeInt64
)

type Schema struct {
	Type        SchemaType
	Description string
	Required    bool
}
