package game

import (
	"github.com/geemili/maze-rogue/generate"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type GameData struct {
	Width, Height int
	Tiles         []int
}

var game GameData

func Init() {
	rooms := generate.PlaceRooms(40, 30, 100, 4, 8) // Place rooms between 3x3 and 5x5 in a 40 x 30 grid of tiles
	bakedRooms := generate.BakeRooms(rooms, 40, 30) // Render rooms down to a grid
	maze := generate.MakeMazes(bakedRooms, 40, 30)  // Finish up by generating mazes between rooms
	game = GameData{40, 30, maze}
}

func Render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Color4f(0.5, 0, 0.5, 1)

	gl.Begin(gl.TRIANGLES)
	for i := 0; i < game.Width; i++ {
		for j := 0; j < game.Height; j++ {
			switch game.Tiles[(j*game.Width)+i] {
			case 1:
				// Left Triangle
				gl.Vertex3f(float32(i), float32(j), 0)
				gl.Vertex3f(float32(i), float32(j+1), 0)
				gl.Vertex3f(float32(i+1), float32(j), 0)
				// Right Triangle
				gl.Vertex3f(float32(i+1), float32(j+1), 0)
				gl.Vertex3f(float32(i+1), float32(j), 0)
				gl.Vertex3f(float32(i), float32(j+1), 0)
			default:
			}
		}
	}
	gl.End()
}

func OnKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	switch k {
	case glfw.KeyLeft:
	case glfw.KeyRight:
	case glfw.KeyUp:
	case glfw.KeyDown:
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}
