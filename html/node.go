package html

type Node interface {
	String() string
	PrettyString(indent string) string

	Bytes() []byte
	PrettyBytes(indent string) []byte
}
