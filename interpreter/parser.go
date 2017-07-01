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

func Parse(data []byte) Program {
	var cycleStack stack
	var prg Program

	curOp := Operation{}

	for _, ch := range data {
		switch ch {
		case '+':
			switch curOp.opcode {
			case PLUS:
				curOp.count += 1
			default:
				prg.push(curOp)
				curOp = Operation{opcode:PLUS, count: 1}
			}
		case '-':
			switch curOp.opcode {
			case PLUS:
				curOp.count -= 1
			default:
				prg.push(curOp)
				curOp = Operation{opcode:PLUS, count: -1}
			}
		case '>':
			switch curOp.opcode {
			case MOVE:
				curOp.count += 1
			default:
				prg.push(curOp)
				curOp = Operation{opcode:MOVE, count: 1}
			}
		case '<':
			switch curOp.opcode {
			case MOVE:
				curOp.count -= 1
			default:
				prg.push(curOp)
				curOp = Operation{opcode:MOVE, count: -1}
			}
		case '[':
			prg.push(curOp)
			cycleStack.Push(prg.pc())
			prg.push(Operation{opcode: CYCLE_OP})
			curOp = Operation{opcode:NOP}
		case ']':
			prg.push(curOp)
			pos := cycleStack.Pop()
			prg[pos].addr = prg.pc()
			prg.push(Operation{opcode: CYCLE_CLOSE, addr: pos})
			curOp = Operation{opcode:NOP}
		case '.':
			prg.push(curOp)
			prg.push(Operation{opcode: PRINT})
			curOp = Operation{opcode:NOP}
		case ',':
			prg.push(curOp)
			prg.push(Operation{opcode: READ})
			curOp = Operation{opcode:NOP}
		}
	}

	prg.push(curOp)

	return prg
}
