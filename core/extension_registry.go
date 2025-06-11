package core

import (
	"fmt"
	"slices"
)

type ExtensionRegistry []Extension

func (self ExtensionRegistry) Has(name string) bool {
	return slices.ContainsFunc(self, func(ext Extension) bool {
		return ext.Name() == name
	})
}

func (self ExtensionRegistry) Get(name string) Extension {
	i := slices.IndexFunc(self, func(ext Extension) bool {
		return ext.Name() == name
	})

	if i < 0 {
		return nil
	}

	return self[i]
}

func (self *ExtensionRegistry) Add(ext Extension) error {
	if self.Has(ext.Name()) {
		return fmt.Errorf("extension '%s' is already registered", ext.Name())
	}

	arr := *self
	arr = append(arr, ext)
	*self = arr
	return nil
}
