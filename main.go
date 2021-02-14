package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"math"
	"runtime"
)

var jsvm *JsVM
var green = 0.2

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
		window.MakeContextCurrent()
		drawOpenGL()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// drawOpenGL demonstrates the direct use of OpenGL commands
func drawOpenGL() {
	_, err := jsvm.V8ctx.RunScript("log(\"GL.ES_VERSION_2_0 =\", GL.ES_VERSION_2_0)", "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error log %v", err))
	}

	_, err = jsvm.V8ctx.RunScript("log(\"GL.COLOR_BUFFER_BIT =\", GL.COLOR_BUFFER_BIT)", "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error log %v", err))
	}

	_, err = jsvm.V8ctx.RunScript("log(\"GL.DEPTH_BUFFER_BIT =\", GL.DEPTH_BUFFER_BIT)", "main.js")
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
