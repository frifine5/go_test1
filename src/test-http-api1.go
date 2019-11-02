package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//第一个参数是接口名，第二个参数 http handle func
	http.HandleFunc("/", helloWorld)
	//服务器要监听的主机地址和端口号
	http.ListenAndServe(":8082", nil)
}

// http handle func
func helloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("get on time: "+ (time.Now()).String())

	// 返回字符串 "Hello world"
	fmt.Fprint(rw, "Hello world")

}