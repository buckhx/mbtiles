package mbtiles

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Connect(path string) sql.DB {
	db, err := sql.Open("sqlite3", path)
	check(err)
	return db
}

func GetTileset(db sql.DB) *Tileset {
	rows, err := db.Query("select name, value from metadata")
	check(err)
	defer rows.Close()
	metadata = make(map[string]string)
	for rows.Next() {
		var name, value string
		rows.Scan(&name, &value)
		metadata[name] = value
	}
	return &Tileset{metadata}
}
