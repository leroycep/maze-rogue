package generate

const (
	hwall         = 76
	vwall         = 59
	period        = 226
	leftDownElbow = 251
	leftUpElbow   = 157
	rightUpElbow  = 12
)

func BakeForTileset(tiles []int, w, h int) []int {
	baked := make([]int, w*h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			switch {
			case getTile(i, j, tiles, w, h) == 0 &&
				getTile(i-1, j, tiles, w, h) == 0 &&
				getTile(i, j-1, tiles, w, h) == 0 &&
				(getTile(i-1, j-1, tiles, w, h) != 0 || (getTile(i+1, j, tiles, w, h) != 0 && getTile(i, j+1, tiles, w, h) != 0)):
				baked[(j*w)+i] = leftDownElbow
			case getTile(i, j, tiles, w, h) == 0 &&
				getTile(i-1, j, tiles, w, h) == 0 &&
				getTile(i, j+1, tiles, w, h) == 0 &&
				(getTile(i-1, j+1, tiles, w, h) != 0 || (getTile(i+1, j, tiles, w, h) != 0 && getTile(i, j-1, tiles, w, h) != 0)):
				baked[(j*w)+i] = leftUpElbow
			case getTile(i, j, tiles, w, h) == 0 &&
				(getTile(i-1, j, tiles, w, h) != 0 || getTile(i+1, j, tiles, w, h) != 0):
				baked[(j*w)+i] = vwall
			case getTile(i, j, tiles, w, h) == 0 &&
				(getTile(i, j-1, tiles, w, h) != 0 || getTile(i, j+1, tiles, w, h) != 0):
				baked[(j*w)+i] = hwall
			case getTile(i, j, tiles, w, h) == 0:
				baked[(j*w)+i] = 0
			default:
				baked[(j*w)+i] = period
			}
		}
	}
	return baked
}
