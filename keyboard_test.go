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
				Key(ModCtrl)<<8 | Key(KeyS),
			},
		},
		{"shift+ctrl+s",
			[]Key{
				(Key(ModShift)|Key(ModCtrl))<<8 | Key(KeyS),
			},
		},
		{"s a v e",
			[]Key{
				Key(KeyS), Key(KeyA), Key(KeyV), Key(KeyE),
			},
		},
		{"up up down down left right left right b a enter",
			[]Key{
				Key(KeyUp), Key(KeyUp), Key(KeyDown), Key(KeyDown),
				Key(KeyLeft), Key(KeyRight), Key(KeyLeft), Key(KeyRight),
				Key(KeyB), Key(KeyA), Key(KeyEnter),
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
