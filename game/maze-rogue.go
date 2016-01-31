package game

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"time"
)

type Tile struct {
	X, Y, Type int
}

var (
	squares []Tile
	tick    chan bool
)

func Init() {
	squares = []Tile{
		Tile{9, 5, 0},
	}
	tick = make(chan bool)
	go func() {
		for {
			tick <- true
			time.Sleep(1 * time.Second)
		}
	}()
}

func Render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	select {
	case <-tick:
		for _, square := range squares {
			if square.Type == 0 {
				square.Type = 1
				newSquare := Tile{square.X, square.Y + 1, 0}
				squares = append(squares, newSquare)
			}
		}
	}

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
