// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

import "sync"

// Keyboard is the interface all backends qualify as.
type Keyboard interface {
	// Bind binds the given key or chain to the specified handler.
	Bind(handler func(), keys ...string)

	// Unbind removes the binding for the given key or chain.
	Unbind(keys string)

	// Clear removes all bindings.
	Clear()

	// Call invokes the handler for the given key or sequence.
	Call(key string)
}

// Base is the base type for all backend implementations.
// It takes care of some common house keeping and ensures they all
// qualify as a Keyboard interface.
type Base struct {
	sync.Mutex
	Bindings map[string]func() // Set of registered bindings.
}

// NewBase creates a new Base instance.
func NewBase() *Base {
	b := new(Base)
	b.Clear()
	return b
}

// Bind binds the given keys or sequences to the specified handler.
func (b *Base) Bind(handler func(), keys ...string) {
	if len(keys) == 0 {
		return
	}

	b.Lock()

	for i := range keys {
		b.Bindings[keys[i]] = handler
	}

	b.Unlock()
}

// Unbind removes the binding for the given key or sequence.
func (b *Base) Unbind(key string) {
	if _, ok := b.Bindings[key]; !ok {
		return
	}

	b.Lock()
	delete(b.Bindings, key)
	b.Unlock()
}

// Clear removes all bindings.
func (b *Base) Clear() {
	b.Lock()
	b.Bindings = make(map[string]func())
	b.Unlock()
}

// Call invokes the handler for the given key or sequence.
func (b *Base) Call(key string) {
	if f, ok := b.Bindings[key]; ok {
		f()
	}
}
