package main // 定义当前文件的包名

// 引入包，引入需要使用的其它包
import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"net/url"
	"time"
	//"encoding/json"
)

func main(){	// 主函数，一般为程序的主入口（启动时第一个执行的函数），如果有init()函数，会先执行init
	fmt.Println("main start...")	// 表达式/执行语句
	fmt.Println("--------------------->")

	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()

	// process to-do:


	//GetData()
	httpPost()
}




// 单行注释
/*
	多行注释
*/


/*
当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
*/

func GetData() {
	client := &http.Client{}
	resp, err := client.Get("http://www.baidu.com/")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}


/*
post 有三种
*/

// http.post请求：

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}


// http.postForm:
func PostData() {
	//client := &http.Client{}
	resp, err := http.PostForm("https://www.pgyer.com/apiv2/app/view", url.Values{"appKey": {"62c99290f0cb2c567cb153c1fba75d867e"},
		"_api_key": {"584f29517115df2034348b0c06b3dc57"}, "buildKey": {"22d4944d06354c8dcfb16c4285d04e41"}})
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}

// http.do
func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}


