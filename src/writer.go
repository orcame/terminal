package "terminal"

type writer interface{
	color(colorCode string) *writer
	bcolor(colorCode string)
	intensity() *writer
	underline() *writer
	title(v string) *writer
	reset() *writer
	print(v string) * writer
	nl() *writer
}
