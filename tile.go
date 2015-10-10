package mbtiles

type Tile struct {
	z, x, y int
	data    []byte
}

func EmptyTile(z, x, y int) (tile *Tile) {
	tile = new(Tile)
	tile.x = x
	tile.y = y
	tile.z = z
	return
	//return &Tile{z, x, y}
}
