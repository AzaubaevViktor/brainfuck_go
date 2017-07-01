package interpreter

import "fmt"

type Operation struct{
	Opcode int
	Count int
	Addr int
}

func (O Operation) String() string {
	opcode := "NOP"
	switch O.Opcode {
	case PLUS:
		if O.Count > 0 {
			opcode = "+"
		} else {
			opcode = "-"
		}
	case MOVE:
		if O.Count > 0 {
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
	return fmt.Sprintf("{%s %d *%d}", opcode, O.Count, O.Addr)
}

func Parse(data []byte) (ops []Operation) {
	var cycleStack stack

	for position, ch := range data {
		switch ch {
		case '+':
			ops = append(ops, Operation{Opcode:PLUS, Count:1})
		case '-':
			ops = append(ops, Operation{Opcode:PLUS, Count:-1})
		case '>':
			ops = append(ops, Operation{Opcode:MOVE, Count:1})
		case '<':
			ops = append(ops, Operation{Opcode:MOVE, Count:-1})
		case '[':
			ops = append(ops, Operation{Opcode:CYCLE_OP})
			fmt.Print(position)
			cycleStack.Push(position)
		case ']':
			pos := cycleStack.Pop()
			fmt.Print(pos)
			ops = append(ops, Operation{Opcode:CYCLE_CLOSE, Addr:pos})
			ops[pos].Addr = position
		case '.':
			ops = append(ops, Operation{Opcode:PRINT})
		case ',':
			ops = append(ops, Operation{Opcode:READ})
		}
	}

	return
}
