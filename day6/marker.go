package day6

import mapset "github.com/deckarep/golang-set/v2"

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
	checked := mapset.NewThreadUnsafeSet[rune]()
	for _, r := range s {
		if added := checked.Add(r); !added {
			return false
		}
	}
	return true
}
