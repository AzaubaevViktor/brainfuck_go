package interpreter

import "fmt"

type Operation struct{
	opcode int
	count  int
	addr   ProgramAddress
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

func Parse(data []byte) (ops []Operation) {
	var cycleStack stack

	position := ProgramAddress(0)

	for _, ch := range data {
		switch ch {
		case '+':
			ops = append(ops, Operation{opcode: PLUS, count: 1})
		case '-':
			ops = append(ops, Operation{opcode: PLUS, count: -1})
		case '>':
			ops = append(ops, Operation{opcode: MOVE, count: 1})
		case '<':
			ops = append(ops, Operation{opcode: MOVE, count: -1})
		case '[':
			ops = append(ops, Operation{opcode: CYCLE_OP})
			cycleStack.Push(position)
		case ']':
			pos := cycleStack.Pop()
			ops = append(ops, Operation{opcode: CYCLE_CLOSE, addr: pos})
			ops[pos].addr = position
		case '.':
			ops = append(ops, Operation{opcode: PRINT})
		case ',':
			ops = append(ops, Operation{opcode: READ})
		default:
			position -= 1
		}

		position += 1
	}

	return
}
