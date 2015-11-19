package mbtiles

import (
	"database/sql"
)

type Tileset struct {
	db *sql.DB
	m  *Metadata
}

func ReadTileset(path string) (ts *Tileset, err error) {
	db, err := dBConnect(path)
	if err != nil {
		return
	}
	md, err := dBReadMetadata(db)
	if err != nil {
		return
	}
	ts = &Tileset{db, md}
	return
}

func (ts *Tileset) ReadTile(x, y, z int) (tile *Tile, err error) {
	tile = EmptyTile(z, x, y)
	if err = dBReadTile(tile, ts.db); err == sql.ErrNoRows {
		// if the row was empty, just keep an empty tile
		// and don't throw an error
		err = nil
	}
	return
}

// MBTiles use SW origin(0,0), SlippyMaps use NW
// See: http://gis.stackexchange.com/questions/116288/mbtiles-and-slippymap-tilenames
func (ts *Tileset) ReadSlippyTile(x, y, z int) (tile *Tile, err error) {
	y = (1<<uint(z) - 1) - y
	tile, err = ts.ReadTile(x, y, z)
	return
}

func (ts *Tileset) Metadata() *Metadata {
	return ts.m
}

func (ts *Tileset) Close() {
	ts.db.Close()
}
