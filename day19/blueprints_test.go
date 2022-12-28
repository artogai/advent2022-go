package day19

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func IgnoreMaxOutput24(t *testing.T) {
	time := 24
	cutoff := 0
	bps := readBlueprints("blueprints.txt")
	outs := make([]chan int, len(bps))
	for i, bp := range bps {
		outs[i] = make(chan int)
		go maxOutput(bp, time, cutoff, outs[i])
	}

	maxOutCoeff := 0
	for i, out := range outs {
		maxOut := <-out
		fmt.Println(i, " out = ", maxOut)
		maxOutCoeff += (i + 1) * maxOut
	}

	fmt.Println(maxOutCoeff)
}

func IgnoreMaxOutput32(t *testing.T) {
	time := 32
	cutoff := 6
	bps := readBlueprints("blueprints.txt")[:3]
	outs := make([]chan int, len(bps))
	for i, bp := range bps {
		outs[i] = make(chan int)
		go maxOutput(bp, time, cutoff, outs[i])
	}

	product := 1
	for i, out := range outs {
		maxOut := <-out
		fmt.Println(i, " out = ", maxOut)
		product *= maxOut
	}

	fmt.Println(product)
}

func readBlueprints(filename string) []blueprint {
	lines := file.ReadFileLines(filename)
	blueprints := make([]blueprint, 0, len(lines))
	for _, line := range lines {
		blueprints = append(blueprints, parseBlueprint(line))
	}
	return blueprints
}

func parseBlueprint(line string) blueprint {
	arr := strings.Split(line, ": ")
	arr = strings.Split(arr[1], ". ")
	s := strings.TrimPrefix(arr[0], "Each ore robot costs ")
	s = strings.TrimSuffix(s, " ore")
	oreCost, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid ore cost")
	}
	s = strings.TrimPrefix(arr[1], "Each clay robot costs ")
	s = strings.TrimSuffix(s, " ore")
	clayCost, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid clay cost")
	}
	s = strings.TrimPrefix(arr[2], "Each obsidian robot costs ")
	s = strings.TrimSuffix(s, " clay")
	arr2 := strings.Split(s, " ore and ")
	obsidianCost1, err := strconv.Atoi(arr2[0])
	if err != nil {
		panic("invalid obsidian 1 cost")
	}
	obsidianCost2, err := strconv.Atoi(arr2[1])
	if err != nil {
		panic("invalid obsidian 2 cost")
	}
	s = strings.TrimPrefix(arr[3], "Each geode robot costs ")
	s = strings.TrimSuffix(s, " obsidian.")
	arr2 = strings.Split(s, " ore and ")
	geodeCost1, err := strconv.Atoi(arr2[0])
	if err != nil {
		panic("invalid geode 1 cost")
	}
	geodeCost2, err := strconv.Atoi(arr2[1])
	if err != nil {
		panic("invalid geode 2 cost")
	}
	return blueprint{
		oreRobotPrice:      resources{oreCost, 0, 0, 0},
		clayRobotPrice:     resources{clayCost, 0, 0, 0},
		obsidianRobotPrice: resources{obsidianCost1, obsidianCost2, 0, 0},
		geodeRobotPrice:    resources{geodeCost1, 0, geodeCost2, 0},
	}
}
