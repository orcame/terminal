package terminal

import(
	"fmt"
	"io"
	"bytes"
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
	t.setter.setStyle(string(code),t)
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

func (t *writer) Bold() *writer{
	return t.setStyle('!')
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

func (t *writer) complieStyle(input *bytes.Buffer)(style, str *bytes.Buffer){
	style=bytes.NewBufferString("")
	str=bytes.NewBufferString("")
	c,_,err :=input.ReadRune()
	if err==nil{
		if uint(c)=='{'{
			for {
				c,_,err = input.ReadRune()
				if err !=nil{
					style,str=str,style
					break
				}
				if c=='}'{
					break
				}else{					
					style.WriteRune(c)
				}
			}
		}else{
			str.WriteRune(c)
		}
	}
	return
}

func (t *writer) Fprint(str ...interface{}) *writer{
	t.setter.save()
	for _,s := range str{
		if v,ok:=s.(string);ok{
			input:=bytes.NewBufferString(v)
			buffer:=bytes.NewBufferString("")
			for {
				c, _, err := input.ReadRune()
				if err != nil {
					break
				}
				switch uint(c) {
					case '@':
						style,str:=t.complieStyle(input)
						if style.Len()>=0{							
							t.Print(buffer.String())
							buffer.Reset()
							t.setter.setStyle(style.String(),t)
						}
						if str.Len()>0{
							buffer.Write(str.Bytes())
						}
					default:
						buffer.WriteRune(c)
				}
			}
			if buffer.Len()>0{
				t.Print(buffer.String())
				buffer.Reset()
			}
		}
	}
	t.setter.resetStyle(t)
	return t
}
