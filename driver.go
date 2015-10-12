package mbtiles

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func check(err error) {
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}
}

func dBConnect(path string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", path)
	return
}

func dBReadMetadata(db *sql.DB) (md *Metadata, err error) {
	rows, err := db.Query("select name, value from metadata")
	if err != nil {
		md = nil
		return
	}
	defer rows.Close()
	attrs := make(map[string]string)
	for rows.Next() {
		var name, value string
		rows.Scan(&name, &value)
		attrs[name] = value
	}
	md = &Metadata{attrs}
	return
}

func dBReadTile(tile *Tile, db *sql.DB) (err error) {
	stmt := "select tile_data from tiles where zoom_level=%d and tile_column=%d and tile_row=%d"
	q := fmt.Sprintf(stmt, tile.Z, tile.X, tile.Y)
	row := db.QueryRow(q)
	var blob []byte
	err = row.Scan(&blob)
	tile.Data = blob
	return
}
