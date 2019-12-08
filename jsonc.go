MIT License

Copyright (c) 2019 Muhammad Muzzammil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

package jsonc

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ToJSON returns JSON equivalent of JSON with comments
func ToJSON(r io.Reader) []byte {
	ioutil.ReadAll(r)
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
