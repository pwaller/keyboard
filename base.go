// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

import (
	"fmt"
	"time"
)

// Binding represents a single key/sequence binding.
type Binding struct {
	Keys    []Key  // Keys involved in the binding. This includes modifiers.
	Handler func() // Handler function to be called when binding is triggered.
}

// Base is the base type for all backend implementations.
// It takes care of some common house keeping and ensures they all
// qualify as a Keyboard interface.
type Base struct {
	Bindings map[string]*Binding // Set of registered bindings.
	record   []Key               // List of recorded keypresses.
	stamp    time.Time           // Last key press.
	timeout  time.Duration       // Timeout for sequence resets.
}

// NewBase creates a new Base instance.
func NewBase() *Base {
	b := new(Base)
	b.timeout = time.Second >> 1
	b.Clear()
	return b
}

// SetTimeout sets the timeout for sequence resets in nanoseconds
func (b *Base) SetTimeout(d int64) { b.timeout = time.Duration(d) }

// RecordKeyDown records the given key press.
func (b *Base) RecordKeyDown(key Key, mods Modifier) {
	if time.Since(b.stamp) > b.timeout {
		b.record = b.record[:0]
	}

	b.stamp = time.Now()
	b.record = append(b.record, Key(mods)<<8|key)

	fmt.Printf("%04x\n", b.record)
}

// RecordKeyUp records the given key release.
func (b *Base) RecordKeyUp(key Key, mods Modifier) {

}

// Bind binds the given keys or sequences to the specified handler.
func (b *Base) Bind(handler func(), keys ...string) {
	if len(keys) == 0 {
		return
	}

	for i := range keys {
		b.Bindings[keys[i]] = &Binding{
			Keys:    parseKeys(keys[i]),
			Handler: handler,
		}
	}
}

// Unbind removes the binding for the given key or sequence.
func (b *Base) Unbind(key string) {
	if _, ok := b.Bindings[key]; ok {
		delete(b.Bindings, key)
	}
}

// Clear removes all bindings.
func (b *Base) Clear() {
	b.Bindings = make(map[string]*Binding)
}

// Call invokes the handler for the given key or sequence.
func (b *Base) Call(key string) {
	if b, ok := b.Bindings[key]; ok {
		b.Handler()
	}
}
