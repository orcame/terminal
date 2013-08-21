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
		'!':0x0, 	//bold todo:the bold value is wrong.
	}
)

type Setter struct{
	style uint
}

func (s *Setter) setStyle(code uint,t *writer) {
	if c,ok:=colorCodes[code];ok{
		s.style =s.style|c
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle ,_,_:=proc.Call(t.handle, uintptr(s.style))
		closeHandle:=kernel32.NewProc("CloseHandle")
		closeHandle.Call(handle)
	}
}

func (s *Setter) resetStyle(t *writer){
	s.style=0
	s.setStyle('w',t)
	s.setStyle('D',t)
}

func (s *Setter) setTitle(t *writer){
	//todo:set titlle.
}