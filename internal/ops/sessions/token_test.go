package sessions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestToken is a test function.
func TestToken(t *testing.T) {
	type user struct {
		ID   int64
		Name string
	}

	u := user{ID: 1, Name: "John"}

	token, err := encode(u)
	require.NoError(t, err)

	var u2 user

	err = decode(token, &u2)
	require.NoError(t, err)

	require.Equal(t, u, u2)
}
