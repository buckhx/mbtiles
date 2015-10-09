package mbtiles

type Tile struct {
	z, x, y int32
	data    []byte
}
