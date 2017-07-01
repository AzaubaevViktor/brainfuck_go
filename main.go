package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	PLUS = 1
	MOVE = 2
	CYCLE_OP = 4
	CYCLE_CLOSE = 5
	PRINT = 6
	READ = 7
)

type Operation struct{
	opcode int
	count int
	addr int
}

func (O Operation) String() string {
	opcode := "NOP"
	switch O.opcode {
	case PLUS:
		if O.count > 0 {
			opcode = "+"
		} else {
			opcode = "-"
		}
	case MOVE:
		if O.count > 0 {
			opcode = ">"
		} else {
			opcode = "<"
		}
	case CYCLE_OP:
		opcode = "["
	case CYCLE_CLOSE:
		opcode = "]"
	case PRINT:
		opcode = "."
	case READ:
		opcode = ","
	}
	return fmt.Sprintf("{%s %d *%d}", opcode, O.count, O.addr)
}

type Stack []int

func (A *Stack) Push(value int) {
	*A = append(*A, value)
}

func (A *Stack) Pop() int {
	var last int
	last , *A = (*A)[len(*A) - 1], (*A)[:len(*A) - 1]
	return last
}

func main() {
	dat, err := ioutil.ReadFile("bf/hello.bf")
	check(err)

	var ops []Operation

	var cycleStack Stack


	for position, ch := range dat {
		switch ch {
		case '+':
			ops = append(ops, Operation{opcode:PLUS, count:1})
		case '-':
			ops = append(ops, Operation{opcode:PLUS, count:-1})
		case '>':
			ops = append(ops, Operation{opcode:MOVE, count:1})
		case '<':
			ops = append(ops, Operation{opcode:MOVE, count:-1})
		case '[':
			ops = append(ops, Operation{opcode:CYCLE_OP})
			fmt.Print(position)
			cycleStack.Push(position)
		case ']':
			pos := cycleStack.Pop()
			fmt.Print(pos)
			ops = append(ops, Operation{opcode:CYCLE_CLOSE, addr:pos})
			ops[pos].addr = position
		case '.':
			ops = append(ops, Operation{opcode:PRINT})
		case ',':
			ops = append(ops, Operation{opcode:READ})
		}
	}

	fmt.Print(ops)

}