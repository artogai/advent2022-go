package day20

type Message struct {
	values []int
	idxs   []int
	length int
}

func NewMessage(init []int) *Message {
	values := make([]int, len(init))
	idxs := make([]int, len(init))
	for i := 0; i < len(init); i++ {
		values[i] = init[i]
		idxs[i] = i
	}
	return &Message{
		values: values,
		idxs:   idxs,
		length: len(init),
	}
}

func (m *Message) Decrypt() {
	for pos := 0; pos < m.length; pos++ {
		idx := m.findIndex(pos)
		m.mixAt(idx)
	}
}

func (m *Message) Get(idx int) int {
	zeroIdx := m.findValueIndex(0)
	realIdx := (zeroIdx + idx) % len(m.values)
	return m.values[realIdx]
}

func (m *Message) findValueIndex(v int) int {
	for i := 0; i < m.length; i++ {
		if m.values[i] == v {
			return i
		}
	}
	panic("not found")
}

func (m *Message) findIndex(pos int) int {
	for i := 0; i < m.length; i++ {
		if m.idxs[i] == pos {
			return i
		}
	}
	panic("not found")
}

func (m *Message) mixAt(idx int) {
	v := m.values[idx]
	if v > 0 {
		swapF(m.values, idx, v)
		swapF(m.idxs, idx, v)
	} else if v < 0 {
		swapB(m.values, idx, -v)
		swapB(m.idxs, idx, -v)
	}
}

func swapF(arr []int, idx, cnt int) {
	for i := 0; i < cnt; i++ {
		if idx+1 >= len(arr) {
			swap(arr, idx, 0)
			idx = 0
			continue
		}
		swap(arr, idx, idx+1)
		idx++
	}
}

func swapB(arr []int, idx, cnt int) {
	for i := 0; i < cnt; i++ {
		if idx-1 < 0 {
			swap(arr, idx, len(arr)-1)
			idx = len(arr) - 1
			continue
		}
		swap(arr, idx, idx-1)
		idx--
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
