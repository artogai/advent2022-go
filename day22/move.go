package day22

import (
	"fmt"
	"strconv"
)

type move interface{ isMove() bool }

type rotate int
type walk int

const (
	L rotate = iota
	R rotate = iota
)

func ParseMoves(s string) []move {
	ms := []move{}
	buff := []rune{}

	for _, c := range s {
		switch c {
		case 'L':
			w := walk(emitBuff(&buff))
			ms = append(ms, w, L)
		case 'R':
			w := walk(emitBuff(&buff))
			ms = append(ms, w, R)
		default:
			buff = append(buff, c)
		}
	}
	w := walk(emitBuff(&buff))
	ms = append(ms, w)
	return ms
}

func emitBuff(buff *[]rune) int {
	s := string(*buff)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprint("parseBuff: ", err))
	}
	*buff = []rune{}
	return i
}

func (r rotate) isMove() bool { return true }
func (w walk) isMove() bool   { return true }
