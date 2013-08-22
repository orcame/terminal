package terminal

import(
	"syscall"
)

var(
	colorCodes=map[uint]uint{
		'd':0x0, //# Black
		'b':0x01,//# Blue
		'g':0x02,//# Green
		'r':0x04,//# Red
		'c':0x03,//# Cyan
		'p':0x05,//# Purple
		'y':0x06,//# Yellow
		'w':0x07,//# White

		'D':0x0, // # Black
		'B':0x10,// # Blue
		'G':0x20,// # Green
		'R':0x40,// # Red
		'C':0x30,// # Cyan
		'P':0x50,// # Purple
		'Y':0x60,// # Yellow
		'W':0x70,// # White

		'i':0x08,	//foreground Intensity
		'I':0x80,	//background Intensity
		'_':0x8000,	//underline
		//'!':0x0, 	//bold todo:the bold value is wrong.
	}
)

type Setter struct{
	attr uint
	fcolor uint
	bcolor uint
	saved uint
}

func (s *Setter) save(){
	s.saved = s.attr|s.fcolor|s.bcolor
}

func (s *Setter) setStyle(codes string,t *writer) {
	for _,c:=range codes{
		if st,ok :=colorCodes[uint(c)];ok{
			switch uint(c) {
				case 'd','b','g','r','c','p','y','w':
					s.fcolor=st
				case 'D','B','G','R','C','P','Y','W':
					s.bcolor=st
				default:
					s.attr=s.attr|st
			}			
		}
	}
	style:=s.attr|s.fcolor|s.bcolor
	setStyleToConsole(style,t)
}

func setStyleToConsole(style uint, t *writer){
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle ,_,_:=proc.Call(t.handle, uintptr(style))
	closeHandle:=kernel32.NewProc("CloseHandle")
	closeHandle.Call(handle)
}

func (s *Setter) resetStyle(t *writer){
	if s.saved !=0{
		setStyleToConsole(s.saved,t)
		s.saved	=0
	}else{
		s.attr=0x0
		s.fcolor=colorCodes['w']
		s.bcolor=colorCodes['D']
		setStyleToConsole(s.attr|s.fcolor|s.bcolor,t)
	}
}

func (s *Setter) setTitle(t *writer){
	//todo:set titlle.
}