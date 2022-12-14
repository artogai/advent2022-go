package day14

import (
	"advent2022/file"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	c := read(1000, true, file.ReadLines("cave_sample.txt")...)
	for c.dropSand(0, 500) {
	}

	//774

	fmt.Println(c.countSand())
}

func TestXxx2(t *testing.T) {
	c := read(1000, true, file.ReadLines("cave.txt")...)
	for c.dropSand(0, 500) {
	}

	//774

	fmt.Println(c.countSand())
}
