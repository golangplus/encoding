package jsonp

import (
	"encoding/json"
	"os"
)

// MarshalToFile encodes a data into a JSON file.
func MarshalToFile(v interface{}, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(v)
}

// UnmarshalFile decodes a JSON from a file.
func UnmarshalFile(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}

// MarshalIgnoreError call json.Marshal and ignores the error.
// Use this function when you are sure the error will not happen.
func MarshalIgnoreError(v interface{}) []byte {
	res, _ := json.Marshal(v)
	return res
}
