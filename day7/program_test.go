package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "program.txt"

func TestSizeDirs(t *testing.T) {
	require.Equal(t, 1443806, SizeDirs(path, 100000))
}

func TestFindMinDeleteDirSize(t *testing.T) {
	require.Equal(t, 942298, FindMinDeleteDirSize(path, 70000000, 30000000))
}
