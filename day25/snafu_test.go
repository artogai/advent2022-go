package day25

import (
	"advent2022/file"
	"fmt"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	r := 0
	for _, s := range file.ReadLines("snafu.txt") {
		r += decode(s)
	}

	fmt.Println(r)
	fmt.Println(encode(r))
}
