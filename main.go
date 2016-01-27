package main

import (
	"github.com/geemili/maze-rogue/context/move_player"
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
	player  *model.Beast
	window  *glfw.Window
	rooms   []*model.Area
	enemies []*model.Beast
)

func main() {
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

	player = &model.Beast{3, 6, 6}
	rooms = []*model.Area{
		&model.Area{5, 5, 6, 4},
		&model.Area{11, 6, 3, 1},
		&model.Area{14, 3, 5, 9},
	}
	enemies = []*model.Beast{
		&model.Beast{1, 15, 4},
		&model.Beast{1, 15, 7},
		&model.Beast{1, 17, 9},
	}

	gl.Ortho(0, 40, 0, 30, -1, 3)

	for !window.ShouldClose() {
		render()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, room := range rooms {
		view.RenderRoom(room)
	}
	for _, enemy := range enemies {
		view.RenderEnemy(enemy)
	}
	view.RenderPlayer(player)
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	roomList := make([]move_player.Room, len(rooms))
	for idx, room := range rooms {
		roomList[idx] = room
	}
	switch k {
	case glfw.KeyLeft:
		move_player.MovePlayer(player, roomList, -1, 0)
	case glfw.KeyRight:
		move_player.MovePlayer(player, roomList, 1, 0)
	case glfw.KeyUp:
		move_player.MovePlayer(player, roomList, 0, 1)
	case glfw.KeyDown:
		move_player.MovePlayer(player, roomList, 0, -1)
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}
