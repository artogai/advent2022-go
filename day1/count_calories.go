package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func CountCalories(file string) int {
	return maxCalories(readInventories(file))
}

func maxCalories(inventories [][]int) int {
	max_weight := 0
	for _, inventory := range inventories {
		weight := 0
		for _, item := range inventory {
			weight += item
		}
		if weight > max_weight {
			max_weight = weight
		}
	}
	return max_weight
}

func readInventories(file string) [][]int {
	inventories := make([][]int, 0)
	inventory := make([]int, 0)

	f, err := os.Open(file)
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
