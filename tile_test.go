package mbtiles

import (
	"testing"
)

func TestSniffType(t *testing.T) {
	var sniffs = []struct {
		in   []byte
		want Format
	}{
		{[]byte{}, EMPTY},
		{[]byte{0x00}, UNKNOWN},
		{[]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, PNG},
		{[]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0xFF, 0xFF}, PNG},
		{[]byte{0xFF, 0xD8, 0xFF, 0xD9}, JPG},
		{[]byte{0xFF, 0xD8, 0x00, 0xFF, 0xD9}, JPG},
	}
	for _, sniff := range sniffs {
		sniffed := (&Tile{0, 0, 0, sniff.in}).SniffType()
		if sniffed != sniff.want {
			t.Errorf("%s SniffType() wanted %s, got %s", sniff.in, sniff.want, sniffed)
		}
	}
}

/*
   case len(t.Data) < 2:
           f = EMPTY
   case r.DeepEqual(t.Data[:8], []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}):
           f = PNG
   case r.DeepEqual(t.Data[:2], []byte{0xFF, 0xD8}) && r.DeepEqual(t.Data[len(t.Data)-2:], []byte{0xFF, 0xD9}):
           f = JPG
   case r.DeepEqual(t.Data[:4], []byte{0x47, 0x49, 0x46, 0x38}) && (t.Data[4] == 0x39 || t.Data[4] == 0x37) && t.Data[5] == 0x61:
           f = GIF
   case r.DeepEqual(t.Data[:4], []byte{0x52, 0x49, 0x46, 0x46}) && r.DeepEqual(t.Data[8:12], []byte{0x57, 0x45, 0x42, 0x50}):
           f = WEBP
   case r.DeepEqual(t.Data[:2], []byte{0x78, 0x9C}):
           f = PBF_DF
   case r.DeepEqual(t.Data[:2], []byte{0x1F, 0x88}):
           f = PBF_GZ
   default:
           f = UNKNOWN
   }
*/
