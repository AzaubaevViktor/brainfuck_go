package interpreter

import "fmt"

const (
	OP_PRINT = 6
	OP_READ = 7
)

type MemoryCell uint8
type ProgramAddress int
type MemoryAddress int


type Executable interface {
	toS(int) string
}

func getIndent(indent int) string {
	ind := ""
	for i := 0; i < indent; i++ {
		ind += "  "
	}
	return ind
}

// =====================

type Operation uint8
func (O Operation) toS(indent int) string {
	s := getIndent(indent)
	switch O {
	case OP_PRINT:
		s += "<OP_PRINT>"
	case OP_READ:
		s += "<OP_READ>"
	default:
		s += "<OP_NOP>"
	}
	return s
}


// =====================
type Program []Executable

func (P *Program) push(e Executable) {
	if _, isModifier := e.(Modifier); !isModifier || e.(Modifier).active {
		*P = append(*P, e)
	}
}

func (P Program) toS(indent int) string {
	s := "<P \n"

	for _, v := range P {
		s += getIndent(indent) + v.toS(indent) + "\n"
	}
	return s + ">"
}

func (P Program) String() string {
	return P.toS(0)
}

// ===============
type Modifier struct {
	mem map[int]uint8
	dMP int
	active bool
}

func NewModifier() Modifier {
	return Modifier{mem: map[int]uint8{}}
}

func (M *Modifier) add(val int) {
	M.mem[M.dMP] += uint8(val)
	M.active = true
}

func (M *Modifier) move(val int) {
	M.dMP += val
	M.active = true
}

func (M Modifier) String() string {
	if M.active {
		line := ""
		if M.dMP > 0 {
			line = "-->"
		} else if M.dMP < 0 {
			line = "<-"
		}

		return fmt.Sprintf("<M(%s%d) %v>", line, M.dMP, M.mem)
	} else {
		return fmt.Sprintf("<M Empty>")
	}
}

func (M Modifier) toS(indent int) string {
	return getIndent(indent) + M.String()
}

// ====================
type Cycle struct {
	prg Program
}

func (C Cycle) String() string {
	return C.toS(0)
}

func (C Cycle) toS(indent int) string {
	return fmt.Sprintf(getIndent(indent) + "<C %s>", C.prg.toS(indent + 1))

}