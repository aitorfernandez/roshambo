package ctxkey

// Key holds a name value to associate with context values.
type Key interface {
	String() string
}

type key struct {
	name string
}

// New creates a unique context Key.
func New(n string) Key {
	return &key{n}
}

func (k key) String() string {
	return k.name
}
