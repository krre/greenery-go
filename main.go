package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()
	screenWidth := videoMode.Width
	screenHeight := videoMode.Height
	screenRatio := 0.8

	width := int(screenRatio * float64(screenWidth))
	height := int(screenRatio * float64(screenHeight))
	x := int((screenWidth - width) / 2)
	y := int((screenHeight - height) / 2)

	window, err := glfw.CreateWindow(width, height, "Greenery", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetPos(x, y)
	window.MakeContextCurrent()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			if mods == glfw.ModControl {
				switch key {
				case glfw.KeyN:
					fmt.Println("New")
				case glfw.KeyO:
					fmt.Println("Open")
				case glfw.KeyS:
					fmt.Println("Save")
				case glfw.KeyD:
					fmt.Println("Save as")
				case glfw.KeyQ:
					window.SetShouldClose(true)
				}
			} else {
				switch key {
				case glfw.KeyLeft:
					fmt.Println("Left")
				case glfw.KeyRight:
					fmt.Println("Right")
				case glfw.KeyUp:
					fmt.Println("Up")
				case glfw.KeyDown:
					fmt.Println("Down")
				}
			}
		}
	})

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	gl.ClearColor(0.19, 0.12, 0.08, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Do OpenGL stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
