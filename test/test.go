package main

import "../src"

func main() {
	terminal.Stdout.Color('c').
		Intensity().
		Bold().
		Print("this is a blue string.").
		Nl(3).
		Color('y').
		Underline().
		Fprint("@{yi}[yellow]Hello"," ","@{ir}World").
		Nl().Fprint("hello @{bi}world.").Nl().Reset()

	//terminal.Stdout.Color('y').Nl().Print("\tThe means of word ").
		//		Color('c').Print("hello").Color('y').Print(" is:").Nl().Reset()
}
