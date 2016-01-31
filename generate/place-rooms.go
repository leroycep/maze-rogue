package generate

import "math/rand"

type Room struct {
	X, Y, W, H int
}

func PlaceRooms(width, height, number, minSize, maxSize int) []Room {
	rooms := []Room{}
	for i := 0; i < number; i++ {
		rw := (rand.Intn((maxSize-minSize)/2)+(minSize/2))*2 - 1
		rh := (rand.Intn((maxSize-minSize)/2)+(minSize/2))*2 - 1
		x := rand.Intn((width-rw)/2) * 2
		y := rand.Intn((height-rh)/2) * 2
		room := Room{x, y, rw, rh}
		canPlace := true
		for _, oroom := range rooms {
			if collides(room, oroom) {
				canPlace = false
				break
			}
		}
		if canPlace {
			rooms = append(rooms, room)
		}
	}
	return rooms
}

func collides(room, oroom Room) bool {
	return (room.X >= oroom.X && room.X <= oroom.X+oroom.W ||
		room.X+room.W >= oroom.X && room.X+room.W <= oroom.X+oroom.W) &&
		(room.Y >= oroom.Y && room.Y <= oroom.Y+oroom.H ||
			room.Y+room.H >= oroom.Y && room.Y+room.H <= oroom.Y+oroom.H)
}
