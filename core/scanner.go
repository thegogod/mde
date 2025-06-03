package core

// Scans tokens one-by-one
// from an array of raw bytes
type Scanner interface {
	Stateful

	Scan() (Token, error)
}
