package mbtiles

import (
	"database/sql"
)

type MBTiles struct {
	db *sql.DB
	ts *Tileset
}

func ReadMBTiles(path string) *MBTiles {
	db := dBConnect(path)
	ts := dBReadTileset(db)
	return &MBTiles{db, ts}
}

func (mbt *MBTiles) ReadTile(x, y, z int) (tile *Tile) {
	tile = EmptyTile(z, x, y)
	dBReadTile(tile, mbt.db)
	return
}
