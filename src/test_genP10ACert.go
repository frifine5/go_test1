
package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"unsafe"
	"math/rand"
	"strings"
)



func SamplePost() {
	/*
	测试：516438BF78454DB487B9B3A6EB7F48B6
正式：0799597FC8F64429B5D65F29178203FD
authuserid
	 */
	authUserId := "5D49044486FA45F5AD006CCCBE9090B8"
	platformName := "统一印章管理平台"

	p10fp := "C:\\Users\\49762\\Desktop\\tmp.txt"
	sign_p10 := ""
	if contents, err := ioutil.ReadFile(p10fp); err == nil{
		sign_p10 = strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("读取p10内容为:\n", sign_p10)
	}else{
		fmt.Println("未找到p10文件/或者p10内容为空")
	}

	//sign_p10 = "MIHvMIGWAgEAMDYxCzAJBgNVBAYTAkNOMScwJQYDVQQDDB7liLbnq6Dns7vnu5/nlJ/kuqfnjq/looPor4HkuaYwWTATBgcqhkjOPQIBBggqgRzPVQGCLQNCAAS8apUOJr1Qz+6s9EB3SpfL4e7NdX6xiHcSyzBLQpapg13yeXaJILIlHOBzHFfmjLQh7jBFGD4e2m0HJzeeScYAMAoGCCqBHM9VAYN1A0gAMEUCIQD7L2JY+65pR/HAeC/gryDdXC0LF75efvATCCSXxfajJAIgAa7q/vhPeLbff3mFqjr/o5qMNPSgysVBkiTwVT5MXjM="

	song := make(map[string]interface{})
	rnd1 := fmt.Sprintf("%08d", rand.Int31())
	song["creditCode"] = "91110108802094766H"  // 使用单位统一社会信用代码
	song["unitName"] = "x公司"+ rnd1    // 单位名称
	song["unitAddress"] = "北京市海淀区首体南路22号楼" // 单位地址
	song["legalName"] = "某军"   // 法人姓名
	song["legalID"] = "110101198003074333" // 法人公民身份号码
	song["legalPhone"] = "13332456789"  // 法人手机号码
	song["provinceCode"] = "110108"    // 省市区编码
	song["prefecturalLevelCity"] = "北京市"    // 地级市(中文)
	song["STProperty"] = "北京市"      // 证书ST属性值如：湖北省
	song["unitProperty"] = "信息安全技术有限公司"    // 证书OU属性值：机关单位/事业/企业（中文）
	song["esID"] = "110108" + rnd1    // 印章唯一赋码
	song["p10"] = sign_p10 // 签名证书申请CSR内容，base64编码
	song["doubleP10"] = sign_p10   // 加密证书CSR内容，base64编码。可以与签名证书一样，加密证书的密钥对最终由KMC生成
	song["authUserId"] = authUserId
	song["platformName"] =  platformName

	bytesData, err := json.Marshal(song)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://221.232.224.75:9080/hbcaLCA/api/v3/certApply.do"
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

	if ioutil.WriteFile("/home/ca/wuhan/go_test_response", respBytes, 0644) == nil{
		fmt.Println("use ioutil.WriteFile write success")
	}


}



func main(){


	SamplePost()


}