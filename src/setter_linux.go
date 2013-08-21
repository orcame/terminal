package terminal

import(
	"fmt"
)

var(
	colorMap = map[uint]uint{
		'd':30,// Black='\e[0;30m'        # Black
		'r':31,// Red='\e[0;31m'          # Red
		'g':32,// Green='\e[0;32m'        # Green
		'y':33,// Yellow='\e[0;33m'       # Yellow
		'b':34,// Blue='\e[0;34m'         # Blue
		'p':35,// Purple='\e[0;35m'       # Purple
		'c':36,// Cyan='\e[0;36m'         # Cyan
		'w':37,// White='\e[0;37m'        # White

		'D':40, // # Black
		'R':41,// # Blue
		'G':42,// # Green
		'Y':44,// # Red
		'B':43,// # Cyan
		'P':45,// # Purple
		'C':46,// # Yellow
		'W':47,// # White

		'i':0x08,	//foreground Intensity
		'I':0x80,	//background Intensity
		'_':4,	//underline
		'!':1, 	//bold todo:the bold value is wrong.
	}
)

type Setter struct{
	fcolor uint
	bcolor uint
}

func (s *Setter) setStyle(code uint,t *writer){
	prefix:="\x1b["
	var style uint=0
	if code =='i'{
		if s.fcolor>0{
			style=s.fcolor+60;
		}
	}else if code =='I'{
		if s.bcolor>0{
			style=s.bcolor+60
		}
	}else if f,ok:=colorMap[code];ok{
		if f>40{
			s.bcolor=f
		}else if f>30{
			s.fcolor=f
		}
		style=f
	}
	if style>0{
		fmt.Fprint(t,prefix+string(style))
	}
}

func (s *Setter) resetStyle(t *writer){
	s.fcolor=0
	s.bcolor=0
	fmt.Fprint(t,"\x1b[0m")
}

func (s *Setter) setTitle(t *writer){
}
