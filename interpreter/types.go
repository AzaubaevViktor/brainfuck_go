package interpreter

import "fmt"

const (
	NOP = 0
	PLUS = 1
	MOVE = 2
	CYCLE_OP = 4
	CYCLE_CLOSE = 5
	PRINT = 6
	READ = 7
)

type MemoryCell uint8
type Program []Operation
type ProgramAddress int
type MemoryAddress int


func (P *Program) push(operation Operation) {
	if NOP != operation.opcode {
		*P = append(*P, operation)
	}
}

func (P Program) pc() ProgramAddress {
	return ProgramAddress(len(P))
}

func (P Program) String() string {
	s := ""
	for i, op := range P {
		s += fmt.Sprintf("%d: %v\n", i, op)
	}
	return s
}