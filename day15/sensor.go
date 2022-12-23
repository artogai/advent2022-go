package day15

import (
	"advent2022/file"
	"advent2022/imath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type coord struct{ x, y int }

type sensor struct {
	coord
	beacon coord
	radius int
}

func CountScanned(y int, filename string) int {
	sensors := read(filename)
	beacons := beacons(sensors)
	scanned := scannedSegments(y, sensors)
	cnt := 0
	for _, r := range scanned {
		cnt += r.y - r.x + 1
	}
	for _, b := range beacons {
		if b.y == y && contains(scanned, b.x) {
			cnt--
		}
	}

	return cnt
}

func FindBeacon(maxC int, filename string) coord {
	sensors := read(filename)
	for y := 0; y < maxC; y++ {
		scanned := scannedSegments(y, sensors)
		notScanned := invert(scanned, maxC)
		if len(notScanned) != 0 {
			return coord{notScanned[0].x, y}
		}
	}
	panic("no beacon found")
}

func invert(scanned []coord, maxX int) []coord {
	inverted := []coord{}
	for i := 0; i < len(scanned); i++ {
		if i == 0 {
			inverted = append(inverted, coord{0, scanned[i].x - 1})
		} else {
			inverted = append(inverted, coord{scanned[i-1].y + 1, scanned[i].x - 1})
		}
	}
	inverted = append(inverted, coord{scanned[len(scanned)-1].y + 1, maxX})

	filtered := []coord{}
	for _, s := range inverted {
		if s.y-s.x >= 0 {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func contains(segments []coord, x int) bool {
	for _, s := range segments {
		if x >= s.x && x <= s.y {
			return true
		}
	}
	return false
}

func beacons(sensors []sensor) []coord {
	beacons := map[coord]bool{}
	for _, s := range sensors {
		beacons[s.beacon] = true
	}
	keys := make([]coord, 0, len(beacons))
	for c := range beacons {
		keys = append(keys, c)
	}
	return keys
}

func scannedSegments(y int, sensors []sensor) []coord {
	scanned := []coord{}
	for _, s := range sensors {
		if y >= s.y-s.radius && y <= s.y+s.radius {
			xradius := s.radius - imath.Abs(y-s.y)
			scanned = append(scanned, coord{s.x - xradius, s.x + xradius})
		}
	}
	slices.SortFunc(scanned, func(i, j coord) bool { return i.x < j.x })
	return joinSegments(scanned)
}

func joinSegments(sortedSegments []coord) []coord {
	joined := []coord{}
	for _, s := range sortedSegments {
		if len(joined) == 0 {
			joined = append(joined, s)
		} else {
			last := joined[len(joined)-1]
			if s.x > last.y {
				joined = append(joined, s)
			} else {
				last.y = imath.Max(last.y, s.y)
				joined[len(joined)-1] = last
			}
		}
	}
	return joined
}

func (c *coord) distance(c0 coord) int {
	return imath.Abs(c.x-c0.x) + imath.Abs(c.y-c0.y)
}

func read(filename string) []sensor {
	sensors := []sensor{}

	for _, row := range file.ReadLines(filename) {
		arr := strings.Split(row, ":")
		self := extract(arr[0])
		beacon := extract(arr[1])
		sensors = append(sensors, sensor{self, beacon, self.distance(beacon)})
	}

	return sensors
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
