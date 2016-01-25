package main

import (
	"fmt"
	"github.com/geemili/maze-rogue/model"
	"github.com/geemili/maze-rogue/view"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

var (
	player *model.Player
	window *glfw.Window
)

func main() {
	fmt.Println("Hello, world!")

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.SetKeyCallback(onKey)

	player = &model.Player{3, -1, -1}

	for !window.ShouldClose() {
		render()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	view.RenderPlayer(player)
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch k {
	case glfw.KeyLeft:
		player.X -= 1
	case glfw.KeyRight:
		player.X += 1
	case glfw.KeyUp:
		player.Y += 1
	case glfw.KeyDown:
		player.Y -= 1
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}
