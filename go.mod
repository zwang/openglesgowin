module github.com/zwang/openglesgowin

go 1.15

require (
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20201108214237-06ea97f0c265
	rogchap.com/v8go v0.5.0
)

// This is needed for using EGL and OpenGL ES on windows
replace github.com/go-gl/gl => ../../go/src/github.com/zwang/gl

// Comment this out to use default v8go, which does not have gl bind from c++ world.
// replace rogchap.com/v8go => ../../go/src/github.com/plato-app/v8go
