package core

// Linear state that can be saved/reverted
type Stateful interface {
	Save()
	Pop()
	Revert()
	RevertAndPop()
}
