package game

import (
	"github.com/geemili/maze-rogue/generate"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

type GameData struct {
	Width, Height    int
	Tiles            []int
	PlayerX, PlayerY int
	Texture          uint32
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
	baked := generate.BakeForTileset(trimmed, width, height)
	texture := newTexture("assets/terminal.png")
	game = GameData{width, height, baked, x, y, texture}
}

func Render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Color4f(1, 1, 1, 1)
	gl.BindTexture(gl.TEXTURE_2D, game.Texture)

	gl.PushMatrix()
	gl.Translatef(20-float32(game.PlayerX), 15-float32(game.PlayerY), 0)
	gl.Begin(gl.TRIANGLES)
	for i := 0; i < game.Width; i++ {
		for j := 0; j < game.Height; j++ {
			renderTile(game.Tiles[(j*game.Width)+i], float32(i), float32(j))
		}
	}
	// Player
	gl.Color4f(1, 0, 0, 1)
	renderTile(4, float32(game.PlayerX), float32(game.PlayerY))
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
	if nx >= 0 && nx < game.Width && ny >= 0 && ny < game.Height && game.Tiles[(ny*game.Width)+nx] == 226 {
		game.PlayerX, game.PlayerY = nx, ny
	}
}

func newTexture(file string) uint32 {
	imgFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func renderTile(id int, x, y float32) {
	var tx, ty, w, h float32 = 0, 0, 0.0625, 0.0625
	tilex := id % 16
	tiley := (id - tilex) / 16
	tx = float32(tilex)
	ty = float32(tiley)
	gl.TexCoord2f(tx/16.0, ty/16.0+h)
	gl.Vertex3f(x, y, 0)
	gl.TexCoord2f(tx/16.0, ty/16.0)
	gl.Vertex3f(x, y+1, 0)
	gl.TexCoord2f(tx/16.0+w, ty/16.0+h)
	gl.Vertex3f(x+1, y, 0)
	gl.TexCoord2f(tx/16.0+w, ty/16.0)
	gl.Vertex3f(x+1, y+1, 0)
	gl.TexCoord2f(tx/16.0+w, ty/16.0+h)
	gl.Vertex3f(x+1, y, 0)
	gl.TexCoord2f(tx/16.0, ty/16.0)
	gl.Vertex3f(x, y+1, 0)
}
