package day11

import (
	"sort"

	"github.com/samber/lo"
)

type monkeys []*monkey

type monkey struct {
	indx      int
	items     []int64
	op        func(int64) int64
	cond      func(int64) int64
	inspected int
}

func (m *monkey) nextItem(reduceLevels bool) (int64, int64) {
	if len(m.items) > 0 {
		newItem := m.op(m.items[0])
		m.items = m.items[1:]
		if reduceLevels {
			newItem = newItem / 3
		} else {
			newItem = newItem % 9699690
		}
		throwIdx := m.cond(newItem)
		m.inspected = m.inspected + 1
		return throwIdx, newItem
	}
	return -1, -1
}

func (m *monkey) addItem(item int64) {
	m.items = append(m.items, item)
}

func (ms monkeys) runSimulation(rounds int, reduceLevels bool) {
	for round := 0; round < rounds; round++ {
		ms.nextRound(reduceLevels)
	}
}

func (ms monkeys) nextRound(reduceLevels bool) {
	for _, m := range ms {
		for {
			idx, item := m.nextItem(reduceLevels)
			if idx != -1 {
				ms.addItem(idx, item)
			} else {
				break
			}
		}
	}
}

func (ms monkeys) businessScore() int {
	inspected := lo.Map(ms, func(m *monkey, _ int) int {
		return m.inspected
	})
	sort.Slice(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})
	return inspected[0] * inspected[1]
}

func (ms monkeys) addItem(idx int64, item int64) {
	ms[idx].addItem(item)
}
