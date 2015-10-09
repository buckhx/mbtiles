package mbtiles

import (
	"database/sql"
)

type MBTiles struct {
	db *sql.DB
	ts *Tileset
}

func ReadMBTiles(path string) *MBTiles {
	db := Connect(path)
	ts := ReadTileset(db)
	return &MBTiles{db, ts}
}

func (mbt *MBTiles) ReadTile(x, y, z int) (tile *Tile) {
	tile = &Tile{z, x, y}
	ReadTile(tile, mbt.db)
	return
}
