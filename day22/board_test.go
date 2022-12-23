package day22

import (
	"advent2022/file"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	boardSamplePath = "data/board_sample.txt"
	movesSamplePath = "data/moves_sample.txt"
	boardPath       = "data/board.txt"
	movesPath       = "data/moves.txt"
)

func TestFlatBoardSample(t *testing.T) {
	moves := ParseMoves(file.ReadFile(movesSamplePath))
	board := ParseFlatBoard(file.ReadFile(boardSamplePath), 12, 16)
	pos := board.StartPosition()
	pos.MoveMany(moves, board)
	require.Equal(t, 6032, pos.Score())
}

func TestBoard(t *testing.T) {
	moves := ParseMoves(file.ReadFile(movesPath))
	board := ParseFlatBoard(file.ReadFile(boardPath), 200, 150)
	pos := board.StartPosition()
	pos.MoveMany(moves, board)
	require.Equal(t, 146092, pos.Score())
}

func TestCubeBoard(t *testing.T) {
	moves := ParseMoves(file.ReadFile(movesPath))
	board := NewCubeBoard(ParseFlatBoard(file.ReadFile(boardPath), 200, 150))
	pos := board.StartPosition()
	pos.MoveMany(moves, board)
	pos = board.OriginalPosition(pos)

	//fmt.Println(board.coordinates.front)
	//fmt.Println(pos.facing)

	// account for position change due to rotation
	// we are on 2nd, upside down so east becomes west
	pos = Position{row: pos.row, column: pos.column, facing: west}
	require.Equal(t, 110342, pos.Score())
}
