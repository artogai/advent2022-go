package day21

import (
	"advent2022/file"
	"strconv"
	"strings"
)

type operator int

const (
	plus   operator = iota
	minus           = iota
	times           = iota
	divide          = iota
)

type id string

type expr interface{ isExpr() bool }
type leaf int
type node struct {
	left, right id
	op          operator
}

type exprs map[id]expr

func NewExprs(filename string) exprs {
	lines := file.ReadLines(filename)
	es := make(exprs)
	for _, l := range lines {
		id, expr := parseExpr(l)
		es[id] = expr
	}
	return es
}

func parseExpr(s string) (id, expr) {
	arr := strings.Split(s, ": ")
	id0 := id(arr[0])
	num, err := strconv.Atoi(arr[1])
	if err == nil {
		return id0, leaf(num)
	}
	nodeStr := arr[1]
	id1 := nodeStr[:4]
	id2 := nodeStr[7:]
	opStr := nodeStr[5:6]
	var op operator
	switch opStr {
	case "+":
		op = plus
	case "-":
		op = minus
	case "*":
		op = times
	case "/":
		op = divide
	default:
		panic("unknown operator")
	}
	return id0, node{id(id1), id(id2), op}
}

func (es exprs) Eval(id id) int {
	switch expr := es[id].(type) {
	case leaf:
		return int(expr)
	case node:
		l := es.Eval(expr.left)
		r := es.Eval(expr.right)
		switch expr.op {
		case plus:
			return l + r
		case minus:
			return l - r
		case times:
			return l * r
		case divide:
			return l / r
		}
	}
	panic("unknown expr type")
}

func (l leaf) isExpr() bool { return true }
func (n node) isExpr() bool { return true }
