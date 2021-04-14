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
	f.SetCellValue("sheet 1","A"+strconv.Itoa(1),"UserName")
	f.SetCellValue("sheet 1","B"+strconv.Itoa(1),"Password")
	f.SetCellValue("sheet 1","C"+strconv.Itoa(1),"UserRole")
	f.SetCellValue("sheet 1","D"+strconv.Itoa(1),"Email")
	f.SetCellValue("sheet 1","E"+strconv.Itoa(1),"MobileNo")
	op :=2
for _,data := range entity{
	userColumn := "A"+strconv.Itoa(op)
	passwordColumn := "B"+strconv.Itoa(op)
	userRoleColumn := "C"+strconv.Itoa(op)
	emailColumn := "D"+strconv.Itoa(op)
	MobNoColumn := "E"+strconv.Itoa(op)
	f.SetCellValue("sheet 1",userColumn, data.UserName)
	f.SetCellValue("sheet 1",passwordColumn, data.Password)
	f.SetCellValue("sheet 1",userRoleColumn, data.UserRole)
	f.SetCellValue("sheet 1",emailColumn, data.Email)
	f.SetCellValue("sheet 1",MobNoColumn, data.MobNo)
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
func CreateExcelForMaster()string{
	log.Println("CreateExcelForMaster()")
	f := excelize.NewFile()
	index:= f.NewSheet("sheet 1")
	f.SetActiveSheet(index)
	return ""
}
