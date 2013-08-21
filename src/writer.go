package terminal

import(
	"fmt"
	"io"
)

type writer struct{
	io.Writer
	handle uintptr
	setter setter
}

func createWriter(w io.Writer,handle uintptr) *writer{
	var s=&Setter{}
	return &writer{
		w,
		handle,
		s,
	}
}

func (t *writer) setStyle(code uint) *writer{
	t.setter.setStyle(code,t)
	return t
}

func (t *writer) resetStyle() *writer{
	t.setter.resetStyle(t)
	return t
}

func (t *writer) Color(colorCode uint) *writer{
	return t.setStyle(colorCode)
}

func (t *writer) Bcolor(colorCode uint) *writer{
	if colorCode>97{
		colorCode -=32
	}
	return t.setStyle(colorCode)
}

func (t *writer) Intensity() *writer{
	return t.setStyle('i')
}

func (t *writer) Bintensity() *writer{
	return t.setStyle('I')
}

func (t *writer) Underline() *writer{
	return t.setStyle('_')
}

func (t *writer) Title(v string) *writer{
	//todo:set the title.
	return t
}

func (t *writer) Reset() *writer{
	return t.resetStyle()
}

func (t *writer) Print(str ...interface{}) *writer{
	fmt.Fprint(t,str...)
	return t
}

func (t *writer) Nl(count ...interface{}) *writer{
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