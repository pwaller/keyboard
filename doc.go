// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Keyboard offers a universal keyboard shortcut binding interface
for various backends. The backends are implemented as separate packages.

Keybindings for an application are specified through a backend implementation.
Apart from individual keys, it can bind sequences on which events have
to be triggered. For example:

Creating a new keyboard (for GLFW in this case):

	import "github.com/jteeuwen/keyboard/glfw"
	...
	kb := glfw.New()

Binding a single key:

	kb.Bind(func() {
		println("'s' was pressed.")
	}, "s")

Binding a single key with modifier:

	kb.Bind(func() {
		println("'ctrl+s' was pressed.")
	}, "ctrl+s")

Binding a single key with multiple modifiers:

	kb.Bind(func() {
		println("'shift+ctrl+s' was pressed.")
	}, "shift+ctrl+s")

Binding multiple shortcuts to the same handler:

	kb.Bind(func() {
		println("'ctrl+s' or 'command+s' was pressed.")
	}, "ctrl+s", "command+s")

Binding a key sequence:

	kb.Bind(func() {
		println("Konami Code!")
	}, "up up down down left right left right b a enter")

*/
package keyboard
