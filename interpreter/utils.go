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

type Iterator func() (byte, bool)

func ByteIterator(D []byte) (Iterator, bool) {
	dataLen := len(D)
	id := 0
	return func() (byte, bool) {
		id++
		if id < dataLen {
			return D[id-1], true
		} else {
			return 0, false
		}
	}, id < dataLen
}

var Debug struct{
	Parser bool
	Interpreter bool
}
