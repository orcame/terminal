package main

import "../src"

func main() {
	terminal.Stdout.Color('b').
		Intensity().
		Bold().
		Print("this is a blue string.").
		Nl(3).
		Color('y').
		Underline().
		Print("[yellow]Hello"," ","World").
		Nl().
		Reset()
}
