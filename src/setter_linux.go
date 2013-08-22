package terminal

import(
	"fmt"
	"strings"
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

		'_':4,	//underline
		'!':1, 	//bold todo:the bold value is wrong.
	}
)

type Setter struct{
	fcolor string
	bcolor string
	attr string
	saved string
}

func (s *Setter) save(){
	s.saved=string(s.fcolor)+string(s.bcolor)+s.attr
}

func (s *Setter) setStyle(code uint,t *writer){
	var style uint=0
	for _,c:=range *codes{
		if st,ok :=colorCodes[c];ok{
			sc:=string(c)
			style=st
			switch c{
				case 'd','b''g','r','c','p','y','w':
					s.fcolor=sc
				case 'D','B''G','R','C','P','Y','W':
					s.bcolor=sc
				default:
					if c== 'i' || c=='I'{
						style+=60
					}
					if !strings.Contains(s.attr,sc){
						s.attr=append(s.attr,sc)
					}
			}
			fmt.Fprint(t,fmt.Sprintf("\033[%dm",style))	
		}
	}	
}

func (s *Setter) resetStyle(t *writer){
	s.fcolor=""
	s.bcolor=""
	s.attr=""
	if len(s.saved)!=0{
		s.setStyle(s.saved,t)
		s.saved=""
	}else{
		fmt.Fprint(t,"\033[0m")
	}
}

func (s *Setter) setTitle(t *writer){
}
