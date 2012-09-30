// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package glfw

import (
	"github.com/jteeuwen/glfw"
	"github.com/jteeuwen/keyboard"
)

// Keyboard is a keyboard backend for libGLFW.
type Keyboard struct {
	*keyboard.Base
	keymap map[int]string
	mods   keyboard.Modifier // Active key modifiers.
}

// New constructs a new keyboard instance.
func New() keyboard.Keyboard {
	kb := new(Keyboard)
	kb.Base = keyboard.NewBase()

	glfw.SetKeyCallback(func(key, state int) {
		kb.onKey(key, state)
	})

	return kb
}

// isMod returns the appropriate modifier value
// if the given key is a known modifier.
func isMod(key int) keyboard.Modifier {
	switch key {
	case glfw.KeyLshift, glfw.KeyRshift:
		return keyboard.ModShift

	case glfw.KeyLctrl, glfw.KeyRctrl:
		return keyboard.ModCtrl

	case glfw.KeyLalt, glfw.KeyRalt:
		return keyboard.ModAlt

	case glfw.KeyLsuper, glfw.KeyRsuper:
		return keyboard.ModSuper
	}

	return 0
}

// onKey handles GLFW key events.
func (kb *Keyboard) onKey(key, state int) {
	if mod := isMod(key); mod > 0 {
		if state == 1 {
			kb.mods |= mod
		} else {
			kb.mods &^= mod
		}

		return
	}

	var k keyboard.Key

	switch key {
	case 'A':
		k = keyboard.KeyA
	case 'B':
		k = keyboard.KeyB
	case 'C':
		k = keyboard.KeyC
	case 'D':
		k = keyboard.KeyD
	case 'E':
		k = keyboard.KeyE
	case 'F':
		k = keyboard.KeyF
	case 'G':
		k = keyboard.KeyG
	case 'H':
		k = keyboard.KeyH
	case 'I':
		k = keyboard.KeyI
	case 'J':
		k = keyboard.KeyJ
	case 'K':
		k = keyboard.KeyK
	case 'L':
		k = keyboard.KeyL
	case 'M':
		k = keyboard.KeyM
	case 'N':
		k = keyboard.KeyN
	case 'O':
		k = keyboard.KeyO
	case 'P':
		k = keyboard.KeyP
	case 'Q':
		k = keyboard.KeyQ
	case 'R':
		k = keyboard.KeyR
	case 'S':
		k = keyboard.KeyS
	case 'T':
		k = keyboard.KeyT
	case 'U':
		k = keyboard.KeyU
	case 'V':
		k = keyboard.KeyV
	case 'W':
		k = keyboard.KeyW
	case 'X':
		k = keyboard.KeyX
	case 'Y':
		k = keyboard.KeyY
	case 'Z':
		k = keyboard.KeyZ
	case '0':
		k = keyboard.Key0
	case '1':
		k = keyboard.Key1
	case '2':
		k = keyboard.Key2
	case '3':
		k = keyboard.Key3
	case '4':
		k = keyboard.Key4
	case '5':
		k = keyboard.Key5
	case '6':
		k = keyboard.Key6
	case '7':
		k = keyboard.Key7
	case '8':
		k = keyboard.Key8
	case '9':
		k = keyboard.Key9
	case '`':
		k = keyboard.KeyGraveAccent
	case ',':
		k = keyboard.KeyComma
	case '.':
		k = keyboard.KeyPeriod
	case '/':
		k = keyboard.KeySlash
	case ';':
		k = keyboard.KeySemicolon
	case '\'':
		k = keyboard.KeyApostrophe
	case '-':
		k = keyboard.KeyMinus
	case '=':
		k = keyboard.KeyEqual
	case '[':
		k = keyboard.KeyLeftBracket
	case ']':
		k = keyboard.KeyRightBracket
	case '\\':
		k = keyboard.KeyBackslash

	case glfw.KeySpace:
		k = keyboard.KeySpace
	case glfw.KeyEsc:
		k = keyboard.KeyEscape
	case glfw.KeyF1:
		k = keyboard.KeyF1
	case glfw.KeyF2:
		k = keyboard.KeyF2
	case glfw.KeyF3:
		k = keyboard.KeyF3
	case glfw.KeyF4:
		k = keyboard.KeyF4
	case glfw.KeyF5:
		k = keyboard.KeyF5
	case glfw.KeyF6:
		k = keyboard.KeyF6
	case glfw.KeyF7:
		k = keyboard.KeyF7
	case glfw.KeyF8:
		k = keyboard.KeyF8
	case glfw.KeyF9:
		k = keyboard.KeyF9
	case glfw.KeyF10:
		k = keyboard.KeyF10
	case glfw.KeyF11:
		k = keyboard.KeyF11
	case glfw.KeyF12:
		k = keyboard.KeyF12
	case glfw.KeyF13:
		k = keyboard.KeyF13
	case glfw.KeyF14:
		k = keyboard.KeyF14
	case glfw.KeyF15:
		k = keyboard.KeyF15
	case glfw.KeyF16:
		k = keyboard.KeyF16
	case glfw.KeyF17:
		k = keyboard.KeyF17
	case glfw.KeyF18:
		k = keyboard.KeyF18
	case glfw.KeyF19:
		k = keyboard.KeyF19
	case glfw.KeyF20:
		k = keyboard.KeyF20
	case glfw.KeyF21:
		k = keyboard.KeyF21
	case glfw.KeyF22:
		k = keyboard.KeyF22
	case glfw.KeyF23:
		k = keyboard.KeyF23
	case glfw.KeyF25:
		k = keyboard.KeyF25
	case glfw.KeyUp:
		k = keyboard.KeyUp
	case glfw.KeyDown:
		k = keyboard.KeyDown
	case glfw.KeyLeft:
		k = keyboard.KeyLeft
	case glfw.KeyRight:
		k = keyboard.KeyRight
	case glfw.KeyTab:
		k = keyboard.KeyTab
	case glfw.KeyEnter:
		k = keyboard.KeyEnter
	case glfw.KeyBackspace:
		k = keyboard.KeyBackspace
	case glfw.KeyInsert:
		k = keyboard.KeyInsert
	case glfw.KeyDel:
		k = keyboard.KeyDelete
	case glfw.KeyPageup:
		k = keyboard.KeyPageUp
	case glfw.KeyPagedown:
		k = keyboard.KeyPageDown
	case glfw.KeyHome:
		k = keyboard.KeyHome
	case glfw.KeyEnd:
		k = keyboard.KeyEnd
	case glfw.KeyKP0:
		k = keyboard.KeyKp0
	case glfw.KeyKP1:
		k = keyboard.KeyKp1
	case glfw.KeyKP2:
		k = keyboard.KeyKp2
	case glfw.KeyKP3:
		k = keyboard.KeyKp3
	case glfw.KeyKP4:
		k = keyboard.KeyKp4
	case glfw.KeyKP5:
		k = keyboard.KeyKp5
	case glfw.KeyKP6:
		k = keyboard.KeyKp6
	case glfw.KeyKP7:
		k = keyboard.KeyKp7
	case glfw.KeyKP8:
		k = keyboard.KeyKp8
	case glfw.KeyKP9:
		k = keyboard.KeyKp9
	case glfw.KeyKPDidivde:
		k = keyboard.KeyKpDivide
	case glfw.KeyKPMultiply:
		k = keyboard.KeyKpMultiply
	case glfw.KeyKPSubtract:
		k = keyboard.KeyKpSubtract
	case glfw.KeyKPAdd:
		k = keyboard.KeyKpAdd
	case glfw.KeyKPDecimal:
		k = keyboard.KeyKpDecimal
	case glfw.KeyKPEqual:
		k = keyboard.KeyKpEqual
	case glfw.KeyKPEnter:
		k = keyboard.KeyKpEnter
	case glfw.KeyKPNumlock:
		k = keyboard.KeyNumLock
	case glfw.KeyCapslock:
		k = keyboard.KeyCapsLock
	case glfw.KeyScrolllock:
		k = keyboard.KeyScrollLock
	case glfw.KeyPause:
		k = keyboard.KeyPause
	case glfw.KeyMenu:
		k = keyboard.KeyMenu
	default:
		k = keyboard.KeyUnknown
	}

	if state == 1 {
		kb.RecordKey(k, kb.mods)
	}
}
