// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package glfw

import (
	"github.com/jteeuwen/glfw"
	"github.com/jteeuwen/keyboard"
)

type Keyboard struct {
	*keyboard.Base
}

// New constructs a new keyboard instance.
func New() keyboard.Keyboard {
	kb := new(Keyboard)
	kb.Base = keyboard.NewBase()

	glfw.SetKeyCallback(func(key, state int) {
		kb.onKey(key, state)
	})

	glfw.SetCharCallback(func(key, state int) {
		kb.onChar(key, state)
	})

	return kb
}

func (kb *Keyboard) onKey(key, state int) {
	println("key", key, state)
}

func (kb *Keyboard) onChar(key, state int) {
	println("char", key, state)
}
