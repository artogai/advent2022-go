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
	getPath := shortestPathMemoized(eds)

	var rec func(valve, int, int, []bool)
	rec = func(v valve, flow int, time int, opened []bool) {
		for i, nextV := range valves {
			if !opened[i] {
				nextTime := time - getPath(v, nextV) - 1
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

	rec(0, 0, 30, make([]bool, len(valves)))
	return maxFlow
}

func countMaxRate2(eds edges, rts rates) int {
	maxFlow := 0
	valves := nonEmptyValves(rts)
	getPath := shortestPathMemoized(eds)

	type pathEnd struct {
		v1    valve
		v2    valve
		time1 int
		time2 int
	}
	bestPaths := make(map[pathEnd]int)

	var rec func(pathEnd, int, []bool)
	rec = func(st pathEnd, flow int, opened []bool) {
		bestFlow, exists := bestPaths[st]
		if exists && bestFlow >= flow {
			return
		} else {
			bestPaths[st] = flow
			bestPaths[pathEnd{st.v2, st.v1, st.time2, st.time1}] = flow
			for i, nextV1 := range valves {
				if !opened[i] {
					for j, nextV2 := range valves {
						if !opened[j] && i != j {
							if exists && bestFlow > flow {
								continue
							} else {
								nextTime1 := st.time1 - getPath(st.v1, nextV1) - 1
								nextTime2 := st.time2 - getPath(st.v2, nextV2) - 1
								if nextTime1 >= 0 || nextTime2 >= 0 {
									nextFlow := flow
									nextOpened := make([]bool, len(opened))
									copy(nextOpened, opened)

									if nextTime1 >= 0 {
										nextFlow += nextTime1 * rts[nextV1]
										nextOpened[i] = true
									}

									if nextTime2 >= 0 {
										nextFlow += nextTime2 * rts[nextV2]
										nextOpened[j] = true
									}

									if nextTime1 >= 0 && nextTime2 >= 0 {
										rec(pathEnd{nextV1, nextV2, nextTime1, nextTime2}, nextFlow, nextOpened)
									} else if nextTime1 >= 0 {
										rec(pathEnd{nextV1, st.v2, nextTime1, st.time2}, nextFlow, nextOpened)
									} else if nextTime2 >= 0 {
										rec(pathEnd{st.v1, nextV2, st.time1, nextTime2}, nextFlow, nextOpened)
									}
								}
							}
						}
					}
				}
			}
		}
		if flow > maxFlow {
			maxFlow = flow
		}
	}

	rec(pathEnd{0, 0, 26, 26}, 0, make([]bool, len(valves)))
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

func shortestPathMemoized(eds edges) func(valve, valve) int {
	paths := make(map[valve]map[valve]int)

	return func(v1 valve, v2 valve) int {
		if _, exists := paths[v1]; !exists {
			paths[v1] = make(map[valve]int)
		}
		if _, exists := paths[v2]; !exists {
			paths[v2] = make(map[valve]int)
		}
		if _, exists := paths[v1][v2]; !exists {
			paths[v1][v2] = shortestPath(v1, v2, eds)
			paths[v2][v1] = paths[v1][v2]
		}
		if res, ok := paths[v1][v2]; ok {
			return res
		} else {
			panic("Path not found")
		}
	}
}

func shortestPath(v1 valve, v2 valve, eds edges) int {
	d := make(map[valve]int)
	for v := range eds {
		d[v] = math.MaxInt / 2
	}
	d[v1] = 0
	for i := 1; i < len(eds); i++ {
		for v, neighbours := range eds {
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
	for _, line := range file.ReadFileLines(filename) {
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
