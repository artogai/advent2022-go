package day1

import (
	"errors"
	"file"
	"log"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func MaxInventoryCalories(filename string) (int, error) {
	return MaxNInventoriesCalories(filename, 1)
}

func MaxNInventoriesCalories(filename string, n int) (int, error) {
	return maxCaloriesTopN(readInventories(filename), n)
}

func maxCaloriesTopN(inventories []int, n int) (int, error) {
	if len(inventories) < n {
		return -1, errors.New("inventories count is less than n")
	}
	sort.Slice(inventories, func(i, j int) bool {
		return inventories[i] > inventories[j]
	})
	return lo.Sum(inventories[0:n]), nil
}

func readInventories(filename string) []int {
	lines := file.ReadLines(filename)
	inventories := []int{}
	inventoryCalories := 0
	for _, line := range lines {
		if line == "" {
			inventories = append(inventories, inventoryCalories)
			inventoryCalories = 0
		} else {
			itemCalories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			inventoryCalories += itemCalories
		}
	}
	inventories = append(inventories, inventoryCalories)
	return inventories
}
