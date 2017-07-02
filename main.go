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
	data, err := ioutil.ReadFile("bf/hello.bf")
	check(err)

	it, _ := i.ByteIterator(data)

	i.Debug.Parser = true

	prg := i.NewParse(it)

	fmt.Println("Program:")
	fmt.Println(prg)

	//fmt.Println("Run:")
	//interpreter := i.NewBFInterpreter(ops)
	//i.Debug = false
	//interpreter.Run()
	//fmt.Println("\n==================== STATS ================")
	//fmt.Printf("Ticks:         %v\nTicks in fact: %d\n",
	//	interpreter.Ticks,
	//	interpreter.TicksFact)
}