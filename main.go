package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	. "excelfix/models"
	"os"
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
	"regexp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tealeg/xlsx"
	"bufio"
)

// ORM 引擎

var engine *xorm.Engine

func init() { //初始化数据库

	// 创建 ORM 引擎与数据库
	var err error
	//engine, err = xorm.NewEngine("mysql", "root:oyhzabcd@/excel?charset=utf8")
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	checkError(err, "Engine Error!")

	//同步结构体与数据表
	err = engine.Sync2(new(GuangZhouSrc), new(NanChangSrc), new(CalBianMa), new(CalCompany))

	checkError(err, "Fail to sync database!")

}

func main() {

	fmt.Println("读取广州数据...")
	readFromExcelGuangZhou()

	fmt.Println("读取南昌数据...")
	readFromExcelNanchang()

	fmt.Println("合并类似公司名称...")
	mergerSimilar()

	fmt.Println("清理编码统计表...")
	engine.Where("1=1").Delete(new(CalBianMa))
	//
	fmt.Println("统计南昌数据...")
	calBianMaNanChang()

	fmt.Println("统计广州数据...")
	calBianMaGuangZhou()

	fmt.Println("按公司统计数据...")
	calCompany()

	fmt.Println("导出Excel...")
	output2Excel("CalBianMa")
	output2Excel("CalCompany")

	fmt.Println("按回车键继续...")
	inputReader := bufio.NewScanner(os.Stdin)
	inputReader.Scan()
	fmt.Println(inputReader.Text())

	defer engine.Close()

}

func mergerSimilar(){
	sql := []string{
		"update guang_zhou_src set A21='台湾东电化贸易股份有限公司' where A21='台湾东电化贸易股份有限公司(TDK TAIWAN ELECTRONICS CORP.)';",

		"update guang_zhou_src set A21='库力索法半导体(苏州)有限公司' where A21='库力索法半导体（苏州）有限公司';",

		"update guang_zhou_src set A21='南央国际贸易（上海）有限公司' where A21='南央国际贸易(上海)有限公司';",

		"update guang_zhou_src set A21='先域微电子技术服务(上海)有限公司深圳分公司' where A21='先域微电子技术服务（上海）有限公司深圳分公司';",

		"update guang_zhou_src set A21='世蕴电子C-ON TECH.CO.,LTD' where A21='世蕴电子C-ON TECH.CO. LTD';",

		"update guang_zhou_src set A21='SDK Co., Ltd.' where A21='SDK CO. LTD';",

		"update guang_zhou_src set A21='HYVISION SYSTEM Inc.' where A21='HyVISION System Inc';",

		"update guang_zhou_src set A21='AP-Tech Co.,Ltd' where A21='AP Tech Co.  Ltd.';",

		"update nan_chang_src set A10='台湾东电化贸易股份有限公司' WHERE A10='台湾东电化股份有限公司';",

		"update guang_zhou_src set A21='台湾东电化贸易股份有限公司' WHERE A21='台湾东电化股份有限公司';",
	}

	for _, y := range sql {
		engine.Exec(y)
	}
}

func readFromExcelGuangZhou() {
	//excel read start
	xlsx, err := excelize.OpenFile("./GuangZhou.xlsx")
	checkError(err, "read xls error")

	rows := xlsx.GetRows("Sheet1")

	engine.Where("1=1").Delete(new(GuangZhouSrc)) //清空数据表

	for k, row := range rows {
		if (k == 0) { //跳过标题行
			continue
		}
		excelMap := make(map[string]interface{}, 0)
		for k, colCell := range row { //生成数据Map
			stringk := strconv.Itoa(k + 1)
			excelMap["A"+stringk] = colCell

		}

		//修改A5的数据类型

		dateRegexp := regexp.MustCompile(`^([\d]+)/([\d]+)/([\d]+) ([\d]+):([\d]+)$`)
		dateParase := dateRegexp.FindStringSubmatch(excelMap["A5"].(string))
		//fmt.Println(dateParase)

		excelMap["Month"] = dateParase[1]
		//fmt.Println(reflect.TypeOf(excelMap["Month"]))

		if (len(dateParase[1]) == 1) { //月份日期补位
			dateParase[1] = "0" + dateParase[1]
		}
		if (len(dateParase[2]) == 1) {
			dateParase[2] = "0" + dateParase[2]
		}
		//整理字符串格式
		dateFormat := "20" + dateParase[3] + "-" + dateParase[1] + "-" + dateParase[2] + " " + dateParase[4] + ":" + dateParase[5] + ":00"

		//local, _ := time.LoadLocation("Local")
		excelMap["A5"], err = time.Parse("2006-01-02 15:04:05", dateFormat) //生成time格式日期
		//其中layout的时间必须是"2006-01-02 15:04:05"这个时间，不管格式如何，时间点一定得是这个，如："Jan 2, 2006 at 3:04pm (MST)"，"2006-Jan-02"等。如换一个时间解析出来的时间就不对了，要特别注意这一点。
		checkError(err, "String to time error!")

		//fmt.Println(excelMap)

		//准备写入数据库
		srcExcl := &GuangZhouSrc{}
		err = srcExcl.FillStruct(excelMap) //map转换成struct   struct名称不能有中文，数字不得在首位，不得有- _之类的连接符、下划线等
		checkError(err, "GuangZhouSrc map2struct error")

		_, err = engine.Insert(srcExcl)
		fmt.Println("GuangZhou inserting...", k)
		checkError(err, "Company insert error")
		//fmt.Println(affected)
		//fmt.Println()
	}
	//excel read end
}

