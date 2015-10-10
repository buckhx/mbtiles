package mbtiles

import (
	"testing"
)

func TestGetters(t *testing.T) {
	attrs := map[string]string{
		"name":        "test-name",
		"type":        "test-type",
		"version":     "v0.0.0-test",
		"description": "test description",
		"format":      "test-format",
		"bounds":      "-179.231086,-14.601813000000002,179.859681,71.441059",
		"attribution": "test attribution @ CC ™ ®",
	}
	m := &Metadata{attrs}
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
	if m.Attribution() != attrs["attribution"] {
		t.Errorf("Getter Error with %s", attrs)
	}
	bnds, err := m.Bounds()
	check(err)
	for _, bnd := range bnds {
		if bnd == 0 {
			t.Errorf("Getter Error with %s", attrs)
		}
	}
}
