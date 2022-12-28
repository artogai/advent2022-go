package day3

import (
	"advent2022/file"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

func MissplacedItemsScore(filename string) int {
	score := 0
	for _, rucksack := range readRucksacks(filename) {
		score += missplacedItemScore(rucksack)
	}
	return score
}

func BadgesScore(filename string) int {
	score := 0
	for _, group := range lo.Chunk(readRucksacks(filename), 3) {
		score += badgeScore(group)
	}
	return score
}

type item rune
type rucksack struct{ left, right mapset.Set[item] }

func missplacedItemScore(r rucksack) int {
	item, exists := r.left.Intersect(r.right).Pop()
	if !exists {
		panic("missplacedItemScore: no missplaced item")
	}
	return score(item)
}

func badgeScore(group []rucksack) int {
	items := group[0].left.Union(group[0].right)
	for _, r := range group[1:] {
		items = items.Intersect(r.left.Union(r.right))
	}
	badgeItem, exists := items.Pop()
	if !exists {
		panic("badgeScore: no badge item")
	}
	return score(badgeItem)
}

func score(i item) int {
	if i >= 65 && i < 91 { // A-Z
		return int(i - 38)
	}
	if i >= 97 && i < 123 { // a-z
		return int(i - 96)
	}
	panic(fmt.Sprintln("unexpected character", i))
}

func readRucksacks(filename string) []rucksack {
	return file.ParseFile(filename, parseRucksack)
}

func parseRucksack(s string) rucksack {
	pivot := len(s) / 2
	left := mapset.NewThreadUnsafeSet([]item(s[:pivot])...)
	right := mapset.NewThreadUnsafeSet([]item(s[pivot:])...)
	return rucksack{left, right}
}
