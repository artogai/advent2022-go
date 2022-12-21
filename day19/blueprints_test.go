package day19

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMaxOutput(t *testing.T) {
	bps := readBlueprints("blueprints.txt")
	maxOutCoeff := 0
	for _, bp := range bps {
		maxOut := maxOutput(bp, 24)
		maxOutCoeff += bp.id * maxOut.geode
		fmt.Println(bp.id, " done, out = ", maxOut.geode)
	}
	fmt.Println(maxOutCoeff)
}

func readBlueprints(filename string) []blueprint {
	lines := file.ReadLines(filename)
	blueprints := make([]blueprint, 0, len(lines))
	for i, line := range lines {
		blueprints = append(blueprints, parseBlueprint(i, line))
	}
	return blueprints
}

func parseBlueprint(i int, line string) blueprint {
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
		id:                 i + 1,
		oreRobotPrice:      resources{oreCost, 0, 0, 0},
		clayRobotPrice:     resources{clayCost, 0, 0, 0},
		obsidianRobotPrice: resources{obsidianCost1, obsidianCost2, 0, 0},
		geodeRobotPrice:    resources{geodeCost1, 0, geodeCost2, 0},
	}
}
