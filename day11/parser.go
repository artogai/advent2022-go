package day11

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"strings"

	"github.com/bytom/bytom/math/checked"
	"github.com/samber/lo"
)

func read(filename string) monkeys {
	monkeysLines := readLines(filename)
	ms := []*monkey{}
	for indx, mLines := range monkeysLines {
		monkey := parseMonkey(indx, mLines)
		ms = append(ms, &monkey)
	}
	return ms
}

func readLines(filename string) [][]string {
	return lo.Map(
		strings.Split(file.ReadFile(filename), "Monkey ")[1:],
		func(s string, _ int) []string {
			strs := strings.Split(strings.TrimSpace(s), "\n")
			return lo.Map(strs, func(s string, _ int) string {
				return strings.TrimSpace(s)
			})
		},
	)
}

func parseMonkey(indx int, strs []string) monkey {
	return monkey{
		indx,
		parseItems(strs[1]),
		parseOp(strs[2]),
		parseCond(strs[3], strs[4], strs[5]),
		0,
	}
}

func parseItems(s string) []int64 {
	s = strings.TrimPrefix(s, "Starting items: ")
	itemsStrs := strings.Split(s, ", ")
	items := make([]int64, len(itemsStrs))
	for i, itemStr := range itemsStrs {
		item, err := strconv.Atoi(itemStr)
		if err != nil {
			panic(err)
		}
		items[i] = int64(item)
	}
	return items
}

func parseOp(s string) func(int64) int64 {
	s = strings.TrimPrefix(s, "Operation: new = ")
	opStr := strings.Split(s, " ")
	op1, err1 := strconv.Atoi(opStr[0])
	op2, err2 := strconv.Atoi(opStr[2])
	op1_64 := int64(op1)
	op2_64 := int64(op2)

	var f func(x, y int64) int64

	switch opStr[1] {
	case "+":
		f = func(x, y int64) int64 {
			res, ok := checked.AddInt64(x, y)
			if !ok {
				panic(fmt.Sprintf("overflow %d+%d", x, y))
			}
			return res
		}
	case "*":
		f = func(x, y int64) int64 {
			res, ok := checked.MulInt64(x, y)
			if !ok {
				panic(fmt.Sprintf("overflow %d*%d", x, y))
			}
			return res
		}
	}
	return func(i int64) int64 {
		x1 := i
		x2 := i
		if err1 == nil {
			x1 = op1_64
		}
		if err2 == nil {
			x2 = op2_64
		}
		return f(x1, x2)
	}
}

func parseCond(sTest, sTrue, sFalse string) func(int64) int64 {
	iTest := extractInt("Test: divisible by ", sTest)
	iTrue := extractInt("If true: throw to monkey ", sTrue)
	iFalse := extractInt("If false: throw to monkey ", sFalse)
	return func(i int64) int64 {
		if i%iTest == 0 {
			return iTrue
		} else {
			return iFalse
		}
	}
}

func extractInt(prefix string, s string) int64 {
	s = strings.TrimPrefix(s, prefix)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int64(i)
}
