module github.com/zwang/openglesgowin

go 1.16

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20201108214237-06ea97f0c265
	rogchap.com/v8go v0.5.0
)

replace rogchap.com/v8go => ../../go/src/github.com/plato-app/v8go

replace github.com/go-gl/glfw/v3.3/glfw => ../glfw/v3.3/glfw