func readFromExcelNanchang() {
	//excel read start
	xlsx, err := excelize.OpenFile("./Nanchang.xlsx")
	checkError(err, "read xls error")

	rows := xlsx.GetRows("Sheet1")
	//fmt.Println(len(rows))

	engine.Where("1=1").Delete(new(NanChangSrc))

	for k, row := range rows {
		if (k == 0) { //跳过标题行
			continue
		}
		excelMap := make(map[string]interface{}, 0)
		for k, colCell := range row {
			stringk := strconv.Itoa(k + 1)
			excelMap["A"+stringk] = colCell
		}

		//fmt.Println(excelMap)
		//准备写入数据库
		srcExcl := &NanChangSrc{}
		err := srcExcl.FillStruct(excelMap) //map转换成struct   struct名称不能有中文，数字不得在首位，不得有- _之类的连接符、下划线等
		checkError(err, "map2struct error")

		//fmt.Println("srcExcl:", srcExcl)

		_, err = engine.Insert(srcExcl)
		fmt.Println("NanChang inserting...", k)
		checkError(err, "Company insert error")

	}
	//excel read end
}

func getResult(tname string) []map[int]string {
	//创建字段中文映射表

	translateList := map[string]string{}
	translateList["Bianma"] = "存货编码"
	translateList["Gongyingshang"] = "供应商"
	translateList["Mingcheng"] = "存货名称"
	translateList["M1shuliang"] = "1月-数量"
	translateList["M1danjia"] = "1月-平均单价"
	translateList["M1caigoujine"] = "1月-采购金额"
	translateList["M1danjiajiangjiajine"] = "1月-单价降价金额"
	translateList["M1danjiajiangjiabili"] = "1月-单价降价比例"
	translateList["M1zongjiajiangjiajine"] = "1月-总价降价金额"
	translateList["M1zongjiajiangjiabili"] = "1月-总价降价比例"
	translateList["M2shuliang"] = "2月-数量"
	translateList["M2danjia"] = "2月-平均单价"
	translateList["M2caigoujine"] = "2月-采购金额"
	translateList["M2danjiajiangjiajine"] = "2月-单价降价金额"
	translateList["M2danjiajiangjiabili"] = "2月-单价降价比例"
	translateList["M2zongjiajiangjiajine"] = "2月-总价降价金额"
	translateList["M2zongjiajiangjiabili"] = "2月-总价降价比例"
	translateList["M3shuliang"] = "3月-数量"
	translateList["M3danjia"] = "3月-平均单价"
	translateList["M3caigoujine"] = "3月-采购金额"
	translateList["M3danjiajiangjiajine"] = "3月-单价降价金额"
	translateList["M3danjiajiangjiabili"] = "3月-单价降价比例"
	translateList["M3zongjiajiangjiajine"] = "3月-总价降价金额"
	translateList["M3zongjiajiangjiabili"] = "3月-总价降价比例"
	translateList["M4shuliang"] = "4月-数量"
	translateList["M4danjia"] = "4月-平均单价"
	translateList["M4caigoujine"] = "4月-采购金额"
	translateList["M4danjiajiangjiajine"] = "4月-单价降价金额"
	translateList["M4danjiajiangjiabili"] = "4月-单价降价比例"
	translateList["M4zongjiajiangjiajine"] = "4月-总价降价金额"
	translateList["M4zongjiajiangjiabili"] = "4月-总价降价比例"
	translateList["M5shuliang"] = "5月-数量"
	translateList["M5danjia"] = "5月-平均单价"
	translateList["M5caigoujine"] = "5月-采购金额"
	translateList["M5danjiajiangjiajine"] = "5月-单价降价金额"
	translateList["M5danjiajiangjiabili"] = "5月-单价降价比例"
	translateList["M5zongjiajiangjiajine"] = "5月-总价降价金额"
	translateList["M5zongjiajiangjiabili"] = "5月-总价降价比例"
	translateList["M6shuliang"] = "6月-数量"
	translateList["M6danjia"] = "6月-平均单价"
	translateList["M6caigoujine"] = "6月-采购金额"
	translateList["M6danjiajiangjiajine"] = "6月-单价降价金额"
	translateList["M6danjiajiangjiabili"] = "6月-单价降价比例"
	translateList["M6zongjiajiangjiajine"] = "6月-总价降价金额"
	translateList["M6zongjiajiangjiabili"] = "6月-总价降价比例"
	translateList["M7shuliang"] = "7月-数量"
	translateList["M7danjia"] = "7月-平均单价"
	translateList["M7caigoujine"] = "7月-采购金额"
	translateList["M7danjiajiangjiajine"] = "7月-单价降价金额"
	translateList["M7danjiajiangjiabili"] = "7月-单价降价比例"
	translateList["M7zongjiajiangjiajine"] = "7月-总价降价金额"
	translateList["M7zongjiajiangjiabili"] = "7月-总价降价比例"
	translateList["M8shuliang"] = "8月-数量"
	translateList["M8danjia"] = "8月-平均单价"
	translateList["M8caigoujine"] = "8月-采购金额"
	translateList["M8danjiajiangjiajine"] = "8月-单价降价金额"
	translateList["M8danjiajiangjiabili"] = "8月-单价降价比例"
	translateList["M8zongjiajiangjiajine"] = "8月-总价降价金额"
	translateList["M8zongjiajiangjiabili"] = "8月-总价降价比例"
	translateList["M9shuliang"] = "9月-数量"
	translateList["M9danjia"] = "9月-平均单价"
	translateList["M9caigoujine"] = "9月-采购金额"
	translateList["M9danjiajiangjiajine"] = "9月-单价降价金额"
	translateList["M9danjiajiangjiabili"] = "9月-单价降价比例"
	translateList["M9zongjiajiangjiajine"] = "9月-总价降价金额"
	translateList["M9zongjiajiangjiabili"] = "9月-总价降价比例"
	translateList["M10shuliang"] = "10月-数量"
	translateList["M10danjia"] = "10月-平均单价"
	translateList["M10caigoujine"] = "10月-采购金额"
	translateList["M10danjiajiangjiajine"] = "10月-单价降价金额"
	translateList["M10danjiajiangjiabili"] = "10月-单价降价比例"
	translateList["M10zongjiajiangjiajine"] = "10月-总价降价金额"
	translateList["M10zongjiajiangjiabili"] = "10月-总价降价比例"
	translateList["M11shuliang"] = "11月-数量"
	translateList["M11danjia"] = "11月-平均单价"
	translateList["M11caigoujine"] = "11月-采购金额"
	translateList["M11danjiajiangjiajine"] = "11月-单价降价金额"
	translateList["M11danjiajiangjiabili"] = "11月-单价降价比例"
	translateList["M11zongjiajiangjiajine"] = "11月-总价降价金额"
	translateList["M11zongjiajiangjiabili"] = "11月-总价降价比例"
	translateList["M12shuliang"] = "12月-数量"
	translateList["M12danjia"] = "12月-平均单价"
	translateList["M12caigoujine"] = "12月-采购金额"
	translateList["M12danjiajiangjiajine"] = "12月-单价降价金额"
	translateList["M12danjiajiangjiabili"] = "12月-单价降价比例"
	translateList["M12zongjiajiangjiajine"] = "12月-总价降价金额"
	translateList["M12zongjiajiangjiabili"] = "12月-总价降价比例"

	Exceldata := []map[int]string{}

	switch tname {
	case "CalCompany":

		var ResultList []CalCompany
		err := engine.Select("*").Find(&ResultList)
		checkError(err, "CalCompany reade error")

		//将标题写入Exceldata
		titleData := map[int]string{}
		titleSrc := reflect.TypeOf(ResultList[0])
		for i := 0; i < titleSrc.NumField(); i++ {
			titleData[i] = translateList[titleSrc.Field(i).Name] //string类型
		}
		Exceldata = append(Exceldata, titleData)

		for k := 1; k < len(ResultList); k++ { //将数据写入Exceldata
			rowData := map[int]string{}
			rowSrc := reflect.ValueOf(ResultList[k])
			//row = sheet.AddRow()
			for i := 0; i < rowSrc.NumField(); i++ {
				switch rowSrc.Field(i).Kind() { //判断数值类型
				case reflect.String:
					rowData[i] = rowSrc.Field(i).String() //reflect.Value 转 string
				case reflect.Int:
					rowData[i] = strconv.FormatInt(rowSrc.Field(i).Int(), 10) //reflect.Int 转 Int 再转string
				}
			}
			Exceldata = append(Exceldata, rowData)
			rowData = nil
		}

		//fmt.Println(Exceldata)
		return Exceldata

	case "CalBianMa":

		var ResultList []CalBianMa
		err := engine.Select("*").Find(&ResultList)
		checkError(err, "CalBianMa reade error")

		//将标题写入Exceldata
		titleData := map[int]string{}
		titleSrc := reflect.TypeOf(ResultList[0])
		for i := 0; i < titleSrc.NumField(); i++ {
			titleData[i] = translateList[titleSrc.Field(i).Name] //string类型
		}
		Exceldata = append(Exceldata, titleData)

		for k := 1; k < len(ResultList); k++ { //将数据写入Exceldata
			rowData := map[int]string{}
			rowSrc := reflect.ValueOf(ResultList[k])
			//row = sheet.AddRow()
			for i := 0; i < rowSrc.NumField(); i++ {
				switch rowSrc.Field(i).Kind() { //判断数值类型
				case reflect.String:
					rowData[i] = rowSrc.Field(i).String() //reflect.Value 转 string
				case reflect.Int:
					rowData[i] = strconv.FormatInt(rowSrc.Field(i).Int(), 10) //reflect.Int 转 Int 再转string
				}
			}
			Exceldata = append(Exceldata, rowData)
			rowData = nil
		}
		return Exceldata
	default:
		Exceldata = nil
		return Exceldata
	}
}

