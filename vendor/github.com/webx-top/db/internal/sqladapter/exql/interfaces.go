package exql

import (
	"github.com/webx-top/db/internal/cache"
)

// Fragment is any interface that can be both cached and compiled.
type Fragment interface {
	cache.Hashable

	compilable
}

type compilable interface {
	Compile(*Template) (string, error)
}

type hasIsEmpty interface {
	IsEmpty() bool
}
