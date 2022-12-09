package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSizeDirs(t *testing.T) {
	require.Equal(t, 1443806, SizeDirs("program.txt", 100000))
}

func TestFindMinDeleteDirSize(t *testing.T) {
	require.Equal(t, 942298, FindMinDeleteDirSize("program.txt", 70000000, 30000000))
}
