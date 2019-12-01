package jsonc

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ToJSON returns JSON equivalent of JSON with comments
func ToJSON(r io.Reader) []byte {
	buf := make([]byte, 512)
	res := []byte{}

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}

		res = append(res, extract(buf[:n])...)
	}

	return res
}

// ReadFromFile reads jsonc file and returns JSONC and JSON encodings
func ReadFromFile(filename string) ([]byte, []byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	jc := data
	j := extract(jc)
	return jc, j, nil
}

// Unmarshal parses the JSONC-encoded data and stores the result in the value pointed to by v.
// Equivalent of calling `json.Unmarshal(jsonc.ToJSON(data), v)`
func Unmarshal(data []byte, v interface{}) error {
	j := extract(data)
	return json.Unmarshal(j, v)
}

// Valid reports whether data is a valid JSONC encoding or not
func Valid(data []byte) bool {
	j := extract(data)
	return json.Valid(j)
}
