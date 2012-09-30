## keyboard

Keyboard offers a universal keyboard shortcut binding interface
for various backends. The backends are implemented as separate packages.

This is based on ccampbell's javascript library [mousetrap][mt].

[mt]: https://github.com/ccampbell/mousetrap

Keybindings for an application are specified through a backend implementation.
Apart from individual keys, it can bind sequences on which events have
to be triggered.


### Examples

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


### Sequences

A sequence binding is triggered by typing the letters in the sequence,
one after another. The package records all keypresses and triggers the
appropriate handler when a known sequence has been completed.

The recording is reset when one takes too long to type a followup character.
The timeout for this reset is configurable.


### Modifiers

This library recognizes four modifier keys: Alt, Control, Shift and Super.
These can be triggered by their respective left and right physical keys.
The super modifier is triggered on the `super` and `command` key presses.
(`command` being used on MacOSX). Therefore `super+s` is functionally
identical to `command+s`.


### Special token shortcuts

This library recognizes string binding literals like `!`, `@`, `#`, `$`, `%`,
etc. The ones that would normally require typing `shift+1`, `shift+2`, etc.
We can use these in our sequence definitions literally. The package maps them
to the appropriate key + modifier definitions. It should be noted that this
assumes a US/QUERTY keyboard layout and will likely not yield the correct
results on different layouts.

For example, with a US/QUERTY layout, the following mappings hold true:

	"!" -> ModShift + Key1
	"#" -> ModShift + Key3
	"_" -> ModShift + KeyMinus
	"{" -> ModShift + KeyLeftBracket
	"}" -> ModShift + KeyRightBracket
	"<" -> ModShift + KeyComma
	">" -> ModShift + KeyPeriod

If reliability and predictable behaviour across different types of host systems
is a prime concern, this feature should probably not be used.


### Backends

A backend is implemented as a separate package which implements the `Keyboard`
interface defined in this package. If a backend requires initialization, the
host application is responsible for doing so before calling any of the keyboard
functions.

For instance, GLFW requires calling of some initialization functions which the
keyboard package can not do for you. Reason being that the initialization
may require information specific to the host application. It is therefore up
to the caller to do the setup first, then bind all the keys.
Refer to `glfw/keyboard_test.go` for an example of this.

Each library implemented by a given backend, may impose its own restrictions
on the kind of keys one can bind. For instance, `termbox-go` has a unique
way of dealing with keyboard input and this comes with some limitations. Make
sure you read the documentation for the library in question to find out what
these restricitons may be.


### Usage

    go get github.com/jteeuwen/keyboard


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

