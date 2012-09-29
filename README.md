## keyboard

Keyboard is a package which offers a universal keyboard event interface
for various backends. The backends are implemented as separate packages.

This is based on ccampbell's javascript library [mousetrap][mt].

[mt]: https://github.com/ccampbell/mousetrap

Keybindings for an application are specified through this package.
Apart from individual keys, it can bind sequences on which events have
to be triggered.


### Examples

Createing a new keyboard:

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


### Chained bindings

A chained key binding is triggered by typing the letters in the chain,
one after another. The package records all keypresses and triggers the
appropriate handler when a known chain has been completed.

The recording is reset when a wrong character is typed, or one takes too
long to type followup characters. The timeout is one second by default,
but this value can be customized.


### Backends

A backend is implemented as a separate package which implements the interface
defined in this package. If a backend requires initialization, the host
application is responsible for doing so before calling any of the keyboard
functions.

For instance, GLFW requires calling of some initialization functions which the
keyboard package can not do for you. Reason being that the initialization
may require information specific to the host application. It is therefore up
to the caller to do the setup first, then bind all the keys.


### Usage

    go get github.com/jteeuwen/keyboard


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

