package mbtiles

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	attrs := map[string]string{
		"name":        "cache",
		"type":        "overlay",
		"version":     "1",
		"description": "some info here",
		"format":      "png",
		"bounds":      "-180,-85.0511,180,85.0511",
	}
	ts, err := ReadTileset("resources/world_countries.mbtiles")
	m := ts.Metadata()
	if m.Name() != attrs["name"] {
		t.Errorf("Getter Error with %s", attrs)
	}
	if m.Type() != attrs["type"] {
		t.Errorf("Getter Error with %s", attrs)
	}
	if m.Version() != attrs["version"] {
		t.Errorf("Getter Error with %s", attrs)
	}
	if m.Description() != attrs["description"] {
		t.Errorf("Getter Error with %s", attrs)
	}
	if m.Format() != attrs["format"] {
		t.Errorf("Getter Error with %s", attrs)
	}
	if m.Attribution() != "" {
		t.Errorf("Getter Error with %s", attrs)
	}
	bnds, err := m.Bounds()
	check(err)
	for _, bnd := range bnds {
		if bnd == 0 {
			t.Errorf("Getter Error with %s", attrs)
		}
	}
	tile, err := ts.ReadTile(0, 0, 0)
	check(err)
	if tile.Z != 0 {
		t.Errorf("Bad tile %s", tile)
	}
}

func TestWrite(t *testing.T) {
	test_path := "resources/test.mbtiles"
	attrs := map[string]string{
		"name":        "test",
		"type":        "overlay",
		"version":     "1",
		"description": "some info here",
		"format":      "png",
		"bounds":      "-180,-85.0511,180,85.0511",
	}
	ts, err := InitTileset(test_path, attrs)
	defer os.Remove(test_path)
	if err != nil {
		t.Errorf("Error initializing tileset %v", err)
	}
	writeTests := []*Tile{
		&Tile{0, 0, 0, []byte{1, 2, 3, 4}},
		//&Tile{1, 1, 1, []byte{}},
	}
	for _, tile := range writeTests {
		_, err := ts.WriteTile(tile.X, tile.Y, tile.Z, tile.Data)
		if err != nil {
			t.Errorf("Error writing tile %v", err)
		}
		out, err := ts.ReadTile(tile.X, tile.Y, tile.Z)
		if err != nil {
			t.Errorf("Error reading tile %v", err)
		}
		if !tile.Equals(out) {
			t.Errorf("Error reading written tile %v -> %v", tile, out)
		}
	}
}
