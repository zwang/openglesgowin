module github.com/zwang/openglesgowin

go 1.15

require (
	gioui.org v0.0.0-20210206065156-6682f75db903
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20201108214237-06ea97f0c265
	golang.org/x/exp v0.0.0-20210201131500-d352d2db2ceb // indirect
	golang.org/x/image v0.0.0-20201208152932-35266b937fa6 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
)

replace github.com/go-gl/gl => ../../go/src/github.com/zwang/gl
