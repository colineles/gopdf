package gopdf


import (
	"bytes"
	"strconv"
	"fmt"
)

type ContentObj struct { //impl IObj
	buffer bytes.Buffer
	stream  bytes.Buffer
	
	//text bytes.Buffer
	getRoot func()(*GoPdf)
}

func (me *ContentObj) Init(funcGetRoot func()(*GoPdf)) {
	me.getRoot = funcGetRoot
}

func (me *ContentObj) Build() {

	streamlen := me.stream.Len()
		
	me.buffer.WriteString("<<\n")
	me.buffer.WriteString("/Length "+strconv.Itoa(streamlen)+"\n")
	me.buffer.WriteString(">>\n")
	me.buffer.WriteString("stream\n")
	me.buffer.Write(me.stream.Bytes())
	me.buffer.WriteString("endstream\n")
}

func (me *ContentObj) GetType() string {
	return "Content"
}

func (me *ContentObj) GetObjBuff() *bytes.Buffer {
	return &(me.buffer)
}


func (me *ContentObj) AppendStream(rectangle *Rect,text string){
	
	
	fontSize := me.getRoot().Curr.Font_Size
	
	x := fmt.Sprintf("%0.2f",me.getRoot().Curr.X)
	y := fmt.Sprintf("%0.2f",me.getRoot().config.PageSize.H  - me.getRoot().Curr.Y - (float64(fontSize) *0.7) )
	
	me.stream.WriteString("BT\n")
	me.stream.WriteString(x+" "+y+" TD\n")
	me.stream.WriteString("/F"+strconv.Itoa( me.getRoot().Curr.Font_FontCount  + 1)+" "+ strconv.Itoa(fontSize)+" Tf\n")
	me.stream.WriteString("("+text+") Tj\n")
	me.stream.WriteString("ET\n")
	if rectangle == nil {
		me.getRoot().Curr.X += StrHelper_GetStringWidth(text,fontSize,me.getRoot().Curr.Font_IFont)
	}else{
		me.getRoot().Curr.X += rectangle.W
	}
	
}


func (me *ContentObj) AppendStreamLine(x1 float64 , y1 float64, x2 float64 , y2 float64){

	h := me.getRoot().config.PageSize.H
	me.stream.WriteString( fmt.Sprintf("%0.2f %0.2f m %0.2f %0.2f l s\n",x1,h - y1,x2,h - y2))
}

func (me *ContentObj) AppendStreamSetLineWidth(w float64){
	
	me.stream.WriteString(fmt.Sprintf("%.2f w\n",w))
	
}


func (me *ContentObj) AppendStreamImage(iindex int,x float64,y float64,rect *Rect){
	
	me.stream.WriteString(fmt.Sprintf("q 85.04 0 0 92.27 28.35 732.61 cm /I%d Do Q\n",iindex+1))
	
	//me.stream.WriteString(fmt.Sprintf("q %.2F 0 0 %.2F %.2F %.2F cm /I%d Do Q",rect.W,rect.H,x,y,iindex+1))
	
}