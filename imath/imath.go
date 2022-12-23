package imath

func Max(is ...int) int {
	max := is[0]
	for _, i := range is {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(is ...int) int {
	min := is[0]
	for _, i := range is {
		if i < min {
			min = i
		}
	}
	return min
}

func Abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}
