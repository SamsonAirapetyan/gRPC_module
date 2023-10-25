package data

import (
	"encoding/json"
	"io"
)

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func FromJSON(i interface{}, w io.Reader) error {
	e := json.NewDecoder(w)
	return e.Decode(i)
}
