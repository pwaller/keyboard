// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Keyboard is a package which offers a universal keyboard event interface
for various backends. The backends are implemented as separate packages,
which register themselves with the keyboard package.

Keybindings for an application are specified through a backend implementation.
Apart from individual keys, it can bind chains on which events have
to be triggered. For example:

Creating a new keyboard:

	import "github.com/jteeuwen/keyboard/glfw"
	...
	kb := glfw.New()

Binding a single key:

	kb.Bind(func() {
		...
	}, "s")

Binding a single key with modifier:

	kb.Bind(func() {
		...
	}, "ctrl+s")

Binding multiple shortcuts to the same handler:

	kb.Bind(func() {
		...
	}, "ctrl+s", "command+s")

Binding a key sequence:

	kb.Bind(func() {
		...
	}, "s a v e")

	kb.Bind(func() {
		...
	}, "up up down down left right left right b a enter")

*/
package keyboard
