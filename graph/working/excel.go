package working

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/shubhacker/gqlgen-todos/graph/entity"
	"log"
	"strconv"
)


func CreateExcel(entity []entity.UserExcel )string{
f := excelize.NewFile()
index := f.NewSheet("sheet 1")
excelize.TitleToNumber("AK")
	op :=1
for _,data := range entity{
	inc := "A"+strconv.Itoa(op)
	f.SetCellValue("sheet 1",inc, data.UserName)
	op += 1
}
//f.SetCellValue("sheet 1","A3", "Hey There!")
//f.SetCellValue("sheet 1", "B1",10)
f.SetActiveSheet(index)
if err := f.SaveAs("Test.xlsx");err != nil{
	log.Println(err.Error())
}
return "Excel Created"
}
