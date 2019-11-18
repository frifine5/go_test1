package main

// 引入依赖包
import (
	"fmt"
	"obs"
	"os"
	"bufio"
)

var ak = "OM65GZAEBJQAX9EFZVAX"
var sk = "d4VGQ3gqxh8uWtbN18wKyUIhe1cbF16PAW2l6MAJ"
var endpoint = "http://172.20.2.10:80"
var bucket = "obs-dc-dzqz"

var config = obs.WithConnectTimeout(15)

// 创建ObsClient结构体
var obsClient, _ = obs.New(ak, sk, endpoint, config)

func main() {
	println("start ...")

	download()
	//listBkt()

	//listObjInBkt("obs-dc-dzqz")

	println("... end.")
}


func download(){
	var key string
	fmt.Println("请输入要下载的obs文件名")
	//使用os.Stdin开启输入流
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		key = in.Text()
		fmt.Println("input data = " + key)
	} else {
		fmt.Println("input err and exit")
		return
	}

	input := &obs.GetObjectInput{}
	input.Bucket = bucket
	input.Key = key

	// 下载对象
	output, err := obsClient.GetObject(input)
	if err == nil {
		defer output.Body.Close()
		// 获取对象自定义元数据
		//fmt.Printf("Metadata:%v\n", output.Metadata)
		p := make([]byte, 1024)
		var readErr error
		var readCount int

		// 读取对象内容
		for {
			readCount, readErr = output.Body.Read(p)
			if readErr != nil {
				fmt.Println(readErr)
				fmt.Println("读取出错")
				return
			}
			if readCount > 0 {
				//fmt.Printf("%s", p[:readCount])
				obsWrite2( key, p)
			}

		}
		fmt.Printf("down %s finish\n", key)
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Printf("Code:%s\n", obsError.Code)
		fmt.Printf("Message:%s\n", obsError.Message)
	} else {
		fmt.Println(err)
	}
}


func listBkt(){
	// 列举桶
	input := &obs.ListBucketsInput{}
	input.QueryLocation = true
	output, err := obsClient.ListBuckets(nil)
	if err == nil {
		fmt.Printf("Owner.ID:%s\n", output.Owner.ID)
		for index, val := range output.Buckets {
			fmt.Printf("Bucket[%d]-Name:%s,CreationDate:%s,Location:%s\n", index, val.Name, val.CreationDate, val.Location)
		}
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println(obsError.Code)
		fmt.Println(obsError.Message)
	} else {
		fmt.Println(err)
	}
}


func listObjInBkt(bkt string){

	input := &obs.ListObjectsInput{}
	input.Bucket = bkt
	input.MaxKeys = 100
	output, err := obsClient.ListObjects(input)
	if err == nil {
		for index, val := range output.Contents {
			lineCont := fmt.Sprintf( "Content[%d]-OwnerId:%s, ETag:%s, Key:%s, LastModified:%s, Size:%d, StorageClass:%s\n",
				index, val.Owner.ID, val.ETag, val.Key, val.LastModified, val.Size, val.StorageClass)
			fmt.Println(lineCont)
			obsWrite1("rsList", lineCont)
		}
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Printf("Code:%s\n", obsError.Code)
		fmt.Printf("Message:%s\n", obsError.Message)
	}
}

func obsWrite1(fp string, content string) {

	fileObj, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) //  |os.O_TRUNC
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)// 应用进程退出
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content + "\n"); err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.")
	}
}
func obsWrite2(fp string, ctx []byte) {

	fileObj, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) //  |os.O_TRUNC
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)// 应用进程退出
	}
	defer fileObj.Close()
	if _, err := fileObj.Write(ctx); err == nil {
	}
}

