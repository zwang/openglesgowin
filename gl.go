package main

import (
	"fmt"

	gl "github.com/go-gl/gl/v3.1/gles2"

	"rogchap.com/v8go"
)

// bind openGL related commands to v8 for windows

const (
	glObjName      = "GL"
	clearColorName = "clearColor"
	clearName      = "clear"
)

// bindGLObject bind methods to "GL" object, which is one of the top level object
// for OpenGL methods
func bindGLObject(iso *v8go.Isolate, global *v8go.ObjectTemplate) error {
	glObj, err := v8go.NewObjectTemplate(iso)
	if err != nil {
		return fmt.Errorf("creating platoObject error: %w", err)
	}

	logFunc, err := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("log: %+v\n", info.Args())
		return nil
	})
	if err != nil {
		return err
	}

	err = global.Set("log", logFunc, v8go.ReadOnly)
	if err != nil {
		return err
	}

	clearColorFunc, err := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("clearColor %+v\n", info.Args())
		args := info.Args()
		gl.ClearColor(float32(args[0].Number()), float32(args[1].Number()), float32(args[2].Number()), float32(args[3].Number()))
		return nil
	})
	if err != nil {
		return err
	}

	err = glObj.Set(clearColorName, clearColorFunc, v8go.ReadOnly)
	if err != nil {
		return err
	}

	clearFunc, err := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("clear %+v\n", info.Args())
		args := info.Args()
		gl.Clear(args[0].Uint32())
		return nil
	})
	if err != nil {
		return err
	}

	err = glObj.Set(clearName, clearFunc, v8go.ReadOnly)
	if err != nil {
		return err
	}

	err = glObj.Set("COLOR_BUFFER_BIT", int32(gl.COLOR_BUFFER_BIT), v8go.ReadOnly)
	if err != nil {
		return err
	}

	err = glObj.Set("DEPTH_BUFFER_BIT", int32(gl.DEPTH_BUFFER_BIT), v8go.ReadOnly)
	if err != nil {
		return err
	}

	err = global.Set(glObjName, glObj, v8go.ReadOnly)
	return err
}
