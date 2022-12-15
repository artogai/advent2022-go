package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountScanned(t *testing.T) {
	require.Equal(t, 5083287, CountScanned(2000000, "sensors.txt"))
}
