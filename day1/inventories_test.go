package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxInventoryCalories(t *testing.T) {
	got, err := MaxInventoryCalories("inventories.txt")
	require.NoError(t, err)
	require.Equal(t, 69693, got)
}

func TestMaxNInventoryCalories(t *testing.T) {
	got, err := MaxNInventoriesCalories("inventories.txt", 3)
	require.NoError(t, err)
	require.Equal(t, 200945, got)
}
