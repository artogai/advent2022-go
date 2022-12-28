package day25

import (
	"advent2022/file"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	r := 0
	for _, s := range file.ReadFileLines("snafu.txt") {
		r += decode(s)
	}
	require.Equal(t, "20==1==12=0111=2--20", encode(r))
}
