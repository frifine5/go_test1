package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
)

type Result struct {
	ErrorCode int  `json:"errorCode"`
	Message string `json:"message"`
	Data  string `json:"data"`
}

func main() {
	//第一个参数是接口名，第二个参数 http handle func
	http.HandleFunc("/mock/", getMockData)
	//服务器要监听的主机地址和端口号
	http.ListenAndServe(":8082", nil)
}

// http handle func
func getMockData(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("get on time: "+ (time.Now()).String())

	path := req.RequestURI
	path = "/mnt" + path

	rs := Result{

	}
	if contents, err := ioutil.ReadFile(path); err == nil{
		result:= strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use ioutil.ReadFile to read a file:\n", result)
		rs.ErrorCode = 0
		rs.Message = "ok"
		rs.Data = result
	}else{
		fmt.Println("Use ioutil.ReadFile >>> null")
		rs.ErrorCode = -1
		rs.Message = "加载mock数据失败"
	}


	// 返回字符串 "Hello world"
	fmt.Fprint(rw, rs)

}