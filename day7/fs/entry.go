package fs

import (
	"strconv"
	"strings"
)

type Entry interface {
	Name() string
	Size() int
	SetParent(p *Directory)
}

func Parse(s string) Entry {
	tupleStr := strings.Split(s, " ")
	first, second := tupleStr[0], tupleStr[1]
	if first == "dir" {
		return NewDirectory(second, nil)
	} else {
		size, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}
		return NewFile(second, size)
	}
}
