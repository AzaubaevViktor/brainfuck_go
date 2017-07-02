package interpreter

import "fmt"

const (
	_format_debug = "CH:%c\nPROG:%v\nMOD:%v\n===============\n"
)

func NewParse(it Iterator) Program {
	prg := Program{}
	modifier := NewModifier()

	for ch, hasNext := it(); hasNext && ch != ']'; ch, hasNext = it() {
		switch ch {
		case '+':
			modifier.add(1)
		case '-':
			modifier.add(-1)
		case '>':
			modifier.move(1)
		case '<':
			modifier.move(-1)
		case '.':
			prg.push(modifier)
			modifier = NewModifier()
			prg.push(Operation(OP_PRINT))
		case ',':
			prg.push(modifier)
			modifier = NewModifier()
			prg.push(Operation(OP_READ))
		case '[':
			prg.push(modifier)
			modifier = NewModifier()
			cycle := Cycle{}
			if Debug.Parser {
				fmt.Printf(_format_debug, ch, prg, modifier)
			}
			cycle.prg = NewParse(it)
			prg.push(cycle)
		}

		if Debug.Parser {
			fmt.Printf(_format_debug, ch, prg, modifier)
		}
	}

	prg.push(modifier)

	if Debug.Parser {
		fmt.Printf(_format_debug, ']', prg, modifier)
	}

	return prg
}
