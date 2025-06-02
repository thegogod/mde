package core

// Scans tokens one-by-one
// from an array of raw bytes
type Scanner interface {
	Scan() (Token, error)
}
