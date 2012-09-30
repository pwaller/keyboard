// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package glfw

import (
	"fmt"
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
		println("pressed 't e s t'")
	}, "t e s t")

	kb.Bind(func() {
		println("pressed '. f 3'")
	}, ". f 3")

	kb.Bind(func() {
		println("pressed '. f #'")
	}, ". f #")

	fmt.Printf("Known key bindings:\n")
	for _, b := range kb.Bindings() {
		fmt.Printf(" - %q\n", b)
	}

	for glfw.WindowParam(glfw.Opened) == 1 {
		glfw.SwapBuffers()
	}
}