func output2Excel(tname string) {

	ResultList := getResult(tname)

	//初始化excel读写库
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	checkError(err, "add sheet error.")

	//先读出字段,并添加到excel文件
	titleData := ResultList[0]
	row = sheet.AddRow()
	for i := 1; i < len(titleData); i++ {
		cell = row.AddCell()
		cell.Value = titleData[i] //string
	}

	for v := 1; v < len(ResultList); v++ {
		valueDate := ResultList[v]
		row = sheet.AddRow()
		for i := 1; i < len(valueDate); i++ {
			cell = row.AddCell()
			valueConvert, err := strconv.ParseFloat(valueDate[i], 64)
			if err == nil {
				cell.SetFloat(valueConvert)
				//fmt.Println(err,"-OK-", valueConvert)
			} else {
				cell.Value = valueDate[i] //string
				//fmt.Println(err,":error:", valueConvert)
			}

		}
		valueDate = nil
	}

	fileName := "result" + tname + ".xlsx"
	err = file.Save(fileName) //写入excel文件
	checkError(err, fileName+" write faild.")
	fmt.Println(fileName, "> Output finished!")
}

//公司分类统计

func calCompany() {
	var err error

	var companyList []CalBianMa //不同的供应商列表

	engine.Where("1=1").Delete(new(CalCompany)) //清空公司分类统计表

	err = engine.Table(new(CalBianMa)).Select("distinct GongYingShang").Find(&companyList)
	checkError(err, "GongYingShang read error")

	for k, company := range companyList {

		calList := make([]CalBianMa, 0) //用Find方法，每次循环要进行清理，否则会累加
		err = engine.Table(new(CalBianMa)).Select("*").Where("GongYingShang = ?", company.Gongyingshang).Find(&calList)
		checkError(err, "calList read error")

		//fmt.Println("len(calList)", len(calList))

		companyTongJi := make(map[string]interface{}, 0)

		for _, detail := range calList {
			detailTitle := reflect.TypeOf(detail)  //数据库字段名称
			detailValue := reflect.ValueOf(detail) //数据库字段内容
			//detailValue.NumField()

			//fmt.Println(detail)
			companyTongJi["Gongyingshang"] = company.Gongyingshang

			for i := 4; i < detailValue.NumField(); i++ { //统计数据

				//fmt.Println(detailTitle.Field(i).Name)
				if (strings.Contains(detailTitle.Field(i).Name, "zongjia") == true || strings.Contains(detailTitle.Field(i).Name, "caigoujine") == true) && (strings.Contains(detailTitle.Field(i).Name, "bili") == false) { //有单价的跳过
					//只累加 采购金额、总价降价金额
					if _, ok := companyTongJi[detailTitle.Field(i).Name]; ok { //map中存在变量
						//fmt.Println("map类型：", reflect.TypeOf(companyTongJi[detailTitle.Field(i).Name]))
						//fmt.Println("参数类型：", reflect.TypeOf(detailValue.Field(i).Interface()))

						switch reflect.TypeOf(detailValue.Field(i).Interface()).Name() {
						case "string":
							float64num, err := strconv.ParseFloat(detailValue.Field(i).String(), 64)
							checkError(err, "string to float64 error")

							switch reflect.TypeOf(companyTongJi[detailTitle.Field(i).Name]).Name() {
							case "string":
								Float64num2, err := strconv.ParseFloat(companyTongJi[detailTitle.Field(i).Name].(string), 64)
								checkError(err, "string to float64 error 2")
								companyTongJi[detailTitle.Field(i).Name] = float64num + Float64num2
							case "float64":
								Float64num2, _ := companyTongJi[detailTitle.Field(i).Name].(float64)
								companyTongJi[detailTitle.Field(i).Name] = float64num + Float64num2
							}
							//fmt.Println(detail.Bianma,detailTitle.Field(i).Name,":",companyTongJi[detailTitle.Field(i).Name])

						case "int":

							companyTongJi[detailTitle.Field(i).Name] = int(detailValue.Field(i).Int()) + int(companyTongJi[detailTitle.Field(i).Name].(int))
						}

					} else { //不存在时候直接赋值
						companyTongJi[detailTitle.Field(i).Name] = detailValue.Field(i).Interface()
						//fmt.Println("Insert")
					}
				}
			}

		}

		for monthNum := 1; monthNum < 13; monthNum++ {
			mm := strconv.FormatInt(int64(monthNum), 10)
			if reflect.TypeOf(companyTongJi["M"+mm+"caigoujine"]).Name() == "float64" {
				if companyTongJi["M"+mm+"caigoujine"].(float64) > 0 {
					companyTongJi["M"+mm+"zongjiajiangjiabili"] = companyTongJi["M"+mm+"zongjiajiangjiajine"].(float64) / companyTongJi["M"+mm+"caigoujine"].(float64)

				} else {
					companyTongJi["M"+mm+"zongjiajiangjiabili"] = 0.00
				}
			} else {
				companyTongJi["M"+mm+"zongjiajiangjiabili"] = 0.00
			}
		}

		for x, y := range companyTongJi { //调整map值的类型与struct一致
			//fmt.Println(x,y)
			//fmt.Println(reflect.TypeOf(y))
			if reflect.TypeOf(y).Name() == "float64" {
				companyTongJi[x] = strconv.FormatFloat(companyTongJi[x].(float64), 'f', 4, 64)
			}
		}

		//fmt.Println(companyTongJi)

		//准备写入数据库
		companyResult := &CalCompany{}
		err := companyResult.FillStruct(companyTongJi) //map转换成struct   struct名称不能有中文，数字不得在首位，不得有- _之类的连接符、下划线等
		checkError(err, "map2struct error")

		companyTongJi = nil

		//fmt.Println("companyResult:", companyResult)

		oldCalResult := &CalCompany{}

		has, err := engine.Select("ID").Where("Gongyingshang =? ", companyResult.Gongyingshang).Get(oldCalResult) //判断是否有记录
		checkError(err, "Check BianMa error")

		//fmt.Println("oldcompanyResult:", oldCalResult)

		//os.Exit(1)
		var affected int64 = 0
		if has { //根据是否有记录情况进行更新或插入操作
			affected, err = engine.Id(oldCalResult.Id).Update(companyResult)
			fmt.Println("Company updating...", k, affected)
			checkError(err, "Company update error")
		} else {
			affected, err = engine.Insert(companyResult)
			fmt.Println("Company inserting...", k, affected)
			checkError(err, "Company insert error")
		}

	}
}

