package view

import (
	"github.com/geemili/maze-rogue/model"
	"github.com/go-gl/gl/v2.1/gl"
)

func RenderRoom(room *model.Area) {
	gl.PushMatrix()
	gl.Translatef(float32(room.X), float32(room.Y), 0)

	gl.Color4f(0.01, 0.5, 0.01, 1)

	gl.Begin(gl.TRIANGLES)
	// Left Triangle
	gl.Vertex3f(-.05, -.05, 1)
	gl.Vertex3f(-.05, float32(room.H)+.05, 1)
	gl.Vertex3f(float32(room.W)+.05, -.05, 1)
	// Right Triangle
	gl.Vertex3f(float32(room.W)+.05, float32(room.H)+.05, 1)
	gl.Vertex3f(float32(room.W)+.05, -.05, 1)
	gl.Vertex3f(-.05, float32(room.H)+.05, 1)
	gl.End()

	gl.PopMatrix()
}
