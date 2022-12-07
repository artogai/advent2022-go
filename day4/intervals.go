package day4

import (
	"file"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func CountIntervalsFullyContains(filename string) int {
	cnt := 0
	for _, a := range readAssignments(filename) {
		if a.first.contains(a.second) || a.second.contains(a.first) {
			cnt += 1
		}
	}
	return cnt
}

type interval struct{ start, end int }
type assignment struct{ first, second interval }

func (i interval) contains(other interval) bool {
	return i.start <= other.start && i.end >= other.end
}

func readAssignments(filename string) []assignment {
	lines := file.ReadLines(filename)
	assignments := lo.Map(lines, func(s string, _ int) assignment { return parseAssignments(s) })
	return assignments
}

func parseAssignments(s string) assignment {
	x := strings.Split(s, ",")
	return assignment{parseInterval(x[0]), parseInterval(x[1])}
}

func parseInterval(s string) interval {
	x := strings.Split(s, "-")
	start, _ := strconv.Atoi(x[0])
	end, _ := strconv.Atoi(x[1])
	return interval{start, end}
}
