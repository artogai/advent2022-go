package day17

type cyclicBuff[A any] struct {
	arr []A
	i   int
}

func newCyclicBuff[A any](arr []A) *cyclicBuff[A] {
	return &cyclicBuff[A]{arr, 0}
}

func (c *cyclicBuff[A]) next() A {
	i := c.i
	c.i += 1
	if c.i == len(c.arr) {
		c.i = 0
	}
	return c.arr[i]
}
