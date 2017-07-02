package main

import (
	//"fmt"
	"io/ioutil"
	i "./interpreter"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("bf/mandelbrot.bf")
	check(err)

	it, _ := i.ByteIterator(data)

	i.Debug.Parser = false

	prg := i.Parse(it)

	if i.Debug.Parser {
		fmt.Println("Program:")
		fmt.Println(prg)
	}

	fmt.Println("======================  Run  ================")
	interpreter := i.NewBFInterpreter()
	i.Debug.Interpreter = false
	interpreter.Run(prg)
	fmt.Println("\n==================== STATS ================")
	fmt.Printf("Ticks:         %v\nTicks in fact: %d\n",
		interpreter.Ticks,
		interpreter.TicksFact)
}