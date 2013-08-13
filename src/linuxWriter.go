package terminal

import(
	"fmt"
	"io"
	"log"
	"os"
)

var codeMap :=map[int][int]{
	'd':30,// Black='\e[0;30m'        # Black
	'r':31,// Red='\e[0;31m'          # Red
	'g':32,// Green='\e[0;32m'        # Green
	'y':33,// Yellow='\e[0;33m'       # Yellow
	'b':34,// Blue='\e[0;34m'         # Blue
	'p':35,// Purple='\e[0;35m'       # Purple
	'c':36,// Cyan='\e[0;36m'         # Cyan
	'w':37// White='\e[0;37m'        # White
}

type linuxWriter struct{
	io.Writer
	currentColor string
}

func getStyle(color string,bold bool,underline bool,intensity bool) string{
	var code := codeMap[color]
	if intensity{
		code +=60
	}
	var a1:=0
	if bold:{
		a1=1
	}else if underline{
		a1=4
	}
	return fmt.Sprintf("\x1b[%d;%d;%dm",
	//todo: get the color string from code.
}

func (t *writer) writeOut(v string){
	if _,err:=io.WriteString(t,v); err!=nil{
		log.Fatal("write error.")
	}
}

func (t *writer) color(colorCode string) *writer{
	if t.currentColor == colorCode{
		return t
	}
	t.writeOut(getColor(colorCode))
	return t
}

func (t *writer) intensity() *writer{

	return t
}

func (t *writer) underline() *writer{

	return t
}

func (t *writer) title(v string) *writer{

	return t
}

func (t *writer) reset() *writer{
	t.currentColor=""
	t.writeOut("\x1b[0m")
	return t
}

func (t *writer) print(v string) * writer{
	return t
}

func (t *writer) nl() *writer{

	return t
}