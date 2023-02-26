package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	// 将以 "/" 开头的 URL 都用 handler 函数处理
	http.HandleFunc("/", handler)
	// 启动服务器监听 8000 端口处的请求
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	// 返回请求的消息头和表单数据
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// go 允许一个简单的语句跟在 if 条件的前面
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "from[%q] = %q\n", k, v)
	}
}