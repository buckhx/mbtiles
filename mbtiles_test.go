package mbtiles

import (
	"testing"
)

func TestRead(t *testing.T) {
	meta := map[string]string{
		"name":        "cache",
		"type":        "overlay",
		"version":     "1",
		"description": "some info here",
		"format":      "png",
		"bounds":      "-180,-85.0511,180,85.0511",
	}
	mbt := CreateMBT("resources/world_countries.mbtiles")
	ts := mbt.ts
	if ts.Name() != meta["name"] {
		t.Errorf("Getter Error with %s", meta)
	}
	if ts.Type() != meta["type"] {
		t.Errorf("Getter Error with %s", meta)
	}
	if ts.Version() != meta["version"] {
		t.Errorf("Getter Error with %s", meta)
	}
	if ts.Description() != meta["description"] {
		t.Errorf("Getter Error with %s", meta)
	}
	if ts.Format() != meta["format"] {
		t.Errorf("Getter Error with %s", meta)
	}
	if ts.Attribution() != "" {
		t.Errorf("Getter Error with %s", meta)
	}
	bnds, err := ts.Bounds()
	check(err)
	for _, bnd := range bnds {
		if bnd == 0 {
			t.Errorf("Getter Error with %s", meta)
		}
	}
}
