package day6

func FindMarkerPos(signal string, windowSize int) int {
	for i := 0; i < len(signal); i++ {
		if i >= windowSize {
			if isUnique(signal[i-windowSize : i]) {
				return i
			}
		}
	}
	return 0
}

func isUnique(s string) bool {
	checked := make(map[rune]bool, len(s))
	for _, r := range s {
		_, exists := checked[r]
		if exists {
			return false
		} else {
			checked[r] = true
		}
	}
	return true
}