//按月统计 第二个表格

func calBianMaGuangZhou() {
	var err error

	var bianMaList []GuangZhouSrc //对应字段A2列表
	var cunHuoList []GuangZhouSrc //存货分月列表

	err = engine.Table(new(GuangZhouSrc)).Select("distinct A2").Find(&bianMaList)

	checkError(err, "Table GuangZhouSrc A12 error!")

	cunHuoBianMaTongJi := make(map[string]interface{})

	for k, bianMa := range bianMaList { //取出所有不同的 存货编码A2  进行统计

		//fmt.Println(bianMa.A2)

		cunHuoBianMaTongJi["Bianma"] = bianMa.A2

		var basePrice float64 = 0.00 //基期单价
		//

		for monthNum := 1; monthNum < 13; monthNum++ { //按月统计 存货编码A2 情况

			//startDate := "2018-" + strconv.Itoa(monthNum) + "-1"
			//endDate := "2018-" + strconv.Itoa(monthNum+1) + "-1"
			//
			//fmt.Println(endDate)

			//engine.ShowSQL(true)
			cunHuoList = make([]GuangZhouSrc, 0) //用Find方法，每次循环要进行清理，否则会累加
			err = engine.Table(new(GuangZhouSrc)).Select("A2,A4,A5,A7,A10,A13,A14,A18,A21,A31").
				Where("A2 = ?", bianMa.A2).
			//Where(" A5 > ?", startDate).
			//Where("A5 < ?", endDate).
			//OrderBy("A5").
				Where(" month = ?", monthNum).
				OrderBy("month").
				Find(&cunHuoList)

			//fmt.Println("cunHuoList:",len(cunHuoList),cunHuoList)
			//fmt.Println(len(cunHuoList))

			checkError(err, "GuangZhouSrc list error!")

			var cunHuoCount int = 0        //数量
			var cunHuoTotal float64 = 0.00 //总价
			var cunHuoPrice float64 = 0.00 //单价

			if len(cunHuoList) > 0 { //当月有购买信息  获取基础信息并计算 数量 总价 单价
				//fmt.Println(startDate, "有数据")
				cunHuoBianMaTongJi["Gongyingshang"] = cunHuoList[0].A21
				cunHuoBianMaTongJi["Mingcheng"] = cunHuoList[0].A7

				for _, cunHuo_Month := range cunHuoList { //单月里每条信息进行统计

					shuliang, err := strconv.Atoi(cunHuo_Month.A13)
					checkError(err, "A13 string to int error")

					jiqidanjia, err := strconv.ParseFloat(cunHuo_Month.A14, 64)
					checkError(err, "A14 string to int float64")

					zhichujine, err := strconv.ParseFloat(cunHuo_Month.A18, 64) //支出金额总价
					checkError(err, "A18 string to int float64")

					//huilv, err := strconv.ParseFloat(cunHuo_Month.A31, 64)
					//checkError(err, "A31 string to int float64")

					//switch cunHuo_Month.A4 { //汇率转换
					//case "JPY": //日元
					//	zhichujine = zhichujine * 0.061
					//	jiqidanjia = jiqidanjia * 0.061
					//case "USD": //美元
					//	zhichujine = zhichujine / huilv
					//	jiqidanjia = jiqidanjia / huilv
					//default:
					//	zhichujine = zhichujine
					//	jiqidanjia = jiqidanjia
					//}

					cunHuoCount += shuliang //累加数量

					cunHuoTotal += zhichujine //累加支出金额

					if (cunHuoCount > 0) {
						cunHuoPrice = cunHuoTotal / float64(cunHuoCount) //计算平均单价
					}
					//读取基期单价
					if basePrice == 0.00 { //如果基期单价为0，则设置当前价格为基期单价
						basePrice = jiqidanjia
					}

				}
			}
			//准备数据

			numberOfMonth := strconv.FormatInt(int64(monthNum), 10) //int64转string

			cunHuoBianMaTongJi["M"+numberOfMonth+"shuliang"] = cunHuoCount
			cunHuoBianMaTongJi["M"+numberOfMonth+"danjia"] = strconv.FormatFloat(cunHuoPrice, 'f', 2, 64)
			cunHuoBianMaTongJi["M"+numberOfMonth+"caigoujine"] = strconv.FormatFloat(cunHuoTotal, 'f', 2, 64)

			danjiajiangjiabili := 0.00  //单价降价比例默认为0
			danjiajiangjiajine := 0.00  //单价降价金额默认为0
			zongjiajiangjiajine := 0.00 //总价降价金额默认为0
			zongjiajiangjiabili := 0.00 //总价降价比例默认为0

			//fmt.Println(basePrice)
			if basePrice > 0.00 { //有基期价格
				//fmt.Println("基期价格：", basePrice)
				if cunHuoCount > 0 { //并且有购买数量
					//fmt.Println("Num：", cunHuoCount)
					danjiajiangjiajine = basePrice - cunHuoPrice                    //计算单价降价金额
					danjiajiangjiabili = danjiajiangjiajine / basePrice             //计算单价降价比例
					zongjiajiangjiajine = float64(cunHuoCount) * danjiajiangjiajine //计算总价降价金额
					zongjiajiangjiabili = zongjiajiangjiajine / cunHuoTotal         //计算总价降价比例
				}
			}

			cunHuoBianMaTongJi["M"+numberOfMonth+"danjiajiangjiajine"] = strconv.FormatFloat(danjiajiangjiajine, 'f', 2, 64)

			cunHuoBianMaTongJi["M"+numberOfMonth+"danjiajiangjiabili"] = strconv.FormatFloat(danjiajiangjiabili, 'f', 4, 64)

			cunHuoBianMaTongJi["M"+numberOfMonth+"zongjiajiangjiajine"] = strconv.FormatFloat(zongjiajiangjiajine, 'f', 2, 64)

			cunHuoBianMaTongJi["M"+numberOfMonth+"zongjiajiangjiabili"] = strconv.FormatFloat(zongjiajiangjiabili, 'f', 4, 64)

		} //单月循环结束

		//fmt.Println(cunHuoBianMaTongJi)

		//准备写入数据库
		calResult := &CalBianMa{}
		err := calResult.FillStruct(cunHuoBianMaTongJi) //map转换成struct   struct名称不能有中文，数字不得在首位，不得有- _之类的连接符、下划线等
		checkError(err, "map2struct error")

		affected, err := engine.Insert(calResult)
		fmt.Println("GuangZhou BianMa inserting...", k, affected)
		checkError(err, "GuangZhou BianMa insert error")

		//不再判断是否存在，直接插入
		//oldCalResult := &CalBianMa{}

		//has, err := engine.Select("ID").Where("BianMa =? ", calResult.Bianma).Get(oldCalResult) //判断是否有记录
		//checkError(err, "Check BianMa error")

		//fmt.Println(oldCalResult)

		//var affected int64 = 0
		//if has { //根据是否有记录情况进行更新或插入操作
		//
		//	affected, err = engine.Id(oldCalResult.Id).Update(calResult)
		//	fmt.Println("updating...")
		//	checkError(err, "BianMa update error")
		//} else {
		//	affected, err = engine.Insert(calResult)
		//	fmt.Println("inserting...")
		//	checkError(err, "BianMa insert error")
		//}
		//
		//fmt.Println(affected)
		//break
		//defer engine.Close() //关闭数据库连接

	}
}

