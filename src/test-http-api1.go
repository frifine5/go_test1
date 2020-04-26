package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//第一个参数是接口名，第二个参数 http handle func
	http.HandleFunc("/a", helloWorld) // 使用‘/’或匹配所有
	http.HandleFunc("/abc", helloWorld2)
	//服务器要监听的主机地址和端口号
	http.ListenAndServe(":8082", nil)
}

// http handle func
func helloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("get on time: "+ (time.Now()).String())

	// 返回字符串 "Hello world"
	fmt.Fprint(rw, " Hello world")

}
func helloWorld2(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("get on time 2: "+ (time.Now()).String())

	// 返回字符串 "Hello world"
	fmt.Fprint(rw, "Hello world - in url 2")

}
