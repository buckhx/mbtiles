package mbtiles

type Tile struct {
	Z, X, Y int
	Data    []byte
}

func EmptyTile(z, x, y int) (tile *Tile) {
	tile = new(Tile)
	tile.X = x
	tile.Y = y
	tile.Z = z
	return
	//return &Tile{z, x, y}
}
