package day13

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"testing"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type signal interface{}
type signalValue int
type signalSeq []signal

func parseSignalValue(buff []byte) signalValue {
	i, err := strconv.Atoi(string(buff))
	if err != nil {
		panic(err)
	}
	return signalValue(i)
}

func parseSignalSeq(str string) signalSeq {
	buff := []byte{}
	emitBuffer := func(res *[]signal) {
		if len(buff) > 0 {
			*res = append(*res, parseSignalValue(buff))
			buff = []byte{}
		}
	}

	var rec func(int, []signal) (signalSeq, int)
	rec = func(i int, res []signal) (signalSeq, int) {
		if i >= len(str) {
			return res, i
		}

		head := str[i]
		switch {
		case head >= '0' && head <= '9':
			buff = append(buff, head)
			return rec(i+1, res)
		case head == '[':
			emitBuffer(&res)
			nextSeq, nextI := rec(i+1, []signal{})
			return rec(nextI+1, append(res, nextSeq))
		case head == ',':
			emitBuffer(&res)
			return rec(i+1, res)
		case head == ']':
			emitBuffer(&res)
			return res, i
		default:
			panic("unsupported char")
		}
	}

	res, _ := rec(0, []signal{})
	return res
}

func areValid(s1, s2 signal) (bool, bool) {
	switch s1 := s1.(type) {
	case signalValue:
		switch s2 := s2.(type) {
		case signalValue:
			if s1 == s2 {
				return true, false
			} else if s1 < s2 {
				return true, true
			} else {
				return false, true
			}
		case signalSeq:
			return areValid(signalSeq{s1}, s2)
		}
	case signalSeq:
		switch s2 := s2.(type) {
		case signalValue:
			//todo: find out why []signal{s2} not working
			return areValid(s1, signalSeq{s2})
		case signalSeq:
			if len(s1) == 0 && len(s2) == 0 {
				return true, false
			} else if len(s1) > 0 && len(s2) == 0 {
				return false, true
			} else if len(s1) == 0 && len(s2) > 0 {
				return true, true
			} else {
				valid, shortCircuit := areValid(s1[0], s2[0])
				if shortCircuit {
					return valid, shortCircuit
				} else {
					if valid {
						return areValid(s1[1:], s2[1:])
					} else {
						return false, true
					}
				}
			}
		}
	}
	panic("unreachable")
}

func TestSignals(t *testing.T) {
	signals := lo.Chunk(file.ReadFileLines("signals.txt"), 3)

	valid := []int{}
	for i, s := range signals {
		s1 := parseSignalSeq(s[0])[0]
		s2 := parseSignalSeq(s[1])[0]
		v, _ := areValid(s1, s2)
		if v {
			valid = append(valid, i+1)
		}

	}
	fmt.Println(lo.Sum(valid))
}

func TestSignals2(t *testing.T) {
	signalsLines := lo.Chunk(file.ReadFileLines("signals.txt"), 3)

	signals := []signal{}

	signals = append(signals, parseSignalSeq("[[2]]")[0])
	signals = append(signals, parseSignalSeq("[[6]]")[0])

	for _, s := range signalsLines {
		s1 := parseSignalSeq(s[0])[0]
		s2 := parseSignalSeq(s[1])[0]

		signals = append(signals, s1)
		signals = append(signals, s2)
	}

	slices.SortFunc(signals, func(a signal, b signal) bool {
		v, _ := areValid(a, b)
		return v
	})

	// for i, s := range signals {
	// 	fmt.Println(i, s)
	// }
	for i, s := range signals {
		if st, ok := s.(signalSeq); ok && len(st) == 1 {
			if s0, ok := st[0].(signalSeq); ok && len(s0) == 1 {
				if s1, ok := s0[0].(signalValue); ok {
					fmt.Println(i, s1)
				}
			}
		}
	}
}
