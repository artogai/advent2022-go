package day21

import (
	"advent2022/file"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type id string

type operator int

const (
	plus   operator = iota
	minus           = iota
	times           = iota
	divide          = iota
	equals          = iota
)

type exprs map[id]expr
type expr interface{ isExpr() bool }

type leaf big.Rat
type unknown struct{ v big.Rat }
type node struct {
	left, right id
	op          operator
}

type eval interface {
	ToString() string
	evalLeaf(leaf, operator) eval
	evalUnknown(unknown, operator) eval
	evalNode(evalNode, operator) eval
	isEval() bool
}
type evalNode struct {
	left, right eval
	op          operator
}

func NewExprs(filename string) exprs {
	lines := file.ReadLines(filename)
	es := make(exprs)
	for _, l := range lines {
		id, expr := parseExpr(l)
		es[id] = expr
	}
	return es
}

func (es exprs) Eval(id id) eval {
	switch expr := es[id].(type) {
	case leaf:
		return expr
	case unknown:
		return expr
	case node:
		l := es.Eval(expr.left)
		r := es.Eval(expr.right)
		// fmt.Println(
		// 	id,
		// 	" = ",
		// 	expr.left,
		// 	expr.op.ToString(),
		// 	expr.right, " = ",
		// 	l.ToString(),
		// 	expr.op.ToString(),
		// 	r.ToString())
		switch r := r.(type) {
		case leaf:
			return l.evalLeaf(r, expr.op)
		case unknown:
			return l.evalUnknown(r, expr.op)
		case evalNode:
			return l.evalNode(r, expr.op)
		}
	}
	panic("unsupported expr type")
}

func (l leaf) evalLeaf(r leaf, op operator) eval {
	li := big.Rat(l)
	ri := big.Rat(r)
	res := big.Rat{}
	switch op {
	case plus:
		return leaf(*res.Add(&li, &ri))
	case minus:
		return leaf(*res.Sub(&li, &ri))
	case times:
		return leaf(*res.Mul(&li, &ri))
	case divide:
		return leaf(*res.Quo(&li, &ri))
	case equals:
		return evalNode{l, r, op}
	}
	panic("unsupported operator")
}

func (l leaf) evalUnknown(r unknown, op operator) eval {
	li := big.Rat(l)
	res := big.Rat{}

	switch op {
	case plus, minus:
		return evalNode{l, r, op}
	case times:
		return newUnknown(*res.Mul(&li, &r.v))
	}
	panic("unsupported operator")
}

func (l unknown) evalLeaf(r leaf, op operator) eval {
	ri := big.Rat(r)
	res := big.Rat{}
	switch op {
	case plus, minus:
		return evalNode{l, r, op}
	case times:
		return newUnknown(*res.Mul(&l.v, &ri))
	case divide:
		return newUnknown(*res.Quo(&l.v, &ri))
	}
	panic("unsupported operator")
}

func (l leaf) evalNode(r evalNode, op operator) eval {
	switch op {
	case plus:
		switch r.op {
		case plus:
			// l + (rl + rr)

			// (l + rl) + rr
			rl, ok1 := r.left.(leaf)
			_, ok2 := r.right.(unknown)

			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rl, plus), r.right, plus}
			}

			// rl + (l + rr)
			_, ok1 = r.left.(unknown)
			rr, ok2 := r.right.(leaf)
			if ok1 && ok2 {
				return evalNode{r.left, l.evalLeaf(rr, plus), plus}
			}
		case minus:
			// l + (rl - rr)

			// (l + rl) - rr
			rl, ok1 := r.left.(leaf)
			_, ok2 := r.right.(unknown)

			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rl, plus), r.right, minus}
			}

			// (l - rr) + rl
			_, ok1 = r.left.(unknown)
			rr, ok2 := r.right.(leaf)
			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rr, minus), r.left, plus}
			}
		}
	case minus:
		switch r.op {
		case plus:
			// l - (rl + rr)

			// (l - rl) - rr
			rl, ok1 := r.left.(leaf)
			_, ok2 := r.right.(unknown)

			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rl, minus), r.right, minus}
			}

			// (l - rr) - rl
			_, ok1 = r.left.(unknown)
			rr, ok2 := r.right.(leaf)

			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rr, minus), r.left, minus}
			}
		case minus:
			// l - (rl - rr)

			// (l - rl) + rr
			rl, ok1 := r.left.(leaf)
			_, ok2 := r.right.(unknown)

			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rl, minus), r.right, plus}
			}

			// (l + rr) - rl
			_, ok1 = r.left.(unknown)
			rr, ok2 := r.right.(leaf)
			if ok1 && ok2 {
				return evalNode{l.evalLeaf(rr, plus), r.left, minus}
			}
		}
	case times:
		switch r.op {
		case plus, minus:
			return evalNode{r.left.evalLeaf(l, times), r.right.evalLeaf(l, times), r.op}
		}
	}
	panic("unsupported operator")
}

