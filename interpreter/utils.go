package interpreter

type stack []ProgramAddress

func (A *stack) Push(value ProgramAddress) {
	*A = append(*A, value)
}

func (A *stack) Pop() ProgramAddress {
	var last ProgramAddress
	last , *A = (*A)[len(*A) - 1], (*A)[:len(*A) - 1]
	return last
}

func abs(i int) uint64 {
	if i < 0 {
		i = -i
	}
	return uint64(i)
}

var Debug bool
