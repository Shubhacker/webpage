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
func CreateExcelForMaster(entity []entity.MasterExcel)(string){
	log.Println("CreateExcelForMaster()")
	f := excelize.NewFile()
	op := 1
	UserSheet:= f.NewSheet("User Sheet")
	VideoSheet := f.NewSheet("Video Sheet")
	BlogSheet := f.NewSheet("Blog Sheet")
	for _,data := range entity{
		Column1 := "A"+strconv.Itoa(op)
		Column2 := "B"+strconv.Itoa(op)
		Column3 := "C"+strconv.Itoa(op)
		Column4 := "D"+strconv.Itoa(op)
		Column5 := "E"+strconv.Itoa(op)
		Column6 := "F"+strconv.Itoa(op)
		f.SetCellValue("User Sheet",Column1,data.UserName)
		f.SetCellValue("User Sheet",Column2,data.Password)
		f.SetCellValue("User Sheet",Column3,data.UserRole)
		f.SetCellValue("User Sheet",Column4,data.Email)
		f.SetCellValue("User Sheet",Column5,data.MobNo)
		f.SetCellValue("Video Sheet",Column1,data.VideoTopic)
		f.SetCellValue("Video Sheet",Column2,data.VideoLink)
		f.SetCellValue("Video Sheet",Column3,data.BookName)
		f.SetCellValue("Video Sheet",Column4,data.BookLink)
		f.SetCellValue("Video Sheet",Column5,data.ToolName)
		f.SetCellValue("Video Sheet",Column6,data.VideoLink)
		f.SetCellValue("Blog Sheet",Column1,data.BlogText)
		f.SetCellValue("Blog Sheet",Column2,data.ReferenceLink)
		op += 1
	}

	f.SetActiveSheet(UserSheet)
	f.SetActiveSheet(VideoSheet)
	f.SetActiveSheet(BlogSheet)
	if err := f.SaveAs("Test.xlsx");err != nil{
		log.Println(err.Error())
	}
	return "Excel Created Successfully"
}
