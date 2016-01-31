package game

import (
	"encoding/json"
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"math/rand"
	"time"
)

type Tile struct {
	X, Y, Type int
}

type Vector struct {
	X, Y int
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
		for i, _ := range squares {
			if squares[i].Type == 0 {
				directions := []Vector{Vector{0, 1}, Vector{0, -1}, Vector{1, 0}, Vector{-1, 0}}
				surrounded := true
				for _, dir := range directions {
					if !IsOccupied(squares[i].X+dir.X, squares[i].Y+dir.Y) {
						surrounded = false
					}
				}
				if surrounded {
					squares[i].Type = 1
				}
				switch rand.Intn(9) {
				case 0:
					idx := rand.Intn(len(directions))
					dir := directions[idx]
					directions = append(directions[:idx], directions[idx:]...)
					newSquare := Tile{squares[i].X + dir.X, squares[i].Y + dir.Y, 0}
					if !IsOccupied(newSquare.X, newSquare.Y) {
						squares[i].Type = 1
						squares = append(squares, newSquare)
					}
					fallthrough
				case 2, 3, 4, 5, 6, 7:
					idx := rand.Intn(len(directions))
					dir := directions[idx]
					newSquare := Tile{squares[i].X + dir.X, squares[i].Y + dir.Y, 0}
					if !IsOccupied(newSquare.X, newSquare.Y) {
						squares[i].Type = 1
						squares = append(squares, newSquare)
					}
				case 8:
					// Do nothing
				}
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

func IsOccupied(x, y int) bool {
	for _, square := range squares {
		if square.X == x && square.Y == y {
			return true
		}
	}
	return false
}

type GameData struct {
	Squares []Tile
}

func OnKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	switch k {
	case glfw.KeyLeft:
	case glfw.KeyRight:
	case glfw.KeyUp:
		output, _ := json.Marshal(GameData{squares})
		fmt.Println(string(output))
	case glfw.KeyDown:
		squares = []Tile{{5, 5, 0}}
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}
