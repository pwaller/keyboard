// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

// Key represents a single key with optional modifiers.
// The upper 8 bits hold the modifiers as bit flags.
// The lower 8 bits hold the key constant: KeyA, KeyB, KeyC, etc.
type Key uint16

// Value returns the key value: KeyA, KeyB, KeyC, etc.
func (k Key) Value() Key { return k & 0xff }

// Mods returns the modifiers for this key: ModAlt, ModShift, ModSuper, ModCtrl.
func (k Key) Mods() Modifier { return Modifier(k >> 8) }

// Keyboard is the interface all backends qualify as.
type Keyboard interface {
	// Bindings returns a list of current bindings.
	Bindings() []*Binding

	// Bind binds the given key or chain to the specified handler.
	Bind(handler func(), keys ...string)

	// Unbind removes the binding for the given key or chain.
	Unbind(keys string)

	// Clear removes all bindings.
	Clear()

	// Call invokes the handler for the given key or sequence.
	Call(key string)

	// RecordKeyDown records the given key press.
	// This expects a valid key constant as defined in this package.
	RecordKey(key Key, mods Modifier)

	// SetTimeout sets the timeout for sequence resets in nanoseconds.
	SetTimeout(int64)

	// Poll can be called with an backend-specific event object
	// in cases where this is necessary. The backend should process the
	// event accordingly. Not all backends will use this.
	Poll(interface{})
}
