package interpreter

import (
	"fmt"
	//"time"
)

type BFInterpreter struct {
	memory []MemoryCell
	pc     ProgramAddress
	prg    Program
	mp	   MemoryAddress
}

func NewBFInterpreter(program Program) *BFInterpreter {
	return &BFInterpreter{memory: make([]MemoryCell, 30000), pc: 0, prg: program}
}

func (I BFInterpreter) cur() MemoryCell {
	return I.memory[I.mp]
}

func (I *BFInterpreter) step() {
	op := I.prg[I.pc]

	switch op.opcode {
	case PLUS:
		I.memory[I.mp] += MemoryCell(op.count)
	case MOVE:
		I.mp += MemoryAddress(op.count)
	case CYCLE_OP:
		if 0 == I.cur() {
			I.pc = op.addr - 1 // Because below +=1
		}
	case CYCLE_CLOSE:
		if 0 != I.cur() {
			I.pc = op.addr // and +1 below
		}
	case PRINT:
		fmt.Printf("%c", I.cur())
	}

	I.pc += 1
}

func (I BFInterpreter) Run() {
	for I.pc < ProgramAddress(len(I.prg)) {
		//fmt.Println(I.prg[I.pc])
		//time.Sleep(100 * time.Millisecond)
		I.step()
		//fmt.Println(I)

	}
}

func (I BFInterpreter) String() string {
	return fmt.Sprintf("%v PC: %d; MP: %d", I.memory, I.pc, I.mp)
}



