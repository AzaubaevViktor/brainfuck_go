package interpreter

import (
	"fmt"
)

type BFInterpreter struct {
	memory []MemoryCell
	mp	   MemoryAddress
	depth  int

	Ticks     uint64
	TicksFact uint64
}

func NewBFInterpreter() *BFInterpreter {
	return &BFInterpreter{memory: make([]MemoryCell, 10000)}
}

func (I BFInterpreter) cur() MemoryCell {
	return I.memory[I.mp]
}

func (I BFInterpreter) debug() bool {
	return Debug.Interpreter
}

func (I *BFInterpreter) Run(prg Program) {
	for _, ex := range prg {
		if I.debug() {
			fmt.Printf(getIndent(I.depth) + "%v:\n", ex)
		}
		switch ex.(type) {
		case Modifier:
			m := ex.(Modifier)
			for addr, d := range m.mem {
				I.memory[I.mp+addr] += d
			}
			I.mp += m.dMP
		case Operation:
			switch ex.(Operation) {
			case OP_PRINT:
				fmt.Printf("%c", I.memory[I.mp])
			}
		case Cycle:
			cycle := ex.(Cycle)
			if I.debug() {
				fmt.Println(getIndent(I.depth) + "Go into cycle!")
			}
			I.depth += 1
			for 0 != I.cur() {
				I.Run(cycle.prg)
			}
			I.depth -= 1
			if I.debug() {
				fmt.Println(getIndent(I.depth) + "Go out of cycle!")
			}
		}
		if I.debug() {
			fmt.Printf(getIndent(I.depth) + "%v\n==================\n", I)
		}
	}
}

func (I BFInterpreter) String() string {
	return fmt.Sprintf("%v MP: %d", I.memory, I.mp)
}



