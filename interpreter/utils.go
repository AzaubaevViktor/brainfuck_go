package interpreter

type stack []int

func (A *stack) Push(value int) {
	*A = append(*A, value)
}

func (A *stack) Pop() int {
	var last int
	last , *A = (*A)[len(*A) - 1], (*A)[:len(*A) - 1]
	return last
}
