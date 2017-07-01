package main

import (
	//"fmt"
	"io/ioutil"
	i "./interpreter"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("bf/mandelbrot.bf")
	check(err)

	ops := i.Parse(data)

	//fmt.Println("Operations:")
	//fmt.Println(ops)

	//fmt.Println("Run:")
	interpreter := i.NewBFInterpreter(ops)

	interpreter.Run()
}