package terminal

import(
	"fmt"
	"io"
	"log"
	"syscall"
)
// STD_INPUT_HANDLE = -10
// STD_OUTPUT_HANDLE= -11
// STD_ERROR_HANDLE = -12

// FOREGROUND_BLACK = 0x0
// FOREGROUND_BLUE = 0x01 # text color contains blue.
// FOREGROUND_GREEN= 0x02 # text color contains green.
// FOREGROUND_RED = 0x04 # text color contains red.

// FOREGROUND_INTENSITY = 0x08 # text color is intensified.
// BACKGROUND_BLUE = 0x10 # background color contains blue.
// BACKGROUND_GREEN= 0x20 # background color contains green.
// BACKGROUND_RED = 0x40 # background color contains red.
// BACKGROUND_INTENSITY = 0x80 # background color is intensified.
var foregroundColorCodes=map[uint]uint{
	'd':0x0,// Black='\e[0;30m'        # Black
	'b':0x01,// Blue='\e[0;34m'         # Blue
	'g':0x02,// Green='\e[0;32m'        # Green
	'r':0x04,// Red='\e[0;31m'          # Red
	'c':0x03,// Cyan='\e[0;36m'         # Cyan
	'p':0x05,// Purple='\e[0;35m'       # Purple
	'y':0x06,// Yellow='\e[0;33m'       # Yellow
	'w':0x07,// White='\e[0;37m'        # White
}
var foregroundIntensity uint =0x08

var backgroundColorCodes=map[uint]uint{
	'd':0x0,// Black='\e[0;30m'        # Black
	'b':0x10,// Blue='\e[0;34m'         # Blue
	'g':0x20,// Green='\e[0;32m'        # Green
	'r':0x40,// Red='\e[0;31m'          # Red
	'c':0x30,// Cyan='\e[0;36m'         # Cyan
	'p':0x50,// Purple='\e[0;35m'       # Purple
	'y':0x60,// Yellow='\e[0;33m'       # Yellow
	'w':0x70,// White='\e[0;37m'        # White
}

var backgroundIntensity uint = 0x80

type win32Writer struct{
	io.Writer
	handle uintptr
	fcolor uint
	bcolor uint
	fintensity bool
	bintensity bool
	underline bool
	currentStyle uint
}

func createWin32Writer(writer io.Writer,handle uintptr) writer{
	return &win32Writer{
		writer,
		handle,
		0,
		0,
		false,
		false,
		false,
		0,
	}
}

func  (t *win32Writer)setStyle()  {
	style:= t.fcolor | t.bcolor
	if t.fintensity{
		style =style|foregroundIntensity
	}
	if t.bintensity{
		style = style|backgroundIntensity
	}
	t.currentStyle = style
	//handler := syscall.Stdout
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle ,_,_:=proc.Call(t.handle, uintptr(style))
	closeHandle:=kernel32.NewProc("CloseHandle")
	closeHandle.Call(handle)
}

func (t *win32Writer) Color(colorCode uint) writer{
	if v,ok:=foregroundColorCodes[colorCode];ok{
		if v != t.fcolor{
			t.fcolor=v
			t.setStyle()
		}
	}else{			
		log.Fatal("Wrong color code.")
	}
	return t
}

func (t *win32Writer) Bcolor(colorCode uint) writer{
	return t
}

func (t *win32Writer) Intensity() writer{
	if !t.fintensity{
		t.fintensity=true
		t.setStyle()
	}
	return t
}

func (t *win32Writer) Underline() writer{

	return t
}

func (t *win32Writer) Title(v string) writer{

	return t
}

func (t *win32Writer) Reset() writer{
	t.fcolor = foregroundColorCodes['w']
	t.bcolor = backgroundColorCodes['d']
	t.fintensity=false
	t.bintensity=false
	t.setStyle()
	return t
}

func (t *win32Writer) Print(str ...interface{}) writer{
	fmt.Fprint(t,str...)
	return t
}

func (t *win32Writer) Nl(count ...interface{}) writer{
	length := 1
	if len(count) > 0 {
		length =0
		for i:=0;i<len(count);i++{
			length+=count[i].(int)
		}		
	}
	for i := 0; i < length; i++ {
		fmt.Fprint(t,"\n")
	}
	return t
}