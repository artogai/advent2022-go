package day24

import (
	"fmt"
	"testing"
)

func TestShortestPathSample(t *testing.T) {
	c := parse("cave_sample.txt")
	caves := genCaves(c)
	start := coordinate{0, 1}
	end := coordinate{5, 6}

	p1 := shortestPath(start, end, caves[0])
	p2 := shortestPath(end, start, caves[p1%len(caves)])
	p3 := shortestPath(start, end, caves[(p1+p2)%len(caves)])
	fmt.Println(p1, p2, p3)
	fmt.Println(p1 + p2 + p3)
}

func TestShortestPath(t *testing.T) {
	c := parse("cave.txt")
	caves := genCaves(c)
	start := coordinate{0, 1}
	end := coordinate{26, 120}

	p1 := shortestPath(start, end, caves[0])
	p2 := shortestPath(end, start, caves[p1%len(caves)])
	p3 := shortestPath(start, end, caves[(p1+p2)%len(caves)])
	fmt.Println(p1, p2, p3)
	fmt.Println(p1 + p2 + p3)
}
