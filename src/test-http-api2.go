package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
)



func main() {
	//第一个参数是接口名，第二个参数 http handle func
	http.HandleFunc("/mock/", getMockData)
	//服务器要监听的主机地址和端口号
	http.ListenAndServe(":8082", nil)
}

// http handle func
func getMockData(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("get on time: "+ (time.Now()).String())

	s, _ := ioutil.ReadAll(req.Body) //把  body 内容读入字符串 s
	fmt.Printf(  "%s\n", s)        //在返回页面中显示内容。


	path := req.RequestURI
	path = "/mnt" + path

	result:="加载mock数据失败"

	if contents, err := ioutil.ReadFile(path); err == nil{
		result= strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("Use ioutil.ReadFile to read a file:\n", result)
	}else{
		fmt.Println("Use ioutil.ReadFile >>> null")
	}


	// 返回字符串 "Hello world"
	fmt.Fprint(rw, result)

}