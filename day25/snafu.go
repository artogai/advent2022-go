package day25

import (
	"math"
	"strings"
)

func encode(i int) string {
	var res strings.Builder
	n := findN(i) - 1

	for {
		k := int(math.Round(float64(i) / math.Pow(5, float64(n))))
		i -= k * int(math.Pow(5, float64(n)))

		switch k {
		case 2:
			res.WriteRune('2')
		case 1:
			res.WriteRune('1')
		case 0:
			res.WriteRune('0')
		case -1:
			res.WriteRune('-')
		case -2:
			res.WriteRune('=')
		}

		if n == 0 {
			break
		}
		n--
	}

	return res.String()
}

func findN(i int) int {
	n := 0
	for {
		max := 0
		for i := 0; i < n; i++ {
			max += 2 * int(math.Pow(5, float64(i)))
		}
		min := 1 * int(math.Pow(5, float64(n-1)))
		for i := 0; i < n-1; i++ {
			min += -2 * int(math.Pow(5, float64(i)))
		}
		if i >= min && i <= max {
			return n
		}
		n += 1
	}
}

func decode(s string) int {
	res := 0
	for i, c := range s {
		pos := int(math.Pow(5, float64(len(s)-i-1)))
		switch c {
		case '2':
			res += 2 * pos
		case '1':
			res += 1 * pos
		case '0':
			res += 0 * pos
		case '-':
			res += -1 * pos
		case '=':
			res += -2 * pos
		}
	}
	return res
}
