package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"time"
	"strconv"
)

func  pxls2Arr(fp string ) []string{

	xlFile, err := xlsx.OpenFile(fp)
	if err != nil{
		println("open err")
		return nil
	}
	//获取行数
	length := len(xlFile.Sheets[0].Rows)
	fmt.Printf("xlsx共%d行\n", length)
	//开辟除表头外的行数的数组内存
	resourceArr := make([]string, length)
	//遍历sheet
	for _, sheet := range xlFile.Sheets {

		//遍历每一行
		for rowIndex, row := range sheet.Rows {
			//跳过第一行表头信息
	/*		if rowIndex == 0 {
				// for _, cell := range row.Cells {
				//  text := cell.String()
				//  fmt.Printf("%s\n", text)
				// }
				continue
			}
	*/

			//遍历每一个单元
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					//如果是每一行的第一个单元格
					if cellIndex == 0 {
						resourceArr[rowIndex] = text
					}
				}
			}
		}
	}
/*
	fmt.Println(len(resourceArr))
	for i,x:= range resourceArr {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	*/

	return resourceArr
}


func addaWrite(fp string){
	xlFile, err := xlsx.OpenFile(fp)
	if err != nil{
		println("open err")
		return
	}
	//获取行数

	shLen := len(xlFile.Sheets)
	fmt.Printf(	"now file sheet = %d\n",shLen)
	if(shLen<1){
		// add sheet
		xlFile.AddSheet("record")
		println("add sheet")
	}

	//xlFile, err = xlsx.OpenFile(fp)
	sheet2 := xlFile.Sheets[0]
	row1 := sheet2.AddRow()
	cell1 := row1.AddCell()
	cell1.SetValue("testBlkObjKey")
	cell2 := row1.AddCell()
	cell2.SetValue( 0 )
	cell3 := row1.AddCell()
	nt := time.Now().Format("2006-01-02 15:04:05")
	cell3.SetValue( nt )
	fmt.Println(nt)
	xlFile.Save(fp)



}


func loopWrite(fp string){

	xlFile, err := xlsx.OpenFile(fp)
	if err != nil {
		fmt.Println(err)

		return
	}
	shLen := len(xlFile.Sheets)
	fmt.Printf("xlsx 共%d页", shLen)

	sheet := xlFile.Sheets[0]
	if sheet == nil {
		sheet, err =xlFile.AddSheet("record")
		if err != nil {
			fmt.Println(err)

		}

	}

	for i := 0; i < 10; i++ {
		w1(sheet, xlFile, fp, i)
	}



}

func w1(sheet *xlsx.Sheet, xlFile *xlsx.File, fp string, i int){
	row1 := sheet.AddRow()
	cell1 := row1.AddCell()
	cell1.SetValue("testBlkObjKey-" + strconv.Itoa(i))
	cell2 := row1.AddCell()
	cell2.SetValue( i )
	cell3 := row1.AddCell()
	nt := time.Now().Format("2006-01-02 15:04:05")
	cell3.SetValue( nt )
	fmt.Println(nt)
	xlFile.Save(fp)
	time.Sleep(time.Millisecond * 500)
}


func main() {

	var xlsFp = "C:\\Users\\49762\\Desktop\\1.xlsx"
	//var arr = pxls2Arr(xlsFp)

	// 读写
	//addaWrite(xlsFp)

	//loopWrite(xlsFp)


	fmt.Println(xlsFp)

	nt := time.Now().Format("20060102150405")
	fmt.Println(nt)




}
