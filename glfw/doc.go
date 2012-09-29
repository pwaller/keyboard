// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
This is a keyboard backend for libglfw.

GLFW should be initialized by the host application. This package hooks into
the Key/Char callbacks and expects the host application to take care of
appropriate calls to `glfw.PollEvents()`.
*/
package glfw
