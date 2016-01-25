package view

import (
	"github.com/geemili/maze-rogue/model"
	"github.com/go-gl/gl/v2.1/gl"
)

func RenderPlayer(player *model.Player) {
	gl.PushMatrix()
	gl.Translatef(float32(player.X), float32(player.Y), 0)

	gl.Color4f(0.5, 0, 0.5, 1)

	gl.Begin(gl.TRIANGLES)
	gl.Vertex3f(0, 0, 0)
	gl.Vertex3f(0, 1, 0)
	gl.Vertex3f(1, 0, 0)
	gl.Vertex3f(1, 1, 0)
	gl.Vertex3f(1, 0, 0)
	gl.Vertex3f(0, 1, 0)
	gl.End()

	gl.PopMatrix()
}
