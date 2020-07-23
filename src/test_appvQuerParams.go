
package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"unsafe"
	"strconv"
)


func ApprvPost() {

	// 登录
	lgi := make(map[string]interface{})
	lgi["passWord"] = "ddb5dae482528387954e1a3bc7af2ca3"
	//lgi["userAccount"] = "宜宾市质监局"
	lgi["userAccount"] = "宜宾市驻京联络处"


	bytesData, err := json.Marshal(lgi)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://10.100.4.50:10088/user/webLogin"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	//if ioutil.WriteFile("C:\\Users\\49762\\Desktop\\go_test_response", respBytes, 0644) == nil{
	//	fmt.Println("use ioutil.WriteFile write success")
	//}

	// {"code":0,"message":"操作成功","data":
	// {"access_token":"95538881-699a-47e2-9d1c-8e46da9ada56","phone":14,"userAccount":14,"id":14,"userName":14},
	// "method":null,"requestId":null}
	resq :=  Resq{}
	err2 := json.Unmarshal(respBytes, &resq)
	if err2!=nil{
		fmt.Println(err2)
		return
	}

	accessToken := resq.Data.AccessToken
	fmt.Println("accessToken = " + accessToken)


}

type ResqDataBody struct {
	AccessToken string `json:"access_token"`
	Phone int
	UserAccount int
	UserName int
}

type Resq struct {
	Code int
	Message string
	Data ResqDataBody `json:"data"`
	Method string
	RequestId string
}

type Class struct{
	Name string
	Grade int
}

func main(){

	//ApprvPost()

 	s :="[{\"Name\":\"1班\",\"Grade\":1},{\"Name\":\"2班\",\"Grade\":1},{\"Name\":\"3班\",\"Grade\":1},{\"Name\":\"1班\",\"Grade\":2},{\"Name\":\"2班\",\"Grade\":2}]";

 	clss :=  []*Class{}
 	err := json.Unmarshal( []byte(s), &clss)
 	if err !=nil{
		fmt.Println(err)
		return
	}

	for i:=0; i< len(clss); i++{
		fmt.Print( strconv.Itoa(clss[i].Grade) +"年级" + clss[i].Name  + "\t")
	}

	fmt.Println()
	s2, e := json.Marshal(clss)
	if e !=nil{
		fmt.Println(e)
	}else{
		fmt.Println(string(s2))
	}

}