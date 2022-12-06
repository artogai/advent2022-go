package day1

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func CountCalories(filename string) int {
	return CountCaloriesTopN(filename, 1)
}

func CountCaloriesTopN(filename string, n int) int {
	return maxCaloriesTopN(readInventories(filename), n)
}

func maxCaloriesTopN(inventories [][]int, n int) int {
	caloriesByInventory := make([]int, 0, len(inventories))
	for _, inventory := range inventories {
		weight := 0
		for _, item := range inventory {
			weight += item
		}
		caloriesByInventory = append(caloriesByInventory, weight)
	}

	sort.Slice(caloriesByInventory, func(i, j int) bool {
		return caloriesByInventory[i] > caloriesByInventory[j]
	})

	sum := 0
	for _, v := range caloriesByInventory[0:n] {
		sum += v
	}
	return sum
}

func readInventories(filename string) [][]int {
	inventories := make([][]int, 0)
	inventory := make([]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inventories = append(inventories, inventory)
			inventory = make([]int, 0)
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			inventory = append(inventory, calories)
		}
	}

	inventories = append(inventories, inventory)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inventories
}
