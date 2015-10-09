package mbtiles

import (
	"strings"
)

const (
	NAME = "name"
	TYPE = "type"
	VERS = "version"
	DESC = "description"
	FMT  = "format"
	BNDS = "bounds"
	ATTR = "attribution"
)

type Tileset struct {
	metadata map[string]string
}

func (ts *Tileset) Metadata() map[string]string {
	return ts.metadata
}

func (ts *Tileset) Name() string {
	return ts.metadata[NAME]
}

func (ts *Tileset) Type() string {
	return ts.metadata[TYPE]
}

func (ts *Tileset) Version() string {
	return ts.metadata[VERS]
}

func (ts *Tileset) Description() string {
	return ts.metadata[DESC]
}

func (ts *Tileset) Format() string {
	return ts.metadata[FMT]
}

func (ts *Tileset) Bounds() ([4]Coordinate, error) {
	var bnds [4]Coordinate
	for i, coordstr := range strings.Split(ts.metadata[BNDS], ",") {
		c, err := ParseCoordinate(coordstr)
		if err != nil {
			return [4]Coordinate{}, err
		}
		bnds[i] = c
	}
	return bnds, nil
}

func (ts *Tileset) Attribution() string {
	return ts.metadata[ATTR]
}
