package day3

import (
	"advent2022/file"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

func CalcMissplacedItemsScore(filename string) int {
	score := 0
	for _, rucksack := range readRucksacks(filename) {
		score += calcMissplacedItemScore(rucksack)
	}
	return score
}

func CalcBadgesScore(filename string) int {
	groups := lo.Chunk(readRucksacks(filename), 3)
	score := lo.Reduce(
		groups,
		func(agg int, group []rucksack, _ int) int { return agg + calcGroupBadgeScore(group) },
		0,
	)
	return score
}

type rucksack struct{ left, right string }

func calcMissplacedItemScore(r rucksack) int {
	s1 := toSet(r.left)
	s2 := toSet(r.right)
	missplacedItem, _ := s1.Intersect(s2).Pop()
	return score(missplacedItem)
}

func calcGroupBadgeScore(group []rucksack) int {
	items := lo.Map(group, func(r rucksack, _ int) mapset.Set[rune] {
		return toSet(r.left).Union(toSet(r.right))
	})
	badgeItem, _ := lo.Reduce(
		items[1:],
		func(agg mapset.Set[rune], r mapset.Set[rune], _ int) mapset.Set[rune] {
			return agg.Intersect(r)
		},
		items[0],
	).Pop()
	return score(badgeItem)
}

func readRucksacks(filename string) []rucksack {
	lines := file.ReadLines(filename)
	rucksacks := lo.Map(lines, func(s string, _ int) rucksack { return parseRucksack(s) })
	return rucksacks
}

func parseRucksack(s string) rucksack {
	pivot := len(s) / 2
	return rucksack{s[:pivot], s[pivot:]}
}

func score(r rune) int {
	// A-Z
	if r >= 65 && r < 91 {
		return int(r - 38)
	} else if r >= 97 && r < 123 { // a-z
		return int(r - 96)
	} else {
		panic(fmt.Errorf("unexpected character %d", r))
	}
}

func toSet(s string) mapset.Set[rune] {
	set := mapset.NewThreadUnsafeSet[rune]()
	for _, r := range s {
		set.Add(r)
	}
	return set
}
