package mbtiles

import (
	"database/sql"
)

type MBTiles struct {
	db *sql.DB
	ts *Tileset
}

func CreateMBT(path string) *MBTiles {
	db := Connect(path)
	ts := GetTileset(db)
	return &MBTiles{db, ts}
}
