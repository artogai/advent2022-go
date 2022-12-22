package day22

import (
	"advent2022/file"
	"fmt"
	"testing"
)

func TestBoardSample(t *testing.T) {
	moves := ParseMoves(file.ReadFile("moves_sample.txt"))
	board := ParseBoard(file.ReadLines("board_sample.txt"), 12, 16)
	pos := Position{row: 0, column: 8, facing: east}
	pos.MoveMany(moves, board)
	fmt.Println(pos)
	fmt.Println(score(pos))
}

func TestBoard(t *testing.T) {
	moves := ParseMoves(file.ReadFile("moves.txt"))
	board := ParseBoard(file.ReadLines("board.txt"), 200, 150)
	pos := Position{row: 0, column: 50, facing: east}
	pos.MoveMany(moves, board)
	fmt.Println(pos)
	fmt.Println(score(pos))
}

func score(pos Position) int {
	return (pos.row+1)*1000 + (pos.column+1)*4 + (int(pos.facing)+3)%4
}
