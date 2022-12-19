package day17

import (
	"advent2022/file"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	jets := []jetDir(file.ReadFile("jets.txt"))
	c := newCave(7, 5000, []jetDir(jets))
	c.drop(2022)
	//c.print()
	fmt.Println(c.height - c.top)
}
