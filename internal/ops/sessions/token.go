package sessions

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

// decode go binary decoder, where argument v interface{} should be of pointer type.
func decode(str string, v any) error {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return fmt.Errorf("failed base64 Decode: %w", err)
	}

	buf := bytes.NewBuffer(b)

	d := gob.NewDecoder(buf)

	return d.Decode(v)
}

// encode go binary to string.
func encode(u any) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	if err := e.Encode(u); err != nil {
		return "", fmt.Errorf(`failed gob Encode: %w`, err)
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}
