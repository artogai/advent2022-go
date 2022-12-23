package day22

import "advent2022/matrix"

type cube[A any] struct {
	front, right, back, left, top, bottom matrix.Square[A]
}

func (c *cube[A]) rotateRight() {
	c.front, c.right, c.back, c.left = c.left, c.front, c.right, c.back

	c.top.RotateAntiClockwise()
	c.bottom.RotateClockwise()
}

func (c *cube[A]) rotateLeft() {
	c.front, c.right, c.back, c.left = c.right, c.back, c.left, c.front

	c.top.RotateClockwise()
	c.bottom.RotateAntiClockwise()
}

func (c *cube[A]) rotateUp() {
	c.back.RotateClockwise()
	c.back.RotateClockwise()
	c.front, c.top, c.back, c.bottom = c.bottom, c.front, c.top, c.back
	c.back.RotateClockwise()
	c.back.RotateClockwise()

	c.left.RotateAntiClockwise()
	c.right.RotateClockwise()
}

func (c *cube[A]) rotateDown() {
	c.back.RotateClockwise()
	c.back.RotateClockwise()
	c.front, c.top, c.back, c.bottom = c.top, c.back, c.bottom, c.front
	c.back.RotateClockwise()
	c.back.RotateClockwise()

	c.left.RotateClockwise()
	c.right.RotateAntiClockwise()
}

func (c *cube[A]) edgeSize() int {
	return c.front.Size()
}
