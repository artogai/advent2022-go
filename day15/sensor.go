package day15

import (
	"advent2022/file"
	m "advent2022/math"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type coord struct{ x, y int }

type sensor struct {
	coord
	beacon coord
	radius int
}

func CountScanned(y int, filename string) int {
	sensors, minC, maxC, maxDist := read(filename)
	cnt := 0
	for x := minC.x - maxDist; x <= maxC.x+maxDist; x++ {
		if isScanned(x, y, sensors) {
			cnt++
		}
	}
	return cnt
}

func isScanned(x, y int, sensors []sensor) bool {
	c := coord{x, y}
	for _, s := range sensors {
		if c == s.beacon {
			return false
		}
	}
	for _, s := range sensors {
		if c.distance(s.coord) <= s.radius {
			return true
		}
	}
	return false
}

func (c *coord) distance(c0 coord) int {
	return m.Abs(c.x-c0.x) + m.Abs(c.y-c0.y)
}

func read(filename string) (sensors []sensor, minC, maxC coord, maxDist int) {
	sensors = []sensor{}
	minC = coord{math.MaxInt, math.MaxInt}
	maxC = coord{math.MinInt, math.MinInt}
	maxDist = 0

	for _, row := range file.ReadLines(filename) {
		arr := strings.Split(row, ":")
		self := extract(arr[0])
		beacon := extract(arr[1])
		dist := self.distance(beacon)

		sensors = append(sensors, sensor{self, beacon, dist})
		minC.x = m.Min(minC.x, self.x, beacon.x)
		maxC.x = m.Max(maxC.x, self.x, beacon.x)
		minC.y = m.Min(minC.y, self.y, beacon.y)
		maxC.y = m.Max(maxC.y, self.y, beacon.y)
		if dist > maxDist {
			maxDist = dist
		}
	}

	return
}

var (
	xReg = regexp.MustCompile(`x=(-?\d+)`)
	yReg = regexp.MustCompile(`y=(-?\d+)`)
)

func extract(s string) coord {
	xGrp := xReg.FindStringSubmatch(s)
	yGrp := yReg.FindStringSubmatch(s)
	x, err := strconv.Atoi(xGrp[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yGrp[1])
	if err != nil {
		panic(err)
	}
	return coord{x, y}
}
