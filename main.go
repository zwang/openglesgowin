// SPDX-License-Identifier: Unlicense OR MIT

// GLFW doesn't build on OpenBSD and FreeBSD.
// +build !openbsd,!freebsd,!android,!ios,!js

// The glfw example demonstrates integration of Gio into a foreign
// windowing and rendering library, in this case GLFW
// (https://www.glfw.org).
//
// See the go-glfw package for installation of the native
// dependencies:
//
// https://github.com/go-gl/glfw
package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"math"
	"runtime"
)

// desktopGL is true when the (core, desktop) OpenGL should
// be used, false for OpenGL ES.
const desktopGL = runtime.GOOS == "darwin"

var jsvm *JsVM

func init() {
	// Required by the OpenGL threading model.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()
	// Gio assumes a sRGB back buffer.
	glfw.WindowHint(glfw.SRGBCapable, glfw.True)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)
	glfw.WindowHint(glfw.CocoaRetinaFramebuffer, glfw.True)
	glfw.WindowHint(glfw.ContextCreationAPI, glfw.EGLContextAPI)
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)

	window, err := glfw.CreateWindow(800, 600, "Gio + GLFW", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()

	jsvm, err = NewJsVM()
	if err != nil {
		log.Fatalf("unable to get new jsVM. %v", err)
	}

	registerCallbacks(window)
	for !window.ShouldClose() {
		glfw.PollEvents()
		drawOpenGL()
		window.SwapBuffers()
	}
}

var (
	green float64 = 0.2
)

// drawOpenGL demonstrates the direct use of OpenGL commands
// to draw non-Gio content below the Gio UI.
func drawOpenGL() {
	_, err := jsvm.V8ctx.RunScript("log(\"GL.ES_VERSION_2_0 =\", GL.ES_VERSION_2_0)", "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error log %v", err))
	}

	_, err = jsvm.V8ctx.RunScript(fmt.Sprintf("GL.clearColor(0, %f, 0, 1)", float32(green)), "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error clearColor %v", err))
	}
	_, err = jsvm.V8ctx.RunScript(fmt.Sprint("GL.clear(GL.COLOR_BUFFER_BIT | GL.DEPTH_BUFFER_BIT)"), "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error clear %v", err))
	}
}

func registerCallbacks(window *glfw.Window) {
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		// log.Printf("mouse cursor: (%f,%f)", xpos, ypos)
	})
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			green += 0.1
			green, _ = math.Frexp(green)
		}
		log.Printf("mouse button: %v action %v mods %v", button, action, mods)
	})
}
