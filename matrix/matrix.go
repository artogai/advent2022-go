package matrix

type Square[A any] [][]A

func NewSquare[A any](n int) Square[A] {
	mx := make([][]A, n)
	for i := range mx {
		mx[i] = make([]A, n)
	}
	return mx
}

func (mx Square[A]) RotateClockwise() {
	n := len(mx)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := mx[i][j]
			mx[i][j] = mx[n-1-j][i]
			mx[n-1-j][i] = mx[n-1-i][n-1-j]
			mx[n-1-i][n-1-j] = mx[j][n-1-i]
			mx[j][n-1-i] = temp
		}
	}
}

func (mx Square[A]) RotateAntiClockwise() {
	n := len(mx)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := mx[i][j]
			mx[i][j] = mx[j][n-1-i]
			mx[j][n-1-i] = mx[n-1-i][n-1-j]
			mx[n-1-i][n-1-j] = mx[n-1-j][i]
			mx[n-1-j][i] = temp
		}
	}
}

func (mx Square[A]) Size() int {
	return len(mx)
}
