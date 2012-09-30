// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

import (
	"testing"
)

func TestParse(t *testing.T) {
	list := []struct {
		S string
		R []Key
	}{
		{"s",
			[]Key{
				Key(KeyS),
			},
		},
		{"ctrl+s",
			[]Key{
				Key(ModCtrl)<<8 | KeyS,
			},
		},
		{"shift+ctrl+s",
			[]Key{
				Key(ModShift|ModCtrl)<<8 | KeyS,
			},
		},
		{"s a v e",
			[]Key{
				KeyS, KeyA, KeyV, KeyE,
			},
		},
		{"up up down down left right left right b a enter",
			[]Key{
				KeyUp, KeyUp, KeyDown, KeyDown, KeyLeft, KeyRight, KeyLeft,
				KeyRight, KeyB, KeyA, KeyEnter,
			},
		},
		{"! { }",
			[]Key{
				Key(ModShift)<<8 | Key1,
				Key(ModShift)<<8 | KeyLeftBracket,
				Key(ModShift)<<8 | KeyRightBracket,
			},
		},
	}

	for i, test := range list {
		list := parseKeys(test.S)

		if len(list) != len(test.R) {
			t.Fatalf("%d: %q: length mismatch. Want %d, have %d.",
				i, test.S, len(test.R), len(list))
		}

		for k := range list {
			if list[k] != test.R[k] {
				t.Fatalf("%d: %q: Want %d, have %d.",
					i, test.S, test.R[k], list[k])
			}
		}
	}
}
