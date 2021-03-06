package main

import (
	"fmt"

	"rogchap.com/v8go"
)

type JsVM struct {
	Isolate *v8go.Isolate
	V8ctx   *v8go.Context
	Global  *v8go.ObjectTemplate
}

func NewJsVM() (*JsVM, error) {
	iso, err := v8go.NewIsolate()
	if err != nil {
		return nil, fmt.Errorf("newIsolate error: %w", err)
	}

	global, err := v8go.NewObjectTemplate(iso)
	if err != nil {
		return nil, fmt.Errorf("newObjectTemplate error: %w", err)
	}

	logFunc, err := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("log: %+v\n", info.Args())
		return nil
	})

	if err != nil {
		return nil, err
	}

	err = global.Set("log", logFunc, v8go.ReadOnly)
	if err != nil {
		return nil, err
	}

	// creates a new V8 context with specified Isolate and ObjectTemplate
	v8Ctx, err := v8go.NewContext(iso, global)
	if err != nil {
		return nil, err
	}

	return &JsVM{
		Isolate: iso,
		V8ctx:   v8Ctx,
		Global:  global,
	}, nil
}
