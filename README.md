# Go-GLFW + v8go + OpenGLES in V8

This is a demo app trying to demonstrate using go-glfw to create a window do simple OpenGL ES Commands to draw screen

In `main` branch, the OpenGL ES Commands are called in javascript in v8go (they are bound to v8 in C++ world) to draw UI in the glfw window created in golang world. However, it keeps getting `gl error: 1282` for simple gl commands like glClear and glClearColor in Windows platform.

In `bindGLInGolang` branch, the OpenGL ES Commands are also called in javascript in v8go (but they are bound to v8 in Golang world) to draw UI in the glfw window created in golang world; and this works correctly.

## Dependencies: (see go.mod for correct folder path structure)

1. github.com/go-gl/glfw/v3.3/glfw
2. github.com/plato-app/v8go (for main branch for binding OpenGL ES in C++ world)
3. https://github.com/rogchap/v8go (for bindGLInGolang branch for binding OpenGL ES in Golang world)
4. V8 for windows (see below)

#### V8 for Windows:

While no prebuilt static V8 library is included for Windows, MSYS2 provides a package containing a dynamically linked V8 library that works.

To set this up:

1. Install MSYS2 (https://www.msys2.org/)
2. Add the Mingw-w64 bin to your PATH environment variable (C:\msys64\mingw64\bin by default)
3. Open MSYS2 MSYS and execute `pacman -S mingw-w64-x86_64-toolchain mingw-w64-x86_64-v8`
4. This will allow building projects that depend on `v8go`, but, in order to actually run them, you will need to copy the `snapshot_blob.bin` file from the Mingw-w64 bin folder to your program's working directory (which is typically wherever main.go is)
5. V8 requires 64-bit on Windows, therefore will not work on 32-bit systems.
