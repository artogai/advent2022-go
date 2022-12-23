package day22

import "fmt"

type Position struct {
	row, column int
	facing      direction
}

type direction int

const (
	north direction = iota
	east  direction = iota
	south direction = iota
	west  direction = iota
)

func (p *Position) MoveMany(ms []move, b board) {
	for _, m := range ms {
		p.move(m, b)
	}
}

func (p *Position) Score() int {
	return (p.row+1)*1000 + (p.column+1)*4 + (int(p.facing)+3)%4
}

func (p *Position) move(m move, b board) {
	switch m := m.(type) {
	case rotate:
		p.rotate(m)
	case walk:
		p.walk(int(m), b)
	default:
		panic(fmt.Sprint("move: unknown move ", m))
	}
}

func (p *Position) rotate(r rotate) {
	switch r {
	case L:
		p.facing = (p.facing + 3) % 4
	case R:
		p.facing = (p.facing + 1) % 4
	}
}

func (p *Position) walk(distance int, b board) {
	for i := 0; i < distance; i++ {
		b.step(p)
	}
}
