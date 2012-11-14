// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

type Modifier uint8

// Known modifier flags.
const (
	ModAlt Modifier = 1 << iota
	ModCtrl
	ModShift
	ModSuper
)

// Known key codes.
const (
	KeyUnknown Key = iota
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeySpace
	KeyApostrophe // '
	KeyComma
	KeyMinus
	KeyPeriod
	KeySlash
	KeySemicolon
	KeyEqual
	KeyLeftBracket
	KeyBackslash
	KeyRightBracket
	KeyGraveAccent
	KeyEscape
	KeyEnter
	KeyTab
	KeyBackspace
	KeyInsert
	KeyDelete
	KeyRight
	KeyLeft
	KeyDown
	KeyUp
	KeyPageUp
	KeyPageDown
	KeyHome
	KeyEnd
	KeyCapsLock
	KeyScrollLock
	KeyNumLock
	KeyPrintScreen
	KeyPause
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyKp0
	KeyKp1
	KeyKp2
	KeyKp3
	KeyKp4
	KeyKp5
	KeyKp6
	KeyKp7
	KeyKp8
	KeyKp9
	KeyKpDecimal
	KeyKpDivide
	KeyKpMultiply
	KeyKpSubtract
	KeyKpAdd
	KeyKpEnter
	KeyKpEqual
	KeyLeftShift
	KeyLeftControl
	KeyLeftAlt
	KeyLeftSuper
	KeyRightShift
	KeyRightControl
	KeyRightAlt
	KeyRightSuper
	KeyMenu
	KeyWorld1 // = Non-Us = #1
	KeyWorld2 // = Non-Us = #2
)

// KeyFromName returns a key identifier from the given name.
func KeyFromName(name string) Key {
	k := KeyUnknown

	switch name {
	case "0":
		return Key0
	case "1":
		return Key1
	case "2":
		return Key2
	case "3":
		return Key3
	case "4":
		return Key4
	case "5":
		return Key5
	case "6":
		return Key6
	case "7":
		return Key7
	case "8":
		return Key8
	case "9":
		return Key9
	case "a":
		return KeyA
	case "b":
		return KeyB
	case "c":
		return KeyC
	case "d":
		return KeyD
	case "e":
		return KeyE
	case "f":
		return KeyF
	case "g":
		return KeyG
	case "h":
		return KeyH
	case "i":
		return KeyI
	case "j":
		return KeyJ
	case "k":
		return KeyK
	case "l":
		return KeyL
	case "m":
		return KeyM
	case "n":
		return KeyN
	case "o":
		return KeyO
	case "p":
		return KeyP
	case "q":
		return KeyQ
	case "r":
		return KeyR
	case "s":
		return KeyS
	case "t":
		return KeyT
	case "u":
		return KeyU
	case "v":
		return KeyV
	case "w":
		return KeyW
	case "x":
		return KeyX
	case "y":
		return KeyY
	case "z":
		return KeyZ
	case "space":
		return KeySpace
	case "'":
		return KeyApostrophe
	case ",":
		return KeyComma
	case "-":
		return KeyMinus
	case ".":
		return KeyPeriod
	case "/":
		return KeySlash
	case ";":
		return KeySemicolon
	case "=":
		return KeyEqual
	case "[":
		return KeyLeftBracket
	case "\\":
		return KeyBackslash
	case "]":
		return KeyRightBracket
	case "`":
		return KeyGraveAccent
	case "escape":
		return KeyEscape
	case "enter":
		return KeyEnter
	case "tab":
		return KeyTab
	case "backspace":
		return KeyBackspace
	case "insert":
		return KeyInsert
	case "delete":
		return KeyDelete
	case "right":
		return KeyRight
	case "left":
		return KeyLeft
	case "down":
		return KeyDown
	case "up":
		return KeyUp
	case "pageup":
		return KeyPageUp
	case "pagedown":
		return KeyPageDown
	case "home":
		return KeyHome
	case "end":
		return KeyEnd
	case "capslock":
		return KeyCapsLock
	case "scrolllock":
		return KeyScrollLock
	case "numlock":
		return KeyNumLock
	case "printsreen":
		return KeyPrintScreen
	case "pause":
		return KeyPause
	case "f1":
		return KeyF1
	case "f2":
		return KeyF2
	case "f3":
		return KeyF3
	case "f4":
		return KeyF4
	case "f5":
		return KeyF5
	case "f6":
		return KeyF6
	case "f7":
		return KeyF7
	case "f8":
		return KeyF8
	case "f9":
		return KeyF9
	case "f10":
		return KeyF10
	case "f11":
		return KeyF11
	case "f12":
		return KeyF12
	case "f13":
		return KeyF13
	case "f14":
		return KeyF14
	case "f15":
		return KeyF15
	case "f16":
		return KeyF16
	case "f17":
		return KeyF17
	case "f18":
		return KeyF18
	case "f19":
		return KeyF19
	case "f20":
		return KeyF20
	case "f21":
		return KeyF21
	case "f22":
		return KeyF22
	case "f23":
		return KeyF23
	case "f24":
		return KeyF24
	case "f25":
		return KeyF25
	case "kp0":
		return KeyKp0
	case "kp1":
		return KeyKp1
	case "kp2":
		return KeyKp2
	case "kp3":
		return KeyKp3
	case "kp4":
		return KeyKp4
	case "kp5":
		return KeyKp5
	case "kp6":
		return KeyKp6
	case "kp7":
		return KeyKp7
	case "kp8":
		return KeyKp8
	case "kp9":
		return KeyKp9
	case "kpdecimal":
		return KeyKpDecimal
	case "kpdivide":
		return KeyKpDivide
	case "kpmultiply":
		return KeyKpMultiply
	case "kpsubtract":
		return KeyKpSubtract
	case "kpadd":
		return KeyKpAdd
	case "kpenter":
		return KeyKpEnter
	case "kpequal":
		return KeyKpEqual
	case "menu":
		return KeyMenu
	case "world1":
		return KeyWorld1
	case "world2":
		return KeyWorld2

	// These keys all refer to modified keys (ModShift + KeyXXX).
	// These assume a US/QUERTY keyboard layout.

	case "~":
		k = KeyGraveAccent
	case "!":
		k = Key1
	case "@":
		k = Key2
	case "#":
		k = Key3
	case "$":
		k = Key4
	case "%":
		k = Key5
	case "^":
		k = Key6
	case "&":
		k = Key7
	case "*":
		k = Key8
	case "(":
		k = Key9
	case ")":
		k = Key0
	case "_":
		k = KeyMinus
	case "+":
		k = KeyEqual
	case "{":
		k = KeyLeftBracket
	case "}":
		k = KeyRightBracket
	case "|":
		k = KeyBackslash
	case ":":
		k = KeySemicolon
	case "\"":
		k = KeyApostrophe
	case "<":
		k = KeyComma
	case ">":
		k = KeyPeriod
	case "?":
		k = KeySlash
	}

	return Key(ModShift)<<8 | k
}
