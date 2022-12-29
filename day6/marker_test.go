package day6

import (
	"advent2022/file"
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "signal.txt"

func TestFindPacketMarkerPos(t *testing.T) {
	l := 4
	pos := FindMarkerPos("bvwbjplbgvbhsrlpgdmjqwftvncz", l)
	require.Equal(t, 5, pos)

	pos = FindMarkerPos("nppdvjthqldpwncqszvftbrmjlhg", l)
	require.Equal(t, 6, pos)

	pos = FindMarkerPos("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", l)
	require.Equal(t, 10, pos)

	pos = FindMarkerPos("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", l)
	require.Equal(t, 11, pos)

	str := file.ReadFile(path)
	pos = FindMarkerPos(str, l)
	require.Equal(t, 1760, pos)
}

func TestFindMessageMarkerPos(t *testing.T) {
	l := 14
	pos := FindMarkerPos("mjqjpqmgbljsphdztnvjfqwrcgsmlb", l)
	require.Equal(t, 19, pos)

	pos = FindMarkerPos("bvwbjplbgvbhsrlpgdmjqwftvncz", l)
	require.Equal(t, 23, pos)

	pos = FindMarkerPos("nppdvjthqldpwncqszvftbrmjlhg", l)
	require.Equal(t, 23, pos)

	pos = FindMarkerPos("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", l)
	require.Equal(t, 29, pos)

	pos = FindMarkerPos("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", l)
	require.Equal(t, 26, pos)

	str := file.ReadFile(path)
	pos = FindMarkerPos(str, l)
	require.Equal(t, 2974, pos)
}
