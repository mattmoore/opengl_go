package carbon

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
  "github.com/go-gl/mathgl/mgl32"
	"os"
  "strings"
)

type Graphics struct {
	GraphicsOptions *GraphicsOptions
	Window          *sdl.Window
	Context         sdl.GLContext
}

type GraphicsOptions struct {
	Width  int
	Height int
}

func (g *Graphics) Init() {
	var err error = nil

	g.Window, err = sdl.CreateWindow("Carbon", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, g.GraphicsOptions.Width, g.GraphicsOptions.Height, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		panic(err)
	}

	err = gl.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init gl: %s\n", err)
		panic(err)
	}

	g.Context, err = sdl.GL_CreateContext(g.Window)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create context: %s\n", err)
		panic(err)
	}

//  // Configure the vertex and fragment shaders
//  program, err := newProgram(vertexShader, fragmentShader)
//  if err != nil {
//    panic(err)
//  }
//
//  gl.UseProgram(program)
//
//  projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(g.GraphicsOptions.Width) / float32(g.GraphicsOptions.Height), 0.1, 10.0)
//  projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
//  gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

//  camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
//  cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
//  gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])
}

var vertexShader = `
#version 120
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
varying vec3 vert;
varying vec2 vertTexCoord;
varying vec2 fragTexCoord;
void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 120
uniform sampler2D tex;
varying vec2 fragTexCoord;
varying vec4 outputColor;
void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
  vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
  if err != nil {
    return 0, err
  }

  fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
  if err != nil {
    return 0, err
  }

  program := gl.CreateProgram()

  gl.AttachShader(program, vertexShader)
  gl.AttachShader(program, fragmentShader)
  gl.LinkProgram(program)

  var status int32
  gl.GetProgramiv(program, gl.LINK_STATUS, &status)
  if status == gl.FALSE {
    var logLength int32
    gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

    return 0, fmt.Errorf("failed to link program: %v", log)
  }

  gl.DeleteShader(vertexShader)
  gl.DeleteShader(fragmentShader)

  return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
  shader := gl.CreateShader(shaderType)

  csources, free := gl.Strs(source)
  gl.ShaderSource(shader, 1, csources, nil)
  free()
  gl.CompileShader(shader)

  var status int32
  gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    var logLength int32
    gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

    return 0, fmt.Errorf("failed to compile %v: %v", source, log)
  }

  return shader, nil
}
