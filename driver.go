package mbtiles

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Connect(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	check(err)
	return db
}

func ReadTileset(db *sql.DB) *Tileset {
	rows, err := db.Query("select name, value from metadata")
	check(err)
	defer rows.Close()
	metadata := make(map[string]string)
	for rows.Next() {
		var name, value string
		rows.Scan(&name, &value)
		metadata[name] = value
	}
	return &Tileset{metadata}
}

func ReadTile(tile *Tile, db *sql.DB) (tile *Tile) {
	q := fmt.Sprintf("select tile_data from tiles where zoom_level=%d and tile_column=%d and tile_row=%d", tile.z, tile.x, tile.y)
	row, err := db.QueryRow(q)
	check(err)
	var blob []byte
	rows.Scan(blob)
	tile.data = blob
	return
}
