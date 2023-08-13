package sbs1

import (
	"encoding/csv"
	"errors"
	"io"
)

const (
	timeFormat = "2006/01/02T15:04:05.999"
)

var (
	ErrUnkownMessageType       = errors.New("unknown message type")
	ErrUnknownTransmissionType = errors.New("unknown transmission type")
)

func NewReader(r io.Reader) *Reader {
	var csvr = csv.NewReader(r)

	return &Reader{csvr: csvr}
}

type Reader struct {
	csvr *csv.Reader
}

func (r *Reader) Read() (*Message, error) {
	fields, err := r.csvr.Read()

	if err != nil {
		return nil, err
	}

	return Parse(fields)
}
