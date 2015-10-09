package mbtiles

import (
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestGetters(t *testing.T) {
	meta := map[string]string{
		"name":        "test-name",
		"type":        "test-type",
		"version":     "v0.0.0-test",
		"description": "test description",
		"format":      "test-format",
		"bounds":      "-179.231086,-14.601813000000002,179.859681,71.441059",
		"attribution": "test attribution @ CC ™ ®",
	}
	ts := &Tileset{meta}
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
	if ts.Attribution() != meta["attribution"] {
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
