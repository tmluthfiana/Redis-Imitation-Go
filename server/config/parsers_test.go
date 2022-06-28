package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssertByteStreamLengthError(t *testing.T) {
	err := assertNonEmptyStream([]byte{})
	require.Error(t, err)
}

func TestAssertByteStreamLengthNoError(t *testing.T) {
	err := assertNonEmptyStream([]byte{'y'})
	require.NoError(t, err)
}

func TestParseRequestSuccess(t *testing.T) {
	res, err := ParseRequest([]byte("Hello\r\n"))
	require.NoError(t, err)
	require.Equal(t, res[0], "Hello")
}