func (l evalNode) evalLeaf(r leaf, op operator) eval {
	switch op {
	case plus:
		switch l.op {
		case plus:
			// (ll + lr) + r

			// (ll + r) + lr
			ll, ok1 := l.left.(leaf)
			_, ok2 := l.right.(unknown)
			if ok1 && ok2 {
				return evalNode{ll.evalLeaf(r, plus), l.right, plus}
			}

			// ll + (lr + r)
			_, ok1 = l.left.(unknown)
			lr, ok2 := l.right.(leaf)
			if ok1 && ok2 {
				return evalNode{l.left, lr.evalLeaf(r, plus), plus}
			}
		case minus:
			// (ll - lr) + r

			// (ll + r) - lr
			ll, ok1 := l.left.(leaf)
			_, ok2 := l.right.(unknown)
			if ok1 && ok2 {
				return evalNode{ll.evalLeaf(r, plus), l.right, minus}
			}

			// ll + (r - lr)
			_, ok1 = l.left.(unknown)
			lr, ok2 := l.right.(leaf)
			if ok1 && ok2 {
				return evalNode{l.left, r.evalLeaf(lr, minus), plus}
			}
		}
	case minus:
		switch l.op {
		case plus:
			// (ll + lr) - r

			// (ll - r) + lr
			ll, ok1 := l.left.(leaf)
			_, ok2 := l.right.(unknown)
			if ok1 && ok2 {
				return evalNode{ll.evalLeaf(r, minus), l.right, plus}
			}

			// ll + (lr - r)
			_, ok1 = l.left.(unknown)
			lr, ok2 := l.right.(leaf)
			if ok1 && ok2 {
				return evalNode{l.left, lr.evalLeaf(r, minus), plus}
			}
		case minus:
			// (ll - lr) - r

			// (ll - r) - lr
			ll, ok1 := l.left.(leaf)
			_, ok2 := l.right.(unknown)
			if ok1 && ok2 {
				return evalNode{ll.evalLeaf(r, minus), l.right, minus}
			}

			// ll - (lr + r)
			_, ok1 = l.left.(unknown)
			lr, ok2 := l.right.(leaf)
			if ok1 && ok2 {
				return evalNode{l.left, lr.evalLeaf(r, plus), minus}
			}
		}
	case times, divide:
		switch l.op {
		case plus, minus:
			return evalNode{l.left.evalLeaf(r, op), l.right.evalLeaf(r, op), l.op}
		}
	case equals:
		return evalNode{l, r, op}
	}

	panic("unsupported operator")
}

func (l unknown) evalUnknown(r unknown, op operator) eval {
	return evalNode{l, r, op}
}

func (l unknown) evalNode(r evalNode, op operator) eval {
	return evalNode{l, r, op}
}

func (l evalNode) evalUnknown(r unknown, op operator) eval {
	return evalNode{l, r, op}
}

func (l evalNode) evalNode(r evalNode, op operator) eval {
	return evalNode{l, r, op}
}

func newUnknown(v big.Rat) unknown {
	return unknown{v}
}

func (en evalNode) ToString() string {
	return fmt.Sprintf("(%s %s %s)", en.left.ToString(), en.op.ToString(), en.right.ToString())
}

func (l leaf) ToString() string {
	li := big.Rat(l)
	return fmt.Sprintf("%v", li.RatString())
}

func (un unknown) ToString() string {
	return fmt.Sprint("x ", un.v.RatString())
}

func (op operator) ToString() string {
	switch op {
	case plus:
		return "+"
	case minus:
		return "-"
	case times:
		return "*"
	case divide:
		return "/"
	case equals:
		return "="
	}
	panic("unknown operator")
}

func (l leaf) isExpr() bool     { return true }
func (n node) isExpr() bool     { return true }
func (u unknown) isExpr() bool  { return true }
func (l leaf) isEval() bool     { return true }
func (n evalNode) isEval() bool { return true }
func (u unknown) isEval() bool  { return true }

func parseExpr(s string) (id, expr) {
	arr := strings.Split(s, ": ")
	id0 := id(arr[0])
	num, err := strconv.Atoi(arr[1])
	if err == nil {
		return id0, leaf(*big.NewRat(int64(num), 1))
	}
	nodeStr := arr[1]
	if nodeStr == "x" {
		return id0, unknown{*big.NewRat(1, 1)}
	}
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
	case "=":
		op = equals
	default:
		panic("unknown operator")
	}
	return id0, node{id(id1), id(id2), op}
}
