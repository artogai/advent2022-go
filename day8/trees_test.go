package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountVisible(t *testing.T) {
	require.Equal(t, 21, CountVisible("trees_sample.txt"))
	require.Equal(t, 1818, CountVisible("trees.txt"))
}
