package interpreter

const (
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
