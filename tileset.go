package mbtiles

import (
	"database/sql"
)

type Tileset struct {
	db *sql.DB
	m  *Metadata
}

func ReadTileset(path string) *Tileset {
	db := dBConnect(path)
	m := dBReadMetadata(db)
	return &Tileset{db, m}
}

func (ts *Tileset) ReadTile(x, y, z int) (tile *Tile) {
	tile = EmptyTile(z, x, y)
	dBReadTile(tile, ts.db)
	return
}

// MBTiles use SW origin(0,0), SlippyMaps use NW
// See: http://gis.stackexchange.com/questions/116288/mbtiles-and-slippymap-tilenames
func (ts *Tileset) ReadSlippyTile(x, y, z int) (tile *Tile) {
	y = (1<<uint(z) - 1) - y
	tile = ts.ReadTile(x, y, z)
	return
}

func (ts *Tileset) Metadata() *Metadata {
	return ts.m
}

func (ts *Tileset) Close() {
	ts.db.Close()
}
