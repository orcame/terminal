package main

import "../src"

func main() {
	terminal.Stdout.Color('y').
		Intensity().
		Print("this is a blue string.").
		Nl(3).
		Print("Hello"," ","World").
		Reset()
}