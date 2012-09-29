// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Keyboard is a package which offers a universal keyboard event interface
for various backends. The backends are implemented as separate packages,
which register themselves with the keyboard package.

Keybindings for an application are specified through this package.
Apart from individual keys, it can bind chains on which events have
to be triggered. For example:

Binding a single key:

	err := keyboard.Bind("s", func() {
		...
	})

Binding a single key with modifier:

	err := keyboard.Bind("ctrl+s", func() {
		...
	})

Binding a single key chain:

	err := keyboard.Bind("s a v e", func() {
		...
	})

*/
package keyboard
