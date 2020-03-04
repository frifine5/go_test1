package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"bufio"
	"os"
)


type Seal struct {
	sealName        sql.NullString `db:"SEAL_NAME"`
	sealType         sql.NullString `db:"SEAL_TYPE_CODE"`
	sealTypeName     sql.NullString `db:"SEAL_TYPE_NAME"`
	sealCode         sql.NullString `db:"SEAL_CODE"`
	taskId           sql.NullString `db:"TASK_ID"`
	recordStatusName sql.NullString `db:"RECORD_STATUS"`
	sealStatus       sql.NullString `db:"SEAL_STATUS"`
	pubStatus        sql.NullString `db:"SEAL_PUBLISH_STATUS"`
	applyUnitName    sql.NullString `db:"APPLY_UNIT_NAME"`
	useUnitName      sql.NullString `db:"USE_UNIT_NAME_CN"`
	useUnitCode      sql.NullString `db:"USE_UNIT_CODE"`
	madeTime         sql.NullString `db:"MAKE_SEAL_TIME"`
	validStart       sql.NullString `db:"VALID_START_TIME"`
	validEnd         sql.NullString `db:"VALID_END_TIME"`
	operator         sql.NullString `db:"OPERATOR"`
}

func sealPrepareQuery(db *sql.DB, mount int) []Seal {
	stmt, err := db.Prepare("select SEAL_NAME , SEAL_TYPE_CODE , SEAL_TYPE_NAME , SEAL_CODE , TASK_ID , RECORD_STATUS , SEAL_STATUS ,  SEAL_PUBLISH_STATUS , APPLY_UNIT_NAME , USE_UNIT_NAME_CN , USE_UNIT_CODE , MAKE_SEAL_TIME , VALID_START_TIME  , VALID_END_TIME ,  OPERATOR from makeseal_seal") //" limit ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	defer rows.Close()


	seals := make([]Seal, 0)
	var i int = 0
	for rows.Next() {
		var seal Seal
		err := rows.Scan(&seal.sealName, &seal.sealType, &seal.sealTypeName, &seal.sealCode, &seal.taskId, &seal.recordStatusName, &seal.sealStatus, &seal.pubStatus, &seal.applyUnitName, &seal.useUnitName, &seal.useUnitCode, &seal.madeTime, &seal.validStart, &seal.validEnd, &seal.operator)
		if err != nil {
			panic(err)
			continue
		}
		//fmt.Printf("seal: %#v\n", seal)
		seals = append(seals, seal)
		i++
	}

	return seals
}

func main() {

	fmt.Println("请输入数据库配置,格式 >>>\t用户名:密码@tcp(数据库地址:端口)/库名")
	var cofg1 string
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		cofg1 = in.Text()
		fmt.Println("input data = " + cofg1)
	} else {
		fmt.Println("input err and exit")
		return
	}
	dns := cofg1
	//dns = "root:123456@tcp(mysql.gmdev.baiwang-inner.com:3306)/seal_make_sm2"

	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	seals := sealPrepareQuery(db, 10)
	fmt.Printf("共查询到%d条记录\n", len(seals))
/*	for  i, e := range seals{
		fmt.Printf("seal: %#v, %#v\n",i, e)
	}*/

	xlFile := xlsx.NewFile()

	//获取行数
	shLen := len(xlFile.Sheets)
	fmt.Printf("now file sheet = %d\n", shLen)
	if (shLen < 1) {
		// add sheet
		xlFile.AddSheet("record")
		println("add sheet")
	}
	sheet2 := xlFile.Sheets[0]
	rowTitle := sheet2.AddRow()
	cell0 := rowTitle.AddCell()
	cell0.SetValue("序号")
	cell1 := rowTitle.AddCell()
	cell1.SetValue("印章名称")
	cell2 := rowTitle.AddCell()
	cell2.SetValue("印章类型编号")
	cell3 := rowTitle.AddCell()
	cell3.SetValue("印章类型")
	cell4 := rowTitle.AddCell()
	cell4.SetValue("印章编码")
	cell5 := rowTitle.AddCell()
	cell5.SetValue("备案编码")
	cell6 := rowTitle.AddCell()
	cell6.SetValue("备案状态")
	cell7 := rowTitle.AddCell()
	cell7.SetValue("印章状态")
	cell8 := rowTitle.AddCell()
	cell8.SetValue("发布状态")
	cell9 := rowTitle.AddCell()
	cell9.SetValue("上级单位")
	cell10 := rowTitle.AddCell()
	cell10.SetValue("用章单位名称")
	cell11 := rowTitle.AddCell()
	cell11.SetValue("统一社会信用代码")
	cell12 := rowTitle.AddCell()
	cell12.SetValue("制作时间")
	cell13 := rowTitle.AddCell()
	cell13.SetValue("启用时间")
	cell14 := rowTitle.AddCell()
	cell14.SetValue("终止时间")
	cell15 := rowTitle.AddCell()
	cell15.SetValue("经办人")


	nt := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(nt)

	for  i, seal := range seals{
		//fmt.Printf("seal: %#v, %#v\n",i, seal)
		eachRow := sheet2.AddRow()
		cell0 := eachRow.AddCell()
		num := i + 1
		cell0.SetValue(num)
		cell1 := eachRow.AddCell()
		cell1.SetValue(seal.sealName.String)
		cell2 := eachRow.AddCell()
		cell2.SetValue(seal.sealType.String)
		cell3 := eachRow.AddCell()
		cell3.SetValue(seal.sealTypeName.String)
		cell4 := eachRow.AddCell()
		cell4.SetValue(seal.sealCode.String)
		cell5 := eachRow.AddCell()
		cell5.SetValue(seal.taskId.String)
		cell6 := eachRow.AddCell()
		cell6.SetValue(seal.recordStatusName.String)
		cell7 := eachRow.AddCell()
		cell7.SetValue(seal.sealStatus.String)
		cell8 := eachRow.AddCell()
		cell8.SetValue(seal.pubStatus.String)
		cell9 := eachRow.AddCell()
		cell9.SetValue(seal.applyUnitName.String)
		cell10 := eachRow.AddCell()
		cell10.SetValue(seal.useUnitName.String)
		cell11 := eachRow.AddCell()
		cell11.SetValue(seal.useUnitCode.String)
		cell12 := eachRow.AddCell()
		cell12.SetValue(seal.madeTime.String)
		cell13 := eachRow.AddCell()
		cell13.SetValue(seal.validStart.String)
		cell14 := eachRow.AddCell()
		cell14.SetValue(seal.validEnd.String)
		cell15 := eachRow.AddCell()
		cell15.SetValue(seal.operator.String)
	}

	var xlsFp2 = "./seallist.xlsx"
	xlFile.Save(xlsFp2)

	fmt.Println("finish business. OutFile is './seallist.xlsx'")

}
