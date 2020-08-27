package sessions

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/pkg/errors"
)

// decode go binary decoder, where argument v interface{} should be of pointer type.
func decode(str string, v interface{}) error {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return errors.Wrap(err, `failed base64 Decode`)
	}

	buf := bytes.NewBuffer(b)

	d := gob.NewDecoder(buf)

	return d.Decode(v)
}

// encode go binary to string.
func encode(u interface{}) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	if err := e.Encode(u); err != nil {
		return "", errors.Wrap(err, `failed gob Encode`)
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}
