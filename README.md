This is a golang package used for colorful output. 

Cross platform win32/linux. 

example:

```go 

terminal.Stdout.Color('r').
	Intensity().
	Print('this is a red text')

terminal.Stdout.Fprint("@{gi} Green and Intensity Text","@{b}","Blue Text","@{yiB}Yellow text with Blue background")

```