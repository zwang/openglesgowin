package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"io/ioutil"
	"log"
	"math"
	"runtime"
)

var green = 0.2
var cursorXPos = 0.0
var cursorYPos = 0.0

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

	window, err := glfw.CreateWindow(640, 480, "Gio + GLFW", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.MakeContextCurrent()

	fileName := "js/touchDemo/main.js"
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf(fileName))
	if err != nil {
		log.Printf("unable to read file %s: %v", fileName, err)
	}
	err = window.RunScript(string(fileBytes), fileName)
	if err != nil {
		log.Printf("error runScript in %s: %v", fileName, err)
	}
	err = window.RunScript("initSurface();", "main.js")
	if err != nil {
		log.Printf("error initSurface() in %s: %v", fileName, err)
	}

	registerCallbacks(window)
	logGLInfo(window)
	for !window.ShouldClose() {
		window.MakeContextCurrent()
		err = window.RunScript("drawFrame();", "main.js")
		if err != nil {
			log.Printf("error drawFrame in %s: %v", fileName, err)
		}
		// drawOpenGL(window)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// drawOpenGL demonstrates the direct use of OpenGL commands
func drawOpenGL(window *glfw.Window) {
	err := window.RunScript(fmt.Sprintf("GL.clearColor(0, %f, 0, 1)", float32(green)), "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error clearColor %v", err))
	}
	err = window.RunScript(fmt.Sprint("GL.clear(GL.COLOR_BUFFER_BIT | GL.DEPTH_BUFFER_BIT)"), "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error clear %v", err))
	}
}

func logGLInfo(window *glfw.Window) {
	err := window.RunScript("plato.log(\"GL.ES_VERSION_2_0 =\", GL.ES_VERSION_2_0)", "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error log %v", err))
	}

	err = window.RunScript("plato.log(\"GL.COLOR_BUFFER_BIT =\", GL.COLOR_BUFFER_BIT)", "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error log %v", err))
	}

	err = window.RunScript("plato.log(\"GL.DEPTH_BUFFER_BIT =\", GL.DEPTH_BUFFER_BIT)", "main.js")
	if err != nil {
		fmt.Println(fmt.Sprintf("error log %v", err))
	}
}

func registerCallbacks(window *glfw.Window) {
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		// log.Printf("mouse cursor: (%f,%f)", xpos, ypos)
		cursorXPos = xpos
		cursorYPos = ypos
	})
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			green += 0.1
			green, _ = math.Frexp(green)
		}
		onClickCode := fmt.Sprintf("onClick(%f,%f)", cursorXPos, cursorYPos)
		err := window.RunScript(onClickCode, "main.js")
		if err != nil {
			log.Printf("error runScript %s: %v", onClickCode, err)
		}
		log.Printf("mouse button: %v action %v mods %v", button, action, mods)
	})
}
