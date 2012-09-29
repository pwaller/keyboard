// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package glfw

import (
	"github.com/jteeuwen/glfw"
	"testing"
)

func Test(t *testing.T) {
	err := glfw.Init()
	if err != nil {
		t.Fatal(err)
	}

	defer glfw.Terminate()

	err = glfw.OpenWindow(800, 600, 8, 8, 8, 0, 0, 0, glfw.Windowed)
	if err != nil {
		t.Fatal(err)
	}

	defer glfw.CloseWindow()

	kb := New()
	kb.Bind(func() {
		println("pressed s")
	}, "s")

	kb.Bind(func() {
		println("pressed ctrl+s or command+s")
	}, "ctrl+s", "command+s")

	kb.Bind(func() {
		println("pressed 's a v e'")
	}, "s a v e")

	for glfw.WindowParam(glfw.Opened) == 1 {
		glfw.PollEvents()
	}
}
