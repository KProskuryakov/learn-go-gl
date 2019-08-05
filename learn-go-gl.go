package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"

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

	// compile the vertex shader
	var vertexShader = gl.CreateShader(gl.VERTEX_SHADER)
	cVertexShaderSource, free := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, cVertexShaderSource, nil)
	free()
	gl.CompileShader(vertexShader)
	// test success of shader compilation
	var success int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vertexShader, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to compile vertex shader: %v", log))
	}

	// compile the fragment shader
	var fragmentShader = gl.CreateShader(gl.FRAGMENT_SHADER)
	cFragmentShaderSource, free := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, cFragmentShaderSource, nil)
	free()
	gl.CompileShader(fragmentShader)

	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fragmentShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fragmentShader, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to compile fragment shader: %v", log))
	}

	// compile the fragment shader
	var fragmentShader2 = gl.CreateShader(gl.FRAGMENT_SHADER)
	cFragmentShaderSource, free = gl.Strs(fragmentShaderYellowSource)
	gl.ShaderSource(fragmentShader2, 1, cFragmentShaderSource, nil)
	free()
	gl.CompileShader(fragmentShader2)

	gl.GetShaderiv(fragmentShader2, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(fragmentShader2, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(fragmentShader2, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to compile fragment shader: %v", log))
	}

	// link shaders
	var shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to link shader program: %v", log))
	}

	// link shaders
	var shaderProgram2 = gl.CreateProgram()
	gl.AttachShader(shaderProgram2, vertexShader)
	gl.AttachShader(shaderProgram2, fragmentShader2)
	gl.LinkProgram(shaderProgram2)

	gl.GetProgramiv(shaderProgram2, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram2, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram2, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to link shader program: %v", log))
	}

	// delete the shaders when we're done loading them
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	gl.DeleteShader(fragmentShader2)

	// init vao and vbo
	var vao, vbo [2]uint32
	gl.GenVertexArrays(2, &vao[0])
	gl.GenBuffers(2, &vbo[0])

	gl.BindVertexArray(vao[0])
	// put vertices in array buffer

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo[0])
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// set vertex attribute pointers
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(vao[1])

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo[1])
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices2)*4, gl.Ptr(vertices2), gl.STATIC_DRAW)

	// set vertex attribute pointers
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	// render loop
	for !window.ShouldClose() {
		// input
		processInput(window)

		// renderArrays
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		timeVal := glfw.GetTime()
		greenVal := float32((math.Sin(timeVal) / 2.0) + 0.5)
		fmt.Printf("%f", greenVal)
		vertexColorLoc := gl.GetUniformLocation(shaderProgram2, gl.Str("ourColor\x00"))

		// draw
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(vao[0])
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		gl.UseProgram(shaderProgram2)
		gl.Uniform4f(vertexColorLoc, 0, greenVal, 0, 1)
		gl.BindVertexArray(vao[1])
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.BindVertexArray(0)

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

var vertices = []float32{
	0.5, 0.5, 0.0, // top right
	0.5, -0.5, 0.0, // bottom right
	-0.5, 0.5, 0.0, // top left
}

var vertices2 = []float32{
	0.5, -0.5, 0.0, // bottom right
	-0.5, -0.5, 0.0, // bottom left
	-0.5, 0.5, 0.0, // top left
}

var indices = []uint32{
	0, 1, 3, // first triangle
}
var indices2 = []uint32{
	1, 2, 3, // second triangle
}

const vertexShaderSource = `
#version 330 core
layout (location = 0) in vec3 aPos;

void main()
{
	gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}
`

const fragmentShaderSource = `
#version 330 core
out vec4 FragColor;

void main()
{
	FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
`

const fragmentShaderYellowSource = `
#version 330 core
out vec4 FragColor;

uniform vec4 ourColor;

void main()
{
	FragColor = ourColor;
}
`
