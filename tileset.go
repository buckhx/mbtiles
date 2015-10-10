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

func (ts *Tileset) Metadata() *Metadata {
	return ts.m
}

func (ts *Tileset) Close() {
	ts.db.Close()
}
