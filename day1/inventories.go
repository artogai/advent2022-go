package day1

import (
	"advent2022/file"
	"errors"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func TopOne(filename string) (int, error) {
	return SumTop(filename, 1)
}

func SumTop(filename string, n int) (int, error) {
	return sumTop(readInventories(filename), n)
}

func sumTop(arr []int, n int) (int, error) {
	if len(arr) < n {
		return -1, errors.New("arr count is less than n")
	}
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	sort.Slice(arrCopy, func(i, j int) bool {
		return arrCopy[i] > arrCopy[j]
	})
	return lo.Sum(arrCopy[:n]), nil
}

func readInventories(filename string) []int {
	lines := file.ReadFileLines(filename)
	inventories := []int{}
	inventoryCalories := 0
	for _, line := range lines {
		if line == "" {
			inventories = append(inventories, inventoryCalories)
			inventoryCalories = 0
		} else {
			itemCalories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			inventoryCalories += itemCalories
		}
	}
	inventories = append(inventories, inventoryCalories)
	return inventories
}
