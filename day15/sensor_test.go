package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountScannedSample(t *testing.T) {
	require.Equal(t, 26, CountScanned(10, "sensors_sample.txt"))
}

func TestCountScanned(t *testing.T) {
	require.Equal(t, 5083287, CountScanned(2000000, "sensors.txt"))
}

func TestFindBeaconSample(t *testing.T) {
	require.Equal(t, coord{14, 11}, FindBeacon(20, "sensors_sample.txt"))
}

func TestFindBeacon(t *testing.T) {
	require.Equal(t, coord{3283509, 3205729}, FindBeacon(4000000, "sensors.txt"))
	require.Equal(t, 13134039205729, 3283509*4000000+3205729)
}
