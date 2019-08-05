package shaders

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Shader Program ID
type Shader struct {
	ProgramID uint32
}

// MakeShader creates program from the vertex and fragment shaders specified
func MakeShader(vertexPath, fragmentPath string) (shader Shader) {
	vertexShader := compileShader(gl.VERTEX_SHADER, vertexPath)
	fragmentShader := compileShader(gl.FRAGMENT_SHADER, fragmentPath)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	var success int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to link shader program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	shader = Shader{ProgramID: shaderProgram}
	return
}

func compileShader(xtype uint32, path string) uint32 {
	shader := gl.CreateShader(xtype)
	csource, free := readFile(path)
	gl.ShaderSource(shader, 1, csource, nil)
	free()

	gl.CompileShader(shader)
	// test success of shader compilation
	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		panic(fmt.Errorf("failed to compile shader from path %v: %v", path, log))
	}

	return shader
}

// Use sets the current Program to the one specified
func (shader *Shader) Use() {
	gl.UseProgram(shader.ProgramID)
}

// readFile reads the file with a null term at the end
func readFile(path string) (cstrs **uint8, free func()) {
	data, err := ioutil.ReadFile(path)
	check(err)
	return gl.Strs(string(data) + "\x00")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
