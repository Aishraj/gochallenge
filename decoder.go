package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strings"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	p := &Pattern{}

	if err := binary.Read(r, binary.LittleEndian, &p.Type); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Size); err != nil {
		return nil, err
	}

	//Header validaton is complete. now we parse the actual file
	//lets read the actual

	if err := binary.Read(r, binary.LittleEndian, &p.Version); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Tempo); err != nil {
		return nil, err
	}

	//Now we parse the data
	for true {
		track := Track{}
		err := track.Decode(r)
		if err != nil {
			break
		}
		p.Tracks = append(p.Tracks, track)
	}
	return p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	Type    [6]byte
	Size    [8]byte
	Version [32]byte
	Tempo   [4]byte
	Tracks  []Track
}

// String converts a Pattern to a string
func (p Pattern) String() string {
	return fmt.Sprintf(`Saved with HW Version: %s
Tempo: %g
%s
`, bytes.Trim(p.Version[:], "\x00"), float32frombytes(p.Tempo[:]), p.printTracks())
}

func float32frombytes(tempo []byte) float32 {
	bits := binary.LittleEndian.Uint32(tempo)
	float := math.Float32frombits(bits)
	return float
}

func (p Pattern) printTracks() string {
	var buffer []string
	for _, track := range p.Tracks {
		buffer = append(buffer, track.String())
	}
	return strings.Join(buffer, "\n")
}
