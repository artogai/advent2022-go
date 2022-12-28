package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "inventories.txt"

func TestTopOne(t *testing.T) {
	got, err := TopOne(path)
	require.NoError(t, err)
	require.Equal(t, 69693, got)
}

func TestSumTop(t *testing.T) {
	got, err := SumTop(path, 3)
	require.NoError(t, err)
	require.Equal(t, 200945, got)
}
