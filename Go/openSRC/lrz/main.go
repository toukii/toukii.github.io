package main

import (
	"text/template"
	"github.com/toukii/goutils"
	"net/http"
	"fmt"
	"encoding/base64"
	"os"
	"strings"
	"bufio"
	"time"
)

func main() {
	my:=&myMux{}
	mux:= http.NewServeMux()
	mux.Handle("/", my)

	http.ListenAndServe(":8080",mux)
}

type myMux struct {

}

func (m *myMux)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	if req.Method=="POST" {
		if "/upload"==req.RequestURI {
			m.upload(w,req)
		}else{
			m.streamUpload(w,req)
		}
		return
	}
	tpl,err:=template.ParseFiles("index.html")
	if goutils.CheckErr(err){
		fmt.Fprint(w, err)
		return
	}
	tpl.Execute(w,nil)
}

func (m *myMux) upload(w http.ResponseWriter, req * http.Request)  {
	start:=time.Now()
	req.ParseForm()
	data:=req.FormValue("imgdata")
	b64data := data[strings.IndexByte(data, ',')+1:]
	enc:=base64.StdEncoding
	bs,derr:=enc.DecodeString(b64data)
	if goutils.CheckErr(derr) {
		fmt.Fprint(w, derr)
		return
	}
	gerr:=goutils.WriteFile("upload.png",bs)
	if goutils.CheckErr(gerr) {
		fmt.Fprint(w,gerr)
		return
	}
	fmt.Fprint(w,"success")
	fmt.Println("cost:",time.Now().Sub(start))
}


func (m *myMux) streamUpload(w http.ResponseWriter, req * http.Request)  {
	start:=time.Now()
	enc:=base64.StdEncoding
	file,err:=os.OpenFile("stream.png",os.O_CREATE|os.O_WRONLY,0644)
	defer file.Close()
	if goutils.CheckErr(err) {
		fmt.Fprint(w, err)
		return
	}

	try:=make([]byte,120)
	peekR:=bufio.NewReader(req.Body)
	peekR.ReadBytes(',')
	//fmt.Println(goutils.ToString(peekBs))
	allbytes:=make([]byte,0,req.ContentLength)
	for{
		n1,err1:=peekR.Read(try)
		if n1>0 {
			allbytes = append(allbytes,try[:n1]...)
		}
		if goutils.CheckErr(err1) {
			break
		}
	}

	//bs,err2:=enc.DecodeString(goutils.ToString(allbytes))
	//if !goutils.CheckErr(err2) {
	//	file.Write(bs)
	//}

	dst:=make([]byte,len(allbytes)*2)
	n3,err:=enc.Decode(dst,allbytes)
	if !goutils.CheckErr(err) {
		file.Write(dst[:n3])
	}
	fmt.Fprint(w,"success")
	fmt.Println("cost:",time.Now().Sub(start))
}


func (m *myMux) streamUpload1(w http.ResponseWriter, req * http.Request)  {

	enc:=base64.StdEncoding
	//r:=base64.NewDecoder(enc,req.Body)

	file,err:=os.OpenFile("stream.png",os.O_CREATE|os.O_WRONLY,0644)
	defer file.Close()
	if goutils.CheckErr(err) {
		fmt.Fprint(w, err)
		return
	}

	try:=make([]byte,120)
	peekR:=bufio.NewReader(req.Body)
	_,err=peekR.ReadBytes(',')
	goutils.CheckErr(err)
	//fmt.Println(goutils.ToString(peekBs))
	dst:=make([]byte,240)
	start:=time.Now()
	allbytes:=make([]byte,0,req.ContentLength)
	fmt.Println(cap(allbytes))
	for{
		n1,err1:=peekR.Read(try)
		//n1,err1:=r.Read(try)
		//fmt.Print(try)
		//fmt.Println(n1,err1)
		if n1>0 {
			n2,err2:=enc.Decode(dst,try[:n1])
			if goutils.CheckErr(err2) {
				fmt.Println(n1,n2,try)
				fmt.Println(dst)
			}
			if n2>0 && !goutils.CheckErr(err2) {
				file.Write(dst[:n2])
			}

			//bs,err2:=enc.DecodeString(goutils.ToString(try[:n1]))
			//if !goutils.CheckErr(err2) {
			//	file.Write(bs)
			//}
			allbytes = append(allbytes,try[:n1]...)
		}
		if goutils.CheckErr(err1) {
			break
		}
	}

	//bs,err2:=enc.DecodeString(goutils.ToString(allbytes))
	//if !goutils.CheckErr(err2) {
	//	file.Write(bs)
	//}

	//dst:=make([]byte,len(allbytes)*2)
	//n3,err:=enc.Decode(dst,allbytes)
	//if !goutils.CheckErr(err) {
	//	file.Write(dst[:n3])
	//}

	fmt.Println("cost:",time.Now().Sub(start))

	/*r2:=io.TeeReader(r,file)
	buf:=make([]byte,10)
	for{
		n,rerr:=r2.Read(buf)
		fmt.Print(n)
		if goutils.CheckErr(rerr)|| n<=0 {
			break;
		}
	}*/
	//fmt.Println(data)
	fmt.Println("success")
	fmt.Fprint(w,"success")
}