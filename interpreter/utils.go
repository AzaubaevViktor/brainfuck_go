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

var Debug bool
