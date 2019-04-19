package main

import(
	"fmt"

	"io/ioutil"
	"strings"
	"os"
	"strconv"
	"bufio"
)


func main(){
	fmt.Println("测试文件读写")
	fmt.Println()

	f1p := "/home/tmp"
	ftxt := "测试写文件的一行数据___+_+!___abcdefghijklmnopqrstuvwxyz0123456789."

	//TestWrite1(f1p, ftxt)
	TestWrite2(f1p, ftxt)

	fmt.Println("测试写完成>>>>>>>>>>>>>>读下数据")
	CommonRead(f1p)

}

func CommonRead(fp string){
	// ioutil.ReadFile
	if contents, err := ioutil.ReadFile(fp); err == nil{
		result := string(contents)
		//result = strings.Replace(string(contents), "\n", "", 1)

		fmt.Println("Use ioutil.ReadFile to read a file:\n", result)
	}else{
		fmt.Println("Use ioutil.ReadFile >>> null")
	}
}

func TestRead(fp string){
	// ioutil.ReadFile
	if contents, err := ioutil.ReadFile(fp); err == nil{
		result:= strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use ioutil.ReadFile to read a file:\n", result)
	}else{
		fmt.Println("Use ioutil.ReadFile >>> null")
	}

	// osutil
	if fileObj,err := os.Open(fp);err == nil {
		//if fileObj,err := os.OpenFile(name,os.O_RDONLY,0644); err == nil {
		defer fileObj.Close()
		if contents,err := ioutil.ReadAll(fileObj); err == nil {
			result := strings.Replace(string(contents),"\n","",1)
			fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file :\n",result)
		}
	}else{
		fmt.Println("Use os.Open family functions and ioutil.ReadAll >>> null")
	}

	// FileRead
	if fileObj,err := os.Open(fp);err == nil {
		defer fileObj.Close()
		//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
		buf := make([]byte, 1024)
		if n,err := fileObj.Read(buf);err == nil {
			fmt.Println("The number of bytes read:"+strconv.Itoa(n),"Buf length:"+strconv.Itoa(len(buf)))
			result := strings.Replace(string(buf),"\n","",1)
			fmt.Println("Use os.Open and File's Read method to read a file:\n",result)
		}else{
			fmt.Println("Use os.Open family functions and File's Read >>> null")
		}
	}else{
		fmt.Println("Use os.Open family functions and File's Read >>> null")
	}

	// BufioRead
	if fileObj,err := os.Open(fp);err == nil {
		defer fileObj.Close()
		//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
		reader := bufio.NewReader(fileObj)
		//使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
		if result,err := reader.ReadString(byte('.'));err == nil {
			fmt.Println("使用ReadString by delim相关方法读取内容:\n",result)
		}
		//注意:上述ReadString已经将buffer中的数据读取出来了，下面将不会输出内容
		//需要注意的是，因为是将文件内容读取到[]byte中，因此需要对大小进行一定的把控
		buf := make([]byte,1024)
		//读取Reader对象中的内容到[]byte类型的buf中
		if n,err := reader.Read(buf); err == nil {
			fmt.Println("The number of bytes read:"+strconv.Itoa(n))
			//这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
			fmt.Println("Use bufio.NewReader and os.Open read file contents to a []byte:\n",string(buf))
		}
	}else{
		fmt.Println("Use bufio.NewReader and os.Open read file >>> null")
	}

}


// ioutil.WriteFile 覆盖原内容
func TestWrite1(fp string, content string){
	data := []byte(content)
	// ioutil.WriteFile
	if ioutil.WriteFile(fp, data, 0644) == nil{
		fmt.Println("use ioutil.WriteFile write success")
	}
}


func TestWrite2(fp string, content string){

	fileObj,err := os.OpenFile(fp,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644) //  |os.O_TRUNC
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _,err := fileObj.WriteString(content+"\n");err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.")
	}

	data := []byte(content +"\n")
	if _,err := fileObj.Write(data);err == nil {
		fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.")
	}

}

