//存货编码 按月统计 第一个表格
func calBianMaNanChang() {

	var err error

	var bianMaList []NanChangSrc
	var cunHuoList []NanChangSrc

	err = engine.Table(new(NanChangSrc)).Select("distinct A12").Find(&bianMaList)
	checkError(err, "Table src A12 error!")

	cunHuoBianMaTongJi := make(map[string]interface{})

	//fmt.Println("bianMaList:",bianMaList)

	for k, bianMa := range bianMaList { //取出所有不同的 存货编码A12  进行统计
		//fmt.Println(k, bianMa.A12)

		cunHuoBianMaTongJi["Bianma"] = bianMa.A12 //获取编码

		var basePrice float64 = 0.00 //基期单价

		for monthNum := 1; monthNum < 13; monthNum++ { //按月统计 存货编码A12 情况

			var cunHuoCount int = 0        //数量
			var cunHuoTotal float64 = 0.00 //总价
			var cunHuoPrice float64 = 0.00 //单价

			cunHuoList = make([]NanChangSrc, 0) //用Find方法循环要每次进行清理

			monthName := strconv.FormatInt(int64(monthNum), 10) + "月" //原始数据为 "X月"

			err = engine.Table(new(NanChangSrc)).Select("A7,A10,A12,A13,A17,A20,A21").
				Where("A12 = ?", bianMa.A12).
				Where("A7=?", monthName).
				OrderBy("A7").Find(&cunHuoList)

			checkError(err, "src list error!")

			//fmt.Println(len(cunHuoList))

			if len(cunHuoList) > 0 {
				cunHuoBianMaTongJi["Gongyingshang"] = cunHuoList[0].A10
				cunHuoBianMaTongJi["Mingcheng"] = cunHuoList[0].A13
			}

			for _, cunHuo_Month := range cunHuoList { //单月统计

				A17Int, err := strconv.Atoi(cunHuo_Month.A17)
				checkError(err, "A17 conver error")
				A21Float64, err := strconv.ParseFloat(cunHuo_Month.A21, 64)
				checkError(err, "A21 conver error")
				cunHuoCount += int(A17Int) //累加数量
				cunHuoTotal += A21Float64  //累加支出金额

			}

			if cunHuoCount > 0 {
				cunHuoPrice = cunHuoTotal / float64(cunHuoCount) //计算平均单价
				if basePrice == 0.00 { //如果基期单价为0，则设置当前价格为基期单价
					basePrice = cunHuoPrice
				}

			} else {
				cunHuoPrice = 0.00
			}

			numberOfMonth := strconv.FormatInt(int64(monthNum), 10) //int64转string

			cunHuoBianMaTongJi["M"+numberOfMonth+"shuliang"] = cunHuoCount
			cunHuoBianMaTongJi["M"+numberOfMonth+"danjia"] = strconv.FormatFloat(cunHuoPrice, 'f', 2, 64)
			cunHuoBianMaTongJi["M"+numberOfMonth+"caigoujine"] = strconv.FormatFloat(cunHuoTotal, 'f', 2, 64)

			danjiajiangjiabili := 0.00  //单价降价比例默认为0
			danjiajiangjiajine := 0.00  //单价降价金额默认为0
			zongjiajiangjiajine := 0.00 //总价降价金额默认为0
			zongjiajiangjiabili := 0.00 //总价降价比例默认为0

			//fmt.Println(basePrice)
			if basePrice != 0.00 { //当有基期价格
				if cunHuoCount > 0 { //当月有购买数量
					danjiajiangjiajine = basePrice - cunHuoPrice                    //计算单价降价金额
					danjiajiangjiabili = danjiajiangjiajine / basePrice             //计算单价降价比例
					zongjiajiangjiajine = float64(cunHuoCount) * danjiajiangjiajine //计算总价降价金额
					zongjiajiangjiabili = zongjiajiangjiajine / cunHuoTotal         //计算总价降价比例
				}
			}

			cunHuoBianMaTongJi["M"+numberOfMonth+"danjiajiangjiajine"] = strconv.FormatFloat(danjiajiangjiajine, 'f', 2, 64)

			cunHuoBianMaTongJi["M"+numberOfMonth+"danjiajiangjiabili"] = strconv.FormatFloat(danjiajiangjiabili, 'f', 4, 64)

			cunHuoBianMaTongJi["M"+numberOfMonth+"zongjiajiangjiajine"] = strconv.FormatFloat(zongjiajiangjiajine, 'f', 2, 64)

			cunHuoBianMaTongJi["M"+numberOfMonth+"zongjiajiangjiabili"] = strconv.FormatFloat(zongjiajiangjiabili, 'f', 4, 64)

		}
		//fmt.Println(cunHuoBianMaTongJi)

		//fmt.Println("cunHuoBianMaTongJi:",cunHuoBianMaTongJi)
		calResult := &CalBianMa{}
		err := calResult.FillStruct(cunHuoBianMaTongJi) //map转换成struct   struct名称不能有中文，数字不得在首位，不得有- _之类的连接符、下划线等
		checkError(err, "map2struct error")


		_, err = engine.Insert(calResult)
		fmt.Println("NanChang BianMa inserting...", k)
		checkError(err, "NanChang BianMa insert error")

		//break
	}
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, " >>> Detail:", err)
		os.Exit(3)
	}
}
