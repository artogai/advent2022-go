package day16

import (
	"advent2022/file"
	"math"
	"strconv"
	"strings"
)

type valve int
type edges map[valve][]valve
type rates map[valve]int

func countMaxRate(eds edges, rts rates) int {
	maxFlow := 0
	valves := nonEmptyValves(rts)
	opened := make([]bool, len(valves))
	paths := make(map[valve]map[valve]int)

	var rec func(valve, int, int, []bool)
	rec = func(v valve, flow int, time int, opened []bool) {
		for i, nextV := range valves {
			if !opened[i] {
				if _, exists := paths[v]; !exists {
					paths[v] = make(map[valve]int)
				}
				if _, exists := paths[v][nextV]; !exists {
					paths[v][nextV] = shortestPath(v, nextV, &eds)
				}
				nextTime := time - paths[v][nextV] - 1
				if nextTime >= 0 {
					nextFlow := flow + nextTime*rts[nextV]
					nextOpened := make([]bool, len(opened))
					copy(nextOpened, opened)
					nextOpened[i] = true
					rec(nextV, nextFlow, nextTime, nextOpened)
				}
			}
		}
		if flow > maxFlow {
			maxFlow = flow
		}
	}

	rec(0, 0, 30, opened)
	return maxFlow
}

func nonEmptyValves(rts rates) []valve {
	vs := []valve{}
	for v, rate := range rts {
		if rate != 0 {
			vs = append(vs, v)
		}
	}
	return vs
}

func shortestPath(v1 valve, v2 valve, eds *edges) int {
	d := make(map[valve]int)
	for v := range *eds {
		d[v] = math.MaxInt / 2
	}
	d[v1] = 0
	for i := 1; i < len(*eds); i++ {
		for v, neighbours := range *eds {
			for _, n := range neighbours {
				if d[n] > d[v]+1 {
					d[n] = d[v] + 1
				}
			}
		}
	}
	return d[v2]
}

func parseValves(filename string) (edges, rates) {
	eds := make(edges)
	rts := make(rates)
	for _, line := range file.ReadLines(filename) {
		lineArr := strings.Split(line, "; ")
		num, rate := parseValve(lineArr[0])
		neighbours := parseNeighbours(lineArr[1])
		eds[num] = neighbours
		rts[num] = rate
	}
	return eds, rts
}

func parseValve(s string) (num valve, rate int) {
	num = parseValveNum(s[6:8])
	rate, err := strconv.Atoi(s[23:])
	if err != nil {
		panic(err)
	}
	return num, rate
}

func parseNeighbours(s string) []valve {
	s = strings.TrimPrefix(s, "tunnels lead to valves ")
	s = strings.TrimPrefix(s, "tunnel leads to valve ")
	names := strings.Split(s, ", ")
	res := make([]valve, len(names))
	for i, name := range names {
		res[i] = parseValveNum(name)
	}
	return res
}

func parseValveNum(s string) valve {
	i0 := int(s[0] - 'A')
	i1 := int(s[1]-'A') * 26
	return valve(i0 + i1)
}
