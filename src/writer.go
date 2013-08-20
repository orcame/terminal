package terminal

type writer interface{
	Color(colorCode uint) writer
	Bcolor(colorCode uint) writer
	Intensity() writer
	Underline() writer
	Title(title string) writer
	Reset() writer
	Print(str ...interface{}) writer
	Nl(count ...interface{}) writer
}
