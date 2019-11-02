package main

import (
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"time"
	"github.com/tealeg/xlsx"
	"strconv"
)

type rs1 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
	s :="{\"code\": -1,\"message\": \"解析失败\"}"

	rtn :=  rs1{}
	err := json.Unmarshal( []byte(s), &rtn)
	if err !=nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("code=%d, message=%s\n",  rtn.Code, rtn.Message)
 */

func req(key string) *rs1 {

	url0 := "http://localhost:10090/makesealapi/api/gdata/pdf/oldprehash?"
	u, _ := url.Parse(url0)
	q := u.Query()
	q.Set("fileId", key)
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return nil

	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil

	}
	fmt.Printf("%s", result)
	rtn := rs1{}
	err = json.Unmarshal([]byte(result), &rtn)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &rtn

}

func t() *rs1 {

	s := "{\"code\": -1,\"message\": \"解析失败\"}"

	rtn := rs1{}
	err := json.Unmarshal([]byte(s), &rtn)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &rtn
}

func wt2Xlsx(sheet *xlsx.Sheet, xlFile *xlsx.File, fp string, rs *rs1, key string) {
	if rs == nil {
		return
	}
	row1 := sheet.AddRow()
	cell1 := row1.AddCell()
	cell1.SetValue(key)
	cell2 := row1.AddCell()
	cell2.SetValue(rs.Code)
	cell3 := row1.AddCell()
	nt := time.Now().Format("2006-01-02 15:04:05")
	cell3.SetValue(nt)
	fmt.Println(nt)
	xlFile.Save(fp)

}

func getBatchArr(xlFile *xlsx.File, st int, mount int) []string {

	resourceArr := make([]string, mount)
	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		//遍历每一行
		for i := 0; i < mount; i++ {
			row := sheet.Rows[st+i]
			//遍历每一个单元
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					//如果是每一行的第一个单元格
					if cellIndex == 0 {
						resourceArr[i] = text
					}
				}
			}
		}
	}
	return resourceArr
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func doOneFile(fp0 string, fp string) {
	// 读取xlsx获取批执行的key列表
	st := 0
	xlFile0, err := xlsx.OpenFile(fp0)
	if err != nil {
		println(err)
		return
	}
	//获取行数
	length := len(xlFile0.Sheets[0].Rows)
	fmt.Printf("xlsx共%d行\n", length)

	// 记录
	xlFile, err := xlsx.OpenFile(fp)
	if err != nil {
		println("open " + fp + " err")
		return
	}

	sht0 := xlFile.Sheets[0]
	fmt.Println(sht0)

	ti := If(length%100 == 0, length/100, length/100+1).(int)
	// start index
	for a := 0; a < ti; a++ {
		st = 100 * a
		mount := If(a == ti-1, length%100, 100).(int)
		if length == 100 {
			mount = 100
		}
		keys := getBatchArr(xlFile0, st, mount)

		// 循环间隔调用http接口

		for _, k := range keys {
			if k == "" {
				continue
			}


			time.Sleep(time.Millisecond * 100)
			fmt.Println(k)
			rtn := req(k) // 解析返回结构并存储到记录表中
			if rtn == nil {
				fmt.Println("空对象")
			}
			fmt.Printf("code=%d, message=%s\n", rtn.Code, rtn.Message)
			wt2Xlsx(sht0, xlFile, fp, rtn, k)

		}

	}
	// 当页输入完毕后，添加两个空行

	row1 := sht0.AddRow()
	cell1 := row1.AddCell()
	cell1.SetValue(fp0 + " - 操作结束，共" + strconv.Itoa(length) + "行")
	cell2 := row1.AddCell()
	cell2.SetValue(0)
	cell3 := row1.AddCell()
	nt := time.Now().Format("2006-01-02 15:04:05")
	cell3.SetValue(nt)
	row2 := sht0.AddRow()
	cell21 := row2.AddCell()
	cell21.SetValue("--")
	xlFile.Save(fp)

}

func getFileList(dir string) []string {
	list2, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	flist := []string{}
	for _, v := range list2 {
		if !v.IsDir() {
			flist = append(flist, v.Name())
		}
	}
	fmt.Printf("共%d个文件", len(flist))
	for _, v := range flist {
		fmt.Printf(v + "  ")
	}
	return flist
}

func main() {

	fpdir := "list"
	rndfile := "record.xlsx"
	flist := getFileList(fpdir)
	if flist == nil {
		return
	}

	for _, fp0 := range flist {
		fp0 = fpdir + "/" + fp0
		doOneFile(fp0, rndfile)
		if ioutil.WriteFile(fpdir+fp0+"-finish-record",
			[]byte(  "finish time: "+ time.Now().Format("2006-01-02 15:04:05")), 0644) == nil {
			fmt.Println("\nfinish one file and record success...")
		}
		time.Sleep(time.Second * 10) // 歇息10秒
	}

}
