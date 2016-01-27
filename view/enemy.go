package view

import (
	"github.com/geemili/maze-rogue/model"
	"github.com/go-gl/gl/v2.1/gl"
)

func RenderEnemy(enemy *model.Beast) {
	gl.PushMatrix()
	gl.Translatef(float32(enemy.X), float32(enemy.Y), 0)

	gl.Color4f(1, 0.1, 0.1, 1)

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
