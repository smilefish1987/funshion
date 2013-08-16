package main

import (
	"io"
	"log"
	"net/http"
)

/*
 * helloHandler是http.HandlerFunc类型的实例，有两个必要的参数http.ResponseWriter和http.Request
 * http.ResponseWriter用户包装处理http服务端的响应信息
 * *http.Request 表示的是此次http请求的一个数据结构体，代表一个客户端
 */
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello,world!")
}

func main() {
	//http.HandlerFunc()用于分发请求，即针对某一路径请求将其映射到指定的业务逻辑处理方法中
	http.HandleFunc("/hello", helloHandler)
	//http.ListenAndServer()用于监听8080端口，接受并调用内部程序来处理连接到此端口的请求
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
