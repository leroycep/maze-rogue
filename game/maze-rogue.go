package game

import (
	"github.com/geemili/maze-rogue/generate"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type GameData struct {
	Width, Height    int
	Tiles            []int
	PlayerX, PlayerY int
}

var game GameData

func Init() {
	width, height, regionid := 80, 60, 1

	gl.Ortho(0, 40, 0, 30, -1, 3)

	rooms := generate.PlaceRooms(width, height, 100, 2, 10)                    // Place rooms between 3x3 and 5x5 in a 40 x 30 grid of tiles
	bakedRooms, regionid := generate.BakeRooms(rooms, width, height, regionid) // Render rooms down to a grid
	maze, regionid := generate.MakeMazes(bakedRooms, width, height, regionid)  // Finish up by generating mazes between rooms
	connect, regionid := generate.ConnectRooms(maze, width, height, regionid)
	trimmed := generate.TrimPaths(connect, width, height)
	x, y := 0, 0
LOOP:
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if trimmed[(j*width)+i] != 0 {
				x, y = i, j
				break LOOP
			}
		}
	}
	game = GameData{width, height, trimmed, x, y}
}

func Render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Translatef(20-float32(game.PlayerX), 15-float32(game.PlayerY), 0)
	gl.Begin(gl.TRIANGLES)
	for i := 0; i < game.Width; i++ {
	LOOP:
		for j := 0; j < game.Height; j++ {
			switch game.Tiles[(j*game.Width)+i] {
			case 0:
				continue LOOP
			case -1:
				gl.Color4f(1, 1, 1, 1)
			case 1:
				gl.Color4f(0.3, 0.4, 0.5, 1)
			case 2:
				gl.Color4f(0.7, 0.4, 0.3, 1)
			case 3:
				gl.Color4f(0.2, 0.7, 0.3, 1)
			case 4:
				gl.Color4f(0.1, 0.3, 0.7, 1)
			case 5:
				gl.Color4f(0.2, 0.3, 0.6, 1)
			case 6:
				gl.Color4f(0.3, 0.3, 0.5, 1)
			default:
				gl.Color4f(0.5, 0, 0.5, 1)
			}
			// Left Triangle
			gl.Vertex3f(float32(i), float32(j), 0)
			gl.Vertex3f(float32(i), float32(j+1), 0)
			gl.Vertex3f(float32(i+1), float32(j), 0)
			// Right Triangle
			gl.Vertex3f(float32(i+1), float32(j+1), 0)
			gl.Vertex3f(float32(i+1), float32(j), 0)
			gl.Vertex3f(float32(i), float32(j+1), 0)
		}
	}
	// Player
	gl.Color4f(1, 0, 0, 1)
	gl.Vertex3f(float32(game.PlayerX), float32(game.PlayerY), 0)
	gl.Vertex3f(float32(game.PlayerX), float32(game.PlayerY+1), 0)
	gl.Vertex3f(float32(game.PlayerX+1), float32(game.PlayerY), 0)
	gl.Vertex3f(float32(game.PlayerX+1), float32(game.PlayerY+1), 0)
	gl.Vertex3f(float32(game.PlayerX+1), float32(game.PlayerY), 0)
	gl.Vertex3f(float32(game.PlayerX), float32(game.PlayerY+1), 0)
	gl.End()
	gl.PopMatrix()
}

func OnKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	nx, ny := game.PlayerX, game.PlayerY
	switch k {
	case glfw.KeyLeft:
		nx--
	case glfw.KeyRight:
		nx++
	case glfw.KeyUp:
		ny++
	case glfw.KeyDown:
		ny--
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
	if nx >= 0 && nx < game.Width && ny >= 0 && ny < game.Height && game.Tiles[(ny*game.Width)+nx] != 0 {
		game.PlayerX, game.PlayerY = nx, ny
	}
}
