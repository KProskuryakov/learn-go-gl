package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// https://github.com/go-gl/glfw
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// window dimensions for both glfw and our framebuffers
const width int = 800
const height int = 600

func main() {
	// init glfw for window handling
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	// call terminate when we're done
	defer glfw.Terminate()

	// configure glfw to use OpenGL ver 3.3
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	// Set CORE PROFILE which disables backwards-compatible features that are unneeded
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// May not be necessary, dunno
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create our glfw window object
	window, err := glfw.CreateWindow(width, height, "LearnOpenGL", nil, nil)
	if err != nil {
		panic(err)
	}
	// Make it the main context
	window.MakeContextCurrent()

	// Init opengl, otherwise horrible segfaults
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Set initial viewport size for normalized coordinate calculation
	gl.Viewport(0, 0, int32(width), int32(height))

	// callback when the window is resized
	callback := func(window *glfw.Window, newWidth int, newHeight int) {
		gl.Viewport(0, 0, int32(newWidth), int32(newHeight))
	}
	window.SetFramebufferSizeCallback(callback)

	// set color to use every time buffer is cleared
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	// render loop
	for !window.ShouldClose() {
		// input
		processInput(window)

		// render
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// double buffer system
		window.SwapBuffers()
		// check for input events
		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
