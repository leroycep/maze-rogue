package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Tile struct {
	X, Y int
}

var squares []Tile

func Init() {
	squares = []Tile{
		Tile{5, 5},
		Tile{5, 6},
		Tile{6, 5},
		Tile{6, 6},
		Tile{9, 10},
	}
}

func Render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, square := range squares {
		gl.PushMatrix()
		gl.Translatef(float32(square.X), float32(square.Y), 0)

		gl.Color4f(0.5, 0, 0.5, 1)

		gl.Begin(gl.TRIANGLES)
		// Left Triangle
		gl.Vertex3f(0, 0, 0)
		gl.Vertex3f(0, 1, 0)
		gl.Vertex3f(1, 0, 0)
		// Right Triangle
		gl.Vertex3f(1, 1, 0)
		gl.Vertex3f(1, 0, 0)
		gl.Vertex3f(0, 1, 0)
		gl.End()

		gl.PopMatrix()
	}
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
