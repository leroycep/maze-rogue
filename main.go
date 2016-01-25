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

var player *model.Player

func main() {
	fmt.Println("Hello, world!")

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	player = &model.Player{3, 0, 0}

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
