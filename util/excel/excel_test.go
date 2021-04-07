package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

func TestExcelFile(t *testing.T) {
	// 创建控制文件
	f := excelize.NewFile()
	sheet := "测试表单"
	// 设置文件单元 (?? 好像可有可无)
	sheetFile := f.NewSheet(sheet)

	//设置默认工作表单
	f.SetActiveSheet(sheetFile)

	// 填充数据
	f.SetCellValue(sheet, "A1", "ceshi")
	f.SetCellValue(sheet, "A2", "ceshi2")
	f.SetCellValue(sheet, "AA2", "ceshi2")

	f.SetCellValue(sheet, "B2", "=HYPERLINK(\"./image.jpg\",\"一班\")")

	// 插入图表
	if err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		println(err.Error())
		return
	}

	// 插入图片
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		println(err.Error())
	}
	// 在工作表中插入图片，并设置图片的缩放比例
	if err := f.AddPicture("Sheet1", "D2", "image.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		println(err.Error())
	}
	// 在工作表中插入图片，并设置图片的打印属性
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`); err != nil {
		println(err.Error())
	}

	// 保存文件
	if err := f.SaveAs("../../test.xlsx"); err != nil {
		t.Fatal(err)
	}

	// 判断文件
	// 读取文件
	// 删除
}
