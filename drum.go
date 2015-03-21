// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Track struct {
	id    int32
	name  []uint8
	beats [16]byte
}

func (t *Track) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &t.id); err != nil {
		return err
	}
	var nameSize uint8

	if err := binary.Read(r, binary.LittleEndian, &nameSize); err != nil {
		return err
	}

	t.name = make([]uint8, nameSize)
	if err := binary.Read(r, binary.LittleEndian, &t.name); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &t.beats); err != nil {
		return err
	}

	return nil

}

func (data Track) String() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprint("(", data.id, ") ", string(data.name), "\t"))
	for i := 0; i < len(data.beats); i++ {
		if i%4 == 0 {
			buff.WriteString("|")
		}

		if data.beats[i] == 1 {
			buff.WriteString("x")
		} else {
			buff.WriteString("-")
		}
	}
	buff.WriteString("|")

	return buff.String()
}
