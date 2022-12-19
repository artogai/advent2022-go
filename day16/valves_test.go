package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountMaxRate(t *testing.T) {
	eds, rts := parseValves("valves.txt")
	require.Equal(t, 2114, countMaxRate(eds, rts))
}

func TestCountMaxRate2Sample(t *testing.T) {
	eds, rts := parseValves("valves_sample.txt")
	require.Equal(t, 1707, countMaxRate2(eds, rts))
}

func TestCountMaxRate2(t *testing.T) {
	eds, rts := parseValves("valves.txt")
	require.Equal(t, 2666, countMaxRate2(eds, rts))
}
