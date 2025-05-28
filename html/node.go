package html

type Node interface {
	String() string
	PrettyString(indent string) string
}
