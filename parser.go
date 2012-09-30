// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package keyboard

import (
	"strings"
)

// parseKeys reads the given key/sequence definition and
// turns it into a list of key identifiers.
func parseKeys(keys string) []Key {
	keys = strings.TrimSpace(keys)
	if len(keys) == 0 {
		return nil
	}

	keys = strings.ToLower(keys)
	words := strings.Split(keys, " ")
	list := make([]Key, 0, len(words))

	for i := range words {
		list = append(list, parseKey(words[i]))
	}

	return list
}

// parseKey reads the given key/sequence definition and
// turns it into a key identifier.
func parseKey(key string) Key {
	var value Key
	var mod Modifier

	words := strings.Split(key, "+")

	for _, w := range words {
		switch w {
		case "shift":
			mod |= ModShift
		case "alt":
			mod |= ModAlt
		case "ctrl":
			mod |= ModCtrl
		case "super", "command":
			mod |= ModSuper

		default:
			value = KeyFromName(w)
		}
	}

	return Key(mod)<<8 | value
}
