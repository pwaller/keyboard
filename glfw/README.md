## glfw

This is a keyboard backend for libglfw.

GLFW should be initialized by the host application. This package hooks into
the Key/Char callbacks and expects the host application to take care of
appropriate calls to `glfw.PollEvents()`. Refer to `keyboard_test.go` for
an example.


### Dependencies

    go get github.com/go-gl/glfw


### Usage

    go get github.com/jteeuwen/keyboard/glfw


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

