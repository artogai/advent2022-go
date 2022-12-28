package day20

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"testing"
)

func TestDecrypt(t *testing.T) {
	numbers := readNumbers("encrypted.txt")
	message := NewMessage(numbers)
	message.Decrypt()
	res := message.Get(1000) + message.Get(2000) + message.Get(3000)
	fmt.Println(res)
}

func TestDecrypt2(t *testing.T) {
	numbers := readNumbers("encrypted.txt")
	for i := range numbers {
		numbers[i] *= 811589153
	}
	message := NewMessage(numbers)
	for i := 0; i < 10; i++ {
		message.Decrypt()
	}
	res := message.Get(1000) + message.Get(2000) + message.Get(3000)
	fmt.Println(res)
}

func readNumbers(filename string) []int {
	lines := file.ReadFileLines(filename)
	arr := make([]int, len(lines))
	for i, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		arr[i] = v
	}
	return arr
}
